%{
// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
%}

%yyc c
%yyn c = lx.Next()
%yym lx.Mark()

%{
package gc

func (lx *lexer) scan() int {
		c := lx.Enter()
%}

/* Non ASCII character classes */
eof			\x80
other			\x81
ltlt			\x82
gtgt			\x83

any_to_eol		[^\x80\n\r]*
big_u_value		\\U{hex_digit}{hex_digit}{hex_digit}{hex_digit}{hex_digit}{hex_digit}{hex_digit}{hex_digit}
byte_value		{octal_byte_value}|{hex_byte_value}
comment			"/*"([^*\x80]|\*+[^*/\x80])*\*+\/
decimal_digit		[0-9]
decimal_lit		[1-9]{decimal_digit}*
decimals		{decimal_digit}+
escaped_char		\\[abfnrtv'\x22\\]
exponent		[eE][+-]?{decimals}*
float_lit		{decimals}"."{decimals}?{exponent}?|{decimals}{exponent}|"."{decimals}{exponent}?
hex_byte_value		\\x{hex_digit}?{hex_digit}?
hex_digit		[0-9a-fA-F]
hex_lit			0[xX]{hex_digit}+
identifier		{letter}({letter}|[0-9])*
imaginary_lit		({decimals}|{float_lit})"i"
int_lit			{decimal_lit}|{octal_lit}|{hex_lit}
interpreted_string_lit	\x22({str_unicode_value}|{byte_value})*\x22
letter			[a-zA-Z_]|{other}
little_u_value		\\u{hex_digit}{hex_digit}{hex_digit}{hex_digit}
octal_byte_value	\\{octal_digit}{octal_digit}?{octal_digit}?
octal_digit		[0-7]
octal_lit		0{octal_digit}*
raw_string_lit		`[^`\x80]*`
rune_lit		'({rune_unicode_value}?|{byte_value})'
rune_unicode_value	[^\x80\n\r\\]|{little_u_value}|{big_u_value}|{escaped_char}
str_unicode_value	[^\x22\x80\n\r\\]|{little_u_value}|{big_u_value}|{escaped_char}
string_lit		{raw_string_lit}|{interpreted_string_lit}

%%
		c = lx.Rule0()

 
  // White space except newline.
[ \t\r]+

 // A general comment containing no newlines acts like a space.
"/*"([^*\x80\n]|\*+[^*/\x80\n])*\*+\/
		lx.checkComment()

 // Any other comment acts like a newline.
"//"{any_to_eol}[\n\r\x80]	|
"/*"([^*\x80]|\*+[^*/\x80])*\*+\/
		lx.checkComment()
		return '\n'

"!="		return NEQ
"%="		return MOD_ASSIGN
"&&"		return ANDAND
"&="		return AND_ASSIGN
"&^"		return ANDNOT
"&^="		return ANDNOT_ASSIGN
"*="		return MUL_ASSIGN
"++"		return INC
"+="		return ADD_ASSIGN
"--"		return DEC
"-="		return SUB_ASSIGN
"..."		return DDD
"/="		return DIV_ASSIGN
":="		return COLAS
"<-"		return COMM
"<<"		return LSH
"<<="		return LSH_ASSIGN
"<="		return LEQ
"=="		return EQ
">="		return GEQ
">>"		return RSH
">>="		return RSH_ASSIGN
"^="		return XOR_ASSIGN
"|="		return OR_ASSIGN
"||"		return OROR
{gtgt}		return GTGT
{ltlt}		return LTLT

"break"		return BREAK
"case"		return CASE
"chan"		return CHAN
"const"		return CONST
"continue"	return CONTINUE
"default"	return DEFAULT
"defer"		return DEFER
"else"		return ELSE
"fallthrough"	return FALLTHROUGH
"for"		return FOR
"func"		return FUNC
"go"		return GO
"goto"		return GOTO
"if"		return IF
"import"	return IMPORT
"interface"	return INTERFACE
"map"		return MAP
"package"	return PACKAGE
"range"		return RANGE
"return"	return RETURN
"select"	return SELECT
"struct"	return STRUCT
"switch"	return SWITCH
"type"		return TYPE
"var"		return VAR

{float_lit}	return FLOAT_LIT
{identifier}	return IDENTIFIER
{imaginary_lit}	return IMAG_LIT
{int_lit}	return INT_LIT
{rune_lit}	return CHAR_LIT
{string_lit}	return STRING_LIT

{int_lit}p[0-9]*
		return BAD_FLOAT_LIT

%%
		if c, ok := lx.Abort(); ok {
			return c
		}

		goto yyAction
}
