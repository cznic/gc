// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate.go
//go:generate golex -o scanner.go scanner.l
//go:generate stringer -type ScopeKind
//go:generate stringer -type ValueKind
//go:generate go run generate.go -2

// Package gc is a Go compiler front end. Work in progess (5.34%).
//
// CGO
//
// CGO is not supported.
package gc

// Opt is a LoadPackages option.
type Opt func(*Context) error

// EnableGenerics enables experimental support for generics. See also
//
//  https://github.com/golang/go/issues/15292
func EnableGenerics() Opt {
	return func(c *Context) error {
		c.options.enableGenerics = true
		return nil
	}
}

// ErrLimit sets the maximum number of errors reported.
func ErrLimit(n int32) Opt {
	return func(c *Context) error {
		c.options.errLimit = n
		return nil
	}
}

// LoadPackage loads package importPath, consisting of sourceFiles, as well as
// any transitively imported packages, and returns a mapping from import paths
// to all loaded packages and an error, if any. The error, if non nil, is
// possibly a scanner.ErrorList instance.
//
// The importPath argument is not checked in any way. It's only recorded in the
// relevant Package instance.  This allows to load packages consisting of
// arbitrary files from arbitrary places.
//
// The behavior of the function may be optionally amended by passing options in
// opts.
func LoadPackage(goos, goarch, goroot, importPath string, gopaths, tags, sourceFiles []string, opts ...Opt) (map[string]*Package, error) {
	c, err := NewContext(goos, goarch, goroot, gopaths, tags, opts...)
	if err != nil {
		return nil, err
	}

	return c.loadPackage(importPath, sourceFiles)
}

// LoadPackages loads all packages in importPaths, as well as any transitively
// imported packages, and returns a mapping from import paths to all loaded
// packages and an error, if any. The error, if non nil, is possibly a
// scanner.ErrorList instance.
//
// The behavior of the function may be optionally amended by passing options in
// opts.
func LoadPackages(goos, goarch, goroot string, gopaths, tags, importPaths []string, opts ...Opt) (map[string]*Package, error) {
	c, err := NewContext(goos, goarch, goroot, gopaths, tags, opts...)
	if err != nil {
		return nil, err
	}

	return c.loadPackages(importPaths)
}
