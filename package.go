// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"bufio"
	"bytes"
	"fmt"
	"go/token"
	"io"
	"os"
	"sort"

	"github.com/cznic/golex/lex"
	"github.com/cznic/mathutil"
)

// Package is a Go package.
type Package struct {
	Context     *Context
	Directory   string
	Files       []*File // ASTs.
	ImportPath  string
	Inits       []*FuncDeclaration // All init() functions.
	Name        string
	Scope       *Scope // PackageScope or UniverseScope.
	SourceFiles []string
	avoid       map[int]token.Pos // Names declared in file scope.
	name        int
	namedBy     string
}

func newPackage(c *Context, importPath, dir string) *Package {
	return &Package{
		Context:    c,
		Directory:  dir,
		ImportPath: importPath,
		Scope:      newScope(PackageScope, c.universe),
		avoid:      map[int]token.Pos{},
	}
}

// Load files in p.SourceFiles.
func (p *Package) load() error {
	sort.Strings(p.SourceFiles)
	for _, v := range p.SourceFiles {
		if err := p.loadFile(v); err != nil {
			return err
		}
	}

	if len(p.Files) == 0 && !p.Context.options.disableNoBuildableFilesError {
		return fmt.Errorf("package %s: no buildable Go source files in %s", p.ImportPath, p.Directory)
	}

	//TODO p.post()
	return nil
}

func (p *Package) loadFile(nm string) error {
	f, err := os.Open(nm)
	if err != nil {
		return err
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}

	if !fi.Mode().IsRegular() {
		return fmt.Errorf("%s: not a regular file", nm)
	}

	sz := fi.Size()
	if sz > mathutil.MaxInt {
		return fmt.Errorf("%s: file too big", nm)
	}

	return p.loadReader(nm, int(sz), f)
}

func (p *Package) loadReader(nm string, sz int, r io.Reader) error {
	rr, ok := r.(io.RuneReader)
	if !ok {
		rr = bufio.NewReader(r)
	}

	lx, err := newLexer(nm, sz, rr, p)
	if err != nil {
		return err
	}

	yyParse(lx)
	if p.Context.test != nil { // Want all errorcheck directives.
		for lx.scan() != lex.RuneEOF {
		}
	}
	return nil
}

func (p *Package) loadString(nm, src string) error {
	return p.loadReader(nm, len(src), bytes.NewBufferString(src))
}
