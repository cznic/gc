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
	"path/filepath"
	"sort"

	"github.com/cznic/golex/lex"
	"github.com/cznic/mathutil"
)

// Package is a Go package.
type Package struct {
	Context       *Context
	Directory     string
	Files         []*File // ASTs.
	ImportPath    string
	Inits         []*FuncDeclaration // All init() functions.
	Name          string
	Scope         *Scope // PackageScope or UniverseScope.
	SourceFiles   []string
	avoid         map[int]token.Pos // Names declared in file scope.
	forwardTypes  Bindings
	importPath    int
	ipBase        int
	name          int
	namedBy       string
	parseOnlyName bool
}

func newPackage(c *Context, importPath, dir string) *Package {
	if isRelativeImportPath(importPath) {
		panic("internal error")
	}

	return &Package{
		Context:    c,
		Directory:  dir,
		ImportPath: importPath,
		Scope:      newScope(PackageScope, c.universe),
		avoid:      map[int]token.Pos{},
		importPath: dict.SID(importPath),
		ipBase:     dict.SID(filepath.Base(importPath)),
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

	c := p.Context
	if len(p.Files) == 0 && !c.options.disableNoBuildableFilesError {
		return fmt.Errorf("package %s: no buildable Go source files in %s", p.ImportPath, p.Directory)
	}

	var a []string
	for k := range p.forwardTypes {
		a = append(a, string(dict.S(k)))
	}
	sort.Strings(a)
	for _, v := range a {
		nm := dict.SID(v)
		d := p.forwardTypes[nm]
		ex := c.universe.Bindings[nm]
		if _, ok := ex.(*TypeDeclaration); ok {
			c.err(d, "cannot define new methods on non-local type %s", v)
			continue
		}

		c.err(d, "undefined %s", v)
	}

	p.Scope.check(c)
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
