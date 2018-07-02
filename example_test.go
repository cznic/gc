// Copyright 2018 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"strings"

	"github.com/edsrzf/mmap-go"
)

func example(t interface{}, src string) {
	ts := fmt.Sprintf("%T", t)
	if strings.HasPrefix(ts, "*") {
		ts = "&" + ts[1:]
	}
	l := newLexer(nil, nil)
	y := newParser(nil, nil)
	src = fmt.Sprintf("package p\n%s\n", src)
	b := mmap.MMap(src)
	c, err := newTestContext()
	c.tweaks.example = true
	if err != nil {
		panic(err)
	}

	p := newPackage(c, "", "", nil)
	sf := newSourceFile(p, "", nil, b)
	l.init(sf.File, b)
	y.init(sf, l)
	y.file()
	a := []string{pretty2(sf.TopLevelDecls)}
	for _, v := range sf.TopLevelDecls {
		a = append(a, pretty2(v.Type()))
	}
	s := strings.Join(a, "\n")
	if *oTrc {
		fmt.Printf("==== %s\n%s\n", ts, s)
	}
	i := strings.Index(s, ts)
	if i < 0 {
		fmt.Printf("TODO %s\n", ts) //TODOOK
		return
	}

	for i > 0 && s[i-1] != '\n' {
		i--
	}
	var np int
	s = s[i:]
out:
	for k, v := range s {
		switch v {
		case ' ', '·':
			// ok
		default:
			np = k
			break out
		}
	}
	a = strings.Split(s, "\n")
next:
	for j := 1; j < len(a); j++ {
		for k, v := range a[j] {
			switch v {
			case ' ', '·':
				// ok
			default:
				if k == np {
					a = a[:j+1]
					for i, v := range a {
						a[i] = v[np:]
					}
					fmt.Println(strings.Join(a, "\n"))
					return
				}

				continue next
			}
		}
	}
	fmt.Println(s)
}

func ExampleCompLitExpr_expr() {
	example(&CompLitExprExpr{}, "var v = []int{42}")
	// Output:
	// Expr: &gc.CompLitExprExpr{
	// · Expr: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:15: "42",
	// · },
	// },
}

func ExampleCompLitExpr_keyValList() {
	example(&CompLitExprKeyValList{}, "var v = [][]int{{1: 42, 3: 314}}")
	// Output:
	// Expr: &gc.CompLitExprKeyValList{
	// · List: []*gc.KeyValKeyedExpr{ // len 2
	// · · 0: &gc.KeyValKeyedExpr{
	// · · · Key: &gc.CompLitExprExpr{
	// · · · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · · · Literal: 2:18: "1",
	// · · · · },
	// · · · },
	// · · · Expr: &gc.CompLitExprExpr{
	// · · · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · · · Literal: 2:21: "42",
	// · · · · },
	// · · · },
	// · · },
	// · · 1: &gc.KeyValKeyedExpr{
	// · · · Key: &gc.CompLitExprExpr{
	// · · · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · · · Literal: 2:25: "3",
	// · · · · },
	// · · · },
	// · · · Expr: &gc.CompLitExprExpr{
	// · · · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · · · Literal: 2:28: "314",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleKeyVal_expr() {
	example(&KeyValExpr{}, "var v = []int{42}")
	// Output:
	// 0: &gc.KeyValExpr{
	// · Expr: &gc.CompLitExprExpr{
	// · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · Literal: 2:15: "42",
	// · · },
	// · },
	// },
}

func ExampleKeyVal_keyedExpr() {
	example(&KeyValKeyedExpr{}, "var v = []int{1: 42}")
	// Output:
	// 0: &gc.KeyValKeyedExpr{
	// · Key: &gc.CompLitExprExpr{
	// · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · Literal: 2:15: "1",
	// · · },
	// · },
	// · Expr: &gc.CompLitExprExpr{
	// · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · Literal: 2:18: "42",
	// · · },
	// · },
	// },
}

func ExampleExprOrType_expr() {
	example(&ExprOrTypeExpr{}, "var v = foo(42)")
	// Output:
	// 0: &gc.ExprOrTypeExpr{
	// · Expr: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:13: "42",
	// · },
	// },
}

func ExampleExprOrType_type() {
	example(&ExprOrTypeType{}, "var v = make([]int, 42)")
	// Output:
	// 0: &gc.ExprOrTypeType{
	// },
}

func ExamplePrimaryExpr_parenExpr() {
	example(&PrimaryExprParenExpr{}, "var v = (42)")
	// Output:
	// Initializer: &gc.PrimaryExprParenExpr{
	// · Expr: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:10: "42",
	// · },
	// },
}

func ExamplePrimaryExpr_ident() {
	example(&PrimaryExprIdent{}, "var v = foo")
	// Output:
	// Initializer: &gc.PrimaryExprIdent{
	// · Ident: 2:9: "foo",
	// },
}

func ExamplePrimaryExpr_funcLiteral() {
	example(&PrimaryExprFuncLiteral{}, "var v = func(){}")
	// Output:
	// Initializer: &gc.PrimaryExprFuncLiteral{
	// · Body: &gc.StmtBlock{
	// · },
	// },
}

func ExamplePrimaryExpr_conversion() {
	example(&PrimaryExprConversion{}, "var v = ([]byte)(nil)")
	// Output:
	// Initializer: &gc.PrimaryExprConversion{
	// · Expr: &gc.PrimaryExprIdent{
	// · · Ident: 2:18: "nil",
	// · },
	// },
}

func ExamplePrimaryExpr_intLiteral() {
	example(&PrimaryExprIntLiteral{}, "var v = 42")
	// Output:
	// Initializer: &gc.PrimaryExprIntLiteral{
	// · Literal: 2:9: "42",
	// },
}

func ExamplePrimaryExpr_floatLiteral() {
	example(&PrimaryExprFloatLiteral{}, "var v = 42.314")
	// Output:
	// Initializer: &gc.PrimaryExprFloatLiteral{
	// · Literal: 2:9: "42.314",
	// },
}

func ExamplePrimaryExpr_imagLiteral() {
	example(&PrimaryExprImagLiteral{}, "var v = 42i")
	// Output:
	// Initializer: &gc.PrimaryExprImagLiteral{
	// · Literal: 2:9: "42i",
	// },
}

func ExamplePrimaryExpr_runeLiteral() {
	example(&PrimaryExprRuneLiteral{}, "var v = 'a'")
	// Output:
	// Initializer: &gc.PrimaryExprRuneLiteral{
	// · Literal: 2:9: "'a'",
	// },
}

func ExamplePrimaryExpr_stringLiteral() {
	example(&PrimaryExprStringLiteral{}, `var v = "foo"`)
	// Output:
	// Initializer: &gc.PrimaryExprStringLiteral{
	// · Literal: 2:9: "\"foo\"",
	// },
}

func ExamplePrimaryExpr_compositeLiteral() {
	example(&PrimaryExprCompositeLiteral{}, "var v = []int{42}")
	// Output:
	// Initializer: &gc.PrimaryExprCompositeLiteral{
	// · List: []*gc.KeyValExpr{ // len 1
	// · · 0: &gc.KeyValExpr{
	// · · · Expr: &gc.CompLitExprExpr{
	// · · · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · · · Literal: 2:15: "42",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExamplePrimaryExpr_call() {
	example(&PrimaryExprCall{}, "var v = foo(42)")
	// Output:
	// Initializer: &gc.PrimaryExprCall{
	// · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "foo",
	// · },
	// · Args: []*gc.ExprOrTypeExpr{ // len 1
	// · · 0: &gc.ExprOrTypeExpr{
	// · · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · · Literal: 2:13: "42",
	// · · · },
	// · · },
	// · },
	// },
}

func ExamplePrimaryExpr_typeSwitchGuard() {
	example(&PrimaryExprTypeSwitchGuard{}, `
func foo(x interface{}) {
	switch x.(type) {
	case []int:
		foo()
	case [42]string:
		bar()
	default:
		qux()
	}
}	
`)
	// Output:
	// Expr: &gc.PrimaryExprTypeSwitchGuard{
	// · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · Ident: 4:9: "x",
	// · },
	// },
}

func ExamplePrimaryExpr_typeAssertion() {
	example(&PrimaryExprTypeAssertion{}, "var v, ok = foo.(int)")
	// Output:
	// Initializer: &gc.PrimaryExprTypeAssertion{
	// · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · Ident: 2:13: "foo",
	// · },
	// · Assert: &gc.ExprOrTypeExpr{
	// · · Expr: &gc.PrimaryExprIdent{
	// · · · Ident: 2:18: "int",
	// · · },
	// · },
	// },
}

func ExamplePrimaryExpr_selector() {
	example(&PrimaryExprSelector{}, "var v = foo().bar")
	// Output:
	// Initializer: &gc.PrimaryExprSelector{
	// · PrimaryExpr: &gc.PrimaryExprCall{
	// · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · Ident: 2:9: "foo",
	// · · },
	// · },
	// · Selector: 2:15: "bar",
	// },
}

func ExamplePrimaryExpr_index() {
	example(&PrimaryExprIndex{}, "var v = foo[42]")
	// Output:
	// Initializer: &gc.PrimaryExprIndex{
	// · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "foo",
	// · },
	// · Index: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:13: "42",
	// · },
	// },
}

func ExamplePrimaryExpr_simpleSlice() {
	example(&PrimaryExprSimpleSlice{}, "var v = foo[24:42]")
	// Output:
	// Initializer: &gc.PrimaryExprSimpleSlice{
	// · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "foo",
	// · },
	// · Lo: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:13: "24",
	// · },
	// · Hi: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:16: "42",
	// · },
	// },
}

func ExamplePrimaryExpr_fullSlice() {
	example(&PrimaryExprFullSlice{}, "var v = foo[24:42:314]")
	// Output:
	// Initializer: &gc.PrimaryExprFullSlice{
	// · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "foo",
	// · },
	// · Lo: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:13: "24",
	// · },
	// · Hi: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:16: "42",
	// · },
	// · Max: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:19: "314",
	// · },
	// },
}

func ExampleUnaryExpr_receive() {
	example(&UnaryExprReceive{}, "var v = <-foo")
	// Output:
	// Initializer: &gc.UnaryExprReceive{
	// · Expr: &gc.PrimaryExprIdent{
	// · · Ident: 2:11: "foo",
	// · },
	// },
}

func ExampleUnaryExpr_not() {
	example(&UnaryExprNot{}, "var v = !foo")
	// Output:
	// Initializer: &gc.UnaryExprNot{
	// · Expr: &gc.PrimaryExprIdent{
	// · · Ident: 2:10: "foo",
	// · },
	// },
}

func ExampleUnaryExpr_addr() {
	example(&UnaryExprAddr{}, "var v = &foo")
	// Output:
	// Initializer: &gc.UnaryExprAddr{
	// · Expr: &gc.PrimaryExprIdent{
	// · · Ident: 2:10: "foo",
	// · },
	// },
}

func ExampleUnaryExpr_deref() {
	example(&UnaryExprDeref{}, "var v = *foo")
	// Output:
	// Initializer: &gc.UnaryExprDeref{
	// · Expr: &gc.PrimaryExprIdent{
	// · · Ident: 2:10: "foo",
	// · },
	// },
}

func ExampleUnaryExpr_pos() {
	example(&UnaryExprPos{}, "var v = +0")
	// Output:
	// Initializer: &gc.UnaryExprPos{
	// · Expr: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:10: "0",
	// · },
	// },
}

func ExampleUnaryExpr_neg() {
	example(&UnaryExprNeg{}, "var v = -0")
	// Output:
	// Initializer: &gc.UnaryExprNeg{
	// · Expr: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:10: "0",
	// · },
	// },
}

func ExampleUnaryExpr_xor() {
	example(&UnaryExprXor{}, "var v = ^0")
	// Output:
	// Initializer: &gc.UnaryExprXor{
	// · Expr: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:10: "0",
	// · },
	// },
}

func ExampleExpr_gTR() {
	example(&ExprGTR{}, "var v = > b")
	// Output:
	// Initializer: &gc.ExprGTR{
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:11: "b",
	// · },
	// },
}

func ExampleExpr_nEQ() {
	example(&ExprNEQ{}, "var v = a != b")
	// Output:
	// Initializer: &gc.ExprNEQ{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "b",
	// · },
	// },
}

func ExampleExpr_lAND() {
	example(&ExprLAND{}, "var v = a && b")
	// Output:
	// Initializer: &gc.ExprLAND{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "b",
	// · },
	// },
}

func ExampleExpr_aNDNOT() {
	example(&ExprANDNOT{}, "var v = a &^ b")
	// Output:
	// Initializer: &gc.ExprANDNOT{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "b",
	// · },
	// },
}

func ExampleExpr_sHL() {
	example(&ExprSHL{}, "var v = a << b")
	// Output:
	// Initializer: &gc.ExprSHL{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "b",
	// · },
	// },
}

func ExampleExpr_lEQ() {
	example(&ExprLEQ{}, "var v = a <= b")
	// Output:
	// Initializer: &gc.ExprLEQ{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "b",
	// · },
	// },
}

func ExampleExpr_eQL() {
	example(&ExprEQL{}, "var v = a == b")
	// Output:
	// Initializer: &gc.ExprEQL{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "b",
	// · },
	// },
}

func ExampleExpr_gEQ() {
	example(&ExprGEQ{}, "var v = a >= b")
	// Output:
	// Initializer: &gc.ExprGEQ{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "b",
	// · },
	// },
}

func ExampleExpr_sHR() {
	example(&ExprSHR{}, "var v = a >> b")
	// Output:
	// Initializer: &gc.ExprSHR{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "b",
	// · },
	// },
}

func ExampleExpr_lOR() {
	example(&ExprLOR{}, "var v = a || b")
	// Output:
	// Initializer: &gc.ExprLOR{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "b",
	// · },
	// },
}

func ExampleExpr_rEM() {
	example(&ExprREM{}, "var v = a % b")
	// Output:
	// Initializer: &gc.ExprREM{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:13: "b",
	// · },
	// },
}

func ExampleExpr_aND() {
	example(&ExprAND{}, "var v = a & b")
	// Output:
	// Initializer: &gc.ExprAND{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:13: "b",
	// · },
	// },
}

func ExampleExpr_mUL() {
	example(&ExprMUL{}, "var v = a * b")
	// Output:
	// Initializer: &gc.ExprMUL{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:13: "b",
	// · },
	// },
}

func ExampleExpr_aDD() {
	example(&ExprADD{}, "var v = a + b")
	// Output:
	// Initializer: &gc.ExprADD{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:13: "b",
	// · },
	// },
}

func ExampleExpr_sUB() {
	example(&ExprSUB{}, "var v = a - b")
	// Output:
	// Initializer: &gc.ExprSUB{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:13: "b",
	// · },
	// },
}

func ExampleExpr_qUO() {
	example(&ExprQUO{}, "var v = a / b")
	// Output:
	// Initializer: &gc.ExprQUO{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:13: "b",
	// · },
	// },
}

func ExampleExpr_lSS() {
	example(&ExprLSS{}, "var v = a < b")
	// Output:
	// Initializer: &gc.ExprLSS{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:13: "b",
	// · },
	// },
}

func ExampleExpr_xOR() {
	example(&ExprXOR{}, "var v = a ^ b")
	// Output:
	// Initializer: &gc.ExprXOR{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:13: "b",
	// · },
	// },
}

func ExampleExpr_oR() {
	example(&ExprOR{}, "var v = a | b")
	// Output:
	// Initializer: &gc.ExprOR{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:9: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:13: "b",
	// · },
	// },
}

func ExampleParamType_identType() {
	example(&ParamTypeIdentType{}, "func foo(a int)")
	// Output:
	// 0: &gc.ParamTypeIdentType{
	// · Ident: 2:10: "a",
	// · Type: &gc.NamedType{
	// · · Name: "int",
	// · },
	// },
}

func ExampleParamType_type() {
	example(&ParamTypeType{}, "func foo(int)")
	// Output:
	// 0: &gc.ParamTypeType{
	// · Type: &gc.NamedType{
	// · · Name: "int",
	// · },
	// },
}

func ExampleParamTypeList_type() {
	example(ParamTypeListType{}, "func foo(int, string)")
	// Output:
	// ParameterList: &gc.ParamTypeListType{
	// · List: []*gc.ParamTypeType{ // len 2
	// · · 0: &gc.ParamTypeType{
	// · · · Type: &gc.NamedType{
	// · · · · Name: "int",
	// · · · },
	// · · },
	// · · 1: &gc.ParamTypeType{
	// · · · Type: &gc.NamedType{
	// · · · · Name: "string",
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleParamTypeList_identType() {
	example(ParamTypeListIdentType{}, "func foo(i int, s string)")
	// Output:
	// ParameterList: &gc.ParamTypeListIdentType{
	// · List: []*gc.ParamTypeIdentType{ // len 2
	// · · 0: &gc.ParamTypeIdentType{
	// · · · Ident: 2:10: "i",
	// · · · Type: &gc.NamedType{
	// · · · · Name: "int",
	// · · · },
	// · · },
	// · · 1: &gc.ParamTypeIdentType{
	// · · · Ident: 2:17: "s",
	// · · · Type: &gc.NamedType{
	// · · · · Name: "string",
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleSimpleStmt_decl() {
	example(&SimpleStmtDecl{}, "func foo() { var a, b int }")
	// Output:
	// 0: &gc.SimpleStmtDecl{
	// · List: []*gc.VarDecl{ // len 2
	// · · 0: &gc.VarDecl{
	// · · },
	// · · 1: &gc.VarDecl{
	// · · },
	// · },
	// },
}

func ExampleSimpleStmt_assignment() {
	example(&SimpleStmtAssignment{}, "func foo() { a, b = c, d }")
	// Output:
	// 0: &gc.SimpleStmtAssignment{
	// · LHS: []*gc.PrimaryExprIdent{ // len 2
	// · · 0: &gc.PrimaryExprIdent{
	// · · · Ident: 2:14: "a",
	// · · },
	// · · 1: &gc.PrimaryExprIdent{
	// · · · Ident: 2:17: "b",
	// · · },
	// · },
	// · RHS: []*gc.PrimaryExprIdent{ // len 2
	// · · 0: &gc.PrimaryExprIdent{
	// · · · Ident: 2:21: "c",
	// · · },
	// · · 1: &gc.PrimaryExprIdent{
	// · · · Ident: 2:24: "d",
	// · · },
	// · },
	// },
}

func ExampleSimpleStmt_expr() {
	example(&SimpleStmtExpr{}, "func foo() { bar() }")
	// Output:
	// 0: &gc.SimpleStmtExpr{
	// · Expr: &gc.PrimaryExprCall{
	// · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · Ident: 2:14: "bar",
	// · · },
	// · },
	// },
}

func ExampleSimpleStmt_remAssign() {
	example(&SimpleStmtRemAssign{}, "func foo() { a %= b }")
	// Output:
	// 0: &gc.SimpleStmtRemAssign{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:19: "b",
	// · },
	// },
}

func ExampleSimpleStmt_andAssign() {
	example(&SimpleStmtAndAssign{}, "func foo() { a &= b }")
	// Output:
	// 0: &gc.SimpleStmtAndAssign{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:19: "b",
	// · },
	// },
}

func ExampleSimpleStmt_andNotAssign() {
	example(&SimpleStmtAndNotAssign{}, "func foo() { a &^= b }")
	// Output:
	// 0: &gc.SimpleStmtAndNotAssign{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:20: "b",
	// · },
	// },
}

func ExampleSimpleStmt_mulAssign() {
	example(&SimpleStmtMulAssign{}, "func foo() { a *= b }")
	// Output:
	// 0: &gc.SimpleStmtMulAssign{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:19: "b",
	// · },
	// },
}

func ExampleSimpleStmt_inc() {
	example(&SimpleStmtInc{}, "func foo() { a++ }")
	// Output:
	// 0: &gc.SimpleStmtInc{
	// · Expr: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// },
}

func ExampleSimpleStmt_addAssign() {
	example(&SimpleStmtAddAssign{}, "func foo() { a += b }")
	// Output:
	// 0: &gc.SimpleStmtAddAssign{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:19: "b",
	// · },
	// },
}

func ExampleSimpleStmt_dec() {
	example(&SimpleStmtDec{}, "func foo() { a-- }")
	// Output:
	// 0: &gc.SimpleStmtDec{
	// · Expr: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// },
}

func ExampleSimpleStmt_subAssign() {
	example(&SimpleStmtSubAssign{}, "func foo() { a -= b }")
	// Output:
	// 0: &gc.SimpleStmtSubAssign{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:19: "b",
	// · },
	// },
}

func ExampleSimpleStmt_quoAssign() {
	example(&SimpleStmtQuoAssign{}, "func foo() { a /= b }")
	// Output:
	// 0: &gc.SimpleStmtQuoAssign{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:19: "b",
	// · },
	// },
}

func ExampleSimpleStmt_shlAssign() {
	example(&SimpleStmtShlAssign{}, "func foo() { a <<= b }")
	// Output:
	// 0: &gc.SimpleStmtShlAssign{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:20: "b",
	// · },
	// },
}

func ExampleSimpleStmt_shrAssign() {
	example(&SimpleStmtShrAssign{}, "func foo() { a >>= b }")
	// Output:
	// 0: &gc.SimpleStmtShrAssign{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:20: "b",
	// · },
	// },
}

func ExampleSimpleStmt_xorAssign() {
	example(&SimpleStmtXorAssign{}, "func foo() { a ^= b }")
	// Output:
	// 0: &gc.SimpleStmtXorAssign{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:19: "b",
	// · },
	// },
}

func ExampleSimpleStmt_orAssign() {
	example(&SimpleStmtOrAssign{}, "func foo() { a |= b }")
	// Output:
	// 0: &gc.SimpleStmtOrAssign{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:14: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:19: "b",
	// · },
	// },
}

func ExampleSimpleStmt_define() {
	example(&SimpleStmtDefine{}, "func foo() { a, b := c, d }")
	// Output:
	// 0: &gc.SimpleStmtDefine{
	// · LHS: []*gc.PrimaryExprIdent{ // len 2
	// · · 0: &gc.PrimaryExprIdent{
	// · · · Ident: 2:14: "a",
	// · · },
	// · · 1: &gc.PrimaryExprIdent{
	// · · · Ident: 2:17: "b",
	// · · },
	// · },
	// · RHS: []*gc.PrimaryExprIdent{ // len 2
	// · · 0: &gc.PrimaryExprIdent{
	// · · · Ident: 2:22: "c",
	// · · },
	// · · 1: &gc.PrimaryExprIdent{
	// · · · Ident: 2:25: "d",
	// · · },
	// · },
	// },
}

func ExampleSimpleStmt_assingnment() {
	example(&SimpleStmtAssignment{}, "func foo() { a, b = c, d }")
	// Output:
	// 0: &gc.SimpleStmtAssignment{
	// · LHS: []*gc.PrimaryExprIdent{ // len 2
	// · · 0: &gc.PrimaryExprIdent{
	// · · · Ident: 2:14: "a",
	// · · },
	// · · 1: &gc.PrimaryExprIdent{
	// · · · Ident: 2:17: "b",
	// · · },
	// · },
	// · RHS: []*gc.PrimaryExprIdent{ // len 2
	// · · 0: &gc.PrimaryExprIdent{
	// · · · Ident: 2:21: "c",
	// · · },
	// · · 1: &gc.PrimaryExprIdent{
	// · · · Ident: 2:24: "d",
	// · · },
	// · },
	// },
}

func ExampleSimpleStmt_send() {
	example(&SimpleStmtSend{}, "func f() { a <- b }")
	// Output:
	// 0: &gc.SimpleStmtSend{
	// · LHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:12: "a",
	// · },
	// · RHS: &gc.PrimaryExprIdent{
	// · · Ident: 2:17: "b",
	// · },
	// },
}

func ExampleStmt_break() {
	example(&StmtBreak{}, `
func foo() {
out:
	for range expr {
		break out
	}
}
`)
	// Output:
	// 0: &gc.StmtBreak{
	// · Label: 6:9: "out",
	// },
}

func ExampleStmt_continue() {
	example(&StmtContinue{}, `
func foo() {
next:
	for range expr {
		continue next
	}
}
`)
	// Output:
	// 0: &gc.StmtContinue{
	// · Label: 6:12: "next",
	// },
}

func ExampleStmt_labeled() {
	example(&StmtLabeled{}, "func foo() { label: bar() }")
	// Output:
	// 0: &gc.StmtLabeled{
	// · Label: 2:14: "label",
	// · Stmt: &gc.SimpleStmtExpr{
	// · · Expr: &gc.PrimaryExprCall{
	// · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · Ident: 2:21: "bar",
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleStmt_defer() {
	example(&StmtDefer{}, "func foo() { defer bar() }")
	// Output:
	// 0: &gc.StmtDefer{
	// · Expr: &gc.PrimaryExprCall{
	// · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · Ident: 2:20: "bar",
	// · · },
	// · },
	// },
}

func ExampleStmt_fallthrough() {
	example(StmtFallthrough{}, `
func foo() {
	switch x {
	case 1:
		fallthrough
	default:
	}
}
`)
	// Output:
	// Body: []*gc.StmtFallthrough{ // len 1
	// },
}

func ExampleStmt_forRange() {
	example(&StmtForRange{}, `
func foo() {
	for range expr {
		bar()
	}
}
`)
	// Output:
	// 0: &gc.StmtForRange{
	// · Expr: &gc.PrimaryExprIdent{
	// · · Ident: 4:12: "expr",
	// · },
	// · Body: &gc.StmtBlock{
	// · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 5:3: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleStmt_forRangeAssign() {
	example(&StmtForRange{}, `
func foo() {
	for i, v = range expr {
		bar()
	}
}
`)
	// Output:
	// 0: &gc.StmtForRange{
	// · ExprList: []*gc.PrimaryExprIdent{ // len 2
	// · · 0: &gc.PrimaryExprIdent{
	// · · · Ident: 4:6: "i",
	// · · },
	// · · 1: &gc.PrimaryExprIdent{
	// · · · Ident: 4:9: "v",
	// · · },
	// · },
	// · Expr: &gc.PrimaryExprIdent{
	// · · Ident: 4:19: "expr",
	// · },
	// · Body: &gc.StmtBlock{
	// · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 5:3: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleStmt_forRangeDefine() {
	example(&StmtForRange{}, `
func foo() {
	for i, v := range expr {
		bar()
	}
}
`)
	// Output:
	// 0: &gc.StmtForRange{
	// · ExprList: []*gc.PrimaryExprIdent{ // len 2
	// · · 0: &gc.PrimaryExprIdent{
	// · · · Ident: 4:6: "i",
	// · · },
	// · · 1: &gc.PrimaryExprIdent{
	// · · · Ident: 4:9: "v",
	// · · },
	// · },
	// · Expr: &gc.PrimaryExprIdent{
	// · · Ident: 4:20: "expr",
	// · },
	// · Body: &gc.StmtBlock{
	// · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 5:3: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Define: true,
	// },
}

func ExampleStmt_for() {
	example(&StmtFor{}, `
func foo() {
	for i := 0; i < 10; i++ {
		bar()
	}
}
`)
	// Output:
	// 0: &gc.StmtFor{
	// · Stmt: &gc.SimpleStmtDefine{
	// · · LHS: []*gc.PrimaryExprIdent{ // len 1
	// · · · 0: &gc.PrimaryExprIdent{
	// · · · · Ident: 4:6: "i",
	// · · · },
	// · · },
	// · · RHS: []*gc.PrimaryExprIntLiteral{ // len 1
	// · · · 0: &gc.PrimaryExprIntLiteral{
	// · · · · Literal: 4:11: "0",
	// · · · },
	// · · },
	// · },
	// · Stmt2: &gc.SimpleStmtExpr{
	// · · Expr: &gc.ExprLSS{
	// · · · LHS: &gc.PrimaryExprIdent{
	// · · · · Ident: 4:14: "i",
	// · · · },
	// · · · RHS: &gc.PrimaryExprIntLiteral{
	// · · · · Literal: 4:18: "10",
	// · · · },
	// · · },
	// · },
	// · Stmt3: &gc.SimpleStmtInc{
	// · · Expr: &gc.PrimaryExprIdent{
	// · · · Ident: 4:22: "i",
	// · · },
	// · },
	// · Body: &gc.StmtBlock{
	// · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 5:3: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleStmt_go() {
	example(&StmtGo{}, "func foo() { go bar() }")
	// Output:
	// 0: &gc.StmtGo{
	// · Expr: &gc.PrimaryExprCall{
	// · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · Ident: 2:17: "bar",
	// · · },
	// · },
	// },
}

func ExampleStmt_goto() {
	example(&StmtGoto{}, "func foo() { label: goto label }")
	// Output:
	// Stmt: &gc.StmtGoto{
	// · Ident: 2:26: "label",
	// },
}

func ExampleStmt_if() {
	example(&StmtIf{}, `
func foo() {
	if x {
		y()
	}
}	
`)
	// Output:
	// 0: &gc.StmtIf{
	// · IfHeader: &gc.IfHeader1{
	// · · Stmt: &gc.SimpleStmtExpr{
	// · · · Expr: &gc.PrimaryExprIdent{
	// · · · · Ident: 4:5: "x",
	// · · · },
	// · · },
	// · },
	// · Body: &gc.StmtBlock{
	// · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 5:3: "y",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleStmt_ifAssign() {
	example(&StmtIf{}, `
func foo() {
	if x, y = bar(); x {
		y()
	}
}	
`)
	// Output:
	// 0: &gc.StmtIf{
	// · IfHeader: &gc.IfHeader2{
	// · · Stmt: &gc.SimpleStmtAssignment{
	// · · · LHS: []*gc.PrimaryExprIdent{ // len 2
	// · · · · 0: &gc.PrimaryExprIdent{
	// · · · · · Ident: 4:5: "x",
	// · · · · },
	// · · · · 1: &gc.PrimaryExprIdent{
	// · · · · · Ident: 4:8: "y",
	// · · · · },
	// · · · },
	// · · · RHS: []*gc.PrimaryExprCall{ // len 1
	// · · · · 0: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 4:12: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · Stmt2: &gc.SimpleStmtExpr{
	// · · · Expr: &gc.PrimaryExprIdent{
	// · · · · Ident: 4:19: "x",
	// · · · },
	// · · },
	// · },
	// · Body: &gc.StmtBlock{
	// · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 5:3: "y",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleStmt_ifDefine() {
	example(&StmtIf{}, `
func foo() {
	if x, y := bar(); x {
		y()
	}
}	
`)
	// Output:
	// 0: &gc.StmtIf{
	// · IfHeader: &gc.IfHeader2{
	// · · Stmt: &gc.SimpleStmtDefine{
	// · · · LHS: []*gc.PrimaryExprIdent{ // len 2
	// · · · · 0: &gc.PrimaryExprIdent{
	// · · · · · Ident: 4:5: "x",
	// · · · · },
	// · · · · 1: &gc.PrimaryExprIdent{
	// · · · · · Ident: 4:8: "y",
	// · · · · },
	// · · · },
	// · · · RHS: []*gc.PrimaryExprCall{ // len 1
	// · · · · 0: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 4:13: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · Stmt2: &gc.SimpleStmtExpr{
	// · · · Expr: &gc.PrimaryExprIdent{
	// · · · · Ident: 4:20: "x",
	// · · · },
	// · · },
	// · },
	// · Body: &gc.StmtBlock{
	// · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 5:3: "y",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleStmt_ifElseIf() {
	example(&StmtIf{}, `
func foo() {
	if x, y := bar(); x {
		y()
	} else if z {
		w()
	}
}	
`)
	// Output:
	// 0: &gc.StmtIf{
	// · IfHeader: &gc.IfHeader2{
	// · · Stmt: &gc.SimpleStmtDefine{
	// · · · LHS: []*gc.PrimaryExprIdent{ // len 2
	// · · · · 0: &gc.PrimaryExprIdent{
	// · · · · · Ident: 4:5: "x",
	// · · · · },
	// · · · · 1: &gc.PrimaryExprIdent{
	// · · · · · Ident: 4:8: "y",
	// · · · · },
	// · · · },
	// · · · RHS: []*gc.PrimaryExprCall{ // len 1
	// · · · · 0: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 4:13: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · Stmt2: &gc.SimpleStmtExpr{
	// · · · Expr: &gc.PrimaryExprIdent{
	// · · · · Ident: 4:20: "x",
	// · · · },
	// · · },
	// · },
	// · Body: &gc.StmtBlock{
	// · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 5:3: "y",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · ElseIf: []gc.ElseIf{ // len 1
	// · · 0: gc.ElseIf{
	// · · · IfHeader: &gc.IfHeader1{
	// · · · · Stmt: &gc.SimpleStmtExpr{
	// · · · · · Expr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 6:12: "z",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · · Body: &gc.StmtBlock{
	// · · · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · · · 0: &gc.SimpleStmtExpr{
	// · · · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · · · Ident: 7:3: "w",
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleStmt_ifElseIfElse() {
	example(&StmtIf{}, `
func foo() {
	if x, y := bar(); x {
		y()
	} else if z {
		w()
	} else {
		qux()
	}
}	
`)
	// Output:
	// 0: &gc.StmtIf{
	// · IfHeader: &gc.IfHeader2{
	// · · Stmt: &gc.SimpleStmtDefine{
	// · · · LHS: []*gc.PrimaryExprIdent{ // len 2
	// · · · · 0: &gc.PrimaryExprIdent{
	// · · · · · Ident: 4:5: "x",
	// · · · · },
	// · · · · 1: &gc.PrimaryExprIdent{
	// · · · · · Ident: 4:8: "y",
	// · · · · },
	// · · · },
	// · · · RHS: []*gc.PrimaryExprCall{ // len 1
	// · · · · 0: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 4:13: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · Stmt2: &gc.SimpleStmtExpr{
	// · · · Expr: &gc.PrimaryExprIdent{
	// · · · · Ident: 4:20: "x",
	// · · · },
	// · · },
	// · },
	// · Body: &gc.StmtBlock{
	// · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 5:3: "y",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · ElseIf: []gc.ElseIf{ // len 1
	// · · 0: gc.ElseIf{
	// · · · IfHeader: &gc.IfHeader1{
	// · · · · Stmt: &gc.SimpleStmtExpr{
	// · · · · · Expr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 6:12: "z",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · · Body: &gc.StmtBlock{
	// · · · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · · · 0: &gc.SimpleStmtExpr{
	// · · · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · · · Ident: 7:3: "w",
	// · · · · · · · },
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · Body2: &gc.StmtBlock{
	// · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 9:3: "qux",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleStmt_return() {
	example(&StmtReturn{}, `func foo() (int, string) { return 42, "314" }`)
	// Output:
	// 0: &gc.StmtReturn{
	// · ExprList: []*gc.PrimaryExprIntLiteral{ // len 2
	// · · 0: &gc.PrimaryExprIntLiteral{
	// · · · Literal: 2:35: "42",
	// · · },
	// · · 1: &gc.PrimaryExprStringLiteral{
	// · · · Literal: 2:39: "\"314\"",
	// · · },
	// · },
	// },
}

func ExampleStmt_select() {
	example(&StmtSelect{}, `
func foo() {
	select {
	case <-a:
		x()
	case <-b:
		y()
	default:
		z()
	}
}	
`)
	// Output:
	// 0: &gc.StmtSelect{
	// · List: []gc.CaseBlock{ // len 3
	// · · 0: gc.CaseBlock{
	// · · · ExprOrTypeList: []*gc.ExprOrTypeExpr{ // len 1
	// · · · · 0: &gc.ExprOrTypeExpr{
	// · · · · · Expr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 5:9: "a",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · · Body: []*gc.SimpleStmtExpr{ // len 1
	// · · · · 0: &gc.SimpleStmtExpr{
	// · · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · · Ident: 6:3: "x",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · 1: gc.CaseBlock{
	// · · · ExprOrTypeList: []*gc.ExprOrTypeExpr{ // len 1
	// · · · · 0: &gc.ExprOrTypeExpr{
	// · · · · · Expr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 7:9: "b",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · · Body: []*gc.SimpleStmtExpr{ // len 1
	// · · · · 0: &gc.SimpleStmtExpr{
	// · · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · · Ident: 8:3: "y",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · 2: gc.CaseBlock{
	// · · · Body: []*gc.SimpleStmtExpr{ // len 1
	// · · · · 0: &gc.SimpleStmtExpr{
	// · · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · · Ident: 10:3: "z",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · · Default: true,
	// · · },
	// · },
	// },
}

func ExampleStmt_switch() {
	example(&StmtSwitch{}, `
func foo() {
	switch x := y; 2*y {
	case 42:
		x()
	case 314:
		y()
	default:
		z()
	}
}	
`)
	// Output:
	// 0: &gc.StmtSwitch{
	// · IfHeader: &gc.IfHeader2{
	// · · Stmt: &gc.SimpleStmtDefine{
	// · · · LHS: []*gc.PrimaryExprIdent{ // len 1
	// · · · · 0: &gc.PrimaryExprIdent{
	// · · · · · Ident: 4:9: "x",
	// · · · · },
	// · · · },
	// · · · RHS: []*gc.PrimaryExprIdent{ // len 1
	// · · · · 0: &gc.PrimaryExprIdent{
	// · · · · · Ident: 4:14: "y",
	// · · · · },
	// · · · },
	// · · },
	// · · Stmt2: &gc.SimpleStmtExpr{
	// · · · Expr: &gc.ExprMUL{
	// · · · · LHS: &gc.PrimaryExprIntLiteral{
	// · · · · · Literal: 4:17: "2",
	// · · · · },
	// · · · · RHS: &gc.PrimaryExprIdent{
	// · · · · · Ident: 4:19: "y",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · List: []gc.CaseBlock{ // len 3
	// · · 0: gc.CaseBlock{
	// · · · ExprOrTypeList: []*gc.ExprOrTypeExpr{ // len 1
	// · · · · 0: &gc.ExprOrTypeExpr{
	// · · · · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · · · · Literal: 5:7: "42",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · · Body: []*gc.SimpleStmtExpr{ // len 1
	// · · · · 0: &gc.SimpleStmtExpr{
	// · · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · · Ident: 6:3: "x",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · 1: gc.CaseBlock{
	// · · · ExprOrTypeList: []*gc.ExprOrTypeExpr{ // len 1
	// · · · · 0: &gc.ExprOrTypeExpr{
	// · · · · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · · · · Literal: 7:7: "314",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · · Body: []*gc.SimpleStmtExpr{ // len 1
	// · · · · 0: &gc.SimpleStmtExpr{
	// · · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · · Ident: 8:3: "y",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · 2: gc.CaseBlock{
	// · · · Body: []*gc.SimpleStmtExpr{ // len 1
	// · · · · 0: &gc.SimpleStmtExpr{
	// · · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · · Ident: 10:3: "z",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · · Default: true,
	// · · },
	// · },
	// },
}

func ExampleStmt_block() {
	example(&StmtBlock{}, `
func foo() {
	{
		foo()
		bar()
	}
}	
`)
	// Output:
	// Body: &gc.StmtBlock{
	// · List: []*gc.StmtBlock{ // len 1
	// · · 0: &gc.StmtBlock{
	// · · · List: []*gc.SimpleStmtExpr{ // len 2
	// · · · · 0: &gc.SimpleStmtExpr{
	// · · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · · Ident: 5:3: "foo",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · · 1: &gc.SimpleStmtExpr{
	// · · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · · Ident: 6:3: "bar",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleIfHeader_simple() {
	example(&IfHeader1{}, "func foo() { if x {} }")
	// Output:
	// IfHeader: &gc.IfHeader1{
	// · Stmt: &gc.SimpleStmtExpr{
	// · · Expr: &gc.PrimaryExprIdent{
	// · · · Ident: 2:17: "x",
	// · · },
	// · },
	// },
}

func ExampleIfHeader_assign() {
	example(&IfHeader2{}, "func foo() { if x = y; z {} }")
	// Output:
	// IfHeader: &gc.IfHeader2{
	// · Stmt: &gc.SimpleStmtAssignment{
	// · · LHS: []*gc.PrimaryExprIdent{ // len 1
	// · · · 0: &gc.PrimaryExprIdent{
	// · · · · Ident: 2:17: "x",
	// · · · },
	// · · },
	// · · RHS: []*gc.PrimaryExprIdent{ // len 1
	// · · · 0: &gc.PrimaryExprIdent{
	// · · · · Ident: 2:21: "y",
	// · · · },
	// · · },
	// · },
	// · Stmt2: &gc.SimpleStmtExpr{
	// · · Expr: &gc.PrimaryExprIdent{
	// · · · Ident: 2:24: "z",
	// · · },
	// · },
	// },
}

func ExampleIfHeader_define() {
	example(&IfHeader2{}, "func foo() { if x := y; z {} }")
	// Output:
	// IfHeader: &gc.IfHeader2{
	// · Stmt: &gc.SimpleStmtDefine{
	// · · LHS: []*gc.PrimaryExprIdent{ // len 1
	// · · · 0: &gc.PrimaryExprIdent{
	// · · · · Ident: 2:17: "x",
	// · · · },
	// · · },
	// · · RHS: []*gc.PrimaryExprIdent{ // len 1
	// · · · 0: &gc.PrimaryExprIdent{
	// · · · · Ident: 2:22: "y",
	// · · · },
	// · · },
	// · },
	// · Stmt2: &gc.SimpleStmtExpr{
	// · · Expr: &gc.PrimaryExprIdent{
	// · · · Ident: 2:25: "z",
	// · · },
	// · },
	// },
}

func ExampleCaseBlock_selectAssign() {
	example(CaseBlock{}, `
func foo() {
	select {
	case x, ok = <-c:
		bar()
	}
}
`)
	// Output:
	// List: []gc.CaseBlock{ // len 1
	// · 0: gc.CaseBlock{
	// · · ExprOrTypeList: []*gc.ExprOrTypeExpr{ // len 2
	// · · · 0: &gc.ExprOrTypeExpr{
	// · · · · Expr: &gc.PrimaryExprIdent{
	// · · · · · Ident: 5:7: "x",
	// · · · · },
	// · · · },
	// · · · 1: &gc.ExprOrTypeExpr{
	// · · · · Expr: &gc.PrimaryExprIdent{
	// · · · · · Ident: 5:10: "ok",
	// · · · · },
	// · · · },
	// · · },
	// · · Expr: &gc.UnaryExprReceive{
	// · · · Expr: &gc.PrimaryExprIdent{
	// · · · · Ident: 5:17: "c",
	// · · · },
	// · · },
	// · · Body: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 6:3: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleCaseBlock_selectDefine() {
	example(CaseBlock{}, `
func foo() {
	select {
	case x, ok := <-c:
		bar()
	}
}
`)
	// Output:
	// List: []gc.CaseBlock{ // len 1
	// · 0: gc.CaseBlock{
	// · · ExprOrTypeList: []*gc.ExprOrTypeExpr{ // len 2
	// · · · 0: &gc.ExprOrTypeExpr{
	// · · · · Expr: &gc.PrimaryExprIdent{
	// · · · · · Ident: 5:7: "x",
	// · · · · },
	// · · · },
	// · · · 1: &gc.ExprOrTypeExpr{
	// · · · · Expr: &gc.PrimaryExprIdent{
	// · · · · · Ident: 5:10: "ok",
	// · · · · },
	// · · · },
	// · · },
	// · · Expr: &gc.UnaryExprReceive{
	// · · · Expr: &gc.PrimaryExprIdent{
	// · · · · Ident: 5:18: "c",
	// · · · },
	// · · },
	// · · Body: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 6:3: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · Define: true,
	// · },
	// },
}

func ExampleCaseBlock_switchType() {
	example(CaseBlock{}, `
func foo() {
	switch x.(type) {
	case []int:
		bar()
	}
}
`)
	// Output:
	// List: []gc.CaseBlock{ // len 1
	// · 0: gc.CaseBlock{
	// · · ExprOrTypeList: []*gc.ExprOrTypeType{ // len 1
	// · · · 0: &gc.ExprOrTypeType{
	// · · · },
	// · · },
	// · · Body: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 6:3: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleCaseBlock_switchExpr() {
	example(CaseBlock{}, `
func foo() {
	switch x {
	case 42:
		bar()
	}
}
`)
	// Output:
	// List: []gc.CaseBlock{ // len 1
	// · 0: gc.CaseBlock{
	// · · ExprOrTypeList: []*gc.ExprOrTypeExpr{ // len 1
	// · · · 0: &gc.ExprOrTypeExpr{
	// · · · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · · · Literal: 5:7: "42",
	// · · · · },
	// · · · },
	// · · },
	// · · Body: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 6:3: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleCaseBlock_default() {
	example(CaseBlock{}, `
func foo() {
	switch  x {
	case 42:
		bar()
	default:
		baz()
	}
}
`)
	// Output:
	// List: []gc.CaseBlock{ // len 2
	// · 0: gc.CaseBlock{
	// · · ExprOrTypeList: []*gc.ExprOrTypeExpr{ // len 1
	// · · · 0: &gc.ExprOrTypeExpr{
	// · · · · Expr: &gc.PrimaryExprIntLiteral{
	// · · · · · Literal: 5:7: "42",
	// · · · · },
	// · · · },
	// · · },
	// · · Body: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 6:3: "bar",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · 1: gc.CaseBlock{
	// · · Body: []*gc.SimpleStmtExpr{ // len 1
	// · · · 0: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · Ident: 8:3: "baz",
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · · Default: true,
	// · },
	// },
}

func ExampleArrayType() {
	example(&ArrayType{}, "type T [42]int")
	// Output:
	// &gc.ArrayType{
	// · Expr: &gc.PrimaryExprIntLiteral{
	// · · Literal: 2:9: "42",
	// · },
	// · Element: &gc.NamedType{
	// · · Name: "int",
	// · },
	// }
}

func ExampleSliceType() {
	example(&SliceType{}, "type T []int")
	// Output:
	// &gc.SliceType{
	// · Element: &gc.NamedType{
	// · · Name: "int",
	// · },
	// }
}

func ExampleFunctionType() {
	example(&FunctionType{}, "type T func(int, string) float64")
	// Output:
	// &gc.FunctionType{
	// · ParameterList: &gc.ParamTypeListType{
	// · · List: []*gc.ParamTypeType{ // len 2
	// · · · 0: &gc.ParamTypeType{
	// · · · · Type: &gc.NamedType{
	// · · · · · Name: "int",
	// · · · · },
	// · · · },
	// · · · 1: &gc.ParamTypeType{
	// · · · · Type: &gc.NamedType{
	// · · · · · Name: "string",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · ResultList: &gc.ParamTypeListType{
	// · · List: []*gc.ParamTypeType{ // len 1
	// · · · 0: &gc.ParamTypeType{
	// · · · · Type: &gc.NamedType{
	// · · · · · Name: "float64",
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleMapType() {
	example(&MapType{}, "type T map[int]string")
	// Output:
	// &gc.MapType{
	// · Key: &gc.NamedType{
	// · · Name: "int",
	// · },
	// · Value: &gc.NamedType{
	// · · Name: "string",
	// · },
	// }
}

func ExamplePointerType() {
	example(&PointerType{}, "type T *int")
	// Output:
	// &gc.PointerType{
	// · Element: &gc.NamedType{
	// · · Name: "int",
	// · },
	// }
}

func ExampleChannelType() {
	example(&ChannelType{}, "type T <-chan int")
	// Output:
	// &gc.ChannelType{
	// · Element: &gc.NamedType{
	// · · Name: "int",
	// · },
	// · Dir: RxChan,
	// }
}

func ExampleInterfaceType() {
	example(&InterfaceType{}, "type T interface{ m(int) string }")
	// Output:
	// &gc.InterfaceType{
	// · Methods: map[string]*gc.FunctionType{
	// · · "m": &gc.FunctionType{
	// · · · ParameterList: &gc.ParamTypeListType{
	// · · · · List: []*gc.ParamTypeType{ // len 1
	// · · · · · 0: &gc.ParamTypeType{
	// · · · · · · Type: &gc.NamedType{
	// · · · · · · · Name: "int",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · · ResultList: &gc.ParamTypeListType{
	// · · · · List: []*gc.ParamTypeType{ // len 1
	// · · · · · 0: &gc.ParamTypeType{
	// · · · · · · Type: &gc.NamedType{
	// · · · · · · · Name: "string",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleNamedType() {
	example(&NamedType{}, "type T int")
	// Output:
	// &gc.NamedType{
	// · Name: "int",
	// }
}

func ExampleQualifiedNamedType() {
	example(&QualifiedNamedType{}, "type T foo.Bar")
	// Output:
	// &gc.QualifiedNamedType{
	// · Qualifier: "foo",
	// · Name: "Bar",
	// }
}

func ExampleStructType() {
	example(&StructType{}, "type T struct{ i int; s string }")
	// Output:
	// &gc.StructType{
	// · Fields: []*gc.FieldDeclarationNamed{ // len 2
	// · · 0: &gc.FieldDeclarationNamed{
	// · · · Name: 2:16: "i",
	// · · · Type: &gc.NamedType{
	// · · · · Name: "int",
	// · · · },
	// · · },
	// · · 1: &gc.FieldDeclarationNamed{
	// · · · Name: 2:23: "s",
	// · · · Type: &gc.NamedType{
	// · · · · Name: "string",
	// · · · },
	// · · },
	// · },
	// }
}

func ExampleElseIf() {
	example(ElseIf{}, `
func foo() {
	if x {
		bar()
	} else if y {
		baz()
	} else if y {
		qux()
	}
}
`)
	// Output:
	// ElseIf: []gc.ElseIf{ // len 2
	// · 0: gc.ElseIf{
	// · · IfHeader: &gc.IfHeader1{
	// · · · Stmt: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprIdent{
	// · · · · · Ident: 6:12: "y",
	// · · · · },
	// · · · },
	// · · },
	// · · Body: &gc.StmtBlock{
	// · · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · · 0: &gc.SimpleStmtExpr{
	// · · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · · Ident: 7:3: "baz",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// · 1: gc.ElseIf{
	// · · IfHeader: &gc.IfHeader1{
	// · · · Stmt: &gc.SimpleStmtExpr{
	// · · · · Expr: &gc.PrimaryExprIdent{
	// · · · · · Ident: 8:12: "y",
	// · · · · },
	// · · · },
	// · · },
	// · · Body: &gc.StmtBlock{
	// · · · List: []*gc.SimpleStmtExpr{ // len 1
	// · · · · 0: &gc.SimpleStmtExpr{
	// · · · · · Expr: &gc.PrimaryExprCall{
	// · · · · · · PrimaryExpr: &gc.PrimaryExprIdent{
	// · · · · · · · Ident: 9:3: "qux",
	// · · · · · · },
	// · · · · · },
	// · · · · },
	// · · · },
	// · · },
	// · },
	// },
}

func ExampleFieldDeclaration_embedded() {
	example(&FieldDeclarationEmbedded{}, "type T struct { io.Reader }")
	// Output:
	// 0: &gc.FieldDeclarationEmbedded{
	// · Type: &gc.QualifiedNamedType{
	// · · Qualifier: "io",
	// · · Name: "Reader",
	// · },
	// },
}

func ExampleFieldDeclaration_embeddedPtr() {
	example(&FieldDeclarationEmbedded{}, "type T struct { *int }")
	// Output:
	// 0: &gc.FieldDeclarationEmbedded{
	// · Type: &gc.PointerType{
	// · · Element: &gc.NamedType{
	// · · · Name: "int",
	// · · },
	// · },
	// · Ptr: true,
	// },
}

func ExampleFieldDeclaration_named() {
	example(&FieldDeclarationNamed{}, "type T struct { i int }")
	// Output:
	// 0: &gc.FieldDeclarationNamed{
	// · Name: 2:17: "i",
	// · Type: &gc.NamedType{
	// · · Name: "int",
	// · },
	// },
}
