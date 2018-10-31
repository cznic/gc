// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/cznic/token"
	"github.com/edsrzf/mmap-go"
)

var (
	validOS = map[string]bool{ // Go 1.9
		"android":   true,
		"darwin":    true,
		"dragonfly": true,
		"freebsd":   true,
		"js":        true,
		"linux":     true,
		"nacl":      true,
		"netbsd":    true,
		"openbsd":   true,
		"plan9":     true,
		"solaris":   true,
		"windows":   true,
		"zos":       true,
	}

	archModels = map[string]model{ // Go 1.9
		"386":         {4, 4},
		"amd64":       {8, 8},
		"amd64p32":    {8, 4},
		"arm":         {4, 4},
		"arm64":       {8, 8},
		"arm64be":     {8, 8},
		"armbe":       {4, 4},
		"mips":        {8, 8}, //TODO ?
		"mips64":      {8, 8},
		"mips64le":    {8, 8},
		"mips64p32":   {8, 4},
		"mips64p32le": {8, 4},
		"mipsle":      {8, 8}, //TODO ?
		"ppc":         {4, 4},
		"ppc64":       {8, 8},
		"ppc64le":     {8, 8},
		"risc":        {4, 4},
		"riscv64":     {8, 8},
		"s390":        {4, 4},
		"s390x":       {8, 8},
		"sparc":       {4, 4},
		"sparc64":     {8, 8},
		"wasm":        {4, 4},
	}
)

func universeConsts() map[string]Operand {
	return map[string]Operand{
		"false": UntypedFalse{},
		"iota":  UntypedIntOperand{},
		"true":  UntypedTrue{},
	}
}

func universeFuncs() map[string]Type {
	return map[string]Type{
		"append": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: newSliceType(npos{}, BuiltinGeneric)},
					{Type: newSliceType(npos{}, BuiltinGeneric), Variadic: true},
				},
			},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: newSliceType(npos{}, BuiltinGeneric)},
				},
			},
		),
		"copy": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: newSliceType(npos{}, BuiltinGeneric)},
					{Type: newSliceType(npos{}, BuiltinGeneric)},
				},
			},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: Int},
				},
			},
		),
		"delete": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: newMapType(npos{}, BuiltinGeneric, BuiltinGeneric2)},
					{Type: BuiltinGeneric2},
				},
			},
			nil,
		),
		"len": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: BuiltinGeneric},
				},
			},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: Int},
				},
			},
		),
		"cap": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: BuiltinGeneric},
				},
			},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: Int},
				},
			},
		),
		"make": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: BuiltinGeneric},
					{Type: newSliceType(npos{}, BuiltinGenericInteger), Variadic: true},
				},
			},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: BuiltinGeneric},
				},
			},
		),
		"new": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: BuiltinGeneric},
				},
			},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: newPointerType(npos{}, BuiltinGeneric)},
				},
			},
		),
		"complex": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: BuiltinGenericFloat},
					{Type: BuiltinGenericFloat},
				},
			},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: BuiltinGenericComplex},
				},
			},
		),
		"real": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: BuiltinGenericComplex},
				},
			},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: BuiltinGenericFloat},
				},
			},
		),
		"imag": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: BuiltinGenericComplex},
				},
			},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: BuiltinGenericFloat},
				},
			},
		),
		"close": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: newChannelType(npos{}, TxChan, BuiltinGeneric)},
				},
			},
			nil,
		),
		"panic": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: newInterfaceType(npos{})},
				},
			},
			nil,
		),
		"recover": newFunctionType(
			npos{},
			nil,
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: newInterfaceType(npos{})},
				},
			},
		),
		"print": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: newSliceType(npos{}, BuiltinGeneric), Variadic: true},
				},
			},
			nil,
		),
		"println": newFunctionType(
			npos{},
			&ParamTypeListType{
				List: []*ParamTypeType{
					{Type: newSliceType(npos{}, BuiltinGeneric), Variadic: true},
				},
			},
			nil,
		),
	}
}

func universeTypes() map[string]Type {
	return map[string]Type{
		"ComplexType": BuiltinGenericComplex,
		"FloatType":   BuiltinGenericFloat,
		"IntegerType": BuiltinGenericInteger,
		"Type":        BuiltinGeneric,
		"Type1":       BuiltinGeneric2,

		"bool":       Bool,
		"byte":       Uint8,
		"complex128": Complex128,
		"complex64":  Complex64,
		"float32":    Float32,
		"float64":    Float64,
		"int":        Int,
		"int16":      Int16,
		"int32":      Int32,
		"int64":      Int64,
		"int8":       Int8,
		"rune":       Int32,
		"string":     String,
		"uint":       Uint,
		"uint16":     Uint16,
		"uint32":     Uint32,
		"uint64":     Uint64,
		"uint8":      Uint8,
		"uintptr":    Uintptr,

		"error": func() Type {
			r := newInterfaceType(npos{})
			r.Methods = map[string]*FunctionType{
				"Error": newFunctionType(
					npos{},
					emptyParamTypeList,
					&ParamTypeListType{
						List: []*ParamTypeType{
							{Type: String},
						},
					},
				),
			}
			return r
		}(),
	}
}

func universeVars() map[string]Type {
	return map[string]Type{
		"nil": UntypedNil,
	}
}

func isValidArch(s string) bool {
	_, ok := archModels[s]
	return ok
}

type model struct {
	intBytes int // Size of int.
	ptrBytes int // Size of *T, unsafe.Pointer and uintptr.
}

type tweaks struct {
	errLimit         int
	localImportsPath string

	errchk               bool // misc. tweaks for TestErrchk
	example              bool
	ignoreImports        bool // Test hook.
	ignoreRedeclarations bool
	noChecks             bool // Benchmark hook
	noErrorColumns       bool
}

// Option amends Context.
type Option func(c *Context) error

// IgnoreRedeclarations disables reporting redeclarations as errors.
func IgnoreRedeclarations() Option {
	return func(c *Context) error {
		c.tweaks.ignoreRedeclarations = true
		return nil
	}
}

// NoErrorColumns disable displaying of columns in error messages.
func NoErrorColumns() Option { // gc -C
	return func(c *Context) error {
		c.tweaks.noErrorColumns = true
		return nil
	}
}

// NoErrorLimit disables limit on number of errors reported.
func NoErrorLimit() Option { // gc -e
	return func(c *Context) error {
		c.tweaks.errLimit = -1
		return nil
	}
}

// LocalImportsPath sets the relative path for local imports.
func LocalImportsPath(s string) Option { // gc -D path
	return func(c *Context) error {
		c.tweaks.localImportsPath = s
		return nil
	}
}

// Add directory to import search path.
func addSearchPath(s string) Option { // gc -I path
	return func(c *Context) error {
		c.searchPaths = append(c.searchPaths, s)
		return nil
	}
}

// Context describes the context of loaded packages.
type Context struct {
	goarch      string
	goos        string
	model       model
	packages    map[string]*Package // Key: import path.
	packagesMu  sync.Mutex
	searchPaths []string
	tags        map[string]struct{}
	tweaks      tweaks
	universe    *Scope
}

// NewContext returns a newly created Context. tags are the build tags
// considered when loading packages having build directives (see
// https://golang.org/pkg/go/build/#hdr-Build_Constraints for details).
// searchPaths are examined when looking for a package to load.
func NewContext(goos, goarch string, tags, searchPaths []string, options ...Option) (*Context, error) {
	if !validOS[goos] {
		return nil, fmt.Errorf("unknown operating system: %s", goos)
	}

	model, ok := archModels[goarch]
	if !ok {
		return nil, fmt.Errorf("unknown architecture: %s", goarch)
	}

	tm := make(map[string]struct{}, len(tags))
	for _, v := range tags {
		tm[v] = struct{}{}
	}
	tm[goos] = struct{}{}
	tm[goarch] = struct{}{}
	c := &Context{
		goarch:      goarch,
		goos:        goos,
		model:       model,
		packages:    map[string]*Package{},
		searchPaths: append([]string(nil), searchPaths...),
		tags:        tm,
		tweaks:      tweaks{errLimit: 10},
	}
	for _, o := range options {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	c.universe = newScope(UniverseScope, nil)
	c.universe.Bindings = map[string]Declaration{}
	for k, v := range universeConsts() {
		d := newConstDecl(Token{Val: k}, token.NoPos)
		d.Value = v
		c.universe.Bindings[k] = d
	}
	for k, v := range universeFuncs() {
		d := newFuncDecl(Token{Val: k}, token.NoPos)
		d.setType(v)
		c.universe.Bindings[k] = d
	}
	for k, v := range universeTypes() {
		d := newTypeDecl(Token{Val: k}, token.NoPos)
		d.setType(v)
		c.universe.Bindings[k] = d
	}
	for k, v := range universeVars() {
		d := newVarDecl(Token{Val: k}, token.NoPos, false)
		d.setType(v)
		c.universe.Bindings[k] = d
	}
	pkg, err := c.Load("builtin")
	if err != nil {
		return nil, err
	}

	pkg.Scope = c.universe
	return c, nil
}

/*

Vendor Directories

Go 1.6 includes support for using local copies of external dependencies
to satisfy imports of those dependencies, often referred to as vendoring.

Code below a directory named "vendor" is importable only
by code in the directory tree rooted at the parent of "vendor",
and only using an import path that omits the prefix up to and
including the vendor element.

Here's the example from the previous section,
but with the "internal" directory renamed to "vendor"
and a new foo/vendor/crash/bang directory added:

    /home/user/gocode/
        src/
            crash/
                bang/              (go code in package bang)
                    b.go
            foo/                   (go code in package foo)
                f.go
                bar/               (go code in package bar)
                    x.go
                vendor/
                    crash/
                        bang/      (go code in package bang)
                            b.go
                    baz/           (go code in package baz)
                        z.go
                quux/              (go code in package main)
                    y.go

The same visibility rules apply as for internal, but the code
in z.go is imported as "baz", not as "foo/vendor/baz".

Code in vendor directories deeper in the source tree shadows
code in higher directories. Within the subtree rooted at foo, an import
of "crash/bang" resolves to "foo/vendor/crash/bang", not the
top-level "crash/bang".

Code in vendor directories is not subject to import path
checking (see 'go help importpath').

When 'go get' checks out or updates a git repository, it now also
updates submodules.

Vendor directories do not affect the placement of new repositories
being checked out for the first time by 'go get': those are always
placed in the main GOPATH, never in a vendor subtree.

See https://golang.org/s/go15vendor for details.

*/
func (c *Context) dirForImportPath(position token.Position, importPath string) (string, error) {
	if strings.HasPrefix(importPath, "./") {
		dir := c.tweaks.localImportsPath
		switch dir {
		case "", ".":
			dir = filepath.Dir(position.Filename)
		}
		ok, err := checkDir(dir)
		if !ok || err != nil {
			return dir, err
		}

		return dir, nil
	}

	var a []string
	if importPath != "C" && !strings.HasPrefix(importPath, "/") {
		s := importPath
		a = []string{filepath.Join(s, "vendor")}
		for s != "" {
			s = filepath.Dir(s)
			if s == "." {
				s = ""
			}
			a = append(a, filepath.Join(s, "vendor"))
		}
	}
	a = append(a, "")
	for _, v := range c.searchPaths {
		for _, w := range a {
			dir := filepath.Join(v, w, importPath)
			ok, err := checkDir(dir)
			if ok || err != nil {
				return dir, err
			}
		}
	}

	a = []string{fmt.Sprintf("cannot find package %q in any of:", importPath)}
	for i, v := range c.searchPaths {
		switch i {
		case 0:
			v += " (from GOROOT)"
		default:
			v += " (from GOPATH)"
		}
		a = append(a, "\t"+v)
	}

	return "", fmt.Errorf("%s", strings.Join(a, "\n"))
}

// NumPackages returns the number of loaded packages.
//
// The method is safe for concurrent use by multiple goroutines.
func (c *Context) NumPackages() int {
	c.packagesMu.Lock()
	n := len(c.packages)
	c.packagesMu.Unlock()
	return n
}

// SourceFileForPath searches loaded packages for the one containing the file
// at path and returns the corresponding SourceFile.  The result is nil if the
// file is not considered for build by any of the loaded packages.
//
// The method is safe for concurrent use by multiple goroutines.
func (c *Context) SourceFileForPath(path string) *SourceFile {
	c.packagesMu.Lock()
	defer c.packagesMu.Unlock()

	for _, p := range c.packages {
		for _, f := range p.SourceFiles {
			if f.Path == path {
				return f
			}
		}
	}

	return nil
}

// FilesForImportPath the package directory and a list of source and test files.
func (c *Context) FilesForImportPath(position token.Position, importPath string) (dir string, sourceFiles []string, testFiles []string, err error) {
	if importPath == "C" {
		return "", nil, nil, nil
	}

	if dir, err = c.dirForImportPath(position, importPath); err != nil {
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

			if s := a[len(a)-1]; validOS[s] && s != c.goos {
				continue
			}
		}
		if len(a) > 2 { //  *_GOOS_GOARCH
			if s := a[len(a)-2]; validOS[s] && s != c.goos {
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
	return dir, sourceFiles, testFiles, nil
}

func (c *Context) load(position token.Position, importPath string, syntaxError func(*parser), errList *errorList) *Package {
	//TODO must normalize relative import paths
	c.packagesMu.Lock()

	p := c.packages[importPath]
	if p != nil {
		c.packagesMu.Unlock()
		return p
	}

	p = newPackage(c, importPath, "", errList)
	c.packages[importPath] = p
	c.packagesMu.Unlock()

	go func() {
		dir, files, _, err := c.FilesForImportPath(position, importPath)
		if err != nil {
			errList.forcedAdd(position, "%s", err)
			close(p.ready)
			return
		}

		p.Dir = dir
		p.load(position, files, syntaxError)
	}()

	return p
}

func (c *Context) build(files []string, errList *errorList) *Package {
	dir := filepath.Dir(files[0])
	ip := dir
	for _, v := range c.searchPaths {
		if strings.HasPrefix(dir, v) {
			ip = ip[len(v):]
			break
		}
	}
	p := newPackage(c, ip, "", errList)
	p.Dir = dir
	go func() {
		p.load(token.Position{}, files, nil)
	}()

	return p
}

// Load finds the package in importPath and returns the resulting Package or an
// error if any.
//
// The method is safe for concurrent use by multiple goroutines.
func (c *Context) Load(importPath string) (*Package, error) {
	err := newErrorList(c.tweaks.errLimit, c.tweaks.noErrorColumns)
	p := c.load(token.Position{}, importPath, nil, err).waitFor()
	if err := err.error(); err != nil {
		return nil, err
	}

	return p, nil
}

// Build constructs a package from files and returns the resulting Package or
// an error if any. Build panics if len(files) == 0.
//
// The method is safe for concurrent use by multiple goroutines.
func (c *Context) Build(files []string) (*Package, error) {
	errList := newErrorList(c.tweaks.errLimit, c.tweaks.noErrorColumns)
	p := c.build(files, errList).waitFor()
	if err := errList.error(); err != nil {
		return nil, err
	}

	return p, nil
}

// SourceFile describes a source file.
type SourceFile struct {
	File          *token.File
	ImportSpecs   []*ImportDecl
	InitFunctions []Declaration
	Package       *Package
	Path          string
	Scope         *Scope // File scope.
	TopLevelDecls []Declaration
	enquedChecks  []checker
	f             *os.File  // Underlying src file.
	src           mmap.MMap // Valid only during parsing and checking.
	srcMu         sync.Mutex

	build bool
}

func newSourceFile(pkg *Package, path string, f *os.File, src mmap.MMap) *SourceFile {
	var (
		s *Scope
	)
	if pkg != nil {
		s = newScope(FileScope, pkg.Scope)
	}
	var nm string
	if f != nil {
		nm = f.Name()
	}
	file := token.NewFile(nm, len(src))
	return &SourceFile{
		File:    file,
		Package: pkg,
		Path:    path,
		Scope:   s,
		build:   true,
		f:       f,
		src:     src,
	}
}

func (s *SourceFile) enqueue(c checker) { s.enquedChecks = append(s.enquedChecks, c) }

func (s *SourceFile) checkEnqueued(cs *cstack) {
	for {
		n := len(s.enquedChecks)
		if n == 0 {
			return
		}

		c := s.enquedChecks[n-1]
		s.enquedChecks = s.enquedChecks[:n-1]
		if c == nil {
			continue
		}

		switch c.state() {
		case 0:
			cs.reset().check(c, nil)
		case checking:
			panic("TODO763") // cycle
		case checked:
			// ok
		default:
			panic("internal error 035")
		}
	}
}

func (s *SourceFile) init(pkg *Package, path string) {
	s.Package = pkg
	s.ImportSpecs = s.ImportSpecs[:0]
	s.TopLevelDecls = s.TopLevelDecls[:0]
	s.Path = path
	s.build = true
}

func (s *SourceFile) finit() {
	s.srcMu.Lock()
	if s.src != nil {
		s.src.Unmap()
		s.src = nil
	}
	if s.f != nil {
		s.f.Close()
		s.f = nil
	}
	s.enquedChecks = nil
	s.srcMu.Unlock()
}

// Package describes a package.
type Package struct {
	Dir            string
	ImportPath     string
	ImportedBy     map[string]struct{} // import path: struct{}.
	Imports        map[string]struct{} // import path: struct{}.
	Methods        map[string]Bindings // receiver name: *MethodDecl
	Name           string
	Scope          *Scope // Package scope.
	SourceFiles    []*SourceFile
	ctx            *Context
	errorList      *errorList
	fileScopeNames map[string]Declaration
	importedByMu   sync.Mutex
	nameOrigin     npos
	ready          chan struct{}
}

func newPackage(ctx *Context, importPath, nm string, errorList *errorList) *Package {
	var s *Scope
	if ctx != nil {
		s = newScope(PackageScope, ctx.universe)
	}
	p := &Package{
		ImportPath:     importPath,
		ImportedBy:     map[string]struct{}{},
		Imports:        map[string]struct{}{},
		Methods:        map[string]Bindings{},
		Name:           nm,
		Scope:          s,
		ctx:            ctx,
		errorList:      errorList,
		fileScopeNames: map[string]Declaration{},
		ready:          make(chan struct{}),
	}
	if s != nil {
		s.pkg = p
	}
	return p
}

func (p *Package) isBuiltin() bool { return p != nil && p.ImportPath == "builtin" }

func (p *Package) checkEnqueued(s *cstack) {
	for _, v := range p.SourceFiles {
		v.checkEnqueued(s)
	}
	for _, v := range p.SourceFiles {
		if len(v.enquedChecks) != 0 {
			panic("TODO843")
		}
	}
}

func (p *Package) load(position token.Position, paths []string, syntaxError func(*parser)) {
	returned := false
	defer func() {
		err := recover()
		if !returned && err == nil {
			err = fmt.Errorf("panic(nil)\n%s", debug.Stack())
		}
		if err != nil {
			//dbg("%v\n%s", err, stack())
			p.errorList.forcedAdd(token.Position{}, "%v", err)
		}

		for _, v := range p.SourceFiles {
			v.finit()
		}
		p.fileScopeNames = nil
		close(p.ready)
	}()

	l := newLexer(nil, nil)
	y := newParser(nil, nil)
	l.errHandler = y.err0
	l.commentHandler = y.commentHandler
	y.syntaxError = syntaxError

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			p.errorList.forcedAdd(position, "%s", err)
			returned = true
			return
		}

		fi, err := f.Stat()
		if err != nil {
			f.Close()
			p.errorList.forcedAdd(position, "%s", err)
			returned = true
			return

		}

		var src mmap.MMap
		if fi.Size() > 0 {
			if src, err = mmap.Map(f, 0, 0); err != nil {
				f.Close()
				p.errorList.forcedAdd(position, "%s", err)
				returned = true
				return
			}
		}

		sf := newSourceFile(p, path, f, src)
		l.sourceFile = sf
		l.init(sf.File, src)
		y.init(sf, l)
		y.file()
		switch {
		case sf.build:
			p.SourceFiles = append(p.SourceFiles, sf)
		default:
			sf.finit()
		}
	}
	if p.ctx.tweaks.noChecks {
		returned = true
		return
	}

	//TODO check unbound
	cs := &cstack{p: p}
	for _, sf := range p.SourceFiles {
		for _, d := range sf.TopLevelDecls {
			if x, ok := d.(*MethodDecl); ok {
				cs.reset().check(x, nil)
			}
		}
	}
	for _, sf := range p.SourceFiles {
		for _, d := range sf.TopLevelDecls {
			if _, ok := d.(*MethodDecl); !ok {
				cs.reset().check(d, nil)
			}
		}
	}
	p.checkEnqueued(cs)
	for _, sf := range p.SourceFiles {
		for _, d := range sf.InitFunctions {
			d.(*FuncDecl).checkBody(cs)
		}
	}
	p.checkEnqueued(cs)
	for _, sf := range p.SourceFiles {
		for _, d := range sf.TopLevelDecls {
			switch x := d.(type) {
			case *FuncDecl:
				x.checkBody(cs)
			case *MethodDecl:
				x.checkBody(cs)
			}
		}
	}
	p.checkEnqueued(cs)
	returned = true
}

func (p *Package) waitFor() *Package {
	if p == nil {
		return nil
	}

	<-p.ready
	return p
}
