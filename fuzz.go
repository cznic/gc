// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gofuzz

package gc

import (
	"bytes"
	"fmt"
	"go/scanner"
	gotoken "go/token"
	"strings"

	"github.com/cznic/token"
)

var (
	dotDot = []byte("..")
)

func errString(err error) string {
	var b bytes.Buffer
	scanner.PrintError(&b, err)
	return strings.TrimSpace(b.String())
}

func FuzzLexer(src []byte) int {
	fs2 := gotoken.NewFileSet()
	f := token.NewFile("", len(src))
	f2 := fs2.AddFile("", -1, len(src))
	var ss scanner.Scanner
	l := NewLexer(f, src)
	ss.Init(f2, src, nil, 0)
	for i := 0; ; i++ {
		l.errorCount = 0
		ss.ErrorCount = 0
		off, tok := l.Scan()
		pos := f.Pos(int(off))
		lit := string(l.lit)
		pos2, tok2, lit2 := ss.Scan()
		if tok == gotoken.IDENT {
			for i := 0; i < len(lit); i++ {
				if lit[i] > 0x7f {
					return 0
				}
			}
		}

		if l.errorCount != 0 || tok == gotoken.ILLEGAL || tok == tokenLTLT || tok == tokenGTGT {
			return 0
		}

		if g, e := pos, pos2; g != token.Pos(e) {
			panic(fmt.Errorf(
				"pos[%d] %q(|% x|), %s(%v) %s(%v)",
				i, src, src, f.Position(g), g, f2.Position(e), e,
			))
		}

		if g, e := tok, tok2; g != e {
			if g == tokenBOM && e == gotoken.ILLEGAL {
				return 0
			}

			if g == gotoken.PERIOD && e == gotoken.ILLEGAL && bytes.Contains(src, dotDot) {
				return 0
			}

			panic(fmt.Errorf("tok[%d] %q(|% x|) %s %s", i, src, src, g, e))
		}

		if tok == gotoken.STRING && lit[0] == '`' {
			lit = strings.Replace(lit, "\r", "", -1)
		}
		if lit2 == "" && tok2 != gotoken.EOF {
			lit2 = tok2.String()
		}
		if g, e := lit, lit2; g != e {
			panic(fmt.Errorf("lit[%d] %q(|% x|), %q(|% x|) %q(|% x|)", i, src, src, g, g, e, e))
		}

		if tok == gotoken.EOF {
			return 1
		}
	}
}
