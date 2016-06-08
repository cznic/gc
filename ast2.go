// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"go/token"
	"path/filepath"
	"strings"
	"unicode"
)

// Node represents the interface all AST nodes, xc.Token and lex.Char implement.
type Node interface {
	Pos() token.Pos
}

// ----------------------------------------------------------------- ImportSpec

func (n *ImportSpec) post(lx *lexer) {
	var nm int
	if o := n.IdentifierOpt; o != nil {
		nm = o.Token.Val
	}
	var ip string
	var bl *BasicLiteral
	switch bl = n.BasicLiteral; bl.Case {
	case
		0, //CHAR_LIT
		1, //FLOAT_LIT
		2, //IMAG_LIT
		3: //INT_LIT
		lx.err(bl, "import statement requires a string")
		return
	case 4: //STRING_LIT
		val := bl.val
		if val == nil {
			return
		}

		ip = val.(StringID).String()
		if ip == "" {
			lx.err(bl, "import path is empty")
			return
		}

		if ip[0] == '/' {
			lx.err(bl, "import path cannot be absolute path")
			return
		}

		for _, r := range ip {
			if strings.ContainsAny(ip, "!\"#$%&'()*,:;<=>?[\\]^`{|}") || r == '\ufffd' {
				lx.err(bl, "import path contains invalid character")
				return
			}

			if !unicode.In(r, unicode.L, unicode.M, unicode.N, unicode.P, unicode.S) {
				lx.err(bl, "invalid import path")
				return
			}
		}
	default:
		panic("internal error")
	}
	id := &ImportDeclaration{
		Decl: n,
		name: nm,
		once: lx.oncePackage(bl, ip, lx.pkg.Directory), // Async.
	}
	lx.imports = append(lx.imports, id)
}

// -------------------------------------------------------------- PackageClause

func (n *PackageClause) post(lx *lexer) {
	p := lx.pkg
	t := n.Token2
	nm := t.Val
	switch {
	case p.Name != "":
		if p.name != nm {
			lx.close(t,
				"package %s: found packages %s (%s) and %s (%s) in %s",
				p.ImportPath,
				p.Name, filepath.Base(p.namedBy),
				t.S(), filepath.Base(lx.name),
				p.Directory,
			)
			return
		}
	default:
		p.Name = string(t.S())
		p.name = nm
		p.namedBy = lx.name
	}
}

// ------------------------------------------------------------------- Prologue

func (n *Prologue) post(lx *lexer) {
	pkg := lx.pkg
	for _, v := range lx.imports {
		p := v.once.Value().(*Package)
		v.Package = p
		switch v.Name() {
		case idDot:
			lx.dotImports = append(lx.dotImports, v)
			for _, v := range p.Scope.Bindings {
				p.Scope.mustDeclare(lx, v)
			}
			continue
		case idUnderscore:
			lx.unboundImports = append(lx.unboundImports, v)
			continue
		case 0:
			v.name = p.name
		default:
			v.name = dict.SID(filepath.Base(p.ImportPath))
		}
		if _, ok := pkg.avoid[v.name]; !ok {
			pkg.avoid[v.name] = v.Node().Pos()
		}
		lx.fileScope.mustDeclare(lx, v)
	}
}
