// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"go/token"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/cznic/golex/lex"
	"github.com/cznic/strutil"
	"github.com/cznic/xc"
)

const (
	unicodeSurrogateFirst = 0xd800
	unicodeSurrogateLast  = 0xdfff
)

var (
	printHooks = strutil.PrettyPrintHooks{}

	idAlignof    = dict.SID("Alignof")
	idAppend     = dict.SID("append")
	idBool       = dict.SID("bool")
	idByte       = dict.SID("byte")
	idC          = dict.SID("C")
	idCap        = dict.SID("cap")
	idClose      = dict.SID("close")
	idComplex    = dict.SID("complex")
	idComplex128 = dict.SID("complex128")
	idComplex64  = dict.SID("complex64")
	idCopy       = dict.SID("copy")
	idDelete     = dict.SID("delete")
	idDot        = dict.SID(".")
	idFalse      = dict.SID("false")
	idFloat32    = dict.SID("float32")
	idFloat64    = dict.SID("float64")
	idImag       = dict.SID("imag")
	idInit       = dict.SID("init")
	idInt        = dict.SID("int")
	idInt16      = dict.SID("int16")
	idInt32      = dict.SID("int32")
	idInt64      = dict.SID("int64")
	idInt8       = dict.SID("int8")
	idIota       = dict.SID("iota")
	idLen        = dict.SID("len")
	idMake       = dict.SID("make")
	idNew        = dict.SID("new")
	idNil        = dict.SID("nil")
	idOffsetof   = dict.SID("Offsetof")
	idPanic      = dict.SID("panic")
	idPointer    = dict.SID("Pointer")
	idPrint      = dict.SID("print")
	idPrintln    = dict.SID("println")
	idReal       = dict.SID("real")
	idRecover    = dict.SID("recover")
	idRune       = dict.SID("rune")
	idSizeof     = dict.SID("Sizeof")
	idString     = dict.SID("string")
	idTrue       = dict.SID("true")
	idUint       = dict.SID("uint")
	idUint16     = dict.SID("uint16")
	idUint32     = dict.SID("uint32")
	idUint64     = dict.SID("uint64")
	idUint8      = dict.SID("uint8")
	idUintptr    = dict.SID("uintptr")
	idUnderscore = dict.SID("_")

	todoPanic bool // Dev hook.
	todoTrace bool // Dev hook.
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

	printHooks[reflect.TypeOf((*constValue)(nil))] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		f.Format(prefix)
		f.Format("%s", v)
		f.Format(suffix)
	}

	printHooks[reflect.TypeOf(Bindings(nil))] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		f.Format(prefix)
		b := v.(Bindings)
		f.Format("%T{%i\n", b)
		a := make([]string, 0, len(b))
		for k := range b {
			a = append(a, string(dict.S(k)))
		}
		sort.Strings(a)
		for _, v := range a {
			f.Format("%q: %T,\n", v, b[dict.SID(v)])
		}
		f.Format("%u}")
		f.Format(suffix)
	}

	printHooks[reflect.TypeOf((*Scope)(nil))] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		f.Format(prefix)
		s := v.(*Scope)
		t := fmt.Sprintf("%T", s)
		f.Format("&%s{%i\n", t[1:])
		if len(s.Bindings) != 0 {
			f.Format(strutil.PrettyString(s.Bindings, "Bindings: ", ",\n", printHooks))
		}
		f.Format("Kind: %s,\n", s.Kind)
		if s.Parent != nil {
			f.Format(strutil.PrettyString(s.Parent, "Parent: ", ",\n", printHooks))
		}
		f.Format("%u}")
		f.Format(suffix)
	}
}

func todo(n Node, opt ...bool) { //TODO-
	if todoPanic {
		panic(position(n.Pos()).String())
	}

	if len(opt) != 0 && todoTrace {
		p := position(n.Pos()).String()
		if !strings.HasPrefix(p, "testdata") {
			return
		}

		_, fn, fl, _ := runtime.Caller(1)
		fmt.Fprintf(os.Stderr, "trace %s:%d: ", filepath.Base(fn), fl)
		fmt.Fprintf(os.Stderr, "%s\n", p)
		os.Stderr.Sync()
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

func isRelativeImportPath(p string) bool {
	return strings.HasPrefix(p, "./") || strings.HasPrefix(p, "../")
}

// Adapted from github.com/golang/go: go/ast. See GO-LICENSE.
func isExported(nm int) bool {
	r, _ := utf8.DecodeRune(dict.S(nm))
	return unicode.IsUpper(r)
}

func stringLiteralValue(lx *lexer, t xc.Token) stringValue {
	s := string(t.S())
	value, err := strconv.Unquote(s)
	if err != nil {
		lx.err(t, "%s: %q", err, t.S())
		return nil
	}

	// https://github.com/golang/go/issues/15997
	if b := t.S(); len(b) != 0 && b[0] == '`' {
		value = strings.Replace(value, "\r", "", -1)
	}
	return stringID(dict.SID(value))
}

func isSurrogate(r rune) bool { return r >= unicodeSurrogateFirst && r <= unicodeSurrogateLast }

func uintptrConstValueFromUint64(ctx *context, n uint64) Value {
	switch {
	case n > math.MaxInt64:
		return newConstValue(newIntConst(0, big.NewInt(0).SetUint64(n), ctx.uintptrType, false))
	default:
		return newConstValue(newIntConst(int64(n), nil, ctx.uintptrType, false))
	}
}
