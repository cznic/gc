// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"bytes"
	"go/token"
	"io"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/cznic/golex/lex"
	"github.com/cznic/mathutil"
	"github.com/cznic/xc"
)

const (
	ccEOF = iota + 0x80
	ccOther
	ccLTLT // «
	ccGTGT // »
)

var (
	buildMark           = []byte("// +build")
	errCheckDisbled     = []byte("////")
	errCheckMark1       = []byte("// ERROR ")
	errCheckMark2       = []byte("// GC_ERROR ")
	generalCommentEnd   = []byte("*/")
	generalCommentStart = []byte("/*")
)

func runeClass(r rune) int {
	switch {
	case r == lex.RuneEOF:
		return ccEOF
	case r < 0x80:
		return int(r)
	default:
		return ccOther
	}
}

func runeClassGenerics(r rune) int {
	switch {
	case r == lex.RuneEOF:
		return ccEOF
	case r < 0x80:
		return int(r)
	case r == '«':
		return ccLTLT
	case r == '»':
		return ccGTGT
	default:
		return ccOther
	}
}

type lexer struct {
	*Context
	*lex.Lexer
	build                bool // Whether build tags, if any, satisfied.
	closed               bool // Error limit reached.
	constExpr            *ExpressionList
	dotImports           []*ImportDeclaration
	fileScope            *Scope
	firstConstSpec       bool
	hold                 xc.Token // Lexer state machine ({TX,RX}CHAN).
	holdState            int      // Lexer state machine ({TX,RX}CHAN).
	imports              []*ImportDeclaration
	iota                 int64
	lbr                  bool   // Lexer state machine (BODY).
	lbrStack             []bool // Lexer state machine (BODY).
	lbrace               int    // Lexer state machine (BODY).
	lbraceStack          []int  // Lexer state machine (BODY).
	lexPrev              rune
	lookahead            xc.Token // lookahead.Char == yyParse yychar.
	name                 string
	pkg                  *Package
	resolutionScope      *Scope
	resolutionScopeStack []*Scope
	scanCharPrev         rune // Lexer state machine (semicolon insertion).
	scope                *Scope
	seenPackage          bool // Lexer state machine (check PACKAGE is first).
	switchDecl           []xc.Token
	unboundImports       []*ImportDeclaration
}

func newLexer(nm string, sz int, r io.RuneReader, pkg *Package) (*lexer, error) {
	ctx := pkg.Context
	rc := runeClass
	if ctx.options.enableGenerics {
		rc = runeClassGenerics
	}

	lx := &lexer{
		Context:         ctx,
		build:           true,
		fileScope:       newScope(FileScope, pkg.Scope),
		name:            nm,
		pkg:             pkg,
		resolutionScope: pkg.Scope,
		scope:           pkg.Scope,
	}
	lx0, err := lex.New(
		xc.FileSet.AddFile(nm, -1, sz),
		r,
		lex.ErrorFunc(func(pos token.Pos, msg string) {
			lx.errPos(pos, msg)
		}),
		lex.RuneClass(rc),
		lex.BOMMode(lex.BOMIgnoreFirst),
	)
	if err != nil {
		return nil, ctx.errors(err)
	}

	lx.Lexer = lx0
	return lx, nil
}

func (lx *lexer) close(n Node, format string, arg ...interface{}) {
	lx.err(n, format, arg...)
	lx.closed = true
}

func (lx *lexer) err(n Node, format string, arg ...interface{}) {
	lx.errPos(n.Pos(), format, arg...)
}

func (lx *lexer) errPos(pos token.Pos, format string, arg ...interface{}) {
	lx.closed = lx.pkg.Context.errPos(pos, format, arg...)
}

func (lx *lexer) pushScope() *Scope {
	lx.resolutionScopeStack = append(lx.resolutionScopeStack, lx.resolutionScope)
	old := lx.scope
	lx.scope = newScope(BlockScope, old)
	lx.resolutionScope = lx.scope
	return lx.scope
}

func (lx *lexer) popScope() *Scope {
	old := lx.scope
	if old.Kind == PackageScope {
		lx.err(lx.lookahead, "cannot pop scope")
		return old
	}

	lx.scope = old.Parent
	n := len(lx.resolutionScopeStack)
	lx.resolutionScope = lx.resolutionScopeStack[n-1]
	lx.resolutionScopeStack = lx.resolutionScopeStack[:n-1]
	return lx.scope
}

func (lx *lexer) checkComment() {
	b := lx.TokenBytes(nil)
	if !lx.seenPackage && position(lx.First.Pos()).Column == 1 && bytes.HasPrefix(b, buildMark) {
		lx.buildDirective(b)
		return
	}

	t := lx.pkg.Context.test
	if t != nil && strings.HasPrefix(lx.name, "testdata") {
		lx.errorcheckDirective()
	}
}

func (lx *lexer) buildDirective(b []byte) {
	if !lx.build {
		return
	}

	ctx := lx.pkg.Context
	s := string(b[len(buildMark):])
	s = strings.Replace(s, "\t", " ", -1)
	for _, term := range strings.Split(s, " ") { // term || term
		if term = strings.TrimSpace(term); term == "" {
			continue
		}

		val := true
		for _, factor := range strings.Split(term, ",") { // factor && factor
			if factor = strings.TrimSpace(factor); factor == "" {
				continue
			}

			not := factor[0] == '!'
			if not {
				factor = strings.TrimSpace(factor[1:])
			}

			if factor == "" {
				lx.err(lx.First, "build tags syntax error")
				lx.build = false
				return
			}

			_, ok := ctx.tags[factor]
			if not {
				ok = !ok
			}

			if !ok {
				val = false
				break
			}
		}
		if val {
			return
		}
	}
	lx.build = false
}

func (lx *lexer) errorcheckDirective() {
	t := lx.pkg.Context.test
	if t == nil || !strings.HasPrefix(lx.name, "testdata") {
		return
	}

	ch := lex.NewChar(lx.First.Pos(), ERRCHECK)
	s := lx.TokenBytes(func(buf *bytes.Buffer) {
		for _, c := range lx.Token() {
			if r := c.Rune; r != lex.RuneEOF {
				buf.WriteRune(r)
			}
		}
	})
	if bytes.HasPrefix(s, generalCommentStart) {
		s = s[len(generalCommentStart):]
	}
	if bytes.HasSuffix(s, generalCommentEnd) {
		s = s[:len(s)-len(generalCommentEnd)]
	}
	s = bytes.TrimSpace(s)
	n := len(errCheckMark1)
	i := bytes.LastIndex(s, errCheckMark1)
	j := bytes.LastIndex(s, errCheckDisbled)
	if i < 0 {
		i = bytes.LastIndex(s, errCheckMark2)
		n = len(errCheckMark2)
		if i < 0 {
			return // No check found.
		}
	}

	if j >= 0 && j < i {
		return // Check disabled.
	}

	s = s[i+n:]
	s = bytes.TrimSpace(s)
	t.errChecksMu.Lock()
	t.errChecks = append(t.errChecks, xc.Token{Char: ch, Val: dict.ID(s)})
	t.errChecksMu.Unlock()
}

var colon = []byte{':'}

func (lx *lexer) lineDirective() {
	if position(lx.First.Pos()).Column != 1 {
		return
	}

	line := lx.TokenBytes(nil)
	i := bytes.LastIndex(line, colon)
	n, err := strconv.ParseUint(string(line[i+1:len(line)-1]), 10, mathutil.IntBits-1)
	if err != nil || n <= 0 {
		return
	}

	fn := string(bytes.TrimSpace(line[len("//line "):i]))
	if fn != "" {
		fn = filepath.Clean(fn)
		if !filepath.IsAbs(fn) {
			fn = filepath.Join(lx.pkg.Directory, fn)
		}
	}
	lx.File.AddLineInfo(lx.Offset()-1, fn, int(n))
}

func (lx *lexer) fixLBR() {
	if lx.Token()[0].Rune == '}' {
		lx.lbr = true
		return
	}

	if lx.lbrace != 0 {
		lx.lbrace--
	}
	lx.lbraceStack = append(lx.lbraceStack, lx.lbrace)
	lx.lbrace = 1
}

func (lx *lexer) scanChar() lex.Char {
again:
	if lx.closed {
		return lex.NewChar(lx.First.Pos(), 0)
	}

	r := rune(lx.scan())
	switch {
	case r == '\n', r == lex.RuneEOF:
		switch lx.scanCharPrev {
		case IDENTIFIER, INT_LIT, FLOAT_LIT, IMAG_LIT, CHAR_LIT,
			STRING_LIT, BREAK, CONTINUE, FALLTHROUGH, RETURN, INC,
			DEC, ')', ']', '}', GTGT:
			r = ';'
		default:
			switch r {
			case '\n':
				goto again
			case lex.RuneEOF:
				r = 0
			}
		}
	case r < ' ':
		lx.err(lx.First, "invalid character %#U", r)
		goto again
	case r == BAD_FLOAT_LIT:
		lx.err(lx.First, "malformed floating point constant")
		r = FLOAT_LIT
	}
	lx.scanCharPrev = r
	return lex.NewChar(lx.First.Pos(), r)
}

func (lx *lexer) scanToken() xc.Token {
	c := lx.scanChar()
	var val int
	switch c.Rune {
	case IDENTIFIER:
		val = dict.ID(lx.TokenBytes(func(buf *bytes.Buffer) {
			for i, c := range lx.Token() {
				r := c.Rune
				d := unicode.IsDigit(r)
				switch {
				case d && i == 0:
					lx.err(c, "identifier cannot begin with digit")
				case r >= 0x80 && !unicode.IsLetter(r):
					lx.err(c, "invalid identifier character %#U", r)
				}
				buf.WriteRune(r)
			}
		}))
	case INT_LIT, FLOAT_LIT, IMAG_LIT, CHAR_LIT, STRING_LIT:
		val = dict.ID(lx.TokenBytes(nil))
	}
	t := xc.Token{Char: c, Val: val}
	if !lx.seenPackage {
		if t.Rune != PACKAGE {
			lx.close(t, "package statement must be first")
			t.Rune = 0
		}
		lx.seenPackage = true
	}
	return t
}

// Implements yyLexer.
func (lx *lexer) Error(msg string) {
	msg = strings.Replace(msg, "$end", "EOF", -1)
	switch {
	case lx.pkg.Context.test != nil:
		const s = ", expected "
		parts := strings.Split(msg, ", expected ")
		if p := parts[0]; strings.HasPrefix(p[:len(p)-1], "unexpected '") && strings.HasSuffix(p, "'") {
			p = p[:len(p)-1]
			p = strings.Replace(p, "'", "", 1)
			if len(parts) == 1 {
				msg = p
				break
			}

			msg = p + s + parts[1]
		}
	}
	lx.err(lx.lookahead, "%s", msg)
}

// Implements yyLexer.
func (lx *lexer) Lex(lval *yySymType) int {
	const (
		idle = iota
		holdChan
		holdComm
		holdToken
	)

	var t xc.Token

again:
	switch lx.holdState {
	case idle:
		switch t = lx.scanToken(); t.Rune {
		case CHAN:
			lx.hold = t
			lx.holdState = holdChan
			goto again
		case COMM:
			lx.hold = t
			lx.holdState = holdComm
			goto again
		}
	case holdChan:
		switch t = lx.scanToken(); t.Rune {
		case CHAN:
			t, lx.hold = lx.hold, t
		case COMM:
			t.Rune = TXCHAN
			t, lx.hold = lx.hold, t
			lx.holdState = holdToken
		default:
			t, lx.hold = lx.hold, t
			lx.holdState = holdToken
		}
	case holdComm:
		switch t = lx.scanToken(); t.Rune {
		case CHAN:
			t, lx.hold = lx.hold, t
			t.Rune = RXCHAN
			lx.holdState = holdToken
		default:
			t, lx.hold = lx.hold, t
			lx.holdState = holdToken
		}
	case holdToken:
		t = lx.hold
		lx.holdState = idle
	default:
		panic("internal error")
	}

	r := t.Rune
	switch r {
	case FOR, IF, SELECT, SWITCH:
		lx.lbr = true
	case '(', '[':
		if lx.lbr || len(lx.lbrStack) != 0 {
			lx.lbrStack = append(lx.lbrStack, lx.lbr)
			lx.lbr = false
		}
	case ')', ']':
		if n := len(lx.lbrStack); n > 0 {
			lx.lbr = lx.lbrStack[n-1]
			lx.lbrStack = lx.lbrStack[:n-1]
		}
	case '{':
		if lx.lbrace != 0 {
			lx.lbrace++
		}
		if lx.lbr {
			lx.lbr = false
			r = BODY
		}
	case '}':
		if lx.lbrace > 0 {
			lx.lbrace--
			if lx.lbrace == 0 {
				n := len(lx.lbraceStack)
				lx.lbrace = lx.lbraceStack[n-1]
				lx.lbraceStack = lx.lbraceStack[:n-1]
				lx.lbr = true
			}
		}
	}

	switch lx.lexPrev {
	case CONST:
		lx.constExpr = nil
		lx.firstConstSpec = true
		lx.iota = 0
	case FUNC:
		rs := lx.resolutionScope
		s := lx.pushScope()
		s.isFnScope = true
		s.isMergeScope = true
		lx.resolutionScope = rs
	case IF, FOR, SWITCH: // Implicit blocks.
		lx.pushScope()
	case CASE, DEFAULT: // Implicit blocks.
		lx.pushScope()
		var t xc.Token
		if n := len(lx.switchDecl); n > 0 {
			t = lx.switchDecl[n-1]
		}
		if t.IsValid() {
			lx.scope.declare(lx, newVarDeclaration(-1, t, nil, nil, lx.lookahead.Pos())) //TODO
		}
	}
	lx.lexPrev = r

	lx.lookahead = t
	lval.Token = t
	return int(r)
}

// Implements yyLexerEx.
func (lx *lexer) Reduced(rule, state int, lval *yySymType) (stop bool) {
	t := lx.pkg.Context.test
	if t == nil {
		return false
	}

	if rule != t.exampleRule {
		return false
	}

	switch x := lval.node.(type) {
	case interface {
		fragment() interface{}
	}:
		t.exampleAST = x.fragment()
	default:
		t.exampleAST = x
	}
	return true
}
