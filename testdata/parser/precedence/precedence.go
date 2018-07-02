package main

func main() {
	_ = 1
	//	Rhs: []gc.PrimaryExprIntLiteral{ // len 1
	//	· 0: gc.PrimaryExprIntLiteral{
	//	· · Literal: gc.Token{
	//	· · · File: &token.File{ /* recursive/repetitive pointee not shown */ },
	//	· · · Pos: 34,
	//	· · · Val: "1",
	//	· · },
	//	· },
	//	},

	_ = 1 + 2
	//	Rhs: []gc.ExprADD{ // len 1
	//	· 0: gc.ExprADD{
	//	· · Lhs: gc.PrimaryExprIntLiteral{
	//	· · · · Val: "1",
	//	· · · },
	//	· · },
	//	· · Rhs: gc.PrimaryExprIntLiteral{
	//	· · · · Val: "2",
	//	· · · },
	//	· · },
	//	· },
	//	},

	_ = 1 + 2*3
	//	Rhs: []gc.ExprADD{ // len 1
	//	· 0: gc.ExprADD{
	//	· · Lhs: gc.PrimaryExprIntLiteral{
	//	· · · · Val: "1",
	//	· · · },
	//	· · },
	//	· · Rhs: gc.ExprMUL{
	//	· · · Lhs: gc.PrimaryExprIntLiteral{
	//	· · · · · Val: "2",
	//	· · · · },
	//	· · · },
	//	· · · Rhs: gc.PrimaryExprIntLiteral{
	//	· · · · · Val: "3",
	//	· · · · },
	//	· · · },
	//	· · },
	//	· },
	//	},

	_ = 1*2 + 3
	//	Rhs: []gc.ExprMUL{ // len 1
	//	· 0: gc.ExprMUL{
	//	· · Lhs: gc.PrimaryExprIntLiteral{
	//	· · · · Val: "1",
	//	· · · },
	//	· · },
	//	· · Rhs: gc.ExprADD{
	//	· · · Lhs: gc.PrimaryExprIntLiteral{
	//	· · · · · Val: "2",
	//	· · · · },
	//	· · · },
	//	· · · Rhs: gc.PrimaryExprIntLiteral{
	//	· · · · · Val: "3",
	//	· · · · },
	//	· · · },
	//	· · },
	//	· },
	//	},

	//	Rhs: []gc.ExprADD{ // len 1
	//	· 0: gc.ExprADD{
	//	· · Lhs: gc.ExprMUL{
	//	· · · Lhs: gc.PrimaryExprIntLiteral{
	//	· · · · · Val: "1",
	//	· · · · },
	//	· · · },
	//	· · · Rhs: gc.PrimaryExprIntLiteral{
	//	· · · · · Val: "2",
	//	· · · · },
	//	· · · },
	//	· · },
	//	· · Rhs: gc.PrimaryExprIntLiteral{
	//	· · · · Val: "3",
	//	· · · },
	//	· · },
	//	· },
	//	},

}
