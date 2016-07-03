// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"go/scanner"
	"go/token"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/cznic/xc"
)

const (
	// DefaultIntConstBits is the default maximum untyped integer constant
	// size.
	DefaultIntConstBits = 512
	// DefaultFloatConstPrec is the default maximum untyped floating point
	// constant precision.
	DefaultFloatConstPrec = 512

	builtinSrc = `

package builtin

const (
	true = 0 // Expression does not matter.
	false
	iota 
)

var nil int // Type does not matter.

// Signatures do not matter.
func append()
func cap()
func close()
func complex()
func copy()
func delete()
func imag()
func len()
func make()
func new()
func panic()
func print()
func println()
func real()
func recover()
	
type error interface {
	Error() string
}

`
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
	errChecks   []xc.Token
	errChecksMu sync.Mutex
	exampleAST  interface{}
	exampleRule int
	pkgMap      map[string][]string
}

type contextOptions struct {
	disableNoBuildableFilesError bool
	enableGenerics               bool
	errLimit                     int32
	errLimit0                    int32
}

// Context represents data shared by all packages loaded by LoadPackages.
type Context struct {
	bigIntMaxInt     *big.Int
	bigIntMaxUint    *big.Int
	bigIntMaxUintptr *big.Int
	bigIntMinInt     *big.Int
	boolType         Type
	complex128Type   Type
	falseValue       *constValue
	fileCentral      *xc.FileCentral
	float32Type      Type
	float64Type      Type
	floatConstPrec   uint
	floatMaxInt      float64
	floatMaxUint     float64
	floatMaxUintptr  float64
	floatMinInt      float64
	goarch           string
	goos             string
	gopaths          []string
	goroot           string
	int32Type        Type
	intConstBits     uint
	intType          Type
	maxInt           int64
	maxUint          uint64
	maxUintptr       uint64
	minInt           int64
	model            Model
	nilValue         *nilValue
	options          *contextOptions
	report           *xc.Report
	searchPaths      []string
	stringType       Type
	tags             map[string]struct{}
	test             *testContext
	trueValue        *constValue
	uintptrType      Type
	universe         *Scope
	voidType         Type
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
	model := ArchMap[goarch]
	c := &Context{
		fileCentral:    xc.NewFileCentral(),
		floatConstPrec: DefaultFloatConstPrec,
		goarch:         goarch,
		goos:           goos,
		gopaths:        gopaths,
		goroot:         goroot,
		intConstBits:   DefaultIntConstBits,
		model:          model,
		options:        &contextOptions{errLimit: 10},
		report:         report,
		searchPaths:    searchPaths,
		tags:           map[string]struct{}{},
		universe:       newScope(UniverseScope, nil),
	}
	c.setupLimits()
	for _, v := range tags {
		c.tags[v] = struct{}{}
	}
	c.tags[goos] = struct{}{}
	c.tags[goarch] = struct{}{}
	for _, v := range opts {
		if err := v(c); err != nil {
			return nil, err
		}
	}
	c.options.errLimit0 = c.options.errLimit
	if err := c.declareBuiltins(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Context) setupLimits() {
	switch c.model.IntBytes {
	case 4:
		c.bigIntMaxInt = bigIntMaxInt32
		c.bigIntMaxUint = bigIntMaxUint32
		c.bigIntMinInt = bigIntMinInt32
		c.floatMaxInt = math.MaxInt32
		c.floatMaxUint = math.MaxUint32
		c.floatMinInt = math.MinInt32
		c.maxInt = math.MaxInt32
		c.maxUint = math.MaxUint32
		c.minInt = math.MinInt32
	case 8:
		c.bigIntMaxInt = bigIntMaxInt64
		c.bigIntMaxUint = bigIntMaxUint64
		c.bigIntMinInt = bigIntMinInt64
		c.floatMaxInt = math.MaxInt64
		c.floatMaxUint = math.MaxUint64
		c.floatMinInt = math.MinInt64
		c.maxInt = math.MaxInt64
		c.maxUint = math.MaxUint64
		c.minInt = math.MinInt64
	default:
		panic("invalid model IntBytes")
	}
	switch c.model.PtrBytes {
	case 4:
		c.bigIntMaxUintptr = bigIntMaxUint32
		c.floatMaxUintptr = math.MaxUint32
		c.maxUintptr = math.MaxUint32
	case 8:
		c.bigIntMaxUintptr = bigIntMaxUint64
		c.floatMaxUintptr = math.MaxUint64
		c.maxUintptr = math.MaxUint64
	default:
		panic("invalid model PtrBytes")
	}
}

func (c *Context) declareBuiltins() error {
	for _, v := range []struct {
		name              string
		kind              Kind
		align, fieldAlign int
		size              uint64
	}{
		{"bool", Bool, 1, 1, 1},
		{"complex64", Complex64, 8, 8, 8},
		{"complex128", Complex128, 8, 8, 16},
		{"float32", Float32, 4, 4, 4},
		{"float64", Float64, 8, 8, 8},
		{"int", Int, c.model.IntBytes, c.model.IntBytes, uint64(c.model.IntBytes)},
		{"int8", Int8, 1, 1, 1},
		{"int16", Int16, 2, 2, 2},
		{"int32", Int32, 4, 4, 4},
		{"int64", Int64, 8, 8, 8},
		{"string", String, c.model.PtrBytes, c.model.PtrBytes, uint64(2 * c.model.PtrBytes)},
		{"uint", Uint, c.model.IntBytes, c.model.IntBytes, uint64(c.model.IntBytes)},
		{"uint8", Uint8, 1, 1, 1},
		{"uint16", Uint16, 2, 2, 2},
		{"uint32", Uint32, 4, 4, 4},
		{"uint64", Uint64, 8, 8, 8},
		{"uintptr", Uintptr, c.model.PtrBytes, c.model.PtrBytes, uint64(c.model.PtrBytes)},
	} {
		t := xc.Token{Val: dict.SID(v.name)}
		d := newTypeDeclaration(nil, t, nil)
		d.ctx = c
		base := d.base()
		base.align = v.align
		base.fieldAlign = v.fieldAlign
		base.kind = v.kind
		base.size = v.size
		base.typ = d
		c.universe.declare(nil, d)
	}

	b := c.universe.Bindings
	b[dict.SID("byte")] = b[dict.SID("uint8")]
	b[dict.SID("rune")] = b[dict.SID("int32")]

	c.boolType = b[dict.SID("bool")].(*TypeDeclaration)
	c.complex128Type = b[dict.SID("complex128")].(*TypeDeclaration)
	c.float32Type = b[dict.SID("float32")].(*TypeDeclaration)
	c.float64Type = b[dict.SID("float64")].(*TypeDeclaration)
	c.int32Type = b[dict.SID("int32")].(*TypeDeclaration)
	c.intType = b[dict.SID("int")].(*TypeDeclaration)
	c.stringType = b[dict.SID("string")].(*TypeDeclaration)
	c.uintptrType = b[dict.SID("uintptr")].(*TypeDeclaration)
	c.voidType = newTupleType(&context{Context: c}, nil)

	p := c.newPackage("", "")
	p.Scope = c.universe
	if err := p.loadString("", builtinSrc); err != nil {
		return err
	}

	if t := c.test; t != nil && t.exampleRule != 0 {
		return nil
	}

	b[dict.SID("false")].(*ConstDeclaration).Value = newConstValue(newBoolConst(false, c.boolType, true))
	b[dict.SID("true")].(*ConstDeclaration).Value = newConstValue(newBoolConst(true, c.boolType, true))

	c.falseValue = newConstValue(newBoolConst(false, c.boolType, true))
	c.nilValue = newNilValue()
	c.trueValue = newConstValue(newBoolConst(true, c.boolType, true))
	return nil
}

func (c *Context) newPackage(importPath, directory string) *Package {
	return newPackage(c, importPath, directory)
}

func (c *Context) err(n Node, format string, arg ...interface{}) bool {
	return c.errPos(n.Pos(), format, arg...)
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
// package importPath are to be searched for.
func (c *Context) DirectoryFromImportPath(importPath string) (string, error) {
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
// an error, if any.
func (c *Context) FilesFromImportPath(importPath string) (dir string, sourceFiles []string, testFiles []string, err error) {
	if t := c.test; t != nil {
		if sf, ok := t.pkgMap[importPath]; ok {
			return filepath.Dir(sf[0]), sf, nil, nil
		}
	}

	if importPath == "C" {
		return "", nil, nil, nil
	}

	if dir, err = c.DirectoryFromImportPath(importPath); err != nil {
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

func (c *Context) oncePackage(n Node, importPath string) *xc.Once {
	var pos token.Pos
	if n != nil {
		pos = n.Pos()
	}
	return c.fileCentral.Once(
		importPath,
		func() interface{} {
			dir, sourceFiles, _, err := c.FilesFromImportPath(importPath)
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
		onces[i] = c.oncePackage(nil, v)
	}

	// Wait for all packages to load.
	for _, once := range onces {
		once.Value()
	}

	return c.pkgMap(), c.errors()
}

func (c *Context) isPredeclared(d Declaration) bool {
	return d != nil && d == c.universe.Bindings[d.Name()]
}

// Except for shift operations, if the operands of a binary operation are
// different kinds of untyped constants, the operation and, for non-boolean
// operations, the result use the kind that appears later in this list:
// integer, rune, floating-point, complex.
var untypedArithmeticBinOpTab = [maxKind][maxKind]Kind{
	Int:        {Int: Int},
	Int32:      {Int: Int32, Int32: Int32},
	Float64:    {Int: Float64, Int32: Float64, Float64: Float64},
	Complex128: {Int: Complex128, Int32: Complex128, Float64: Complex128, Complex128: Complex128},
}

func (c *Context) untypedArithmeticBinOpType(a, b Type) Type {
	ak := a.Kind()
	bk := b.Kind()
	if ak > bk {
		ak, bk = bk, ak
	}
	switch untypedArithmeticBinOpTab[bk][ak] {
	case Int32:
		return c.int32Type
	case Int:
		return c.intType
	case Float64:
		return c.float64Type
	case Complex128:
		return c.complex128Type
	default:
		return nil
	}
}

func (c *Context) valueAssignmentFail(n Node, t Type, v Value) {
	switch v.Kind() {
	case ConstValue:
		c.constAssignmentFail(n, t, v.Const())
	default:
		c.err(n, "cannot use type %s as type %s in assignment", v.Type(), t)
	}
}

func (c *Context) constAssignmentFail(n Node, t Type, d Const) {
	if d.Untyped() && d.Type().ConvertibleTo(t) &&
		!(d.Type().Kind() == String && t.Kind() == Slice && (t.Elem().Kind() == Uint8 || t.Elem().Kind() == Int32)) {
		c.constConversionFail(n, t, d)
		return
	}

	c.err(n, "cannot use %s (type %s) as type %s in assignment", d, d.Type(), t)
}

func (c *Context) constConversionFail(n Node, t Type, d Const) {
	switch {
	case t.Kind() == Interface && d.Type().Implements(t):
		// nop
	case !d.Type().ConvertibleTo(t):
		c.err(n, "cannot convert type %s to %s", d.Type(), t)
	case d.Type().FloatingPointType() && t.IntegerType() && !d.Integral():
		if k := d.Type().Kind(); k == Complex64 || k == Complex128 {
			todo(n, true) // ... truncated to real
			break
		}

		c.err(n, "constant %s truncated to integer", d)
	default:
		c.err(n, "constant %s overflows %s", d, t)
	}
}

func (c *Context) arithmeticBinOpShape(a, b Const, n Node) (Type, bool, Const, Const) {
	switch {
	case a.Untyped():
		switch {
		case b.Untyped():
			t := c.untypedArithmeticBinOpType(a.Type(), b.Type())
			d := a.convert(t)
			if d == nil {
				c.constConversionFail(n, t, a)
			}

			e := b.convert(t)
			if e == nil {
				c.constConversionFail(n, t, b)
			}

			if d != nil && e != nil {
				return t, true, d, e
			}
		default: // a.Untyped && !b.Untyped
			t := b.Type()
			if d := a.mustConvert(n, t); d != nil {
				return t, false, d, b.convert(t)
			}
		}
	case b.Untyped(): // !a.Untyped() && b.Untyped()
		t := a.Type()
		if d := b.mustConvert(n, t); d != nil {
			return t, false, a.convert(t), d
		}
	default: // !a.Untyped() && !b.Untyped()
		t := a.Type()
		if b.Type() != t {
			todo(n, true)
			break
		}

		d := a.convert(t)
		if d == nil {
			c.constConversionFail(n, t, a)
		}

		e := b.convert(t)
		if e == nil {
			c.constConversionFail(n, t, b)
		}

		if d != nil && e != nil {
			return t, false, d, e
		}
	}
	return nil, false, nil, nil
}

func (c *Context) mustConvertConst(n Node, t Type, d Const) Const {
	e := d.normalize()
	if e == nil {
		todo(n, true)
		//TODO c.err(n, "constant %s overflows %s", d, d.Type())
		return nil
	}

	if t == nil {
		return e
	}

	if f := e.Convert(t); f != nil {
		if f.Kind() == ConstValue {
			return f.Const()
		}
	}

	c.constConversionFail(n, t, e)
	return nil
}

func (c *Context) compositeLiteralValueFail(n Node, v Value, t Type) bool {
	switch v.Kind() {
	case ConstValue:
		return c.err(n, "cannot use %s (type %s) as type %s in array or slice literal", v.Const(), v.Type(), t)
	default:
		return c.err(n, "cannot use value of type %s as type %s in array or slice literal", v.Type(), t)
	}
}
