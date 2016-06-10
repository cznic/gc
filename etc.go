// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/cznic/golex/lex"
	"github.com/cznic/strutil"
	"github.com/cznic/xc"
)

var (
	printHooks = strutil.PrettyPrintHooks{}

	idC          = dict.SID("C")
	idDot        = dict.SID(".")
	idInit       = dict.SID("init")
	idUnderscore = dict.SID("_")
)

func init() {
	for k, v := range xc.PrintHooks {
		printHooks[k] = v
	}
	lcRT := reflect.TypeOf(lex.Char{})
	lcH := func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		c := v.(lex.Char)
		r := c.Rune
		s := yySymName(int(r))
		if x := s[0]; x >= '0' && x <= '9' {
			s = strconv.QuoteRune(r)
		}
		f.Format("%s%v: %s"+suffix, prefix, position(c.Pos()), s)
	}

	printHooks[lcRT] = lcH
	printHooks[reflect.TypeOf(xc.Token{})] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		t := v.(xc.Token)
		if !t.Pos().IsValid() {
			return
		}

		lcH(f, t.Char, prefix, "")
		if s := dict.S(t.Val); len(s) != 0 {
			f.Format(" %q", s)
		}
		f.Format(suffix)
		return
	}
}

func position(pos token.Pos) token.Position { return xc.FileSet.Position(pos) }

// PrettyString returns pretty formatted strings of things produced by this package.
func PrettyString(v interface{}) string { return strutil.PrettyString(v, "", "", printHooks) }

// Preserves order.
func dedup(a []string) []string {
	var r []string
	m := map[string]struct{}{}
	for _, v := range a {
		if _, ok := m[v]; ok {
			continue
		}

		r = append(r, v)
		m[v] = struct{}{}
	}
	return r
}

func isDirectory(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}

func sanitizeContext(goos, goarch, goroot string, gopaths []string) error {
	if !IsValidOS[goos] {
		return fmt.Errorf("invalid goos: %q", goos)
	}

	if !isValidArch(goarch) {
		return fmt.Errorf("invalid goarch: %q", goarch)
	}

	if goroot == "" {
		return fmt.Errorf("invalid goroot: %q", goroot)
	}

	if !filepath.IsAbs(goroot) {
		return fmt.Errorf("goroot is not an absolute path: %q", goroot)
	}

	ok, err := isDirectory(goroot)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("goroot is not a directory: %q", goroot)
	}

	for _, v := range gopaths {
		if v == "" {
			return fmt.Errorf("invalid gopath: %q", v)
		}

		if !filepath.IsAbs(v) {
			return fmt.Errorf("gopath is not an absolute path: %q", v)
		}

		ok, err := isDirectory(v)
		if err != nil {
			return err
		}

		if !ok {
			return fmt.Errorf("gopath is not a directory: %q", v)
		}
	}

	return nil
}
