// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"github.com/cznic/xc"
)

var (
	_ Declaration = (*ImportDeclaration)(nil)
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
	Node() Node
	Name() int // Name ID.
}

// Bindings map name IDs to declarations.
type Bindings map[int]Declaration

func (b *Bindings) declare(d Declaration) Declaration {
	if *b == nil {
		*b = Bindings{}
	}

	m := *b
	nm := d.Name()
	if ex := m[nm]; ex != nil {
		return ex
	}

	m[nm] = d
	return nil
}

func (b *Bindings) mustDeclare(lx *lexer, d Declaration) {
	ex := b.declare(d)
	if ex == nil {
		return
	}

	lx.err(d.Node(), "%s redeclared, previous declaration at %s", dict.S(d.Name()), position(ex.Node().Pos()))
}

// Scope tracks declarations.
type Scope struct {
	Bindings Bindings
	Kind     ScopeKind
	Parent   *Scope
	Unbound  []Declaration // Declaration with name _.
}

func newScope(kind ScopeKind, parent *Scope) *Scope {
	return &Scope{
		Kind:   kind,
		Parent: parent,
	}
}

func (s *Scope) mustDeclare(lx *lexer, d Declaration) {
	if d.Name() == idUnderscore {
		s.Unbound = append(s.Unbound, d)
		return
	}

	s.Bindings.mustDeclare(lx, d)
}

// ScopeKind is a scope category.
type ScopeKind int

// ImportDeclaration represents a named import declaration of a single package.
// The name comes from an explicit identifier (import foo "bar") when present.
// Otherwise the name in the package clause (package foo) is used.
type ImportDeclaration struct {
	Decl    *ImportSpec // AST link.
	Package *Package
	name    int
	once    *xc.Once
}

// Implements Declaration.
func (n *ImportDeclaration) Node() Node { return n.Decl }

// Implements Declaration.
func (n *ImportDeclaration) Name() int { return n.name }
