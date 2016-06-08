// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"go/scanner"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/cznic/xc"
)

var (
	// IsValidOS lists valid OS types.
	IsValidOS = map[string]bool{ // Go 1.7
		"android":   true,
		"darwin":    true,
		"dragonfly": true,
		"freebsd":   true,
		"linux":     true,
		"nacl":      true,
		"netbsd":    true,
		"openbsd":   true,
		"plan9":     true,
		"solaris":   true,
		"windows":   true,
	}

	// ArchMap maps valid CPU architectures to their respective models.
	ArchMap = map[string]Model{ // Go 1.7
		"386":         {4, 4},
		"amd64":       {8, 8},
		"amd64p32":    {8, 4},
		"arm64":       {8, 8},
		"arm64be":     {8, 8},
		"arm":         {4, 4},
		"armbe":       {4, 4},
		"mips64":      {8, 8},
		"mips64le":    {8, 8},
		"mips64p32":   {8, 4},
		"mips64p32le": {8, 4},
		"mips":        {8, 8}, //TODO ?
		"mipsle":      {8, 8}, //TODO ?
		"ppc64":       {8, 8},
		"ppc64le":     {8, 8},
		"ppc":         {4, 4},
		"s390":        {4, 4},
		"s390x":       {8, 8},
		"sparc64":     {8, 8},
		"sparc":       {4, 4},
	}
)

func isValidArch(s string) bool {
	_, ok := ArchMap[s]
	return ok
}

// Model describes CPU architecture details.
type Model struct {
	IntBytes int // Size of int.
	PtrBytes int // Size of *T, unsafe.Pointer and uintptr.
}

type testContext struct {
	errChecks   []xc.Token  //
	errChecksMu sync.Mutex  //
	exampleAST  interface{} //
	exampleRule int         //
}

type contextOptions struct {
	disableNoBuildableFilesError bool  //
	enableGenerics               bool  //
	errLimit                     int32 //
	errLimit0                    int32 //
}

// Context represents data shared by all packages loaded by LoadPackages.
type Context struct {
	Model       Model               //
	fileCentral *xc.FileCentral     //
	goarch      string              //
	goos        string              //
	gopaths     []string            //
	goroot      string              //
	options     *contextOptions     //
	report      *xc.Report          //
	searchPaths []string            //
	tags        map[string]struct{} //
	test        *testContext        //
	universe    *Scope              //
}

// NewContext returns a newly created Context.
func NewContext(goos, goarch, goroot string, gopaths, tags []string, opts ...Opt) (*Context, error) {
	if err := sanitizeContext(goos, goarch, goroot, gopaths); err != nil {
		return nil, err
	}

	return newContext(goos, goarch, goroot, gopaths, tags, opts...)
}

func newContext(goos, goarch, goroot string, gopaths, tags []string, opts ...Opt) (*Context, error) {
	gopaths = dedup(gopaths)
	searchPaths := []string{filepath.Join(goroot, "src")}
	for _, v := range gopaths {
		searchPaths = append(searchPaths, filepath.Join(v, "src"))
	}
	report := xc.NewReport()
	report.ErrLimit = -1
	c := &Context{
		Model:       ArchMap[goarch],
		fileCentral: xc.NewFileCentral(),
		goarch:      goarch,
		goos:        goos,
		gopaths:     gopaths,
		goroot:      goroot,
		options:     &contextOptions{errLimit: 10},
		report:      report,
		searchPaths: searchPaths,
		tags:        map[string]struct{}{},
		universe:    newScope(UniverseScope, nil),
	}
	for _, v := range tags {
		c.tags[v] = struct{}{}
	}
	for _, v := range opts {
		if err := v(c); err != nil {
			return nil, err
		}
	}
	c.options.errLimit0 = c.options.errLimit
	return c, nil
}

func (c *Context) newPackage(importPath, directory string) *Package {
	return newPackage(c, importPath, directory)
}

func (c *Context) errPos(pos token.Pos, format string, arg ...interface{}) bool {
	if atomic.AddInt32(&c.options.errLimit, -1) < 0 {
		return true // Close
	}

	c.report.Err(pos, format, arg...)
	return false
}

func (c *Context) errors(err ...error) error {
	for _, v := range err {
		switch v.(type) {
		case nil, scanner.ErrorList:
			// nop
		default:
			c.errPos(0, "%s", v)
		}
	}
	return c.report.Errors(false)
}

func (c *Context) clearErrors() {
	c.report.ClearErrors()
	c.options.errLimit = c.options.errLimit0
}

// DirectoryFromImportPath returns the directory where the source files of
// package importPath are to be searched for. Relative import paths are
// computed relative to basePath.
func (c *Context) DirectoryFromImportPath(importPath, basePath string) (string, error) {
	if strings.HasPrefix(importPath, "./") || strings.HasPrefix(importPath, "../") {
		return filepath.Join(basePath, importPath), nil
	}

	for _, v := range c.searchPaths {
		dir := filepath.Join(v, importPath)
		fi, err := os.Stat(dir)
		if err != nil {
			if !os.IsNotExist(err) {
				return "", err
			}

			continue
		}

		if !fi.IsDir() {
			continue
		}

		return dir, nil
	}

	a := []string{fmt.Sprintf("cannot find package %q in any of:", importPath)}
	for _, v := range c.searchPaths {
		a = append(a, "\t"+v)
	}

	return "", fmt.Errorf("%s", strings.Join(a, "\n"))
}

// FilesFromImportPath returns the directory where the source files for package
// importPath are; a list of normal and testing (*_test.go) go source files or
// an error, if any. A relative import path is considered to be relative to
// basePath.
func (c *Context) FilesFromImportPath(importPath, basePath string) (dir string, sourceFiles []string, testFiles []string, err error) {
	if importPath == "C" {
		return "", nil, nil, nil
	}

	if dir, err = c.DirectoryFromImportPath(importPath, basePath); err != nil {
		return "", nil, nil, err
	}

	matches, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		return "", nil, nil, err
	}

	for _, match := range matches {
		b := filepath.Base(match)
		if ex := filepath.Ext(b); ex != "" {
			b = b[:len(b)-len(ex)]
		}
		isTestFile := false
		if strings.HasSuffix(b, "_test") {
			isTestFile = true
			b = b[:len(b)-len("_test")]
		}
		a := strings.Split(b, "_")
		if len(a) > 1 { // *_GOOS or *_GOARCH
			if s := a[len(a)-1]; isValidArch(s) && s != c.goarch {
				continue
			}

			if s := a[len(a)-1]; IsValidOS[s] && s != c.goos {
				continue
			}
		}
		if len(a) > 2 { //  *_GOOS_GOARCH
			if s := a[len(a)-2]; IsValidOS[s] && s != c.goos {
				continue
			}
		}
		switch {
		case isTestFile:
			testFiles = append(testFiles, match)
		default:
			sourceFiles = append(sourceFiles, match)
		}

	}
	if len(sourceFiles) == 0 {
		var err error
		if !c.options.disableNoBuildableFilesError {
			err = fmt.Errorf("package %s: no buildable Go source files in %s", importPath, dir)
		}
		return "", nil, nil, err
	}

	return dir, sourceFiles, testFiles, err
}

func (c *Context) oncePackage(n Node, importPath, basePath string) *xc.Once {
	var pos token.Pos
	if n != nil {
		pos = n.Pos()
	}
	return c.fileCentral.Once(
		importPath,
		func() interface{} {
			dir, sourceFiles, _, err := c.FilesFromImportPath(importPath, basePath)
			p := c.newPackage(importPath, dir)
			if err != nil {
				c.errPos(pos, "%s", err)
				return p
			}

			p.SourceFiles = sourceFiles
			if err := p.load(); err != nil {
				c.errPos(0, "%v", err)
			}
			return p
		},
	)
}

func (c *Context) pkgMap() map[string]*Package {
	m := map[string]*Package{}
	c.fileCentral.Map(func(s string, o *xc.Once) bool {
		m[s] = o.Value().(*Package)
		return true
	})
	return m
}

func (c *Context) loadPackageFiles(importPath, dir string, sourceFiles []string) (map[string]*Package, error) {
	p := c.newPackage(importPath, dir)
	p.SourceFiles = append([]string(nil), sourceFiles...)
	err := p.load()
	return c.pkgMap(), c.errors(err)
}

func (c *Context) loadPackage(importPath string, sourceFiles []string) (map[string]*Package, error) {
	c.clearErrors()
	var dir string
	if len(sourceFiles) != 0 {
		dir = filepath.Dir(sourceFiles[0])
	}
	return c.loadPackageFiles(importPath, dir, sourceFiles)
}

func (c *Context) loadPackages(importPaths []string) (map[string]*Package, error) {
	c.clearErrors()
	importPaths = dedup(importPaths)
	onces := make([]*xc.Once, len(importPaths))
	for i, v := range importPaths {
		onces[i] = c.oncePackage(nil, v, "")
	}

	// Wait for all packages to load.
	for _, once := range onces {
		once.Value()
	}

	return c.pkgMap(), c.errors()
}
