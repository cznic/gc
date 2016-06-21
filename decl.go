// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"bytes"
	"go/token"
	"path/filepath"
	"sort"

	"github.com/cznic/xc"
)

var (
	_ Declaration = (*ConstDeclaration)(nil)
	_ Declaration = (*FuncDeclaration)(nil)
	_ Declaration = (*ImportDeclaration)(nil)
	_ Declaration = (*LabelDeclaration)(nil)
	_ Declaration = (*TypeDeclaration)(nil)
	_ Declaration = (*VarDeclaration)(nil)
)

const (
	gateReady = iota
	gateOpen
	gateClosed
)

// Values of ScopeKind
const (
	UniverseScope ScopeKind = iota
	PackageScope
	FileScope
	BlockScope
)

// Declaration is a named entity, eg. a type, variable, function, etc.
type Declaration interface {
	Node
	Name() int // Name ID.
	ScopeStart() token.Pos
	check(ctx *Context, stack []Declaration, node Node) (stop bool)
	//TODO Exported() bool
}

type declarations []Declaration

func (d declarations) Len() int      { return len(d) }
func (d declarations) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

func (d declarations) Less(i, j int) bool {
	if d[i].Pos() < d[j].Pos() {
		return true
	}

	if d[i].Pos() > d[j].Pos() {
		return false
	}

	return bytes.Compare(dict.S(d[i].Name()), dict.S(d[j].Name())) < 0
}

type fields []*FieldDeclaration

func (f fields) Len() int           { return len(f) }
func (f fields) Less(i, j int) bool { return f[i].Pos() < f[j].Pos() }
func (f fields) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

// Bindings map name IDs to declarations.
type Bindings map[int]Declaration

func (b *Bindings) declare(lx *lexer, d Declaration) {
	if *b == nil {
		*b = Bindings{}
	}
	m := *b
	nm := d.Name()
	if nm == 0 {
		panic("internal error")
	}

	ex := m[nm]
	if ex == nil {
		m[nm] = d
		return
	}

	switch x := d.(type) {
	case *FieldDeclaration:
		lx.err(d, "duplicate field %s, previous declaration at %s", dict.S(d.Name()), position(ex.Pos()))
	case *FuncDeclaration:
		if x.ifaceMethod || x.rx != nil {
			lx.err(d, "duplicate method %s, previous declaration at %s", dict.S(d.Name()), position(ex.Pos()))
			break
		}

		lx.err(d, "%s redeclared, previous declaration at %s", dict.S(d.Name()), position(ex.Pos()))
	case *LabelDeclaration:
		lx.err(d, "label %s already defined, previous declaration at %s", dict.S(d.Name()), position(ex.Pos()))
	default:
		lx.err(d, "%s redeclared, previous declaration at %s", dict.S(d.Name()), position(ex.Pos()))
	}
}

// Scope tracks declarations.
type Scope struct {
	Bindings     Bindings
	Kind         ScopeKind
	Labels       Bindings
	Parent       *Scope
	Unbound      []Declaration // Declarations named _.
	isFnScope    bool
	isMergeScope bool
}

func newScope(kind ScopeKind, parent *Scope) *Scope {
	return &Scope{
		Kind:   kind,
		Parent: parent,
	}
}

func (s *Scope) declare(lx *lexer, d Declaration) {
	nm := d.Name()
	var p *Package
	if lx != nil {
		p = lx.pkg
	}
	if nm == idUnderscore {
		s.Unbound = append(s.Unbound, d)
		return
	}

	switch d.(type) {
	case *ImportDeclaration:
		if s.Kind != FileScope {
			panic("internal error")
		}

		if ex := s.Parent.Bindings[nm]; ex != nil {
			lx.err(d, "%s redeclared as import name, previous declaration at %s", position(ex.Pos()))
			return
		}

		if _, ok := p.avoid[nm]; !ok {
			p.avoid[nm] = d.Pos()
		}
		s.Bindings.declare(lx, d)
	case *LabelDeclaration:
		for s != nil && !s.isFnScope {
			s = s.Parent
		}
		if s != nil {
			s.Labels.declare(lx, d)
			return
		}

		lx.err(d, "label declaration outside of a function")
	default:
		k := s.Kind
		if k == PackageScope { // TLD.
			if ex := p.avoid[nm]; ex != 0 {
				switch {
				case position(d.Pos()).Filename == position(ex).Filename:
					lx.err(d, "%s redeclared in this block\n\tprevious declaration at %s", dict.S(nm), position(ex))
				default:
					lx.errPos(ex, "%s redeclared in this block\n\tprevious declaration at %s", dict.S(nm), position(d.Pos()))
				}
				return
			}
		}

		if s.Kind == PackageScope && nm == idInit {
			if x, ok := d.(*FuncDeclaration); ok {
				p.Inits = append(p.Inits, x)
				return
			}

			lx.err(d, "cannot declare init - must be function")
			return
		}

		s.Bindings.declare(lx, d)
	}
}

func (s *Scope) lookup(t xc.Token, fileScope *Scope) (d Declaration) {
	d, _ = s.lookup2(t, fileScope)
	return d
}

func (s *Scope) lookup2(t xc.Token, fileScope *Scope) (d Declaration, _ *Scope) {
	s0 := s
	for s != nil {
		switch d = s.Bindings[t.Val]; {
		case d == nil:
			if s.Kind == PackageScope {
				if d = fileScope.Bindings[t.Val]; d != nil {
					return d, s
				}
			}
		default:
			if s.Kind != BlockScope || s0 != s || d.ScopeStart() < t.Pos() {
				return d, s
			}
		}
		s = s.Parent
	}
	return nil, nil
}

func (s *Scope) mustLookup(ctx *Context, t xc.Token, fileScope *Scope) (d Declaration) {
	if d = s.lookup(t, fileScope); d == nil {
		ctx.err(t, "undefined: %s", t.S())
	}
	return d
}

func (s *Scope) mustLookupLocalTLDType(ctx *Context, t xc.Token) *TypeDeclaration {
	d, s := s.lookup2(t, nil)
	if s.Kind != PackageScope {
		return nil
	}

	td, ok := d.(*TypeDeclaration)
	if d != nil && !ok {
		ctx.err(t, "%s is not a type", t.S())
	}
	return td
}

func (s *Scope) lookupQI(qi *QualifiedIdent, fileScope *Scope) (d Declaration) {
	switch q, i := qi.q(), qi.i(); {
	case q.IsValid():
		if !isExported(i.Val) {
			return nil
		}

		if p, ok := fileScope.Bindings[q.Val].(*ImportDeclaration); ok {
			return p.Package.Scope.Bindings[i.Val]
		}

		return nil
	default:
		return s.lookup(qi.i(), fileScope)
	}
}

func (s *Scope) mustLookupQI(ctx *Context, qi *QualifiedIdent, fileScope *Scope) (d Declaration) {
	if d = s.lookupQI(qi, fileScope); d == nil {
		ctx.err(qi, "undefined %s", qi.str())
	}
	return d
}

func (s *Scope) mustLookupType(ctx *Context, qi *QualifiedIdent, fileScope *Scope) *TypeDeclaration {
	d := s.lookupQI(qi, fileScope)
	t, ok := d.(*TypeDeclaration)
	if d != nil && !ok {
		ctx.err(qi, "%s is not a type", qi.str())
	}
	return t
}

func (s *Scope) check(ctx *Context) (stop bool) {
	if s.Kind == PackageScope {
		a := make(declarations, 0, len(s.Bindings)+len(s.Unbound))
		for _, d := range s.Bindings {
			a = append(a, d)
		}
		a = append(a, s.Unbound...)
		sort.Sort(a)
		for _, d := range a {
			if d.check(ctx, nil, nil) {
				return true
			}
		}
		for _, d := range a {
			x, ok := d.(*FuncDeclaration)
			if !ok {
				continue
			}

			_ = x //TODO
		}
	}
	return false
}

// ScopeKind is the specific kind of a Scope.
type ScopeKind int

// ConstDeclaration represents a constant declaration.
type ConstDeclaration struct {
	Type       Type
	Value      Value
	expr       *Expression
	guard      gate
	isExported bool
	name       int
	pos        token.Pos
	scopeStart token.Pos
	typ0       *Typ
}

func newConstDeclaration(nm xc.Token, typ0 *Typ, expr *Expression, scopeStart token.Pos) *ConstDeclaration {
	return &ConstDeclaration{
		expr:       expr,
		isExported: isExported(nm.Val),
		name:       nm.Val,
		pos:        nm.Pos(),
		scopeStart: scopeStart,
		typ0:       typ0,
	}
}

func (n *ConstDeclaration) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	stack = append(stack, n)
	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	if n.expr == nil {
		return false
	}

	if n.expr.check(ctx, stack, nil) {
		return true
	}

	if typ0 := n.typ0; typ0 != nil {
		if typ0.check(ctx, stack, nil) {
			return true
		}

		t := typ0.Type
		if t == nil {
			return false
		}

		switch k := t.Kind(); {
		case t.IntegerType(), t.FloatingPointType(), t.ComplexType(), k == Bool, k == String:
			// nop
		default:
			return ctx.err(n.typ0, "invalid constant type %s", t)
		}

		n.Type = typ0.Type
		v := n.expr.Value
		if v == nil {
			return false
		}

		switch v.Kind() {
		case ConstValue:
			n.Value = v.Convert(n.Type)
			if n.Value == nil {
				todo(n) //TODO
			}
		case NilValue:
			todo(n) //TODO
		default:
			todo(n) //TODO
		}
		return
	}

	if v := n.expr.Value; v != nil {
		n.Type = v.Type()
	}
	return false
}

// Pos implements Declaration.
func (n *ConstDeclaration) Pos() token.Pos { return n.pos }

// Name implements Declaration.
func (n *ConstDeclaration) Name() int { return n.name }

// ScopeStart implements Declaration.
func (n *ConstDeclaration) ScopeStart() token.Pos { return n.scopeStart }

// FieldDeclaration represents a struct field
type FieldDeclaration struct {
	guard          gate
	isAnonymous    bool
	isAnonymousPtr bool
	isExported     bool
	name           int
	pos            token.Pos
	qi             *QualifiedIdent
	tag            stringValue
	typ            Type
	typ0           *Typ
}

func newFieldDeclaration(nm xc.Token, typ0 *Typ, isAnonymousPtr bool, qi *QualifiedIdent, tag stringValue) *FieldDeclaration {
	return &FieldDeclaration{
		isAnonymous:    qi != nil,
		isAnonymousPtr: isAnonymousPtr,
		isExported:     isExported(nm.Val),
		name:           nm.Val,
		pos:            nm.Pos(),
		qi:             qi,
		tag:            tag,
		typ0:           typ0,
	}
}

func (n *FieldDeclaration) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	todo(n) //TODO
	return false
}

// Pos implements Declaration.
func (n *FieldDeclaration) Pos() token.Pos { return n.pos }

// Name implements Declaration.
func (n *FieldDeclaration) Name() int { return n.name }

// ScopeStart implements Declaration.
func (n *FieldDeclaration) ScopeStart() token.Pos { return 0 }

// FuncDeclaration represents a function declaration.
type FuncDeclaration struct {
	guard       gate
	ifaceMethod bool
	isExported  bool
	name        int
	pos         token.Pos
	rx          *ReceiverOpt
	sig         *Signature
	typ         Type
}

func newFuncDeclaration(nm xc.Token, rx *ReceiverOpt, sig *Signature, ifaceMethod bool) *FuncDeclaration {
	return &FuncDeclaration{
		ifaceMethod: ifaceMethod,
		isExported:  isExported(nm.Val),
		name:        nm.Val,
		pos:         nm.Pos(),
		rx:          rx,
		sig:         sig,
	}
}

func (n *FuncDeclaration) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	stack = nil
	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	if n.rx.check(ctx, stack, node) || n.sig.check(ctx, stack, node) {
		return true
	}

	var in, out []Type
	if n.rx != nil {
		in = append(in, n.rx.Type)
	}
	isVariadic := false
	for l := n.sig.Parameters.ParameterDeclList; l != nil; l = l.ParameterDeclList {
		switch i := l.ParameterDecl; i.Case {
		case 0: // "..." Typ
			in = append(in, newSliceType(ctx, i.Typ.Type))
			isVariadic = true
		case 1: // IDENTIFIER "..." Typ
			in = append(in, newSliceType(ctx, i.Typ.Type))
			isVariadic = true
		case 2: // IDENTIFIER Typ
			in = append(in, i.Typ.Type)
		case 3: // Typ
			switch {
			case i.isParamName:
				in = append(in, i.typ.Type)
			default:
				in = append(in, i.Typ.Type)
			}
		default:
			panic("internal error")
		}
	}
	switch t := n.sig.Type; {
	case t != nil && t.Kind() == Tuple:
		out = t.Elements()
	default:
		out = []Type{t}
	}
	n.typ = newFuncType(ctx, n.Name(), in, out, n.isExported, isVariadic)
	return false
}

// Pos implements Declaration.
func (n *FuncDeclaration) Pos() token.Pos { return n.pos }

// Name implements Declaration.
func (n *FuncDeclaration) Name() int { return n.name }

// ScopeStart implements Declaration.
func (n *FuncDeclaration) ScopeStart() token.Pos { return n.Pos() }

// ImportDeclaration represents a named import declaration of a single package.
// The name comes from an explicit identifier (import foo "bar") when present.
// Otherwise the name in the package clause (package foo) is used.
type ImportDeclaration struct {
	Package *Package
	name    int
	once    *xc.Once
	pos     token.Pos
}

func newImportDeclaration(p *Package, nm int, pos token.Pos, once *xc.Once) *ImportDeclaration {
	return &ImportDeclaration{
		Package: p,
		name:    nm,
		once:    once,
		pos:     pos,
	}
}

func (n *ImportDeclaration) check(*Context, []Declaration, Node) (stop bool) { return false }

// Pos implements Declaration.
func (n *ImportDeclaration) Pos() token.Pos { return n.pos }

// Name implements Declaration.
func (n *ImportDeclaration) Name() int { return n.name }

// ScopeStart implements Declaration.
func (n *ImportDeclaration) ScopeStart() token.Pos { panic("ScopeStart of ImportDeclaration") }

// LabelDeclaration represents a label declaration.
type LabelDeclaration struct {
	tok xc.Token // AST link.
}

func newLabelDeclaration(tok xc.Token) *LabelDeclaration {
	return &LabelDeclaration{
		tok: tok,
	}
}

func (n *LabelDeclaration) check(*Context, []Declaration, Node) bool { panic("internal error") }

// Pos implements Declaration.
func (n *LabelDeclaration) Pos() token.Pos { return n.tok.Pos() }

// Name implements Declaration.
func (n *LabelDeclaration) Name() int { return n.tok.Val }

// ScopeStart implements Declaration.
func (n *LabelDeclaration) ScopeStart() token.Pos { panic("ScopeStart of LabelDeclaration") }

// ParameterDeclaration represents a function/method parameter declaration.
type ParameterDeclaration struct {
	guard      gate
	isVariadic bool
	name       int
	pos        token.Pos
	scopeStart token.Pos
	typ        Type
	typ0       *Typ
}

func newParamaterDeclaration(nm xc.Token, typ0 *Typ, isVariadic bool, scopeStart token.Pos) *ParameterDeclaration {
	return &ParameterDeclaration{
		isVariadic: isVariadic,
		name:       nm.Val,
		pos:        nm.Pos(),
		scopeStart: scopeStart,
		typ0:       typ0,
	}
}

func (n *ParameterDeclaration) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	todo(n) //TODO
	return false
}

// Pos implements Declaration.
func (n *ParameterDeclaration) Pos() token.Pos { return n.pos }

// Name implements Declaration.
func (n *ParameterDeclaration) Name() int { return n.name }

// ScopeStart implements Declaration.
func (n *ParameterDeclaration) ScopeStart() token.Pos { return n.scopeStart }

// TypeDeclaration represents a type declaration.
type TypeDeclaration struct {
	guard      gate
	isExported bool
	methods    *Scope // Type methods, if any, nil otherwise.
	name       int
	pkg        *Package
	pos        token.Pos
	qualifier  int  // String(): "qualifier.name".
	typ        Type // bar in type foo bar.
	typ0       *Typ // bar in type foo bar.
	typeBase
}

func newTypeDeclaration(lx *lexer, nm xc.Token, typ0 *Typ) *TypeDeclaration {
	var pkgPath, qualifier int
	var ctx *Context
	var pkg *Package
	if lx != nil {
		pkgPath = lx.pkg.importPath
		qualifier = lx.pkg.name
		ctx = lx.Context
		pkg = lx.pkg
	}
	t := &TypeDeclaration{
		isExported: isExported(nm.Val),
		name:       nm.Val,
		pkg:        pkg,
		pos:        nm.Pos(),
		qualifier:  qualifier,
		typ0:       typ0,
		typeBase:   typeBase{ctx: ctx, pkgPath: pkgPath},
	}
	t.typeBase.typ = t
	return t
}

func (n *TypeDeclaration) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	stack = append(stack, n)
	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	if t0 := n.typ0; t0 != nil {
		if t0.check(ctx, stack, node) {
			return true
		}

		if t := t0.Type; t != nil {
			n.typ = t
			n.kind = t.Kind()
			if n.pkgPath == idUnsafe && n.Name() == idPointer {
				if _, err := filepath.Rel(ctx.searchPaths[0], n.pkg.Directory); err == nil {
					n.kind = UnsafePointer
				}
			}
		}
	}
	return false
}

// Pos implements Declaration.
func (n *TypeDeclaration) Pos() token.Pos { return n.pos }

// Name implements Declaration.
func (n *TypeDeclaration) Name() int { return n.name }

// ScopeStart implements Declaration.
func (n *TypeDeclaration) ScopeStart() token.Pos { return n.pos }

func (n *TypeDeclaration) declare(lx *lexer, d Declaration) {
	if n.methods == nil {
		n.methods = newScope(BlockScope, nil)
	}
	nm := d.Name()
	ex := n.methods.Bindings[nm]
	if ex == nil {
		n.methods.declare(lx, d)
		return
	}

	rx := string(dict.S(n.name))
	if d := d.(*FuncDeclaration); d.rx != nil && d.rx.isPtr {
		rx = "(*" + rx + ")"
	}
	lx.err(d, "%s.%s redeclared in this block\n\tprevious declaration at %s", rx, dict.S(nm), position(ex.Pos()))
}

// Identical reports whether this type is identical to u.
func (n *TypeDeclaration) Identical(u Type) bool { return n == u }

func (n *TypeDeclaration) str(w *bytes.Buffer) {
	if b := n.qualifier; b != 0 {
		w.Write(dict.S(b))
		w.WriteByte('.')
	}
	w.Write(dict.S(n.Name()))
}

func (n *TypeDeclaration) String() string {
	var buf bytes.Buffer
	n.str(&buf)
	return buf.String()
}

// VarDeclaration represents a variable declaration.
type VarDeclaration struct {
	Type       Type
	expr       *Expression
	guard      gate
	isExported bool
	name       int
	pos        token.Pos
	scopeStart token.Pos
	tupleIndex int
	typ0       *Typ
}

func newVarDeclaration(tupleIndex int, nm xc.Token, typ0 *Typ, expr *Expression, scopeStart token.Pos) *VarDeclaration {
	return &VarDeclaration{
		expr:       expr,
		isExported: isExported(nm.Val),
		name:       nm.Val,
		pos:        nm.Pos(),
		scopeStart: scopeStart,
		tupleIndex: tupleIndex,
		typ0:       typ0,
	}
}

func (n *VarDeclaration) check(ctx *Context, stack []Declaration, node Node) (stop bool) {
	stack = append(stack, n)
	done, stop := n.guard.check(ctx, stack, node)
	if done || stop {
		return stop
	}

	defer n.guard.done()

	if n.expr.check(ctx, stack, node) || n.typ0.check(ctx, stack, node) {
		return true
	}

	switch {
	case n.typ0 != nil:
		todo(n) //TODO
	default:
		if v := n.expr.Value; v != nil {
			if v.Kind() != ConstValue {
				return ctx.err(n, "not a constant expression")
			}

			c := v.Const()
			if !c.Untyped() {
				n.Type = c.Type()
				break
			}

			todo(n) //TODO
		}
	}
	return false
}

// Pos implements Declaration.
func (n *VarDeclaration) Pos() token.Pos { return n.pos }

// Name implements Declaration.
func (n *VarDeclaration) Name() int { return n.name }

// ScopeStart implements Declaration.
func (n *VarDeclaration) ScopeStart() token.Pos { return n.scopeStart }

func varDecl(lx *lexer, lhs, rhs Node, typ0 *Typ, op string, maxLHS, maxRHS int) {
	var ln []Node
	var names []xc.Token
	switch x := lhs.(type) {
	case *ArgumentList:
		for l := x; l != nil; l = l.ArgumentList {
			n := l.Argument
			ln = append(ln, n)
			names = append(names, n.ident())
		}
	case *ExpressionList:
		for l := x; l != nil; l = l.ExpressionList {
			n := l.Expression
			ln = append(ln, n)
			names = append(names, n.ident())
		}
	case *IdentifierList:
		for l := x; l != nil; l = l.IdentifierList {
			n := l.ident()
			ln = append(ln, n)
			names = append(names, n)
		}
	default:
		panic("internal error")
	}

	m := map[int]struct{}{}
	lhsOk := true
	for i, v := range names {
		if !v.IsValid() {
			lx.err(ln[i], "non-name on left side of %s", op)
			lhsOk = false
			continue
		}

		if val := v.Val; val != idUnderscore {
			if _, ok := m[val]; ok {
				lx.err(v, "%s repeated on left side of %s", v.S(), op)
				lhsOk = false
				names[i] = xc.Token{}
			}
			m[val] = struct{}{}
		}

	}

	var rn []Node
	var exprs []*Expression
	switch x := rhs.(type) {
	case nil:
		// nop
	case *Expression:
		rn = []Node{x}
		exprs = []*Expression{x}
	case *ExpressionList:
		for l := x; l != nil; l = l.ExpressionList {
			n := l.Expression
			rn = append(rn, n)
			exprs = append(exprs, n)
		}
	default:
		panic("internal error")
	}

	if maxLHS > 0 && len(names) > maxLHS {
		lx.err(ln[maxLHS], "too many items on left side of %s", op)
		names = names[:maxLHS]
	}

	if maxRHS >= 0 && len(exprs) > maxRHS {
		lx.err(rn[maxRHS], "too many expressions on right side of %s", op)
		exprs = exprs[:maxRHS]
	}

	scopeStart := lx.lookahead.Pos()
	hasNew := false
	switch len(exprs) {
	case 0:
		// No initializer.
		if op != "" {
			panic("internal error")
		}

		for _, v := range names {
			lx.scope.declare(lx, newVarDeclaration(-1, v, typ0, nil, scopeStart))
		}
	case 1:
		// One initializer.
		//
		// The number of identifiers must match the length of the tuple
		// produced by the initializer, but that is not yet known here.
		for i, v := range names {
			if !v.IsValid() {
				continue
			}

			switch {
			case lx.scope.Bindings[v.Val] != nil:
				if op == ":=" {
					continue
				}
			default:
				if v.Val != idUnderscore {
					hasNew = true
				}
			}
			lx.scope.declare(lx, newVarDeclaration(i, v, typ0, exprs[0], scopeStart))
		}
	default:
		// Initializer list.
		//
		// The number of identifiers and initializers must match
		// exactly, every initializer must be single valued expression,
		// but that is not yet known here.
		for i, v := range names {
			if !v.IsValid() {
				continue
			}

			var e *Expression
			switch {
			case i >= len(exprs):
				lx.err(lx.lookahead, "missing initializer on right side of %s", op)
				return
			default:
				e = exprs[i]
			}

			switch {
			case lx.scope.Bindings[v.Val] != nil:
				if op == ":=" {
					continue
				}
			default:
				if v.Val != idUnderscore {
					hasNew = true
				}
			}
			lx.scope.declare(lx, newVarDeclaration(-1, v, typ0, e, scopeStart))
		}
		if len(exprs) > len(names) {
			lx.err(exprs[len(names)], "extra initializer(s) on right side of %s", op)
		}

	}
	if lhsOk && op == ":=" && !hasNew {
		lx.err(ln[0], "no new variables on left side of %s", op)
	}
}

type gate int

func (g *gate) check(ctx *Context, stack []Declaration, node Node) (done, stop bool) {
	switch *g {
	case gateReady:
		*g = gateOpen
		return false, false
	case gateOpen:
		if len(stack) == 0 {
			return true, false
		}

		d := stack[len(stack)-1]
		var a []Declaration
		for _, v := range stack[:len(stack)-1] {
			if v == d {
				a = append(a, v)
			}
		}
		if len(a) == 0 {
			return true, false
		}

		if node == nil {
			node = d
		}
		var kind string
		switch node.(type) {
		case *ConstDeclaration:
			kind = "constant"
		case
			*FuncDeclaration,
			*ImportDeclaration,
			*LabelDeclaration:
			panic("internal error")
		case *TypeDeclaration:
			kind = "type"
		case *VarDeclaration:
			kind = "variable"
		default:
			panic("internal error")
		}
		if len(a) > 1 {
			todo(node) //TODO
		}
		return true, ctx.err(node, "invalid recursive %s %s", kind, dict.S(d.Name()))
	case gateClosed:
		return true, false
	default:
		panic("internal error")
	}
}

func (g *gate) done() { *g = gateClosed }
