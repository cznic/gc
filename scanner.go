// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

func (lx *lexer) scan() int {
	c := lx.Enter()

	/* Non ASCII character classes */

yystate0:
	yyrule := -1
	_ = yyrule
	c = lx.Rule0()

	// White space except newline.

	goto yystart1

	goto yystate0 // silence unused label error
	goto yyAction // silence unused label error
yyAction:
	switch yyrule {
	case 1:
		goto yyrule1
	case 2:
		goto yyrule2
	case 3:
		goto yyrule3
	case 4:
		goto yyrule4
	case 5:
		goto yyrule5
	case 6:
		goto yyrule6
	case 7:
		goto yyrule7
	case 8:
		goto yyrule8
	case 9:
		goto yyrule9
	case 10:
		goto yyrule10
	case 11:
		goto yyrule11
	case 12:
		goto yyrule12
	case 13:
		goto yyrule13
	case 14:
		goto yyrule14
	case 15:
		goto yyrule15
	case 16:
		goto yyrule16
	case 17:
		goto yyrule17
	case 18:
		goto yyrule18
	case 19:
		goto yyrule19
	case 20:
		goto yyrule20
	case 21:
		goto yyrule21
	case 22:
		goto yyrule22
	case 23:
		goto yyrule23
	case 24:
		goto yyrule24
	case 25:
		goto yyrule25
	case 26:
		goto yyrule26
	case 27:
		goto yyrule27
	case 28:
		goto yyrule28
	case 29:
		goto yyrule29
	case 30:
		goto yyrule30
	case 31:
		goto yyrule31
	case 32:
		goto yyrule32
	case 33:
		goto yyrule33
	case 34:
		goto yyrule34
	case 35:
		goto yyrule35
	case 36:
		goto yyrule36
	case 37:
		goto yyrule37
	case 38:
		goto yyrule38
	case 39:
		goto yyrule39
	case 40:
		goto yyrule40
	case 41:
		goto yyrule41
	case 42:
		goto yyrule42
	case 43:
		goto yyrule43
	case 44:
		goto yyrule44
	case 45:
		goto yyrule45
	case 46:
		goto yyrule46
	case 47:
		goto yyrule47
	case 48:
		goto yyrule48
	case 49:
		goto yyrule49
	case 50:
		goto yyrule50
	case 51:
		goto yyrule51
	case 52:
		goto yyrule52
	case 53:
		goto yyrule53
	case 54:
		goto yyrule54
	case 55:
		goto yyrule55
	case 56:
		goto yyrule56
	case 57:
		goto yyrule57
	case 58:
		goto yyrule58
	case 59:
		goto yyrule59
	case 60:
		goto yyrule60
	case 61:
		goto yyrule61
	case 62:
		goto yyrule62
	case 63:
		goto yyrule63
	}
	goto yystate1 // silence unused label error
yystate1:
	c = lx.Next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate5
	case c == '%':
		goto yystate16
	case c == '&':
		goto yystate18
	case c == '*':
		goto yystate40
	case c == '+':
		goto yystate42
	case c == '-':
		goto yystate45
	case c == '.':
		goto yystate48
	case c == '/':
		goto yystate55
	case c == '0':
		goto yystate65
	case c == ':':
		goto yystate72
	case c == '<':
		goto yystate74
	case c == '=':
		goto yystate79
	case c == '>':
		goto yystate81
	case c == '\'':
		goto yystate23
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	case c == '\u0082':
		goto yystate204
	case c == '\u0083':
		goto yystate205
	case c == '^':
		goto yystate86
	case c == '`':
		goto yystate88
	case c == 'b':
		goto yystate89
	case c == 'c':
		goto yystate94
	case c == 'd':
		goto yystate110
	case c == 'e':
		goto yystate119
	case c == 'f':
		goto yystate123
	case c == 'g':
		goto yystate139
	case c == 'i':
		goto yystate143
	case c == 'm':
		goto yystate158
	case c == 'p':
		goto yystate161
	case c == 'r':
		goto yystate168
	case c == 's':
		goto yystate178
	case c == 't':
		goto yystate194
	case c == 'v':
		goto yystate198
	case c == '|':
		goto yystate201
	case c >= '1' && c <= '9':
		goto yystate71
	case c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'h' || c >= 'j' && c <= 'l' || c == 'n' || c == 'o' || c == 'q' || c == 'u' || c >= 'w' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate2:
	c = lx.Next()
	yyrule = 1
	lx.Mark()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate4
	}

yystate4:
	c = lx.Next()
	yyrule = 5
	lx.Mark()
	goto yyrule5

yystate5:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate6
	case c == '\\':
		goto yystate7
	case c >= '\x01' && c <= '\t' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate5
	}

yystate6:
	c = lx.Next()
	yyrule = 62
	lx.Mark()
	goto yyrule62

yystate7:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '"' || c == '\'' || c >= '0' && c <= '7' || c == '\\' || c == 'a' || c == 'b' || c == 'f' || c == 'n' || c == 'r' || c == 't' || c == 'v' || c == 'x':
		goto yystate5
	case c == 'U':
		goto yystate8
	case c == 'u':
		goto yystate12
	}

yystate8:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate9
	}

yystate9:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate10
	}

yystate10:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate11
	}

yystate11:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate12
	}

yystate12:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate13
	}

yystate13:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate14
	}

yystate14:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate15
	}

yystate15:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate5
	}

yystate16:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate17
	}

yystate17:
	c = lx.Next()
	yyrule = 6
	lx.Mark()
	goto yyrule6

yystate18:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '&':
		goto yystate19
	case c == '=':
		goto yystate20
	case c == '^':
		goto yystate21
	}

yystate19:
	c = lx.Next()
	yyrule = 7
	lx.Mark()
	goto yyrule7

yystate20:
	c = lx.Next()
	yyrule = 8
	lx.Mark()
	goto yyrule8

yystate21:
	c = lx.Next()
	yyrule = 9
	lx.Mark()
	switch {
	default:
		goto yyrule9
	case c == '=':
		goto yystate22
	}

yystate22:
	c = lx.Next()
	yyrule = 10
	lx.Mark()
	goto yyrule10

yystate23:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate26
	case c == '\\':
		goto yystate27
	case c >= '\x01' && c <= '\t' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate24
	}

yystate24:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate25
	}

yystate25:
	c = lx.Next()
	yyrule = 61
	lx.Mark()
	goto yyrule61

yystate26:
	c = lx.Next()
	yyrule = 61
	lx.Mark()
	switch {
	default:
		goto yyrule61
	case c == '\'':
		goto yystate25
	}

yystate27:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '"' || c == '\'' || c == '\\' || c == 'a' || c == 'b' || c == 'f' || c == 'n' || c == 'r' || c == 't' || c == 'v':
		goto yystate24
	case c == 'U':
		goto yystate30
	case c == 'u':
		goto yystate34
	case c == 'x':
		goto yystate38
	case c >= '0' && c <= '7':
		goto yystate28
	}

yystate28:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate25
	case c >= '0' && c <= '7':
		goto yystate29
	}

yystate29:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate25
	case c >= '0' && c <= '7':
		goto yystate24
	}

yystate30:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate31
	}

yystate31:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate32
	}

yystate32:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate33
	}

yystate33:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate34
	}

yystate34:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate35
	}

yystate35:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate36
	}

yystate36:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate37
	}

yystate37:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate24
	}

yystate38:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate25
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate39
	}

yystate39:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate25
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate24
	}

yystate40:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate41
	}

yystate41:
	c = lx.Next()
	yyrule = 11
	lx.Mark()
	goto yyrule11

yystate42:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '+':
		goto yystate43
	case c == '=':
		goto yystate44
	}

yystate43:
	c = lx.Next()
	yyrule = 12
	lx.Mark()
	goto yyrule12

yystate44:
	c = lx.Next()
	yyrule = 13
	lx.Mark()
	goto yyrule13

yystate45:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '-':
		goto yystate46
	case c == '=':
		goto yystate47
	}

yystate46:
	c = lx.Next()
	yyrule = 14
	lx.Mark()
	goto yyrule14

yystate47:
	c = lx.Next()
	yyrule = 15
	lx.Mark()
	goto yyrule15

yystate48:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate49
	case c >= '0' && c <= '9':
		goto yystate51
	}

yystate49:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate50
	}

yystate50:
	c = lx.Next()
	yyrule = 16
	lx.Mark()
	goto yyrule16

yystate51:
	c = lx.Next()
	yyrule = 57
	lx.Mark()
	switch {
	default:
		goto yyrule57
	case c == 'E' || c == 'e':
		goto yystate52
	case c == 'i':
		goto yystate54
	case c >= '0' && c <= '9':
		goto yystate51
	}

yystate52:
	c = lx.Next()
	yyrule = 57
	lx.Mark()
	switch {
	default:
		goto yyrule57
	case c == '+' || c == '-' || c >= '0' && c <= '9':
		goto yystate53
	case c == 'i':
		goto yystate54
	}

yystate53:
	c = lx.Next()
	yyrule = 57
	lx.Mark()
	switch {
	default:
		goto yyrule57
	case c == 'i':
		goto yystate54
	case c >= '0' && c <= '9':
		goto yystate53
	}

yystate54:
	c = lx.Next()
	yyrule = 59
	lx.Mark()
	goto yyrule59

yystate55:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate56
	case c == '/':
		goto yystate62
	case c == '=':
		goto yystate64
	}

yystate56:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate60
	case c == '\n':
		goto yystate57
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= ')' || c >= '+' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate56
	}

yystate57:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate58
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate57
	}

yystate58:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate58
	case c == '/':
		goto yystate59
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '.' || c >= '0' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate57
	}

yystate59:
	c = lx.Next()
	yyrule = 4
	lx.Mark()
	goto yyrule4

yystate60:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate60
	case c == '/':
		goto yystate61
	case c == '\n':
		goto yystate57
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= ')' || c >= '+' && c <= '.' || c >= '0' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate56
	}

yystate61:
	c = lx.Next()
	yyrule = 2
	lx.Mark()
	goto yyrule2

yystate62:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '\n' || c == '\r' || c == '\u0080':
		goto yystate63
	case c >= '\x01' && c <= '\t' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate62
	}

yystate63:
	c = lx.Next()
	yyrule = 3
	lx.Mark()
	goto yyrule3

yystate64:
	c = lx.Next()
	yyrule = 17
	lx.Mark()
	goto yyrule17

yystate65:
	c = lx.Next()
	yyrule = 60
	lx.Mark()
	switch {
	default:
		goto yyrule60
	case c == '.':
		goto yystate51
	case c == '8' || c == '9':
		goto yystate67
	case c == 'E' || c == 'e':
		goto yystate52
	case c == 'X' || c == 'x':
		goto yystate69
	case c == 'i':
		goto yystate54
	case c == 'p':
		goto yystate68
	case c >= '0' && c <= '7':
		goto yystate66
	}

yystate66:
	c = lx.Next()
	yyrule = 60
	lx.Mark()
	switch {
	default:
		goto yyrule60
	case c == '.':
		goto yystate51
	case c == '8' || c == '9':
		goto yystate67
	case c == 'E' || c == 'e':
		goto yystate52
	case c == 'i':
		goto yystate54
	case c == 'p':
		goto yystate68
	case c >= '0' && c <= '7':
		goto yystate66
	}

yystate67:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate51
	case c == 'E' || c == 'e':
		goto yystate52
	case c == 'i':
		goto yystate54
	case c >= '0' && c <= '9':
		goto yystate67
	}

yystate68:
	c = lx.Next()
	yyrule = 63
	lx.Mark()
	switch {
	default:
		goto yyrule63
	case c >= '0' && c <= '9':
		goto yystate68
	}

yystate69:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate70
	}

yystate70:
	c = lx.Next()
	yyrule = 60
	lx.Mark()
	switch {
	default:
		goto yyrule60
	case c == 'p':
		goto yystate68
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate70
	}

yystate71:
	c = lx.Next()
	yyrule = 60
	lx.Mark()
	switch {
	default:
		goto yyrule60
	case c == '.':
		goto yystate51
	case c == 'E' || c == 'e':
		goto yystate52
	case c == 'i':
		goto yystate54
	case c == 'p':
		goto yystate68
	case c >= '0' && c <= '9':
		goto yystate71
	}

yystate72:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate73
	}

yystate73:
	c = lx.Next()
	yyrule = 18
	lx.Mark()
	goto yyrule18

yystate74:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '-':
		goto yystate75
	case c == '<':
		goto yystate76
	case c == '=':
		goto yystate78
	}

yystate75:
	c = lx.Next()
	yyrule = 19
	lx.Mark()
	goto yyrule19

yystate76:
	c = lx.Next()
	yyrule = 20
	lx.Mark()
	switch {
	default:
		goto yyrule20
	case c == '=':
		goto yystate77
	}

yystate77:
	c = lx.Next()
	yyrule = 21
	lx.Mark()
	goto yyrule21

yystate78:
	c = lx.Next()
	yyrule = 22
	lx.Mark()
	goto yyrule22

yystate79:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate80
	}

yystate80:
	c = lx.Next()
	yyrule = 23
	lx.Mark()
	goto yyrule23

yystate81:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate82
	case c == '>':
		goto yystate83
	}

yystate82:
	c = lx.Next()
	yyrule = 24
	lx.Mark()
	goto yyrule24

yystate83:
	c = lx.Next()
	yyrule = 25
	lx.Mark()
	switch {
	default:
		goto yyrule25
	case c == '=':
		goto yystate84
	}

yystate84:
	c = lx.Next()
	yyrule = 26
	lx.Mark()
	goto yyrule26

yystate85:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate86:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate87
	}

yystate87:
	c = lx.Next()
	yyrule = 27
	lx.Mark()
	goto yyrule27

yystate88:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '`':
		goto yystate6
	case c >= '\x01' && c <= '_' || c >= 'a' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate88
	}

yystate89:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'r':
		goto yystate90
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate90:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'e':
		goto yystate91
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate91:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'a':
		goto yystate92
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate92:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'k':
		goto yystate93
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate93:
	c = lx.Next()
	yyrule = 32
	lx.Mark()
	switch {
	default:
		goto yyrule32
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate94:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'a':
		goto yystate95
	case c == 'h':
		goto yystate98
	case c == 'o':
		goto yystate101
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'g' || c >= 'i' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate95:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 's':
		goto yystate96
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate96:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'e':
		goto yystate97
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate97:
	c = lx.Next()
	yyrule = 33
	lx.Mark()
	switch {
	default:
		goto yyrule33
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate98:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'a':
		goto yystate99
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate99:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'n':
		goto yystate100
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate100:
	c = lx.Next()
	yyrule = 34
	lx.Mark()
	switch {
	default:
		goto yyrule34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate101:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'n':
		goto yystate102
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate102:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 's':
		goto yystate103
	case c == 't':
		goto yystate105
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 'u' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate103:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 't':
		goto yystate104
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate104:
	c = lx.Next()
	yyrule = 35
	lx.Mark()
	switch {
	default:
		goto yyrule35
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate105:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'i':
		goto yystate106
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate106:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'n':
		goto yystate107
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate107:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'u':
		goto yystate108
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate108:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'e':
		goto yystate109
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate109:
	c = lx.Next()
	yyrule = 36
	lx.Mark()
	switch {
	default:
		goto yyrule36
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate110:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'e':
		goto yystate111
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate111:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'f':
		goto yystate112
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate112:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'a':
		goto yystate113
	case c == 'e':
		goto yystate117
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate113:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'u':
		goto yystate114
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate114:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'l':
		goto yystate115
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate115:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 't':
		goto yystate116
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate116:
	c = lx.Next()
	yyrule = 37
	lx.Mark()
	switch {
	default:
		goto yyrule37
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate117:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'r':
		goto yystate118
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate118:
	c = lx.Next()
	yyrule = 38
	lx.Mark()
	switch {
	default:
		goto yyrule38
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate119:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'l':
		goto yystate120
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate120:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 's':
		goto yystate121
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate121:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'e':
		goto yystate122
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate122:
	c = lx.Next()
	yyrule = 39
	lx.Mark()
	switch {
	default:
		goto yyrule39
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate123:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'a':
		goto yystate124
	case c == 'o':
		goto yystate134
	case c == 'u':
		goto yystate136
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'n' || c >= 'p' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate124:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'l':
		goto yystate125
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate125:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'l':
		goto yystate126
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate126:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 't':
		goto yystate127
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate127:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'h':
		goto yystate128
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate128:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'r':
		goto yystate129
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate129:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'o':
		goto yystate130
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate130:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'u':
		goto yystate131
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate131:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'g':
		goto yystate132
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate132:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'h':
		goto yystate133
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate133:
	c = lx.Next()
	yyrule = 40
	lx.Mark()
	switch {
	default:
		goto yyrule40
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate134:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'r':
		goto yystate135
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate135:
	c = lx.Next()
	yyrule = 41
	lx.Mark()
	switch {
	default:
		goto yyrule41
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate136:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'n':
		goto yystate137
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate137:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'c':
		goto yystate138
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate138:
	c = lx.Next()
	yyrule = 42
	lx.Mark()
	switch {
	default:
		goto yyrule42
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate139:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'o':
		goto yystate140
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate140:
	c = lx.Next()
	yyrule = 43
	lx.Mark()
	switch {
	default:
		goto yyrule43
	case c == 't':
		goto yystate141
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate141:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'o':
		goto yystate142
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate142:
	c = lx.Next()
	yyrule = 44
	lx.Mark()
	switch {
	default:
		goto yyrule44
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate143:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'f':
		goto yystate144
	case c == 'm':
		goto yystate145
	case c == 'n':
		goto yystate150
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'l' || c >= 'o' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate144:
	c = lx.Next()
	yyrule = 45
	lx.Mark()
	switch {
	default:
		goto yyrule45
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate145:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'p':
		goto yystate146
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate146:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'o':
		goto yystate147
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate147:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'r':
		goto yystate148
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate148:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 't':
		goto yystate149
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate149:
	c = lx.Next()
	yyrule = 46
	lx.Mark()
	switch {
	default:
		goto yyrule46
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate150:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 't':
		goto yystate151
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate151:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'e':
		goto yystate152
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate152:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'r':
		goto yystate153
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate153:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'f':
		goto yystate154
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate154:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'a':
		goto yystate155
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate155:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'c':
		goto yystate156
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate156:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'e':
		goto yystate157
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate157:
	c = lx.Next()
	yyrule = 47
	lx.Mark()
	switch {
	default:
		goto yyrule47
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate158:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'a':
		goto yystate159
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate159:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'p':
		goto yystate160
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate160:
	c = lx.Next()
	yyrule = 48
	lx.Mark()
	switch {
	default:
		goto yyrule48
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate161:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'a':
		goto yystate162
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate162:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'c':
		goto yystate163
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate163:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'k':
		goto yystate164
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate164:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'a':
		goto yystate165
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate165:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'g':
		goto yystate166
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate166:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'e':
		goto yystate167
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate167:
	c = lx.Next()
	yyrule = 49
	lx.Mark()
	switch {
	default:
		goto yyrule49
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate168:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'a':
		goto yystate169
	case c == 'e':
		goto yystate173
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate169:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'n':
		goto yystate170
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate170:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'g':
		goto yystate171
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate171:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'e':
		goto yystate172
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate172:
	c = lx.Next()
	yyrule = 50
	lx.Mark()
	switch {
	default:
		goto yyrule50
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate173:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 't':
		goto yystate174
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate174:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'u':
		goto yystate175
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate175:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'r':
		goto yystate176
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate176:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'n':
		goto yystate177
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate177:
	c = lx.Next()
	yyrule = 51
	lx.Mark()
	switch {
	default:
		goto yyrule51
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate178:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'e':
		goto yystate179
	case c == 't':
		goto yystate184
	case c == 'w':
		goto yystate189
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 's' || c == 'u' || c == 'v' || c >= 'x' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate179:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'l':
		goto yystate180
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate180:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'e':
		goto yystate181
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate181:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'c':
		goto yystate182
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate182:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 't':
		goto yystate183
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate183:
	c = lx.Next()
	yyrule = 52
	lx.Mark()
	switch {
	default:
		goto yyrule52
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate184:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'r':
		goto yystate185
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate185:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'u':
		goto yystate186
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate186:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'c':
		goto yystate187
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate187:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 't':
		goto yystate188
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate188:
	c = lx.Next()
	yyrule = 53
	lx.Mark()
	switch {
	default:
		goto yyrule53
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate189:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'i':
		goto yystate190
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate190:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 't':
		goto yystate191
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate191:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'c':
		goto yystate192
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate192:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'h':
		goto yystate193
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate193:
	c = lx.Next()
	yyrule = 54
	lx.Mark()
	switch {
	default:
		goto yyrule54
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate194:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'y':
		goto yystate195
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z' || c == '\u0081':
		goto yystate85
	}

yystate195:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'p':
		goto yystate196
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate196:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'e':
		goto yystate197
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate197:
	c = lx.Next()
	yyrule = 55
	lx.Mark()
	switch {
	default:
		goto yyrule55
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate198:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'a':
		goto yystate199
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate199:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	switch {
	default:
		goto yyrule58
	case c == 'r':
		goto yystate200
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate200:
	c = lx.Next()
	yyrule = 56
	lx.Mark()
	switch {
	default:
		goto yyrule56
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081':
		goto yystate85
	}

yystate201:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate202
	case c == '|':
		goto yystate203
	}

yystate202:
	c = lx.Next()
	yyrule = 28
	lx.Mark()
	goto yyrule28

yystate203:
	c = lx.Next()
	yyrule = 29
	lx.Mark()
	goto yyrule29

yystate204:
	c = lx.Next()
	yyrule = 31
	lx.Mark()
	goto yyrule31

yystate205:
	c = lx.Next()
	yyrule = 30
	lx.Mark()
	goto yyrule30

yyrule1: // [ \t\r]+
	{

		// A general comment containing no newlines acts like a space.
		goto yystate0
	}
yyrule2: // "/*"([^*\x80\n]|\*+[^*/\x80\n])*\*+\/
	{

		lx.checkComment()
		// Any other comment acts like a newline.
		goto yystate0
	}
yyrule3: // "//"{any_to_eol}[\n\r\x80]
yyrule4: // "/*"([^*\x80]|\*+[^*/\x80])*\*+\/
	{

		lx.checkComment()
		return '\n'
	}
yyrule5: // "!="
	{
		return NEQ
	}
yyrule6: // "%="
	{
		return MOD_ASSIGN
	}
yyrule7: // "&&"
	{
		return ANDAND
	}
yyrule8: // "&="
	{
		return AND_ASSIGN
	}
yyrule9: // "&^"
	{
		return ANDNOT
	}
yyrule10: // "&^="
	{
		return ANDNOT_ASSIGN
	}
yyrule11: // "*="
	{
		return MUL_ASSIGN
	}
yyrule12: // "++"
	{
		return INC
	}
yyrule13: // "+="
	{
		return ADD_ASSIGN
	}
yyrule14: // "--"
	{
		return DEC
	}
yyrule15: // "-="
	{
		return SUB_ASSIGN
	}
yyrule16: // "..."
	{
		return DDD
	}
yyrule17: // "/="
	{
		return DIV_ASSIGN
	}
yyrule18: // ":="
	{
		return COLAS
	}
yyrule19: // "<-"
	{
		return COMM
	}
yyrule20: // "<<"
	{
		return LSH
	}
yyrule21: // "<<="
	{
		return LSH_ASSIGN
	}
yyrule22: // "<="
	{
		return LEQ
	}
yyrule23: // "=="
	{
		return EQ
	}
yyrule24: // ">="
	{
		return GEQ
	}
yyrule25: // ">>"
	{
		return RSH
	}
yyrule26: // ">>="
	{
		return RSH_ASSIGN
	}
yyrule27: // "^="
	{
		return XOR_ASSIGN
	}
yyrule28: // "|="
	{
		return OR_ASSIGN
	}
yyrule29: // "||"
	{
		return OROR
	}
yyrule30: // {gtgt}
	{
		return GTGT
	}
yyrule31: // {ltlt}
	{
		return LTLT
	}
yyrule32: // "break"
	{
		return BREAK
	}
yyrule33: // "case"
	{
		return CASE
	}
yyrule34: // "chan"
	{
		return CHAN
	}
yyrule35: // "const"
	{
		return CONST
	}
yyrule36: // "continue"
	{
		return CONTINUE
	}
yyrule37: // "default"
	{
		return DEFAULT
	}
yyrule38: // "defer"
	{
		return DEFER
	}
yyrule39: // "else"
	{
		return ELSE
	}
yyrule40: // "fallthrough"
	{
		return FALLTHROUGH
	}
yyrule41: // "for"
	{
		return FOR
	}
yyrule42: // "func"
	{
		return FUNC
	}
yyrule43: // "go"
	{
		return GO
	}
yyrule44: // "goto"
	{
		return GOTO
	}
yyrule45: // "if"
	{
		return IF
	}
yyrule46: // "import"
	{
		return IMPORT
	}
yyrule47: // "interface"
	{
		return INTERFACE
	}
yyrule48: // "map"
	{
		return MAP
	}
yyrule49: // "package"
	{
		return PACKAGE
	}
yyrule50: // "range"
	{
		return RANGE
	}
yyrule51: // "return"
	{
		return RETURN
	}
yyrule52: // "select"
	{
		return SELECT
	}
yyrule53: // "struct"
	{
		return STRUCT
	}
yyrule54: // "switch"
	{
		return SWITCH
	}
yyrule55: // "type"
	{
		return TYPE
	}
yyrule56: // "var"
	{
		return VAR
	}
yyrule57: // {float_lit}
	{
		return FLOAT_LIT
	}
yyrule58: // {identifier}
	{
		return IDENTIFIER
	}
yyrule59: // {imaginary_lit}
	{
		return IMAG_LIT
	}
yyrule60: // {int_lit}
	{
		return INT_LIT
	}
yyrule61: // {rune_lit}
	{
		return CHAR_LIT
	}
yyrule62: // {string_lit}
	{
		return STRING_LIT
	}
yyrule63: // {int_lit}p[0-9]*
	{

		return BAD_FLOAT_LIT
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	if c, ok := lx.Abort(); ok {
		return c
	}

	goto yyAction
}
