// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"bytes"
	gotoken "go/token"
	"unicode/utf8"

	"github.com/cznic/token"
)

// Non ASCII character classes.
const (
	classEOF = iota + 0x80
	classNonASCII
	classLTLT
	classGTGT
	classBOM
	classNext
)

const (
	maxTokenToken = gotoken.VAR // 85
)

// Additional tokens.
const (
	tokenNL   = iota + maxTokenToken + 1
	tokenLTLT // «
	tokenGTGT // »
	tokenBODY
	tokenBOM
)

var (
	bom           = []byte("\ufeff")
	lineDirective = []byte("//line")
	nlLit         = []byte{'\n'}

	semiTriggerTokens = [...]bool{
		gotoken.BREAK:       true,
		gotoken.CHAR:        true,
		gotoken.CONTINUE:    true,
		gotoken.DEC:         true,
		gotoken.FALLTHROUGH: true,
		gotoken.FLOAT:       true,
		gotoken.IDENT:       true,
		gotoken.IMAG:        true,
		gotoken.INC:         true,
		gotoken.INT:         true,
		gotoken.RBRACE:      true,
		gotoken.RBRACK:      true,
		gotoken.RETURN:      true,
		gotoken.RPAREN:      true,
		gotoken.STRING:      true,
		tokenGTGT:           true,
	}
)

type lexer struct {
	commentHandler func(off int, lit []byte)
	commentOfs     int
	errHandler     func(position token.Position, msg string, args ...interface{})
	errorCount     int // Number of errors encountered.
	file           *token.File
	fname          string
	lit            []byte
	off            int // Next byte offset.
	prev           gotoken.Token
	sourceFile     *SourceFile
	src            []byte

	b byte // Current byte.
	c byte // Current class.
}

func newLexer(sourceFile *SourceFile, src []byte) *lexer {
	l := &lexer{
		sourceFile: sourceFile,
		src:        src,
	}
	if sourceFile != nil {
		l.file = sourceFile.File
	}
	if bytes.HasPrefix(src, bom) {
		l.off = len(bom)
	}
	l.n()
	return l
}

func (l *lexer) init(file *token.File, src []byte) *lexer {
	l.commentOfs = -1
	l.errorCount = 0
	l.lit = nil
	l.off = 0
	l.prev = tokenNL
	l.src = src
	l.c = classNext
	l.fname = ""
	l.file = file
	if l.sourceFile == nil {
		l.sourceFile = &SourceFile{File: file}
	}
	if bytes.HasPrefix(src, bom) {
		l.off = 3
	}
	l.n()
	return l
}

func (l *lexer) err(position token.Position, msg string, args ...interface{}) {
	l.errorCount++
	if l.errHandler != nil {
		l.errHandler(position, msg, args...)
	}
}

func (l *lexer) pos(off int) token.Pos { return l.file.Pos(off) }

func (l *lexer) position(off int) token.Position { return l.file.Position(l.pos(off)) }

// Token returns the last scanned Token. 'off' must be the offset returned from
// last call to Scan().
func (l *lexer) Token(off int) Token {
	return Token{npos: npos{tpos: l.pos(off), sf: l.sourceFile}, Val: string(l.lit)}
}

// Returns class.
func (l *lexer) n() byte { // n == next
	if l.off == len(l.src) {
		if l.c != classEOF {
		}
		l.c = classEOF
		l.b = 0xff // Invalid UTF-8 byte.
		l.lit = nil
		return l.c
	}

	l.b = l.src[l.off]
	l.off++
	l.c = l.b
	if l.b > 0x7f {
		l.c = classNonASCII
		switch l.b {
		case 0xc2: // {"«","»"}[0]
			if l.off < len(l.src) {
				switch l.src[l.off] {
				case 0xab:
					l.c = classLTLT
				case 0xbb:
					l.c = classGTGT
				}
			}
		case 0xef: // BOM[0]
			if l.off+1 < len(l.src) && l.src[l.off] == 0xbb && l.src[l.off+1] == 0xbf {
				l.err(l.position(l.off-1), "illegal BOM")
				l.c = classBOM
			}
		}
	} else if l.b == 0 {
		l.err(l.position(l.off-1), "illegal character NUL")
	}
	return l.c
}

func (l *lexer) octals(max int) (n int) {
	for max != 0 && l.c >= '0' && l.c <= '7' {
		l.n()
		n++
		max--
	}
	return n
}

func (l *lexer) decimals() (r bool) {
	for l.c >= '0' && l.c <= '9' {
		r = true
		l.n()
	}
	return r
}

func (l *lexer) exponent(off int) (int, gotoken.Token) {
	switch l.c {
	case 'e', 'E':
		switch l.n() {
		case '+', '-':
			l.n()
		}
		if !l.decimals() {
			l.err(l.file.Position(l.file.Pos(off)), "illegal floating-point exponent")
		}
	}
	switch l.c {
	case 'i':
		l.n()
		return off, gotoken.IMAG
	}

	return off, gotoken.FLOAT
}

func (l *lexer) hexadecimals(max int) (n int) {
	for max != 0 && (l.c >= '0' && l.c <= '9' || l.c >= 'a' && l.c <= 'f' || l.c >= 'A' && l.c <= 'F') {
		l.n()
		n++
		max--
	}
	return n
}

func isIdentNext(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' || c >= '0' && c <= '9' || c == classNonASCII
}

func (l *lexer) ident() gotoken.Token {
	for l.c >= 'a' && l.c <= 'z' || l.c >= 'A' && l.c <= 'Z' || l.c == '_' || l.c >= '0' && l.c <= '9' || l.c == classNonASCII {
		l.n()
	}
	return gotoken.IDENT
}

func (l *lexer) skip() rune {
	if c := l.c; c < 0x80 {
		l.n()
		return rune(c)
	}

	if l.c == classEOF {
		return -1
	}

	r, sz := utf8.DecodeRune(l.src[l.off-1:])
	n := sz - 1
	l.off += n
	l.n()
	return r
}

func (l *lexer) stringEscFail() bool {
	switch l.c {
	case '\n':
		l.err(l.position(l.off-1), "illegal character %#U in escape sequence", l.c)
	case '"':
		l.err(l.position(l.off-1), "illegal character %#U in escape sequence", l.c)
		l.n()
		return true
	case '\\':
		l.err(l.position(l.off-1), "illegal character %#U in escape sequence", l.c)
		fallthrough
	case classEOF:
		l.err(l.position(l.off-1), "escape sequence not terminated")
	default:
		l.err(l.position(l.off-1), "illegal character %#U in escape sequence", l.skip())
	}
	return false
}

func (l *lexer) charEscFail() {
	switch l.c {
	case '\n':
		l.err(l.position(l.off-1), "illegal character %#U in escape sequence", l.c)
	case '\\':
		l.err(l.position(l.off-1), "illegal character %#U in escape sequence", l.c)
		l.n()
		fallthrough
	case classEOF:
		l.err(l.position(l.off-1), "escape sequence not terminated")
	default:
		l.err(l.position(l.off-1), "illegal character %#U in escape sequence", l.skip())
	}
}

// Scan returns the next token and its offset.
func (l *lexer) Scan() (off int, tok gotoken.Token) {
skip:
	off, tok = l.scan0()
	if tok == gotoken.COMMENT {
		if l.commentHandler != nil {
			end := l.off
			if l.c != classEOF {
				end--
			}
			l.commentHandler(off, l.src[off:end])
		}
		if l.commentOfs < 0 {
			l.commentOfs = off
		}
		goto skip
	}

	co := l.commentOfs
	l.commentOfs = -1
	if tok == tokenNL || tok == gotoken.EOF {
		if p := int(l.prev); p >= 0 && p < len(semiTriggerTokens) && semiTriggerTokens[l.prev] {
			if co >= 0 {
				off = co
			}
			l.prev = tok
			l.lit = nlLit
			if tok == gotoken.EOF && co < 0 {
				off = l.off
			}
			return off, gotoken.SEMICOLON
		}

		if tok == gotoken.EOF {
			l.lit = nil
			return l.off, gotoken.EOF
		}

		goto skip
	}

	if tok != gotoken.ILLEGAL {
		l.prev = tok
	}
	end := l.off
	if l.c != classEOF {
		end--
	}
	l.lit = l.src[off:end]
	return off, tok
}

func (l *lexer) scan0() (off int, tok gotoken.Token) {
skip:
	off = l.off - 1
	switch l.c {
	case '\t', '\r', ' ':
		l.n()
		goto skip
	case '\n':
		if l.off != len(l.src) {
			l.file.AddLine(l.off)
		}
		l.n()
		return off, tokenNL
	case '!':
		if l.n() == '=' {
			l.n()
			return off, gotoken.NEQ
		}

		return off, gotoken.NOT
	case '"':
		l.n()
	more:
		switch l.c {
		case '\n', classEOF:
			l.err(l.position(l.off-1), "string literal not terminated")
		case '"':
			l.n()
		case '\\':
			switch l.n() {
			case '\n':
				l.err(l.position(l.off-1), "unknown escape sequence")
			case '0', '1', '2', '3', '4', '5', '6', '7':
				if l.octals(3) < 3 && l.stringEscFail() {
					return off, gotoken.STRING
				}
			case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '"':
				l.n()
			case 'u':
				l.n()
				if l.hexadecimals(4) < 4 && l.stringEscFail() {
					return off, gotoken.STRING
				}
			case 'U':
				l.n()
				if l.hexadecimals(8) < 8 && l.stringEscFail() {
					return off, gotoken.STRING
				}
			case 'x':
				l.n()
				if l.hexadecimals(2) < 2 && l.c == classEOF {
					l.err(l.position(l.off-1), "escape sequence not terminated")
				}
			case classEOF:
				l.err(l.position(l.off-1), "escape sequence not terminated")
			default:
				l.err(l.position(l.off-1), "unknown escape sequence")
				l.skip()
			}
			goto more
		default:
			l.skip()
			goto more
		}
		return off, gotoken.STRING
	case '%':
		if l.n() == '=' {
			l.n()
			return off, gotoken.REM_ASSIGN
		}

		return off, gotoken.REM
	case '&':
		switch l.n() {
		case '^':
			if l.n() == '=' {
				l.n()
				return off, gotoken.AND_NOT_ASSIGN
			}

			return off, gotoken.AND_NOT
		case '=':
			l.n()
			return off, gotoken.AND_ASSIGN
		case '&':
			l.n()
			return off, gotoken.LAND
		}

		return off, gotoken.AND
	case '\'':
		switch l.n() {
		case '\n', classEOF:
			l.err(l.position(l.off-1), "invalid character literal (missing closing ')")
			return off, gotoken.CHAR
		case '\'':
			l.err(l.position(l.off-1), "illegal rune literal")
			l.n()
			return off, gotoken.CHAR
		case '\\':
			switch l.n() {
			case '\n':
				l.err(l.position(l.off-1), "unknown escape sequence")
				return off, gotoken.CHAR
			case '0', '1', '2', '3', '4', '5', '6', '7':
				if l.octals(3) < 3 {
					l.charEscFail()
					return off, gotoken.CHAR
				}
			case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '\'':
				l.n()
			case 'u':
				l.n()
				if l.hexadecimals(4) < 4 {
					l.charEscFail()
					return off, gotoken.CHAR
				}
			case 'U':
				l.n()
				if l.hexadecimals(8) < 8 {
					l.charEscFail()
					return off, gotoken.CHAR
				}
			case 'x':
				l.n()
				if l.hexadecimals(2) < 2 {
					l.charEscFail()
					return off, gotoken.CHAR
				}
			case classEOF:
				l.err(l.position(l.off-1), "escape sequence not terminated")
				return off, gotoken.CHAR
			default:
				l.err(l.position(l.off-1), "unknown escape sequence")
				l.skip()
				return off, gotoken.CHAR
			}
		default:
			l.skip()
		}
		switch l.c {
		case '\n', classEOF:
			l.err(l.position(l.off-1), "rune literal not terminated")
		case '\\':
			l.err(l.position(l.off-1), "escape sequence not terminated")
			l.n()
		case '\'':
			l.n()
		default:
			l.err(l.position(l.off-1), "rune literal not terminated")
		search:
			for {
				switch l.c {
				case '\n', classEOF:
					break search
				case '\'':
					l.n()
					break search
				default:
					l.n()
				}
			}
		}
		return off, gotoken.CHAR
	case '(':
		l.n()
		return off, gotoken.LPAREN
	case ')':
		l.n()
		return off, gotoken.RPAREN
	case '*':
		if l.n() == '=' {
			l.n()
			return off, gotoken.MUL_ASSIGN
		}

		return off, gotoken.MUL
	case '+':
		switch l.n() {
		case '=':
			l.n()
			return off, gotoken.ADD_ASSIGN
		case '+':
			l.n()
			return off, gotoken.INC
		}

		return off, gotoken.ADD
	case ',':
		l.n()
		return off, gotoken.COMMA
	case '-':
		switch l.n() {
		case '=':
			l.n()
			return off, gotoken.SUB_ASSIGN
		case '-':
			l.n()
			return off, gotoken.DEC
		}

		return off, gotoken.SUB
	case '.':
		switch l.n() {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.decimals()
			return l.exponent(off)
		case '.':
			switch l.n() {
			case '.':
				l.n()
				return off, gotoken.ELLIPSIS
			default:
				l.off--
				l.c = '.'
				return off, gotoken.PERIOD
			}
		default:
			return off, gotoken.PERIOD
		}
	case '/':
		switch l.n() {
		case '/':
			for l.n() != '\n' && l.c != classEOF {
			}
			return off, gotoken.COMMENT
		case '*':
			var hasNL bool
			for l.n(); l.c != classEOF; l.n() {
				switch l.c {
				case '\n':
					hasNL = true
					if l.off != len(l.src) {
						l.file.AddLine(l.off)
					}
				case '*':
				more2:
					switch l.n() {
					case '*':
						goto more2
					case '\n':
						hasNL = true
						if l.off != len(l.src) {
							l.file.AddLine(l.off - 1)
						}
					case '/':
						l.n()
						if hasNL {
							if l.commentHandler != nil {
								end := l.off
								if l.c != classEOF {
									end--
								}
								l.commentHandler(off, l.src[off:end])
							}
							return off, tokenNL
						}

						return off, gotoken.COMMENT
					}
				}
			}
			l.err(l.position(l.off-1), "comment not terminated")
			return off, gotoken.COMMENT
		case '=':
			l.n()
			return off, gotoken.QUO_ASSIGN
		default:
			return off, gotoken.QUO
		}
	case '0':
		n := l.octals(-1)
		switch l.c {
		case '.':
			l.n()
			l.decimals()
			return l.exponent(off)
		case '8', '9':
			l.decimals()
			switch l.c {
			case '.':
				l.n()
				l.decimals()
				return l.exponent(off)
			case 'e', 'E':
				return l.exponent(off)
			case 'i':
				l.n()
				return off, gotoken.IMAG
			default:
				l.err(l.position(l.off-1), "illegal octal number")
			}
		case 'e', 'E':
			return l.exponent(off)
		case 'i':
			l.n()
			return off, gotoken.IMAG
		case 'x', 'X':
			if n != 1 {
				break
			}

			l.n()
			if l.hexadecimals(-1) == 0 {
				l.err(l.position(l.off-1), "illegal hexadecimal number")
			}
		}

		return off, gotoken.INT
	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		l.decimals()
		switch l.c {
		case '.':
			l.n()
			l.decimals()
			return l.exponent(off)
		case 'e', 'E':
			return l.exponent(off)
		case 'i':
			l.n()
			return off, gotoken.IMAG
		}
		return off, gotoken.INT
	case ':':
		if l.n() == '=' {
			l.n()
			return off, gotoken.DEFINE
		}

		return off, gotoken.COLON
	case ';':
		l.n()
		return off, gotoken.SEMICOLON
	case '<':
		switch l.n() {
		case '<':
			if l.n() == '=' {
				l.n()
				return off, gotoken.SHL_ASSIGN
			}

			return off, gotoken.SHL
		case '=':
			l.n()
			return off, gotoken.LEQ
		case '-':
			l.n()
			return off, gotoken.ARROW
		}

		return off, gotoken.LSS
	case '=':
		if l.n() == '=' {
			l.n()
			return off, gotoken.EQL
		}

		return off, gotoken.ASSIGN
	case '>':
		switch l.n() {
		case '>':
			if l.n() == '=' {
				l.n()
				return off, gotoken.SHR_ASSIGN
			}

			return off, gotoken.SHR
		case '=':
			l.n()
			return off, gotoken.GEQ
		}

		return off, gotoken.GTR
	case '[':
		l.n()
		return off, gotoken.LBRACK
	case ']':
		l.n()
		return off, gotoken.RBRACK
	case '^':
		if l.n() == '=' {
			l.n()
			return off, gotoken.XOR_ASSIGN
		}

		return off, gotoken.XOR
	case '`':
	more3:
		switch l.n() {
		case '`':
			l.n()
			return off, gotoken.STRING
		case '\n':
			if l.off != len(l.src) {
				l.file.AddLine(l.off)
			}
		case classEOF:
			l.err(l.position(l.off-1), "raw string literal not terminated")
			return off, gotoken.STRING
		}
		goto more3
	case 'b':
		if l.n() == 'r' && l.n() == 'e' && l.n() == 'a' && l.n() == 'k' && !isIdentNext(l.n()) {
			return off, gotoken.BREAK
		}

		return off, l.ident()
	case 'c':
		switch l.n() {
		case 'a':
			if l.n() == 's' && l.n() == 'e' && !isIdentNext(l.n()) {
				return off, gotoken.CASE
			}
		case 'h':
			if l.n() == 'a' && l.n() == 'n' && !isIdentNext(l.n()) {
				return off, gotoken.CHAN
			}
		case 'o':
			if l.n() == 'n' {
				switch l.n() {
				case 's':
					if l.n() == 't' && !isIdentNext(l.n()) {
						return off, gotoken.CONST
					}
				case 't':
					if l.n() == 'i' && l.n() == 'n' && l.n() == 'u' && l.n() == 'e' && !isIdentNext(l.n()) {
						return off, gotoken.CONTINUE
					}
				}
			}
		}

		return off, l.ident()
	case 'd':
		if l.n() == 'e' && l.n() == 'f' {
			switch l.n() {
			case 'a':
				if l.n() == 'u' && l.n() == 'l' && l.n() == 't' && !isIdentNext(l.n()) {
					return off, gotoken.DEFAULT
				}
			case 'e':
				if l.n() == 'r' && !isIdentNext(l.n()) {
					return off, gotoken.DEFER
				}
			}
		}

		return off, l.ident()
	case 'e':
		if l.n() == 'l' && l.n() == 's' && l.n() == 'e' && !isIdentNext(l.n()) {
			return off, gotoken.ELSE
		}

		return off, l.ident()
	case 'f':
		switch l.n() {
		case 'a':
			if l.n() == 'l' && l.n() == 'l' && l.n() == 't' && l.n() == 'h' && l.n() == 'r' && l.n() == 'o' && l.n() == 'u' && l.n() == 'g' && l.n() == 'h' && !isIdentNext(l.n()) {
				return off, gotoken.FALLTHROUGH
			}
		case 'o':
			if l.n() == 'r' && !isIdentNext(l.n()) {
				return off, gotoken.FOR
			}
		case 'u':
			if l.n() == 'n' && l.n() == 'c' && !isIdentNext(l.n()) {
				return off, gotoken.FUNC
			}
		}

		return off, l.ident()
	case 'g':
		if l.n() == 'o' {
			if !isIdentNext(l.n()) {
				return off, gotoken.GO
			}

			if l.c == 't' && l.n() == 'o' && !isIdentNext(l.n()) {
				return off, gotoken.GOTO
			}
		}

		return off, l.ident()
	case 'i':
		switch l.n() {
		case 'f':
			if !isIdentNext(l.n()) {
				return off, gotoken.IF
			}
		case 'm':
			if l.n() == 'p' && l.n() == 'o' && l.n() == 'r' && l.n() == 't' && !isIdentNext(l.n()) {
				return off, gotoken.IMPORT
			}
		case 'n':
			if l.n() == 't' && l.n() == 'e' && l.n() == 'r' && l.n() == 'f' && l.n() == 'a' && l.n() == 'c' && l.n() == 'e' && !isIdentNext(l.n()) {
				return off, gotoken.INTERFACE
			}
		}

		return off, l.ident()
	case 'm':
		if l.n() == 'a' && l.n() == 'p' && !isIdentNext(l.n()) {
			return off, gotoken.MAP
		}

		return off, l.ident()
	case 'p':
		if l.n() == 'a' && l.n() == 'c' && l.n() == 'k' && l.n() == 'a' && l.n() == 'g' && l.n() == 'e' && !isIdentNext(l.n()) {
			return off, gotoken.PACKAGE
		}

		return off, l.ident()
	case 'r':
		switch l.n() {
		case 'a':
			if l.n() == 'n' && l.n() == 'g' && l.n() == 'e' && !isIdentNext(l.n()) {
				return off, gotoken.RANGE
			}
		case 'e':
			if l.n() == 't' && l.n() == 'u' && l.n() == 'r' && l.n() == 'n' && !isIdentNext(l.n()) {
				return off, gotoken.RETURN
			}
		}

		return off, l.ident()
	case 's':
		switch l.n() {
		case 'e':
			if l.n() == 'l' && l.n() == 'e' && l.n() == 'c' && l.n() == 't' && !isIdentNext(l.n()) {
				return off, gotoken.SELECT
			}
		case 't':
			if l.n() == 'r' && l.n() == 'u' && l.n() == 'c' && l.n() == 't' && !isIdentNext(l.n()) {
				return off, gotoken.STRUCT
			}
		case 'w':
			if l.n() == 'i' && l.n() == 't' && l.n() == 'c' && l.n() == 'h' && !isIdentNext(l.n()) {
				return off, gotoken.SWITCH
			}
		}

		return off, l.ident()
	case 't':
		if l.n() == 'y' && l.n() == 'p' && l.n() == 'e' && !isIdentNext(l.n()) {
			return off, gotoken.TYPE
		}

		return off, l.ident()
	case 'v':
		if l.n() == 'a' && l.n() == 'r' && !isIdentNext(l.n()) {
			return off, gotoken.VAR
		}

		return off, l.ident()
	case '{':
		l.n()
		return off, gotoken.LBRACE
	case '|':
		switch l.n() {
		case '=':
			l.n()
			return off, gotoken.OR_ASSIGN
		case '|':
			l.n()
			return off, gotoken.LOR
		}

		return off, gotoken.OR
	case '}':
		l.n()
		return off, gotoken.RBRACE
	case classEOF:
		return off, gotoken.EOF
	case classLTLT:
		l.skip()
		return off, tokenLTLT
	case classGTGT:
		l.skip()
		return off, tokenGTGT
	case classBOM:
		l.skip()
		return off, tokenBOM
	default:
		if l.c >= 'a' && l.c <= 'z' || l.c >= 'A' && l.c <= 'Z' || l.c == '_' || l.c == classNonASCII {
			l.n()
			for l.c >= 'a' && l.c <= 'z' || l.c >= 'A' && l.c <= 'Z' || l.c == '_' || l.c >= '0' && l.c <= '9' || l.c == classNonASCII {
				l.n()
			}
			return off, gotoken.IDENT
		}

		switch {
		case l.b < ' ':
			l.err(l.position(l.off-1), "invalid character %U", l.skip())
		default:
			l.err(l.position(l.off-1), "invalid character %#U", l.skip())
		}
		return off, gotoken.ILLEGAL
	}
}
