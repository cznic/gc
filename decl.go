// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"bytes"
	"fmt"
	"go/token"
	"sort"
	"strings"

	"github.com/cznic/xc"
)

var (
	_ Declaration = (*ConstDeclaration)(nil)
	_ Declaration = (*FieldDeclaration)(nil)
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
	check(ctx *context) (stop bool)
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
		if x, ok := d.(*FuncDeclaration); ok {
			lx.lastFuncDeclaration = x
		}
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
	fnType       *Type
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

func (s *Scope) mustLookup(ctx *context, t xc.Token, fileScope *Scope) (d Declaration) {
	if d = s.lookup(t, fileScope); d == nil {
		ctx.err(t, "undefined: %s", t.S())
	}
	return d
}

func (s *Scope) mustLookupLocalTLDType(ctx *context, t xc.Token) *TypeDeclaration {
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

func (s *Scope) mustLookupQI(ctx *context, qi *QualifiedIdent, fileScope *Scope) (d Declaration) {
	if d = s.lookupQI(qi, fileScope); d == nil {
		ctx.err(qi, "undefined: %s", qi.str())
	}
	return d
}

func (s *Scope) mustLookupType(ctx *context, qi *QualifiedIdent, fileScope *Scope) *TypeDeclaration {
	d := s.mustLookupQI(ctx, qi, fileScope)
	t, ok := d.(*TypeDeclaration)
	if d != nil && !ok {
		ctx.err(qi, "%s is not a type", qi.str())
	}
	return t
}

func (s *Scope) check(ctx *context) (stop bool) {
	a := make(declarations, 0, len(s.Bindings)+len(s.Unbound))
	for _, d := range s.Bindings {
		a = append(a, d)
	}
	a = append(a, s.Unbound...)
	sort.Sort(a)
	for _, d := range a {
		if d.check(ctx) {
			return true
		}
	}
	if s.Kind == PackageScope {
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
	Value      Value
	expr       *Expression
	guard      gate
	iota       int64
	isExported bool
	name       int
	pos        token.Pos
	scopeStart token.Pos
	typ0       *Typ
}

func newConstDeclaration(nm xc.Token, typ0 *Typ, expr *Expression, iota int64, scopeStart token.Pos) *ConstDeclaration {
	return &ConstDeclaration{
		expr:       expr,
		iota:       iota,
		isExported: isExported(nm.Val),
		name:       nm.Val,
		pos:        nm.Pos(),
		scopeStart: scopeStart,
		typ0:       typ0,
	}
}

func (n *ConstDeclaration) check(ctx *context) (stop bool) {
	done, stop := n.guard.check(ctx, n)
	if done || stop {
		return stop
	}

	defer n.guard.done(ctx)

	if n.expr == nil {
		return false
	}

	iota := ctx.iota
	ctx.iota = newConstValue(newIntConst(n.iota, nil, ctx.intType, true))
	if n.expr.check(ctx) {
		ctx.iota = iota
		return true
	}

	ctx.iota = iota
	if typ0 := n.typ0; typ0 != nil {
		if typ0.check(ctx) {
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

		v := n.expr.Value
		if v == nil {
			return ctx.err(n.expr, "const initializer is not a constant")
		}

		switch v.Kind() {
		case ConstValue:
			if c := ctx.mustConvertConst(n.expr, t, v.Const()); c != nil {
				n.Value = newConstValue(c)
			}
		case NilValue:
			todo(n, true)
		default:
			todo(n, true)
		}
		return
	}

	if v := n.expr.Value; v != nil {
		switch v.Kind() {
		case ConstValue:
			n.Value = v
		case NilValue:
			ctx.err(n.expr, "const initializer cannot be nil")
		default:
			if t := v.Type(); t != nil && t.Kind() == Ptr {
				ctx.err(n.expr, "invalid constant type")
				break
			}

			ctx.err(n.expr, "const initializer is not a constant")
		}
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
	Type            Type
	fileScope       *Scope
	guard           gate
	isAnonymous     bool
	isAnonymousPtr  bool
	isExported      bool
	name            int
	parent          *StructType
	pos             token.Pos
	qi              *QualifiedIdent
	resolutionScope *Scope // QualifiedIdent
	tag             stringValue
	typ0            *Typ
}

func newFieldDeclaration(nm xc.Token, typ0 *Typ, isAnonymousPtr bool, qi *QualifiedIdent, tag stringValue, fileScope, resolutionScope *Scope) *FieldDeclaration {
	return &FieldDeclaration{
		fileScope:       fileScope,
		isAnonymous:     qi != nil,
		isAnonymousPtr:  isAnonymousPtr,
		isExported:      isExported(nm.Val),
		name:            nm.Val,
		pos:             nm.Pos(),
		qi:              qi,
		resolutionScope: resolutionScope,
		tag:             tag,
		typ0:            typ0,
	}
}

func (n *FieldDeclaration) check(ctx *context) (stop bool) {
	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	switch {
	case n.isAnonymous:
		n.Type, stop = n.qi.checkTypeName(ctx)
		return stop
	default:
		if n.typ0.check(ctx) {
			return true
		}

		n.Type = n.typ0.Type
	}
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
	Type        Type
	guard       gate
	ifaceMethod bool
	isExported  bool
	name        int
	pos         token.Pos
	rx          *ReceiverOpt
	sig         *Signature
	unsafe      bool
}

func newFuncDeclaration(nm xc.Token, rx *ReceiverOpt, sig *Signature, ifaceMethod bool, unsafe bool) *FuncDeclaration {
	return &FuncDeclaration{
		ifaceMethod: ifaceMethod,
		isExported:  isExported(nm.Val),
		name:        nm.Val,
		pos:         nm.Pos(),
		rx:          rx,
		sig:         sig,
		unsafe:      unsafe,
	}
}

func (n *FuncDeclaration) check(ctx *context) (stop bool) {
	stack := ctx.stack
	ctx.stack = nil

	defer func() { ctx.stack = stack }()

	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	if n.rx.check(ctx) || n.sig.check(ctx) {
		return true
	}

	var in []Type
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
	n.Type = newFuncType(ctx, n.Name(), in, n.sig.Type, n.isExported, isVariadic)
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
	bl      *BasicLiteral
	ip      int
	name    int
	once    *xc.Once
	pos     token.Pos
}

func newImportDeclaration(nm int, pos token.Pos, bl *BasicLiteral, ip int) *ImportDeclaration {
	return &ImportDeclaration{
		bl:   bl,
		ip:   ip,
		name: nm,
		pos:  pos,
	}
}

func (n *ImportDeclaration) check(*context) (stop bool) {
	return false
}

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

func (n *LabelDeclaration) check(*context) bool {
	panic("internal error")
}

// Pos implements Declaration.
func (n *LabelDeclaration) Pos() token.Pos { return n.tok.Pos() }

// Name implements Declaration.
func (n *LabelDeclaration) Name() int { return n.tok.Val }

// ScopeStart implements Declaration.
func (n *LabelDeclaration) ScopeStart() token.Pos { panic("ScopeStart of LabelDeclaration") }

// ParameterDeclaration represents a function/method parameter declaration.
type ParameterDeclaration struct {
	Type       Type
	guard      gate
	isVariadic bool
	name       int
	pos        token.Pos
	scopeStart token.Pos
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

func (n *ParameterDeclaration) check(ctx *context) (stop bool) {
	done, stop := n.guard.check(ctx, nil)
	if done || stop {
		return stop
	}

	defer n.guard.done(nil)

	stop = n.typ0.check(ctx)
	n.Type = n.typ0.Type
	return stop
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
		if pkgPath != 0 {
			qualifier = lx.pkg.name
		}
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

func (n *TypeDeclaration) check(ctx *context) (stop bool) {
	done, stop := n.guard.check(ctx, n)
	if done || stop {
		return stop
	}

	defer n.guard.done(ctx)

	if t0 := n.typ0; t0 != nil && t0.Case == 9 { // StructType
		n0 := ctx.node
		ctx.node = t0.StructType.Token2
		defer func() { ctx.node = n0 }()
	}

	t0 := n.typ0
	if t0 == nil {
		return false
	}

	if t0.check(ctx) {
		return true
	}

	t := t0.Type
	if t == nil {
		return false
	}

	k := t.Kind()
	if k == Invalid {
		return false
	}

	n.typ = t
	n.kind = k
	n.align = t.Align()
	n.fieldAlign = t.FieldAlign()
	n.size = t.Size()
	t = n
	if n.methods != nil {
		s := n.methods
		a := make(declarations, 0, len(s.Bindings)+len(s.Unbound))
		for _, d := range s.Bindings {
			a = append(a, d)
		}
		for _, d := range s.Unbound {
			a = append(a, d)
		}
		sort.Sort(a)
		switch {
		case t.Kind() == Ptr:
			for _, m := range a {
				if ctx.err(m, "invalid receiver type %s (%s is a pointer type)", t, t) {
					return true
				}
			}
		default:
			var mta []Method
			var index int
			for _, m := range a {
				if m.check(ctx) {
					return true
				}

				if m.Name() != idUnderscore {
					var pth int
					fd := m.(*FuncDeclaration)
					if !fd.isExported {
						pth = n.pkgPath
					}
					mt := Method{m.Name(), pth, fd.Type, index, false}
					index++
					mta = append(mta, mt)
				}
			}
			n.typeBase.methods = mta
		}
	}
	if n.pkg.unsafe && n.Name() == idPointer {
		n.kind = UnsafePointer
	}

	u := t.UnderlyingType()
	if u.NumMethod() == 0 || t.Kind() == Ptr {
		return false
	}

	m := map[int]struct{}{}
	for i := 0; i < n.NumMethod(); i++ {
		m[n.Method(i).Name] = struct{}{}
	}
	for i := 0; i < u.NumMethod(); i++ {
		mt := u.Method(i)
		if _, ok := m[mt.Name]; ok {
			continue
		}

		m[mt.Name] = struct{}{}
		var mt2 Method
		switch {
		case u.Kind() == Interface:
			if t.Kind() != Interface {
				todo(n) // Merge method of the underlying type, must synthesize new method w/ receiver.
				break
			}

			fallthrough
		default:
			mt2 = *mt
			mt2.merged = true
		}
		mt2.Index = len(n.typeBase.methods)
		n.typeBase.methods = append(n.typeBase.methods, mt2)
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

// ChanDir implements Type.
func (n *TypeDeclaration) ChanDir() ChanDir {
	if n.typ == nil {
		panic("ChanDir of an inappropriate type")
	}

	return n.typ.ChanDir()
}

// Elem implements Type.
func (n *TypeDeclaration) Elem() Type {
	if n.typ == nil {
		panic("Elem of an inappropriate type")
	}

	return n.typ.Elem()
}

// Elements implement Type.
func (n *TypeDeclaration) Elements() []Type {
	if n.typ == nil {
		panic("Elements of an inappropriate type")
	}

	return n.typ.Elements()
}

// Field implements Type.
func (n *TypeDeclaration) Field(i int) *StructField {
	if n.typ == nil {
		panic("Field of an inappropriate type")
	}

	return n.typ.Field(i)
}

// FieldByIndex implements Type.
func (n *TypeDeclaration) FieldByIndex(index []int) StructField {
	if n.typ == nil {
		panic("FieldByIndex of an inappropriate type")
	}

	return n.typ.FieldByIndex(index)
}

// FieldByName implements Type.
// field was not found. The result pointee is read only.
func (n *TypeDeclaration) FieldByName(name int) *StructField {
	if n.typ == nil {
		panic("FieldByName of an inappropriate type")
	}

	return n.typ.FieldByName(name)
}

// FieldByNameFunc implements Type.
func (n *TypeDeclaration) FieldByNameFunc(match func(int) bool) *StructField {
	if n.typ == nil {
		panic("FieldByNameFunc of an inappropriate type")
	}

	return n.typ.FieldByNameFunc(match)
}

// Identical reports whether this type is identical to u.
func (n *TypeDeclaration) Identical(u Type) bool { return n == u }

// In implements Type.
func (n *TypeDeclaration) In(i int) Type {
	if n.typ == nil {
		panic("In of an inappropriate type")
	}

	return n.typ.In(i)
}

// IsVariadic implements Type.
func (n *TypeDeclaration) IsVariadic() bool {
	if n.typ == nil {
		panic("IsVariadic of an inappropriate type")
	}

	return n.typ.IsVariadic()
}

// Key implements Type.
func (n *TypeDeclaration) Key() Type {
	if n.typ == nil {
		panic("Key of an inappropriate type")
	}

	return n.typ.Key()
}

// Len implements Type.
func (n *TypeDeclaration) Len() int64 {
	if n.typ == nil {
		panic("Len of an inappropriate type")
	}

	return n.typ.Len()
}

// NumField implements Type.
func (n *TypeDeclaration) NumField() int {
	if n.typ == nil {
		panic("NumField of an inappropriate type")
	}

	return n.typ.NumField()
}

// NumIn implements Type.
func (n *TypeDeclaration) NumIn() int {
	if n.typ == nil {
		panic("NumIn of an inappropriate type")
	}

	return n.typ.NumIn()
}

// Numeric implements Type.
func (n *TypeDeclaration) Numeric() bool {
	if n.typ == nil {
		switch n.Kind() {
		case Int8, Int16, Int32, Int64, Int, Uint8, Uint16, Uint32, Uint64, Uint, Float32, Float64, Complex64, Complex128, Uintptr:
			return true
		}

		return false
	}

	return n.typ.Numeric()
}

// NumOut implements Type.
func (n *TypeDeclaration) NumOut() int {
	if n.typ == nil {
		panic("NumOut of an inappropriate type")
	}

	return n.typ.NumOut()
}

// Out implements Type.
func (n *TypeDeclaration) Out(i int) Type {
	if n.typ == nil {
		panic("Out of an inappropriate type")
	}

	return n.typ.Out(i)
}

func (n *TypeDeclaration) str(w *bytes.Buffer) {
	if b := n.qualifier; b != 0 {
		w.Write(dict.S(b))
		w.WriteByte('.')
	}
	w.Write(dict.S(n.Name()))
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

func (n *VarDeclaration) check(ctx *context) (stop bool) {
	done, stop := n.guard.check(ctx, n)
	if done || stop {
		return stop
	}

	defer n.guard.done(ctx)

	if n.expr.check(ctx) || n.typ0.check(ctx) {
		return true
	}

	switch {
	case n.typ0 != nil:
		n.Type = n.typ0.Type
		if n.expr == nil { // var v T
			break
		}

		v := n.expr.Value
		if v == nil {
			break
		}

		switch v.Kind() {
		case ConstValue:
			c := v.Const()
			if !c.AssignableTo(n.Type) {
				ctx.constAssignmentFail(n.expr, n.Type, c)
			}
		case NilValue:
			ctx.mustAssignNil(n.expr, n.Type)
		case RuntimeValue:
			if !v.Type().AssignableTo(n.Type) {
				ctx.valueAssignmentFail(n.expr, n.Type, v)
			}
		default:
			//dbg("", v.Kind())
			todo(n)
		}
	default:
		v := n.expr.Value
		if v == nil {
			break
		}

		switch v.Kind() {
		case ConstValue:
			switch c := v.Const(); {
			case c.Untyped():
				t := c.Type()
				if t.Kind() == UntypedBool {
					t = ctx.boolType
				}
				if ctx.mustConvertConst(n, t, c) != nil {
					n.Type = t
				}
			default:
				n.Type = c.Type()
			}
		case RuntimeValue:
			n.Type = v.Type()
		case NilValue:
			return ctx.err(n.expr, "use of untyped nil")
		default:
			todo(n, true)
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
			lx.declarationScope.declare(lx, newVarDeclaration(-1, v, typ0, nil, scopeStart))
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
			case lx.declarationScope.Bindings[v.Val] != nil:
				if op == ":=" {
					continue
				}
			default:
				if v.Val != idUnderscore {
					hasNew = true
				}
			}
			lx.declarationScope.declare(lx, newVarDeclaration(i, v, typ0, exprs[0], scopeStart))
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
			case lx.declarationScope.Bindings[v.Val] != nil:
				if op == ":=" {
					continue
				}
			default:
				if v.Val != idUnderscore {
					hasNew = true
				}
			}
			lx.declarationScope.declare(lx, newVarDeclaration(-1, v, typ0, e, scopeStart))
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

func (g *gate) check(ctx *context, d Declaration) (done, stop bool) {
	switch *g {
	case gateReady:
		if d != nil {
			ctx.stack = append(ctx.stack, d)
		}
		*g = gateOpen
		return false, false
	case gateOpen:
		stack := ctx.stack
		if d != nil {
			stack = append(ctx.stack, d)
		}
		if len(stack) == 0 {
			return true, false
		}

		d = stack[len(stack)-1]
		var a []Declaration
		for _, v := range stack[:len(stack)-1] {
			if v == d {
				a = append(a, v)
			}
		}
		if len(a) == 0 {
			return true, false
		}

		*g = gateClosed
		if f := ctx.errf; f != nil {
			ctx.errf = nil
			return true, f()
		}

		node := ctx.node
		if node == nil {
			node = d
		}
		var prolog string
		switch d.(type) {
		case *ConstDeclaration:
			prolog = "constant definition loop"
		case
			*FuncDeclaration,
			*ImportDeclaration,
			*LabelDeclaration:
			panic("internal error")
		case *TypeDeclaration:
			return true, ctx.err(node, "invalid recursive type %s", dict.S(d.Name()))
		case *VarDeclaration:
			todo(d)
			return true, false
		default:
			panic("internal error")
		}
		for i, v := range stack[:len(stack)-1] {
			if v == d {
				var a []string
				for j, v := range stack[i : len(stack)-1] {
					a = append(a, fmt.Sprintf("\t%s: %s uses %s", position(v.Pos()), dict.S(v.Name()), dict.S(stack[i+j+1].Name())))
				}
				return true, ctx.err(node, "%s\n%s", prolog, strings.Join(a, "\n"))
			}
		}
	case gateClosed:
		return true, false
	}
	panic("internal error")
}

func (g *gate) done(c *context) {
	if c != nil {
		c.pop()
	}
	*g = gateClosed
}

type context struct {
	*Context
	stack []Declaration
	node  Node
	iota  Value
	errf  func() bool
	pkg   *Package
}

func (c *context) pop() { c.stack = c.stack[:len(c.stack)-1] }
