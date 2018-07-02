// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"bytes"
	"fmt"
	"go/scanner"
	gotoken "go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/cznic/token"
)

type errorList struct {
	limit  int
	list   scanner.ErrorList
	mu     sync.Mutex
	noCols bool
}

func newErrorList(limit int, noCols bool) *errorList {
	if limit == 0 {
		limit = 1
	}
	return &errorList{limit: limit, noCols: noCols}
}

func (l *errorList) forcedAdd(position token.Position, msg string, args ...interface{}) {
	if len(args) != 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	l.mu.Lock()
	l.list.Add(gotoken.Position(position), msg)
	l.mu.Unlock()
}

func (l *errorList) add(position token.Position, msg string, args ...interface{}) { //TODO check all callers, use args
	if len(args) != 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	l.mu.Lock()
	if l.limit == 0 {
		l.mu.Unlock()
		panic(fmt.Errorf("%s: %v", position, msg))
	}

	l.limit--
	l.list.Add(gotoken.Position(position), msg)
	l.mu.Unlock()
}

func (l *errorList) error() error {
	l.mu.Lock()
	if len(l.list) == 0 {
		l.mu.Unlock()
		return nil
	}

	l.mu.Unlock()
	return l
}

func (l *errorList) errorString() string {
	var b bytes.Buffer
	for _, v := range l.list {
		switch {
		case l.noCols:
			pos := v.Pos
			s := pos.Filename
			if pos.IsValid() {
				if s != "" {
					s += ":"
				}
				s += fmt.Sprintf("%d", pos.Line)
			}
			if s == "" {
				s = "-"
			}
			b.WriteString(s)
			b.WriteString(": ")
			b.WriteString(v.Msg)
		default:
			b.WriteString(v.Error())
		}
		b.WriteByte('\n')
	}
	s := b.String()
	return s
}

func (l *errorList) Error() string {
	l.mu.Lock()
	s := l.errorString()
	l.mu.Unlock()
	return s
}

func checkDir(dir string) (bool, error) {
	fi, err := os.Stat(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			return false, err
		}

		return false, nil
	}

	if !fi.IsDir() {
		return false, nil
	}

	m, err := filepath.Glob(filepath.Join(dir, "*.go"))
	return len(m) != 0, err
}

func isExported(nm string) bool {
	if nm == "" {
		return false
	}

	switch c := nm[0]; {
	case c >= 'A' && c <= 'Z':
		return true
	case c < 0x80:
		return false
	default:
		r, _ := utf8.DecodeRuneInString(nm)
		return unicode.IsUpper(r)
	}
}

func mustUnqote(s string) string {
	if s == "" {
		return ""
	}
	t, err := strconv.Unquote(s)
	if err != nil {
		panic(err)
	}

	return t
}

type cstack struct {
	list []checker
	p    *Package
}

func (c *cstack) check(x, from checker) {
	if x == nil || x.state() == checked {
		return
	}

	c.list = append(c.list, from)
	l := c.list
	defer func() {
		c.list = l
		x.checked()
	}()

	if x.state() == checking {
		var pos token.Position
		switch x := x.(type) {
		case Declaration:
			pos = x.Position()
		case Type:
			pos = x.Position()
		default:
			c.p.errorList.add(token.Position{}, fmt.Sprintf("TODO171 %T\n%s", x, c.dump()))
			return
		}

		c.p.errorList.add(pos, "%s", c.dump())
		return
	}

	x.checking()
	x.check(c)
}

func (c *cstack) reset() *cstack {
	c.list = c.list[:0]
	return c
}

func unparen(e Expr) Expr {
	switch x := e.(type) {
	case *PrimaryExprParenExpr:
		return x.unparen()
	default:
		return e
	}
}

func (c *cstack) dump() string {
	type p interface{ Position() token.Position }
	type s interface{ str() string }

	var a []string
	for _, v := range c.list {
		if v == nil {
			continue
		}

		var str string
		var pos token.Position
		if x, ok := v.(p); ok {
			pos = x.Position()
		}
		if x, ok := v.(Declaration); ok {
			str = x.Name()
		} else if x, ok := v.(s); ok {
			str = x.str()
		} else {
			str = fmt.Sprintf("%T", v)
		}
		s := fmt.Sprintf("%s: %s", pos, str)
		if len(a) == 0 || a[len(a)-1] != s {
			a = append(a, s)
		}
	}
	return strings.Join(a, "\n")
}

func initializationLoop(d Declaration, stack []checker) string {
	type p interface{ Position() token.Position }
	type s interface{ str() string }

	//TODO a := []string{"initialization loop:", fmt.Sprintf("%s: %s refers to", d.Position(), d.Name())}
	//TODO for i := len(stack) - 1; i >= 0; i-- {
	//TODO 	if stack[i] == d {
	//TODO 		for _, v := range stack[i+1:] {
	//TODO 			if x, ok := v.(Declaration); ok {
	//TODO 				a = append(a, fmt.Sprintf("%s: %s refers to", x.Position(), x.Name()))
	//TODO 			}
	//TODO 		}
	//TODO 		a = append(a, fmt.Sprintf("%s: %s", d.Position(), d.Name()))
	//TODO 		return strings.Join(a, "\n\t")
	//TODO 	}
	//TODO }

	a := []string{fmt.Sprintf("%s: %s refers to", d.Position(), d.Name())}
	for _, v := range stack {
		if v == nil {
			continue
		}

		var str string
		var pos token.Position
		if x, ok := v.(p); ok {
			pos = x.Position()
		}
		if x, ok := v.(Declaration); ok {
			str = x.Name()
		} else if x, ok := v.(s); ok {
			str = x.str()
		} else {
			str = fmt.Sprintf("%T", v)
		}
		s := fmt.Sprintf("%s: %s", pos, str)
		if len(a) == 0 || a[len(a)-1] != s {
			a = append(a, s)
		}
	}
	return fmt.Sprintf("TODO147\n%s", strings.Join(a, "\n\t"))
}

func skipBlank(e Expr) bool {
	x, ok := unparen(e).(*PrimaryExprIdent)
	return ok && x.Ident.Val == "_"
}
