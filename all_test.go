// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/scanner"
	"io"
	"math"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"sync"
	"testing"

	"github.com/cznic/xc"
)

var dbgMu sync.Mutex

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "\tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func callers() []byte { return debug.Stack() }

func dbg(s string, va ...interface{}) {
	dbgMu.Lock()
	defer dbgMu.Unlock()

	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func TODO(...interface{}) string {
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("TODO: %s:%d:\n", path.Base(fn), fl)
}

func use(...interface{}) {}

func init() {
	use(caller, callers, dbg, TODO)
	flag.IntVar(&yyDebug, "yydebug", 0, "")

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("internal error")
	}

	gopaths := filepath.SplitList(os.Getenv("GOPATH"))
	for _, v := range gopaths {
		gp := filepath.Join(v, "src")
		path, err := filepath.Rel(gp, file)
		if err != nil {
			continue
		}

		selfImportPath = filepath.Dir(path)
		return
	}

	panic("internal error")
}

// ============================================================================

var (
	testTags = []string{
		"go1.1",
		"go1.2",
		"go1.3",
		"go1.4",
		"go1.5",
		"go1.6",
		"go1.7",
	}

	oRE            = flag.String("re", "", "regexp")
	selfImportPath string
)

func errStr(err error) string {
	var b bytes.Buffer
	scanner.PrintError(&b, err)
	return b.String()
}

func exampleAST(exampleRule int, src string) interface{} {
	t := &testContext{exampleRule: exampleRule}
	c, err := newContext("", "", "", nil, nil, addTestContext(t), EnableGenerics())
	if err != nil {
		return err
	}

	s := fmt.Sprintf("example%v", exampleRule)
	c.newPackage(s, s).loadString(s+".go", src)
	return t.exampleAST
}

func disableNoBuildableFilesError() Opt {
	return func(c *Context) error {
		c.options.disableNoBuildableFilesError = true
		return nil
	}
}

func addTestContext(t *testContext) Opt {
	return func(c *Context) error {
		c.test = t
		return nil
	}
}

func TestLoad(t *testing.T) {
	var importPaths []string
	root := filepath.Join(runtime.GOROOT(), "src")
	if err := filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				return nil
			}

			if path == root {
				return nil
			}

			if b := filepath.Base(path); b == "testdata" || strings.HasPrefix(b, ".") {
				return filepath.SkipDir
			}

			matches, err := filepath.Glob(filepath.Join(path, "*.go"))
			if err != nil {
				return err
			}

			if len(matches) != 0 {
				importPaths = append(importPaths, path[len(root)+1:])
			}
			return nil

		},
	); err != nil {
		t.Fatal(err)
	}

	importPaths = append(importPaths, selfImportPath)
	sort.Strings(importPaths)
	m, err := LoadPackages(
		runtime.GOOS,
		runtime.GOARCH,
		runtime.GOROOT(),
		filepath.SplitList(os.Getenv("GOPATH")),
		testTags,
		importPaths,
		disableNoBuildableFilesError(),
	)
	t.Log(len(m))
	if err != nil {
		t.Fatal(errStr(err))
	}
}

func Test(t *testing.T) {
	f, err := os.Create("test.log")
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	logw := bufio.NewWriter(f)

	defer logw.Flush()

	c, err := NewContext(
		runtime.GOOS,
		runtime.GOARCH,
		runtime.GOROOT(),
		filepath.SplitList(os.Getenv("GOPATH")),
		testTags,
		EnableGenerics(),
		ErrLimit(math.MaxInt32),
		addTestContext(&testContext{}),
		disableNoBuildableFilesError(),
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if strings.HasSuffix(path, ".dir") {
				return filepath.SkipDir
			}

			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		defer f.Close()

		c.test.errChecks = c.test.errChecks[:0]
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			line := strings.TrimSpace(sc.Text())
			switch {
			case strings.HasPrefix(line, "//"):
				switch line = strings.TrimSpace(line[2:]); {
				case line == "skip":
					return nil
				case strings.HasPrefix(line, "+build"):
					// nop
				case
					line == "build",
					line == "cmpout",
					line == "compile",
					line == "compiledir",
					line == "errorcheckoutput",
					line == "run",
					line == "rundir",
					line == "runoutput",
					line == "true",
					strings.HasPrefix(line, "$"),
					strings.HasPrefix(line, "errorcheckoutput "),
					strings.HasPrefix(line, "run "),
					strings.HasPrefix(line, "runoutput "):

					// N/A for a front end.
				case
					strings.HasPrefix(line, "errorcheck "):
					fmt.Fprintf(logw, "[TODO %q] %s\n", line, path)
					return nil
				case line == "errorcheck":
					errorcheck(t, c, f.Name(), logw)
				case line == "errorcheckdir":
					//TODO errorcheckDir(t, c, f.Name(), logw)
				}
			case line == "":
				return nil
			}
		}
		return sc.Err()
	}); err != nil {
		t.Fatal(err)
	}
}

func errorcheck(t *testing.T, c *Context, fname string, logw io.Writer) {
	if re := *oRE; re != "" {
		ok, err := regexp.MatchString(re, fname)
		if err != nil {
			t.Fatal(err)
		}

		if !ok {
			return
		}

		t.Log(fname)
	}

	ip := filepath.Join(selfImportPath, filepath.Dir(fname))
	_, err := c.loadPackage(ip, []string{fname})
	errorCheckResults(t, c.test.errChecks, err, fname, logw)
}

func qmsg(s string) string {
	return strings.Replace(strings.Replace(s, "\t", "\\t", -1), "\n", "\\n", -1)
}

var errCheckPatterns = regexp.MustCompile(`"([^"]*)"`)

func errorCheckResults(t *testing.T, checks []xc.Token, err error, fname string, logw io.Writer) {
	if *oRE != "" {
		for _, v := range checks {
			t.Log(PrettyString(v))
		}
		if err != nil {
			t.Logf("FAIL\n%s", errStr(err))
		}
	}
	m := map[*scanner.Error]struct{}{}
	if err != nil {
		err := err.(scanner.ErrorList)
		err.RemoveMultiples()
		for _, v := range err {
			p := filepath.ToSlash(v.Pos.Filename)
			if !filepath.IsAbs(p) {
				m[v] = struct{}{}
			}
		}
	}
	n := map[xc.Token]struct{}{}
	for _, v := range checks {
		n[v] = struct{}{}
	}
	var a scanner.ErrorList
	fail := false
	for k := range m {
		line := k.Pos.Line
		for l := range n {
			line2 := position(l.Pos()).Line
			if line2 != line {
				continue
			}

			matched := false
			for _, v := range errCheckPatterns.FindAllSubmatch(l.S(), -1) {
				re := v[1]
				ok, err := regexp.MatchString(string(re), k.Error())
				if err != nil {
					t.Fatal(err)
				}

				if ok {
					matched = true
					break
				}
			}
			switch {
			case matched:
				a = append(a, &scanner.Error{Pos: k.Pos, Msg: fmt.Sprintf("[PASS errorcheck] %s: %s", l.S(), qmsg(k.Msg))})
			default:
				fail = true
				a = append(a, &scanner.Error{Pos: k.Pos, Msg: fmt.Sprintf("[FAIL errorcheck: error does not match] %s: %s", l.S(), k.Msg)})
			}

			delete(m, k)
			delete(n, l)
		}
	}
	if !fail && len(m) == 0 && len(n) == 0 {
		t.Logf("[PASS errorcheck] %v\n", fname)
	}
	for k := range m {
		a = append(a, &scanner.Error{Pos: k.Pos, Msg: fmt.Sprintf("[FAIL errorcheck: extra error] %q", k.Msg)})
	}
	for l := range n {
		a = append(a, &scanner.Error{Pos: position(l.Pos()), Msg: fmt.Sprintf("[FAIL errorcheck: missing error] %s", l.S())})
	}
	a.Sort()
	for _, v := range a {
		fmt.Fprintf(logw, "%s %s\n", v.Msg, v.Pos)
	}
}
