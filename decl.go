// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"go/token"

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
}

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

	switch d.(type) {
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
	pkg := lx.pkg
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

		if _, ok := pkg.avoid[nm]; !ok {
			pkg.avoid[nm] = d.Pos()
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
		if s.Kind == PackageScope { // TLD.
			if ex := lx.pkg.avoid[nm]; ex != 0 {
				lx.err(d, "%s redeclared, previous declaration at %s", dict.S(nm), position(ex))
				return
			}
		}

		if s.Kind == PackageScope && nm == idInit {
			if x, ok := d.(*FuncDeclaration); ok {
				pkg.Inits = append(pkg.Inits, x)
				return
			}

			lx.err(d, "cannot declare init - must be function")
			return
		}

		s.Bindings.declare(lx, d)
	}
}

// ScopeKind is a scope category.
type ScopeKind int

type ConstDeclaration struct {
	name       int
	pos        token.Pos
	scopeStart token.Pos
}

func newConstDeclaration(nm xc.Token, scopeStart token.Pos) *ConstDeclaration {
	return &ConstDeclaration{
		name:       nm.Val,
		pos:        nm.Pos(),
		scopeStart: scopeStart,
	}
}

// Implements Declaration.
func (n *ConstDeclaration) Pos() token.Pos { return n.pos }

// Implements Declaration.
func (n *ConstDeclaration) Name() int { return n.name }

// Implements Declaration.
func (n *ConstDeclaration) ScopeStart() token.Pos { return n.scopeStart }

type FuncDeclaration struct {
	name int
	pos  token.Pos
}

func newFuncDeclaration(nm xc.Token) *FuncDeclaration {
	return &FuncDeclaration{
		name: nm.Val,
		pos:  nm.Pos(),
	}
}

func (n *FuncDeclaration) Pos() token.Pos { return n.pos }

func (n *FuncDeclaration) Name() int { return n.name }

func (n *FuncDeclaration) ScopeStart() token.Pos { panic("ScopeStart of FuncDeclaration") }

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

// Implements Declaration.
func (n *ImportDeclaration) Pos() token.Pos { return n.pos }

// Implements Declaration.
func (n *ImportDeclaration) Name() int { return n.name }

// Implements Declaration.
func (n *ImportDeclaration) ScopeStart() token.Pos { panic("ScopeStart of ImportDeclaration") }

type LabelDeclaration struct {
	tok xc.Token // AST link.
}

func newLabelDeclaration(tok xc.Token) *LabelDeclaration {
	return &LabelDeclaration{
		tok: tok,
	}
}

// Implements Declaration.
func (n *LabelDeclaration) Pos() token.Pos { return n.tok.Pos() }

// Implements Declaration.
func (n *LabelDeclaration) Name() int { return n.tok.Val }

// Implements Declaration.
func (n *LabelDeclaration) ScopeStart() token.Pos { panic("ScopeStart of LabelDeclaration") }

type TypeDeclaration struct {
	name int
	pos  token.Pos
}

func newTypeDeclaration(nm xc.Token) *TypeDeclaration {
	return &TypeDeclaration{
		name: nm.Val,
		pos:  nm.Pos(),
	}
}

// Implements Declaration.
func (n *TypeDeclaration) Pos() token.Pos { return n.pos }

// Implements Declaration.
func (n *TypeDeclaration) Name() int { return n.name }

// Implements Declaration.
func (n *TypeDeclaration) ScopeStart() token.Pos { return n.pos }

type VarDeclaration struct {
	name       int
	pos        token.Pos
	scopeStart token.Pos
}

func newVarDeclaration(nm xc.Token, scopeStart token.Pos) *VarDeclaration {
	return &VarDeclaration{
		name:       nm.Val,
		pos:        nm.Pos(),
		scopeStart: scopeStart,
	}
}

// Implements Declaration.
func (n *VarDeclaration) Pos() token.Pos { return n.pos }

// Implements Declaration.
func (n *VarDeclaration) Name() int { return n.name }

// Implements Declaration.
func (n *VarDeclaration) ScopeStart() token.Pos { return n.scopeStart }
