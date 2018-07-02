// Copyright 2016 The GC Authors. All rights reservedGG.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"

	"github.com/cznic/token"
)

var (
	_ Declaration = (*ConstDecl)(nil)
	_ Declaration = (*FuncDecl)(nil)
	_ Declaration = (*ImportDecl)(nil)
	_ Declaration = (*MethodDecl)(nil)
	_ Declaration = (*TypeDecl)(nil)
	_ Declaration = (*VarDecl)(nil)
)

type checker interface {
	check(*cstack)
	checked()
	checking()
	state() sentinel
}

// DeclarationKind describes a Declaration's Kind.
type DeclarationKind int

// Values of DeclarationKind.
const (
	ConstDeclaration DeclarationKind = iota
	FuncDeclaration
	ImportDeclaration
	MethodDeclaration
	TypeDeclaration
	VarDeclaration
)

// ScopeKind describe a Scope's Kind.
type ScopeKind int

// Values of ScopeKind.
const (
	UniverseScope ScopeKind = iota
	PackageScope
	FileScope
	BlockScope
)

const (
	_ sentinel = iota
	checking
	checked
	checkedFalse
	checkedTrue
)

type sentinel byte

func (s *sentinel) checked()       { *s = checked }
func (s *sentinel) checking()      { *s = checking }
func (s sentinel) state() sentinel { return s }

// ------------------------------------------------------------------- Bindings

// Bindings map names to Declarations.
type Bindings map[string]Declaration

func (b *Bindings) declare(sf *SourceFile, d Declaration) {
	if *b == nil {
		*b = Bindings{}
	}
	m := *b
	nm := d.Name()
	if nm == "" {
		panic(fmt.Errorf("%s: internal error 010: %T.declare - empty declaration name", d.Position(), b))
	}

	ex := m[nm]
	if ex == nil {
		m[nm] = d
		return
	}

	if sf.Package.ctx.tweaks.ignoreRedeclarations {
		return
	}

	switch x := d.(type) {
	case *MethodDecl:
		switch y := ex.(type) {
		case *MethodDecl:
			switch eptr, ptr := y.IsPtrReceiver(), x.IsPtrReceiver(); {
			case eptr != ptr:
				sf.Package.errorList.add(
					d.Position(),
					"method redeclared: %s.%s\n\tmethod(%s) %s\n\tmethod(%s) %s",
					y.Receiver, d.Name(),
					y.Receiver, y.Type(),
					x.Receiver, x.Type(),
				)
			default:
				s := y.receiver()
				if eptr {
					s = "(" + s + ")"
				}
				sf.Package.errorList.add(
					d.Position(),
					"%s.%s redeclared in this block:\n\tprevious declaration at %v",
					s, d.Name(),
					ex.Position(),
				)
			}
		default:
			sf.Package.errorList.add(d.Position(), "TODO109 %T", y)
		}
		return
	case *VarDecl:
		if x.argument {
			sf.Package.errorList.add(d.Position(), "duplicate argument %v", d.Name())
			return
		}
	}

	sf.Package.errorList.add(
		d.Position(),
		"%v redeclared in this block\n\tprevious declaration at %v",
		d.Name(),
		ex.Position(),
	)
}

// ---------------------------------------------------------------- declaration

type declaration struct {
	Token
	typ        Type
	visibility token.Pos

	sentinel
	universe bool
}

func (d *declaration) setType(t Type) { d.typ = t }

// Name implements Declaration.
func (d *declaration) Name() string { return d.Val }

// Visibility implements Declaration.
func (d *declaration) Visibility() token.Pos { return d.visibility }

// Type implements Declaration.
func (d *declaration) Type() Type { return d.typ }

// String implements Declaration.
func (d *declaration) String() string { return d.Val }

// ---------------------------------------------------------------------- Scope

// Scope tracks declarations.
type Scope struct {
	Bindings Bindings
	Kind     ScopeKind
	Labels   Bindings
	Parent   *Scope
	Unbound  []Declaration // Declarations named _. TODO needed?
	pkg      *Package      // Only for Package scope.
}

func newScope(kind ScopeKind, parent *Scope) *Scope {
	return &Scope{
		Kind:   kind,
		Parent: parent,
	}
}

func (s *Scope) declare(sf *SourceFile, d Declaration) {
	nm := d.Name()
	pkg := sf.Package
	if s.Kind == PackageScope {
		sf.TopLevelDecls = append(sf.TopLevelDecls, d)
		if !pkg.ctx.tweaks.ignoreRedeclarations && nm != "_" && d.Kind() != MethodDeclaration {
			switch ex, ok := pkg.fileScopeNames[nm]; {
			case ok && ex.Kind() == ImportDeclaration:
				sf.Package.errorList.add(ex.Position(), "%s: %s redeclared in this block\n\tprevious declaration at %s:", ex.Position(), nm, d.Position())
				return
			}

			pkg.fileScopeNames[nm] = d
		}
	}
	if nm == "_" {
		if s.Kind != FileScope {
			s.Unbound = append(s.Unbound, d)
		}
		return
	}

	switch x := d.(type) {
	case *ConstDecl, *VarDecl, *TypeDecl, *FuncDecl:
		s.Bindings.declare(sf, d)
	case *MethodDecl:
		switch {
		case s.Kind == PackageScope:
			m := sf.Package.Methods
			r := x.receiver()
			if r == "" {
				break
			}

			b := m[r]
			if b == nil {
				b = Bindings{}
				m[r] = b
			}
			b.declare(sf, x)
		default:
			s.Bindings.declare(sf, d)
		}
	case *ImportDecl:
		if s.Kind != FileScope {
			panic("internal error 011")
		}

		switch ex, ok := pkg.fileScopeNames[nm]; {
		case ok && ex.Kind() != ImportDeclaration:
			panic(fmt.Errorf("%s: TODO218 %q %v", d.Position(), nm, ex.Position())) //TODO
		default:
			pkg.fileScopeNames[nm] = d
			s.Bindings.declare(sf, d)
		}
	default:
		panic("internal error 012")
	}
}

// Lookup searches for nm in s or any of its parents. fileScope, if not nil, is
// searched  prior to searching the universe scope when lookup does not find nm
// in package scope and s belongs to pkg.
func (s *Scope) Lookup(pkg *Package, fileScope *Scope, nm Token) Declaration {
	for s0 := s; s != nil; s = s.Parent {
		if d, ok := s.Bindings[nm.Val]; ok && (s.Kind != BlockScope || s != s0 || d.Visibility() < nm.Pos()) {
			return d
		}

		if s.Kind == PackageScope && s.pkg == pkg && fileScope != nil {
			if d, ok := fileScope.Bindings[nm.Val]; ok {
				return d
			}
		}
	}
	return nil
}

// ImportDecl is an import declaration.
type ImportDecl struct {
	ImportPath string // `foo/bar` in `import "foo/bar"`
	Qualifier  string // `baz` in `import baz "foo/bar"`.
	declaration
	pkg *Package // The imported package, if exists.

	Dot bool // The `import . "foo/bar"` variant is used.
}

func newImportDecl(tok Token, off int, dot bool, qualifier, importPath string) *ImportDecl {
	return &ImportDecl{
		Dot:         dot,
		ImportPath:  importPath,
		Qualifier:   qualifier,
		declaration: declaration{Token: tok, visibility: token.NoPos},
	}
}

func (d *ImportDecl) Package() *Package { return d.pkg }
func (d *ImportDecl) check(s *cstack)   { d.checked() }

// ConstDecl describes a constant declaration.
type ConstDecl struct {
	declaration
	Value Operand
	list  []Expr
	index int
}

func newConstDecl(tok Token, visibility token.Pos) *ConstDecl {
	return &ConstDecl{
		declaration: declaration{Token: tok, visibility: visibility},
	}
}

func (d *ConstDecl) check(s *cstack) {
	defer func() { s.check(d.Type(), d) }()

	if p := d.Package(); p.isBuiltin() {
		d.universe = true
		s := p.ctx.universe.Bindings
		switch x := s[d.Token.Val].(type) {
		case *ConstDecl:
			t := x.Value.Type()
			t.setPos(d.pos())
			d.setType(t)
			d.Value = x.Value
			s[d.Token.Val] = d
		default:
			p.errorList.add(d.Position(), "TODO222 %T.check %T %q %v", d, x, d.Token.Val, x)
			return
		}
	}

	if d.index == 0 {
		for _, v := range d.list {
			s.check(v, d)
		}
	}
	if d.Type() == nil {
		//TODO infer type
		return
	}
}

// Kind implements Declaration.
func (d *ConstDecl) Kind() DeclarationKind { return ConstDeclaration }

// FuncDecl describes a function declaration.
type FuncDecl struct {
	declaration
	Body *StmtBlock
}

func newFuncDecl(tok Token, visibility token.Pos) *FuncDecl {
	return &FuncDecl{
		declaration: declaration{Token: tok, visibility: visibility},
	}
}

func (d *FuncDecl) check(s *cstack) {
	defer func() { s.check(d.Type(), d) }()

	p := d.Package()
	if p.isBuiltin() {
		d.universe = true
		s := p.ctx.universe.Bindings
		switch x := s[d.Token.Val].(type) {
		case *FuncDecl:
			t := x.Type()
			t.setPos(d.pos())
			d.setType(t)
			s[d.Token.Val] = d
		default:
			p.errorList.add(d.Position(), "TODO262 %T.check %T %q %v", d, x, d.Token.Val, x)
			return
		}
	}
}

// Kind implements Declaration.
func (d *FuncDecl) Kind() DeclarationKind { return FuncDeclaration }

func (d *FuncDecl) checkBody(s *cstack) {
	if d.Body != nil {
		s.reset().check(d.Body, nil)
	}
}

// MethodDecl describes a method declaration.
type MethodDecl struct {
	declaration
	Receiver Type
	Body     *StmtBlock
}

func newMethodDecl(tok Token, visibility token.Pos) *MethodDecl {
	return &MethodDecl{
		declaration: declaration{Token: tok, visibility: visibility},
	}
}

func (d *MethodDecl) check(s *cstack) {
	defer func() { s.check(d.Type(), d) }()

	sf := d.SourceFile()
	var rx *NamedType
	switch x := d.Receiver.(type) {
	case *NamedType:
		rx = x
	case *PointerType:
		switch y := x.Element.(type) {
		case *NamedType:
			rx = y
		case *QualifiedNamedType:
			sf.Package.errorList.add(y.Position(), "cannot define new methods on non-local type %s", y)
		default:
			sf.Package.errorList.add(x.Position(), "TODO434")
		}
	case nil:
		// nop
	default:
		sf.Package.errorList.add(x.Position(), "TODO437")
	}

	if rx != nil {
		s.check(rx, d)
		switch x := rx.alias().(type) {
		case *NamedType:
			if b := d.SourceFile().Scope.Bindings; b != nil {
				if _, ok := b[x.Declaration().Name()]; ok {
					sf.Package.errorList.add(d.Position(), "cannot define new methods on non-local type %s", x)
				}
			}
		case *QualifiedNamedType:
			sf.Package.errorList.add(d.Position(), "cannot define new methods on non-local type %s", x)
		case TypeKind:
			sf.Package.errorList.add(d.Position(), "cannot define new methods on non-local type %s", x)
			return
		default:
			sf.Package.errorList.add(d.Position(), "invalid receiver type %s (%s is an unnamed type)", x, x)
		}
	}
	sf.Package.Scope.declare(sf, d)
}

// IsPtrReceiver reports whether d's receiver is a pointer.
func (d *MethodDecl) IsPtrReceiver() bool {
	return d.Receiver != nil && d.Receiver.alias().Kind() == Pointer
}

func (d *MethodDecl) receiver() string {
	if d.Receiver == nil {
		return ""
	}

	s := d.Receiver.alias().String()
	if s != "" && s[0] == '*' {
		s = s[1:]
	}
	return s
}

func (d *MethodDecl) checkBody(s *cstack) {
	if d.Body != nil {
		s.reset().check(d.Body, nil)
	}
}

// Kind implements Declaration.
func (d *MethodDecl) Kind() DeclarationKind { return MethodDeclaration }

// TypeDecl describes a type declaration.
type TypeDecl struct {
	declaration

	alias bool
}

func newTypeDecl(tok Token, visibility token.Pos) *TypeDecl {
	return &TypeDecl{
		declaration: declaration{Token: tok, visibility: visibility},
	}
}

func (d *TypeDecl) check(s *cstack) {
	defer func() { s.check(d.Type(), d) }()

	p := d.Package()
	if p.isBuiltin() {
		d.universe = true
		s := p.ctx.universe.Bindings
		switch x := s[d.Token.Val].(type) {
		case *TypeDecl:
			t := x.Type()
			if _, ok := t.(TypeKind); ok && t.isGeneric(nil) {
				delete(s, d.Token.Val)
				return
			}

			t.setPos(d.pos())
			d.setType(t)
			s[d.Token.Val] = d
		default:
			p.errorList.add(d.Position(), "TODO286 %T.check %T %q %v", d, x, d.Token.Val, x)
			return
		}
	}

	if p.ImportPath == "unsafe" && d.Name() == "ArbitraryType" {
		d.setType(BuiltinGeneric)
	}
}

// Kind implements Declaration.
func (d *TypeDecl) Kind() DeclarationKind { return TypeDeclaration }

// VarDecl describes a variable declaration.
type VarDecl struct {
	declaration

	Initializer Expr
	tupleIndex  int

	argument bool // both in and out
}

func newVarDecl(tok Token, visibility token.Pos, argument bool) *VarDecl {
	return &VarDecl{
		argument:    argument,
		declaration: declaration{Token: tok, visibility: visibility},
	}
}

func (d *VarDecl) check(s *cstack) {
	defer func() {
		s.check(d.Type(), d)
	}()

	p := d.Package()
	if p.isBuiltin() {
		d.universe = true
		b := p.ctx.universe.Bindings
		switch x := b[d.Token.Val].(type) {
		case *VarDecl:
			t := x.Type()
			t.setPos(d.pos())
			d.setType(t)
			b[d.Token.Val] = d
		default:
			p.errorList.add(d.Position(), "TODO327 %T.check %T %q %v", d, x, d.Token.Val, x)
			return
		}
	}

	var et Type
	if d.Initializer != nil {
		s.check(d.Initializer, d)
		et = d.Initializer.Type()
	}

	if et == nil {
		return
	}

	switch {
	case d.tupleIndex < 0:
		if et.Kind() == Tuple {
			//TODO panic(fmt.Sprintf("%v: TODO597", d.Position()))
			break
		}

		if d.Type() == nil {
			d.setType(et.defaultType())
			break
		}

		s.check(d.Type(), d)
		if !et.AssignableTo(d.Type()) {
			//dbg("%v: %q", d.Position(), d.Name())
			//dbg("%v: %T(%s)", et.Position(), et, et)
			//dbg("%v: %T(%s)", d.Type().Position(), d.Type(), d.Type())
			//panic("TODO575")
			p.errorList.add(d.Position(), "cannot use %s (type %s) as type %s in assignment", d.Initializer.str(), et.alias(), d.Type())
		}
	default:
		et, ok := et.(*TupleType)
		if !ok {
			d.err("assignment mismatch: more than 1 variable but only 1 value")
			break
		}

		if d.tupleIndex >= len(et.List) {
			panic("TODO623")
		}

		if d.Type() == nil {
			d.setType(et.List[d.tupleIndex].defaultType())
			break
		}

		s.check(d.Type(), d)
		if !et.AssignableTo(d.Type()) {
			panic("TODO632")
		}
	}
}

// Kind implements Declaration.
func (d *VarDecl) Kind() DeclarationKind { return VarDeclaration }

// Declaration is one of *ConstDecl, *FuncDecl, *ImportSpec, *MethodDecl,
// *TypeDecl or *VarDecl.
type Declaration interface {
	//TODO Node
	// Kind returns the Declarations's kind.
	Kind() DeclarationKind

	// Name returns the declared name.
	Name() string

	// Position returns the position of the declaration.
	Position() token.Position

	// Visibility returns the position at which the declaration is visible
	// in its declaration scope or token.NoPos for declarations in package
	// and file scope.
	Visibility() token.Pos

	// Type returns the declaration type. N/A for ImportDeclaration.
	Type() Type

	Package() *Package
	SourceFile() *SourceFile
	String() string

	checker
	err(string, ...interface{})
	err0(token.Position, string, ...interface{})
	setType(Type)
}
