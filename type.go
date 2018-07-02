// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"sort"
	"strings"

	"github.com/cznic/token"
)

var (
	_ Type = (*ArrayType)(nil)
	_ Type = (*ChannelType)(nil)
	_ Type = (*FunctionType)(nil)
	_ Type = (*InterfaceType)(nil)
	_ Type = (*MapType)(nil)
	_ Type = (*NamedType)(nil)
	_ Type = (*PointerType)(nil)
	_ Type = (*QualifiedNamedType)(nil)
	_ Type = (*SliceType)(nil)
	_ Type = (*StructType)(nil)
	_ Type = (*TupleType)(nil)
	_ Type = TypeKind(0)
)

// Type represents a Go type.
type Type interface {
	//TODO Node
	AssignableTo(Type) bool             //TODO Operand method
	ComparableWith(Type) (bool, string) //TODO Operand method
	Identical(Type) bool
	Implements(Type) bool
	IsDefined() bool
	Kind() TypeKind
	Package() *Package
	Position() token.Position
	SourceFile() *SourceFile
	String() string
	UnderlyingType() Type

	alias() Type
	checker
	comparableWith(Type) (bool, string) //TODO Operand method
	defaultType() Type
	err(string, ...interface{})
	err0(token.Position, string, ...interface{})
	has(string, *FunctionType) bool
	identical(map[tpair]sentinel, Type) bool
	identicalInner(map[tpair]sentinel, Type) bool
	implements(t *InterfaceType) bool
	isGeneric(map[Type]bool) bool // Is or refers to BuiltinGeneric*.
	pos() npos
	selector(Token) Type
	setPos(npos)
}

type tpair [2]Type

var (
	typeKinds = [...]string{
		Bool:       "bool",
		Complex128: "complex128",
		Complex64:  "complex64",
		Float32:    "float32",
		Float64:    "float64",
		Int:        "int",
		Int16:      "int16",
		Int32:      "int32",
		Int64:      "int64",
		Int8:       "int8",
		String:     "string",
		Uint:       "uint",
		Uint16:     "uint16",
		Uint32:     "uint32",
		Uint64:     "uint64",
		Uint8:      "uint8",
		Uintptr:    "uintptr",

		UntypedBool:    "untyped bool",
		UntypedComplex: "untyped complex",
		UntypedFloat:   "untyped float",
		UntypedInt:     "untyped int",
		UntypedNil:     "untyped nil",
		UntypedRune:    "untyped rune",
		UntypedString:  "untyped string",

		BuiltinGenericComplex: "generic complex",
		BuiltinGenericFloat:   "generic float",
		BuiltinGenericInteger: "generic int",
		BuiltinGeneric:        "generic type",
		BuiltinGeneric2:       "generic type2",

		Array:             "Array",
		Channel:           "Channel",
		Function:          "Function",
		Interface:         "Interface",
		Map:               "Map",
		Pointer:           "Pointer",
		QualifiedTypeName: "QualifiedTypeName",
		Slice:             "Slice",
		Struct:            "Struct",
		Tuple:             "Tuple",
		TypeName:          "TypeName",
	}

	emptyTuple = &TupleType{}
)

// ChanDir represents channel direction.
type ChanDir int

// Values of type ChanDir
const (
	TxChan ChanDir = 1 << iota
	RxChan
)

// TypeKind represents a particular type kind.
type TypeKind byte

func (n TypeKind) Package() *Package                                { return nil }
func (n TypeKind) SourceFile() *SourceFile                          { return nil }
func (n TypeKind) alias() Type                                      { return n }
func (n TypeKind) check(s *cstack)                                  {}
func (n TypeKind) checked()                                         {}
func (n TypeKind) checking()                                        {}
func (n TypeKind) pos() (r npos)                                    { return r }
func (n TypeKind) has(string, *FunctionType) bool                   { return false }
func (n TypeKind) identical(m map[tpair]sentinel, t Type) bool      { return t != nil && n == t }
func (n TypeKind) identicalInner(m map[tpair]sentinel, t Type) bool { return n.identical(m, t) }
func (n TypeKind) implements(*InterfaceType) bool                   { return false }
func (n TypeKind) setPos(pos npos)                                  {}
func (n TypeKind) state() sentinel                                  { return checked }
func (n TypeKind) err(string, ...interface{})                       { panic(fmt.Errorf("%T.err: unexpected call", n)) }

func (n TypeKind) err0(token.Position, string, ...interface{}) {
	panic(fmt.Errorf("%T.err: unexpected call", n))
}

func (n TypeKind) ComparableWith(t Type) (bool, string) {
	if n == Invalid || t == nil || t == Invalid {
		return true, ""
	}

	if t = t.alias(); t == nil || t == Invalid {
		return true, ""
	}

	var x *typ

	if !x.assignableTo(n, t) && !x.assignableTo(t, n) {
		return false, fmt.Sprintf("mismatched types %s and %s", n, t)
	}

	return n.comparableWith(t)
}

func (n TypeKind) comparableWith(t Type) (bool, string) {
	if n == Invalid || t == Invalid {
		return true, ""
	}

	switch t := t.(type) {
	case TypeKind:
		switch n {
		case Bool:
			switch t {
			case Bool, UntypedBool:
				return true, ""
			}
		case Float32:
			switch t {
			case Float32, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}
		case Float64:
			switch t {
			case Float64, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}
		case Int:
			switch t {
			case Int, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}
		case Int8:
			switch t {
			case Int8, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}
		case Int16:
			switch t {
			case Int16, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}
		case Int32:
			switch t {
			case Int32, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}
		case Int64:
			switch t {
			case Int64, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}
		case Uint:
			switch t {
			case Uint, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}
		case Uint8:
			switch t {
			case Uint8, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}
		case Uint16:
			switch t {
			case Uint16, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}
		case Uint32:
			switch t {
			case Uint32, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}
		case Uint64:
			switch t {
			case Uint64, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}
		case Uintptr:
			switch t {
			case Uintptr, UntypedComplex, UntypedFloat, UntypedInt, UntypedRune:
				return true, ""
			}

		case Function:
			switch t {
			case UntypedNil:
				return true, ""
			}
		case Interface:
			switch t {
			case Interface, UntypedNil:
				return true, ""
			}
		case Map:
			switch t {
			case UntypedNil:
				return true, ""
			}
		case Pointer:
			switch t {
			case Pointer, UntypedNil:
				return true, ""
			}
		case Slice:
			switch t {
			case UntypedNil:
				return true, ""
			}
		case String:
			switch t {
			case String, UntypedString:
				return true, ""
			}
		case UntypedNil:
			switch t {
			case Pointer, Slice, Map, Function, Interface:
				return true, ""
			}
		case UntypedInt, UntypedFloat, UntypedComplex, UntypedRune:
			switch t {
			case UntypedInt, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr, UntypedComplex, UntypedFloat, UntypedRune:
				return true, ""
			}
		case UntypedString:
			switch t {
			case UntypedString, String:
				return true, ""
			}
		default:
			//dbg("\n%s", stack())
			panic(fmt.Sprintf("TODO271 unexpected type %v", n))
		}
	//TODO case *ArrayType:
	//TODO case *ChannelType:
	//TODO case *FunctionType:
	case *InterfaceType:
		return t.comparableWith(n)
	//TODO case *MapType:
	//TODO case *NamedType:
	//TODO case *PointerType:
	//TODO case *QualifiedNamedType:
	//TODO case *SliceType:
	//TODO case *StructType:
	//TODO case *TupleType:
	default:
		//dbg("\n%s", stack())
		panic(fmt.Sprintf("TODO276 unexpected type %v", t))
	}
	return false, ""
}

func (n TypeKind) defaultType() (r Type) {
	switch n {
	case UntypedBool:
		return Bool
	case UntypedComplex:
		return Complex128
	case UntypedFloat:
		return Float64
	case UntypedInt:
		return Int
	case UntypedRune:
		return Int32
	case UntypedString:
		return String
	default:
		return n
	}
}

func (n TypeKind) selector(nm Token) Type { return invalidType }

// IsDefined implements Type.
func (n TypeKind) IsDefined() bool { return true }

// Implements implements Type.
func (n TypeKind) Implements(t Type) bool {
	x, ok := t.UnderlyingType().(*InterfaceType)
	return ok && len(x.Methods) == 0
}

func (n TypeKind) isGeneric(map[Type]bool) bool {
	switch n {
	case
		BuiltinGenericComplex,
		BuiltinGenericFloat,
		BuiltinGenericInteger,
		BuiltinGeneric,
		BuiltinGeneric2:

		return true
	}

	return false
}

// AssignableTo implements Type.
func (n TypeKind) AssignableTo(t Type) bool { return (*typ)(nil).assignableTo(n, t) }

// Identical implements Type.
func (n TypeKind) Identical(t Type) bool {
	if t == nil || t == Invalid || n == t {
		return true
	}

	t = t.alias()
	return t == nil || t == Invalid || n == t
}

// Position implements Type.
func (n TypeKind) Position() (r token.Position) { return r }

// UnderlyingType implements Type.
func (n TypeKind) UnderlyingType() Type { return n }

// Kind implements Type.
func (n TypeKind) Kind() TypeKind { return n }

func (n TypeKind) String() string {
	if s := typeKinds[n]; s != "" {
		return s
	}

	return fmt.Sprintf("%T(%d)", n, int(n))
}

// TypeKind values.
const (
	Invalid TypeKind = iota

	Bool
	Complex128
	Complex64
	Float32
	Float64
	Int
	Int16
	Int32
	Int64
	Int8
	String
	Uint
	Uint16
	Uint32
	Uint64
	Uint8
	Uintptr

	UntypedBool
	UntypedComplex
	UntypedFloat
	UntypedInt
	UntypedNil
	UntypedRune
	UntypedString

	BuiltinGenericComplex // U in 'func complex(r, i T) U'.
	BuiltinGenericFloat   // T in 'func complex(r, i T) U'.
	BuiltinGenericInteger // U in 'make(T, ...U)'.
	BuiltinGeneric        // T in 'func make(T)'.
	BuiltinGeneric2       // U in 'func delete(map[T]U)'.

	Array             // [Expr]Element
	Channel           // chan T
	Function          // func(params) result
	Interface         // interface{...}
	Map               // map[T]U
	Pointer           // *T
	QualifiedTypeName // qualifier.ident
	Slice             // []T
	Struct            // struct{...}
	Tuple             // (T1, T2, ...)
	TypeName          // ident

	maxTypeKind
)

type typ struct {
	node
	outer Type
}

func (n *typ) AssignableTo(t Type) bool {
	//TODO if t := n.outer; t.state() == 0 {
	//TODO 	panic(fmt.Errorf("%v: %[2]p %[2]T %[2]v TODO371\n%[3]s\n====", t.Position(), t, stack()))
	//TODO }

	return t != nil && n.assignableTo(n.outer.alias(), t.alias())
}
func (n *typ) IsDefined() bool   { return false }
func (n *typ) alias() Type       { return n.outer }
func (n *typ) defaultType() Type { return n.outer }
func (n *typ) state() sentinel   { return n.sentinel }

func (n *typ) ComparableWith(t Type) (r bool, _ string) {
	//TODO if n.state() == 0 {
	//TODO 	panic(fmt.Errorf("TODO388\n%s", stack()))
	//TODO }

	if t == nil || t == Invalid {
		return true, ""
	}

	x := n.outer.alias()
	t0 := t
	if t = t.alias(); x == nil || t == nil {
		return true, ""
	}

	// dbg("Comparable(%T(%s), %T(%s))", x, x, t, t)
	// defer func() { dbg("\tComparable(%s, %s) %v", x, t, r) }()

	if !n.assignableTo(x, t) && !n.assignableTo(t, x) {
		return false, fmt.Sprintf("mismatched types %s and %s", n.outer, t0)
	}

	if x.Kind() != Interface && t.Kind() == Interface && x.Implements(t) {
		return true, ""
	}

	if t.Kind() != Interface && x.Kind() == Interface && t.Implements(x) {
		return true, ""
	}

	return x.comparableWith(t)
}

func (n *typ) Identical(t Type) bool { return n.identicalInner(nil, t) }

func (n *typ) identicalInner(m map[tpair]sentinel, t Type) bool {
	//TODO if n.state() == 0 {
	//TODO 	panic(fmt.Errorf("TODO409\n%s", stack()))
	//TODO }

	if t == nil || t == Invalid {
		return true
	}

	x := n.outer.alias()
	t = t.alias()
	return x == nil || t == nil || x.identical(m, t)
}

func (n *typ) Implements(t Type) (r bool) {
	//TODO if n.state() == 0 {
	//TODO 	panic(fmt.Errorf("TODO423\n%s", stack()))
	//TODO }

	if t == nil || t == Invalid {
		return false
	}

	x := n.outer.alias()
	t = t.alias()
	if t == nil || t == Invalid {
		return true
	}

	t = t.UnderlyingType()
	if t == nil || t == Invalid || t.Kind() != Interface {
		return false
	}

	//dbg("Implements(%T(%s), %T(%s))", x, x, t, t)
	//defer func() { dbg("\tImplements(%s, %s) %v", x, t, r) }()

	return x == nil || x == Invalid || t == nil || t == Invalid || x.implements(t.(*InterfaceType))
}

func (n *typ) has(nm string, f *FunctionType) bool {
	//TODO if n.state() == 0 {
	//TODO 	panic(fmt.Errorf("TODO437\n%s", stack()))
	//TODO }

	if t := n.outer.alias(); t != n.outer {
		return t.has(nm, f)
	}

	return false
}

func (n *typ) selector(nm Token) Type {
	panic("TODO268")
	return invalidType
}

func (n *typ) assignableTo(x, t Type) (r bool) {
	if x == nil || x == Invalid || t == nil || t == Invalid {
		return true
	}

	//TODO if n != nil && n.state() == 0 {
	//TODO 	panic(fmt.Errorf("TODO454\n%s", stack()))
	//TODO }

	xu := x.UnderlyingType()
	tu := t.UnderlyingType()

	//dbg("assignableTo(%T(%s), %T(%s))", x, x, t, t)
	//defer func() { dbg("\tassignableTo(%s, %s) %v", x, t, r) }()

	// A value x is assignable to a variable of type T ("x is assignable to
	// T") if one of the following conditions applies:

	// x's type is identical to T.
	if x.Identical(t) {
		return true
	}

	// x's type V and T have identical underlying types and at least one of
	// V or T is not a defined type.
	if xu.Identical(tu) && (!x.IsDefined() || !t.IsDefined()) {
		return true
	}

	// T is an interface type and x implements T.
	if tu.Kind() == Interface && x.Implements(tu) {
		return true
	}

	// x is a bidirectional channel value, T is a channel type, x's type V
	// and T have identical element types, and at least one of V or T is
	// not a defined type.
	if c, ok := x.(*ChannelType); ok {
		if u, ok := t.(*ChannelType); ok &&
			c.Dir == TxChan|RxChan &&
			c.Element.Identical(u.Element) &&
			(!x.IsDefined() || !t.IsDefined()) {
			return true
		}
	}

	// x is the predeclared identifier nil and T is a pointer, function,
	// slice, map, channel, or interface type.
	if x == UntypedNil {
		switch tu.Kind() {
		case Pointer, Function, Slice, Map, Channel, Interface:
			return true
		}
	}

	// x is an untyped constant representable by a value of type T.
	switch x.Kind() {
	case UntypedBool:
		return tu == Bool
	case UntypedComplex:
		return tu == Complex64 || tu == Complex128
	case UntypedFloat:
		return tu == Float32 || tu == Float64
	case UntypedInt, UntypedRune:
		switch tu {
		case
			Complex128,
			Complex64,
			Float32,
			Float64,
			Int,
			Int16,
			Int32,
			Int64,
			Int8,
			Uint,
			Uint16,
			Uint32,
			Uint64,
			Uint8,
			Uintptr:

			return true //TODO check representable
		}
	case UntypedString:
		return tu == String
	}

	return false
}

func (n *typ) implements(t *InterfaceType) bool { return len(t.Methods) == 0 }

func (n *typ) UnderlyingType() Type { return n.outer }

// ArrayType represents '[Expr]Element'.
type ArrayType struct {
	Expr    Expr
	Element Type

	typ
}

func newArrayType(pos npos, expr Expr, elem Type) *ArrayType {
	n := &ArrayType{Expr: expr, Element: elem}
	n.npos = pos
	n.outer = n
	return n
}

func (n *ArrayType) Kind() TypeKind { return Array }
func (n *ArrayType) String() string { return fmt.Sprintf("[%s]%s", n.Expr.str(), n.Element) } //TODO use const value
func (n *ArrayType) str() string    { return fmt.Sprintf("[%s]%s", n.Expr.str(), n.Element) }

func (n *ArrayType) check(s *cstack) {
	s.check(n.Expr, n)
	s.check(n.Element, n)
	//TODO set size
}

func (n *ArrayType) comparableWith(t Type) (bool, string) {
	return n.Element.ComparableWith(t.(*ArrayType).Element)
}

func (n *ArrayType) setPos(pos npos) {
	if n.npos.sf != nil {
		return
	}

	n.npos = pos
	n.Expr.Type().setPos(pos)
	n.Element.setPos(pos)
}

// Two array types are identical if they have identical element types and the
// same array length.
func (n *ArrayType) identical(m map[tpair]sentinel, t Type) bool {
	if n == t {
		return true
	}

	x, ok := t.(*ArrayType)
	return ok && n.Element.identicalInner(m, x.Element) //TODO && n.Expr.Value() != nil && x.Expr.Value() != nil && n.Expr.Value().equal(x.Expr.Value())
}

func (n *ArrayType) isGeneric(m map[Type]bool) bool {
	if m[n] {
		return false
	}

	m[n] = true
	return n.Element.isGeneric(m)
}

// SliceType represents '[]Element'.
type SliceType struct {
	Element Type

	typ
}

func newSliceType(pos npos, elem Type) *SliceType {
	n := &SliceType{Element: elem}
	n.npos = pos
	n.outer = n
	return n
}

func (n *SliceType) Kind() TypeKind { return Slice }
func (n *SliceType) String() string { return fmt.Sprintf("[]%s", n.Element) }

func (n *SliceType) check(s *cstack) {
	n.sf.enqueue(n.Element)
}

func (n *SliceType) comparableWith(t Type) (bool, string) {
	// Slice values are not comparable. However, as a special case, a slice
	// value may be compared to the predeclared identifier nil.
	if t == UntypedNil {
		return true, ""
	}

	return false, "slice can only be compared to nil"
}

func (n *SliceType) setPos(pos npos) {
	if n.npos.sf != nil {
		return
	}

	n.npos = pos
	n.Element.setPos(pos)
}

// Two slice types are identical if they have identical element types.
func (n *SliceType) identical(m map[tpair]sentinel, t Type) bool {
	if n == t {
		return true
	}

	x, ok := t.(*SliceType)
	return ok && n.Element.identicalInner(m, x.Element)
}

func (n *SliceType) isGeneric(m map[Type]bool) bool {
	if m[n] {
		return false
	}

	m[n] = true
	return n.Element.isGeneric(m)
}

// FunctionType represents 'func(Params) Result'.
type FunctionType struct {
	ParameterList ParamTypeList
	ResultList    ParamTypeList

	typ

	Variadic bool
}

func newFunctionType(pos npos, params, result ParamTypeList) *FunctionType {
	if params == nil {
		params = emptyParamTypeList
	}
	if result == nil {
		result = emptyParamTypeList
	}
	n := &FunctionType{ParameterList: params, ResultList: result}
	n.npos = pos
	n.outer = n
	switch x := params.(type) {
	case *ParamTypeListIdentType:
		n.Variadic = len(x.List) > 0 && x.List[len(x.List)-1].Variadic
	case *ParamTypeListType:
		n.Variadic = len(x.List) > 0 && x.List[len(x.List)-1].Variadic
	default:
		panic(fmt.Sprintf("unexpected type %T", x))
	}
	return n
}

func (n *FunctionType) Kind() TypeKind     { return Function }
func (n *FunctionType) String() (r string) { return fmt.Sprintf("func%s", n.str()) }

func (n *FunctionType) check(s *cstack) {
	sf := n.sf
	switch x := n.ParameterList.(type) {
	case *ParamTypeListIdentType:
		for _, v := range x.List {
			sf.enqueue(v.Type)
		}
	case *ParamTypeListType:
		for _, v := range x.List {
			sf.enqueue(v.Type)
		}
	}
	switch x := n.ResultList.(type) {
	case *ParamTypeListIdentType:
		for _, v := range x.List {
			sf.enqueue(v.Type)
		}
	case *ParamTypeListType:
		for _, v := range x.List {
			sf.enqueue(v.Type)
		}
	}
}

func (n *FunctionType) comparableWith(t Type) (bool, string) {
	// Function values are not comparable. However, as a special case, a
	// function value may be compared to the predeclared identifier nil.
	if t == UntypedNil {
		return true, ""
	}

	return false, "func can only be compared to nil"
}

// Parameters returns the TupleType corresponding to n.ParametersList
func (n *FunctionType) Parameters() Type {
	t := n.ParameterList.tuple()
	if t.npos.sf == nil {
		t.setPos(n.npos)
	}
	return t
}

func (n *FunctionType) setPos(pos npos) {
	if n.npos.sf != nil {
		return
	}

	n.npos = pos
	switch x := n.ParameterList.(type) {
	case *ParamTypeListIdentType:
		for _, v := range x.List {
			v.Type.setPos(pos)
		}
	case *ParamTypeListType:
		for _, v := range x.List {
			v.Type.setPos(pos)
		}
	}
	switch x := n.ResultList.(type) {
	case *ParamTypeListIdentType:
		for _, v := range x.List {
			v.Type.setPos(pos)
		}
	case *ParamTypeListType:
		for _, v := range x.List {
			v.Type.setPos(pos)
		}
	}
}

// Result returns n's ResultList type as a TupleType or a non-tuple type if n
// has exactly one result.
func (n *FunctionType) Result() Type {
	l := n.ResultList.tuple()
	if len(l.List) == 1 {
		t := l.List[0]
		if p := t.Position(); !p.IsValid() {
			t.setPos(n.npos)
		}
		return t
	}

	if l.npos.sf == nil {
		l.setPos(n.npos)
	}
	return l
}

func (n *FunctionType) isGeneric(m map[Type]bool) bool {
	if m[n] {
		return false
	}

	m[n] = true
	return n.Parameters().isGeneric(m) || n.Result().isGeneric(m)
}

// Two function types are identical if they have the same number of parameters
// and result values, corresponding parameter and result types are identical,
// and either both functions are variadic or neither is. Parameter and result
// names are not required to match.
func (n *FunctionType) identical(m map[tpair]sentinel, t Type) bool {
	if n == t {
		return true
	}

	x, ok := t.(*FunctionType)
	if !ok {
		return false
	}

	return n.Variadic == x.Variadic &&
		n.ParameterList.len() == x.ParameterList.len() &&
		n.ResultList.len() == x.ResultList.len() &&
		n.Parameters().identicalInner(m, x.Parameters()) &&
		n.Result().identicalInner(m, x.Result())
}

func (n *FunctionType) str() string {
	switch n.ResultList.len() {
	case 0:
		return fmt.Sprintf("(%s)", n.ParameterList.types())
	case 1:
		return fmt.Sprintf("(%s) %s", n.ParameterList.types(), n.Result())
	}

	return fmt.Sprintf("(%s) %s", n.ParameterList.types(), n.Result())
}

// MapType represents 'map[Key]Value'.
type MapType struct {
	Key   Type
	Value Type

	typ
}

func newMapType(pos npos, key, val Type) *MapType {
	n := &MapType{Key: key, Value: val}
	n.npos = pos
	n.outer = n
	return n
}

func (n *MapType) Kind() TypeKind { return Map }
func (n *MapType) String() string { return fmt.Sprintf("map[%s]%s", n.Key, n.Value) }

func (n *MapType) check(s *cstack) {
	n.sf.enqueue(n.Key)
	n.sf.enqueue(n.Value)
}

func (n *MapType) comparableWith(t Type) (bool, string) {
	// Map values are not comparable. However, as a special case, a map
	// value may be compared to the predeclared identifier nil.
	if t == UntypedNil {
		return true, ""
	}

	return false, "map can only be compared to nil"
}

func (n *MapType) setPos(pos npos) {
	if n.npos.sf != nil {
		return
	}

	n.npos = pos
	n.Key.setPos(pos)
	n.Value.setPos(pos)
}

// Two map types are identical if they have identical key and element types.
func (n *MapType) identical(m map[tpair]sentinel, t Type) bool {
	if n == t {
		return true
	}

	x, ok := t.(*MapType)
	return ok && n.Key.identicalInner(m, x.Key) && n.Value.identicalInner(m, x.Value)
}

func (n *MapType) isGeneric(m map[Type]bool) bool {
	if m[n] {
		return false
	}

	m[n] = true
	return n.Key.isGeneric(m) || n.Value.isGeneric(m)
}

// PointerType represents '*Element'.
type PointerType struct {
	Element Type

	typ
}

func newPointerType(pos npos, elem Type) *PointerType {
	n := &PointerType{Element: elem}
	n.npos = pos
	n.outer = n
	return n
}

func (n *PointerType) Kind() TypeKind { return Pointer }
func (n *PointerType) String() string { return fmt.Sprintf("*%s", n.Element) }

func (n *PointerType) check(s *cstack) {
	n.sf.enqueue(n.Element)
}

func (n *PointerType) comparableWith(t Type) (bool, string) {
	return n.Kind().comparableWith(t.UnderlyingType().Kind())
}

func (n *PointerType) setPos(pos npos) {
	if n.npos.sf != nil {
		return
	}

	n.npos = pos
	n.Element.setPos(pos)
}

func (n *PointerType) selector(nm Token) Type {
	t := n.Element.alias()
	if t.Kind() == Pointer {
		panic("TODO612")
	}

	if t.Kind() == Interface {
		panic("TODO617")
	}

	return t.selector(nm)
}

func (n *PointerType) isGeneric(m map[Type]bool) bool {
	if m[n] {
		return false
	}

	m[n] = true
	return n.Element.isGeneric(m)
}

// Two pointer types are identical if they have identical base types.
func (n *PointerType) identical(m map[tpair]sentinel, t Type) bool {
	if n == t {
		return true
	}

	x, ok := t.(*PointerType)
	return ok && n.Element.identicalInner(m, x.Element)
}

func (n *PointerType) has(nm string, f *FunctionType) bool {
	return true //TODO-
}

func (n *PointerType) implements(t *InterfaceType) bool {
	return true //TODO-
}

// ChannelType represents 'chan Element'.
type ChannelType struct {
	Element Type

	typ

	Dir ChanDir
}

func newChannelType(pos npos, dir ChanDir, elem Type) *ChannelType {
	n := &ChannelType{Dir: dir, Element: elem}
	n.npos = pos
	n.outer = n
	return n
}

func (n *ChannelType) Kind() TypeKind                       { return Channel }
func (n *ChannelType) comparableWith(t Type) (bool, string) { return true, "" }

func (n *ChannelType) check(s *cstack) {
	n.sf.enqueue(n.Element)
}

func (n *ChannelType) setPos(pos npos) {
	if n.npos.sf != nil {
		return
	}

	n.npos = pos
	n.Element.setPos(pos)
}

// Two channel types are identical if they have identical element types and the
// same direction.
func (n *ChannelType) identical(m map[tpair]sentinel, t Type) bool {
	if n == t {
		return true
	}

	x, ok := t.(*ChannelType)
	return ok && n.Element.identicalInner(m, x.Element) && n.Dir == x.Dir
}

func (n *ChannelType) isGeneric(m map[Type]bool) bool {
	if m[n] {
		return false
	}

	m[n] = true
	return n.Element.isGeneric(m)
}

func (n *ChannelType) String() string {
	switch n.Dir {
	case RxChan:
		return fmt.Sprintf("<-chan %s", n.Element)
	case TxChan:
		return fmt.Sprintf("chan<- %s", n.Element)
	default:
		return fmt.Sprintf("chan %s", n.Element)
	}
}

// InterfaceType represents 'interface{...}'.
type InterfaceType struct {
	Embedded []Type
	Methods  map[string]*FunctionType
	methods  []*FunctionType
	names    []string

	typ
}

func newInterfaceType(pos npos) *InterfaceType {
	n := &InterfaceType{}
	n.npos = pos
	n.outer = n
	return n
}

func (n *InterfaceType) Kind() TypeKind { return Interface }

func (n *InterfaceType) check(s *cstack) {
	sf := n.sf
	a := make([]string, 0, len(n.Methods))
	for k := range n.Methods {
		a = append(a, k)
	}
	sort.Strings(a)
	for _, k := range a {
		sf.enqueue(n.Methods[k])
	}
	for _, v := range n.Embedded {
		s.check(v, n)
		u := v.UnderlyingType()
		if u.Kind() != Interface {
			sf.Package.errorList.add(v.Position(), "interface contains embedded non-interface %s", v)
			continue
		}

		if x, ok := v.(*NamedType); ok {
			if td := x.Declaration(); td != nil && td.state() == checking {
				continue
			}
		}

		a = a[:0]
		e := u.(*InterfaceType)
		for k := range e.Methods {
			a = append(a, k)
		}
		sort.Strings(a)
		for _, k := range a {
			if _, ok := n.Methods[k]; ok {
				sf.Package.errorList.add(v.Position(), "duplicate method %s", k)
				continue
			}

			if n.Methods == nil {
				n.Methods = map[string]*FunctionType{}
			}
			n.Methods[k] = e.Methods[k]
		}
	}
}

func (n *InterfaceType) implements(t *InterfaceType) bool {
	for k, v := range t.Methods {
		if !n.has(k, v) {
			return false
		}
	}
	return true
}

func (n *InterfaceType) comparableWith(x Type) (bool, string) {
	t := n
	switch xk := x.UnderlyingType().Kind(); xk {
	case
		Interface,
		UntypedNil:

		return true, ""
	default:
		// A value x of non-interface type X and a value t of interface
		// type T are comparable when values of type X are comparable
		// and X implements T.
		if ok, s := x.comparableWith(x); !ok {
			return false, s
		}

		if !x.Implements(t) {
			return false, ""
		}

		return true, ""
	}
}

func (n *InterfaceType) has(nm string, f *FunctionType) bool {
	t := n.Methods[nm]
	return t != nil && t.Identical(f)
}

func (n *InterfaceType) selector(nm Token) Type {
	if f, ok := n.Methods[nm.Val]; ok {
		return f
	}

	return invalidType
}

func (n *InterfaceType) setPos(pos npos) {
	if n.npos.sf != nil {
		return
	}

	n.npos = pos
	for _, v := range n.Embedded {
		v.setPos(pos)
	}
	for _, v := range n.Methods {
		v.setPos(pos)
	}
}

// Two interface types are identical if they have the same set of methods with
// the same names and identical function types. Non-exported method names from
// different packages are always different. The order of the methods is
// irrelevant.
func (n *InterfaceType) identical(m map[tpair]sentinel, t Type) (r bool) {
	if n == t {
		return true
	}

	if m == nil {
		m = map[tpair]sentinel{}
	}
	k := tpair{n, t}
	switch m[k] {
	case 0:
		m[k] = checking

		defer func() {
			if r {
				m[k] = checkedTrue
				return
			}

			m[k] = checkedFalse
		}()
	case checking:
		m[k] = checkedTrue
		return true
	case checkedFalse:
		return false
	case checkedTrue:
		return true
	default:
		panic(fmt.Sprintf("unexpected value %v", m[k]))
	}

	x, ok := t.(*InterfaceType)
	if !ok {
		return false
	}

	if len(n.Methods) != len(x.Methods) {
		return false
	}

	for k, v := range n.Methods {
		w, ok := x.Methods[k]
		if !ok {
			return false
		}

		if !v.identicalInner(m, w) {
			return false
		}

		if !isExported(k) && v.sf.Package != w.sf.Package {
			return false
		}
	}
	return true
}

func (n *InterfaceType) isGeneric(m map[Type]bool) bool {
	if m[n] {
		return false
	}

	m[n] = true
	for _, v := range n.Methods {
		if v.isGeneric(m) {
			return true
		}
	}
	return false
}

func (n *InterfaceType) String() string {
	var a []string
	for _, v := range n.Embedded {
		a = append(a, v.String())
	}
	for i, v := range n.methods {
		a = append(a, fmt.Sprintf("%s%s", n.names[i], v.str()))
	}
	return fmt.Sprintf("interface{%s}", strings.Join(a, "; "))
}

// NamedType represents 'identifier type'.
type NamedType struct {
	Name  string
	scope *Scope

	typ
}

func newNamedType(id Ident) *NamedType {
	n := &NamedType{Name: id.Val, scope: id.Scope()}
	n.npos = id.npos
	n.outer = n
	return n
}

func (n *NamedType) Kind() TypeKind { return TypeName }
func (n *NamedType) Scope() *Scope  { return n.scope }

func (n *NamedType) check(s *cstack) {
	sf := n.SourceFile()
	d := n.Scope().Lookup(sf.Package, sf.Scope, Token{Val: n.Name, npos: n.npos})
	if d == nil {
		sf.Package.errorList.add(n.Position(), "undefined: %v", n)
		return
	}

	td, ok := d.(*TypeDecl)
	if !ok {
		sf.Package.errorList.add(n.Position(), "%s is not a type", n)
		return
	}

	if td.SourceFile().Package == nil {
		panic("TODO987")
	}
	if td.SourceFile().Package == s.p {
		s.check(td, n)
	}
	//TODO set size
}

func (n *NamedType) selector(nm Token) Type {
	if t := n.alias(); t != n {
		return t.selector(nm)
	}

	switch t := n.Type(); t.(type) {
	case nil:
		return nil
	case *StructType:
		if r := t.selector(nm); r != nil && r != Invalid {
			return r
		}
	case *InterfaceType:
		if r := t.selector(nm); r != nil && r != Invalid {
			return r
		}
	}

	m := n.Package().Methods
	if x, ok := m[n.Name][nm.Val].(*MethodDecl); ok {
		return x.Type()
	}

	return invalidType
}

func (n *NamedType) comparableWith(t Type) (bool, string) {
	u := t.UnderlyingType()
	switch x := n.UnderlyingType().(type) {
	case *StructType:
		return x.comparableWith(u)
	default:
		return x.comparableWith(u)
	}
}

func (n *NamedType) String() string {
	d := n.Declaration()
	if d != nil {
		if pkg := n.Package(); pkg != nil {
			if b := n.SourceFile().Scope.Bindings; b != nil {
				if x, ok := b[n.Name].(*TypeDecl); ok {
					return fmt.Sprintf("%s.%s", x.SourceFile().Package.Name, n.Name)
				}
			}
		}
	}

	return n.Name
}

func (n *NamedType) alias() (r Type) {
	d := n.Declaration()
	if d == nil {
		return n
	}

	if d.alias {
		return d.Type().alias()
	}

	if d.universe {
		return d.Type()
	}

	return n
}

func (n *NamedType) setPos(pos npos) {
	if n.SourceFile() != nil {
		return
	}

	n.npos = pos
}

func (n *NamedType) isGeneric(m map[Type]bool) bool {
	if m[n] {
		return false
	}

	m[n] = true
	return n.UnderlyingType().isGeneric(m)
}

// IsDefined implements Type.
func (n *NamedType) IsDefined() bool {
	if t := n.alias(); t != n {
		return t.IsDefined()
	}

	td := n.Declaration()
	return td != nil && !td.universe
}

// A defined type is always different from any other type.
func (n *NamedType) identical(m map[tpair]sentinel, t Type) (r bool) {
	if n == t {
		return true
	}

	var td *TypeDecl
	switch x := t.(type) {
	case *NamedType:
		td = x.Declaration()
	case *QualifiedNamedType:
		td = x.Declaration()
	default:
		return false
	}

	nd := n.Declaration()
	return nd != nil && td != nil && nd == td
}

func (n *NamedType) has(nm string, f *FunctionType) bool {
	if a := n.alias(); a != n {
		return a.has(nm, f)
	}

	if m := n.Package().Methods[n.Name][nm]; m != nil {
		return m.(*MethodDecl).Type().Identical(f)
	}

	return n.Type().has(nm, f)
}

func (n *NamedType) implements(t *InterfaceType) (r bool) {
	// dbg("%T(%s)implements(%s)", n, n, t)
	// defer func() { dbg("%T(%s)implements(%s) %v", n, n, t, r) }()

	if a := n.alias(); a != n {
		return a.implements(t)
	}

	for k, v := range t.Methods {
		if !n.has(k, v) {
			return false
		}
	}
	return true
}

// Type returns U in type T U.
func (n *NamedType) Type() Type {
	if d := n.Declaration(); d != nil {
		return d.Type()
	}

	return nil
}

// Declaration returns the declaration of n.
func (n *NamedType) Declaration() *TypeDecl {
	sf := n.SourceFile()
	if sf == nil {
		return nil
	}

	d := n.Scope().Lookup(sf.Package, sf.Scope, Token{Val: n.Name, npos: n.npos})
	if d == nil {
		return nil
	}

	if td, ok := d.(*TypeDecl); ok {
		return td
	}

	return nil
}

// UnderlyingType implements Type.
func (n *NamedType) UnderlyingType() Type {
	//TODO if n.state() == 0 {
	//TODO 	var s string
	//TODO 	s = fmt.Sprintf("%p %v: %s %p %s", n, n.Position(), n, &n.sentinel, n.sentinel)
	//TODO 	panic(fmt.Errorf("%s\n%v: %q %p TODO1286\n%s\n====", s, n.Position(), n.Ident.Val, n, stack()))
	//TODO }
	if t := n.alias(); t != n {
		return t.UnderlyingType()
	}

	if td := n.Declaration(); td != nil {
		return td.Type().UnderlyingType()
	}

	return invalidType
}

// QualifiedNamedType represents 'qualifier.identifier'.
type QualifiedNamedType struct {
	scope     *Scope
	Qualifier string
	Name      string

	typ
}

func newQualifiedNamedType(q Ident, nm string) *QualifiedNamedType {
	n := &QualifiedNamedType{Qualifier: q.Val, Name: nm, scope: q.Scope()}
	n.npos = q.npos
	n.outer = n
	return n
}

func (n *QualifiedNamedType) Kind() TypeKind { return QualifiedTypeName }
func (n *QualifiedNamedType) Scope() *Scope  { return n.scope }

func (n *QualifiedNamedType) check(s *cstack) {
	if n.Qualifier == "C" { //TODO
		return
	}

	sf := n.SourceFile()
	d := n.Scope().Lookup(sf.Package, sf.Scope, Token{Val: n.Qualifier, npos: n.npos})
	if d == nil {
		sf.Package.errorList.add(n.Position(), "undefined: %s", n.Qualifier)
		return
	}

	id, ok := d.(*ImportDecl)
	if !ok {
		//TODO ~/go/test/fixedbugs/issue18392.go:13: TODO510
		sf.Package.errorList.add(n.Position(), "XTODO510") // not a package qualifier
		return
	}

	if !isExported(n.Name) {
		sf.Package.errorList.add(n.Position(), "cannot refer to unexported name %s", n)
	}

	p := id.Package()
	d = p.Scope.Lookup(p, nil, Token{Val: n.Name, npos: n.npos})
	if d == nil {
		sf.Package.errorList.add(n.Position(), "undefined: %s", n)
		return
	}

	td, ok := d.(*TypeDecl)
	if !ok {
		sf.Package.errorList.add(n.Position(), "%s is not a type", n)
		return
	}

	if td.Type().state() == checked {
		//TODO set size
	}
}

func (n *QualifiedNamedType) comparableWith(t Type) (bool, string) { return true, "" }

func (n *QualifiedNamedType) has(nm string, f *FunctionType) bool {
	if a := n.alias(); a != n {
		return a.has(nm, f)
	}

	td := n.Declaration()
	if td == nil {
		return false
	}

	if m := td.Token.SourceFile().Package.Methods[n.Name][nm]; m != nil {
		return m.(*MethodDecl).Type().Identical(f)
	}

	return n.Type().has(nm, f)
}

func (n *QualifiedNamedType) implements(t *InterfaceType) bool {
	if a := n.alias(); a != n {
		return a.implements(t)
	}

	for k, v := range t.Methods {
		if !n.has(k, v) {
			return false
		}
	}
	return true
}

func (n *QualifiedNamedType) selector(nm Token) Type {
	if t := n.alias(); t != n {
		return t.selector(nm)
	}

	switch t := n.Type(); t.(type) {
	case *StructType:
		if r := t.selector(nm); r != nil {
			return r
		}
	case *InterfaceType:
		if r := t.selector(nm); r != nil {
			return r
		}
	}

	td := n.Declaration()
	if td == nil {
		return invalidType
	}

	m := td.Token.SourceFile().Package.Methods
	if x, ok := m[n.Name][nm.Val].(*MethodDecl); ok {
		return x.Type()
	}

	return invalidType
}

func (n *QualifiedNamedType) alias() Type {
	d := n.Declaration()
	if d == nil {
		return n
	}

	if d.alias {
		return d.Type().alias()
	}

	return n
}

func (n *QualifiedNamedType) isGeneric(m map[Type]bool) bool {
	if m[n] {
		return false
	}

	m[n] = true
	return n.UnderlyingType().isGeneric(m)
}

// A defined type is always different from any other type.
func (n *QualifiedNamedType) identical(m map[tpair]sentinel, t Type) (r bool) {
	if n == t {
		return true
	}

	var td *TypeDecl
	switch x := t.(type) {
	case *NamedType:
		td = x.Declaration()
	case *QualifiedNamedType:
		if n.Qualifier == "C" && x.Qualifier == "C" { //TODO-
			return n.Name == x.Name
		}

		td = x.Declaration()
	default:
		return false
	}

	nd := n.Declaration()
	return nd != nil && td != nil && nd == td
}

// IsDefined implements Type.
func (n *QualifiedNamedType) IsDefined() bool {
	if t := n.alias(); t != n {
		return t.IsDefined()
	}

	td := n.Declaration()
	return td != nil && !td.universe
}

// Type returns U in type T U.
func (n *QualifiedNamedType) Type() Type {
	if d := n.Declaration(); d != nil {
		return d.Type()
	}

	return nil
}

// Declaration returns the declaration of n.
func (n *QualifiedNamedType) Declaration() *TypeDecl {
	sf := n.SourceFile()
	d := n.Scope().Lookup(sf.Package, sf.Scope, Token{Val: n.Qualifier, npos: n.npos})
	if d == nil {
		return nil
	}

	id, ok := d.(*ImportDecl)
	if !ok {
		return nil
	}

	if !isExported(n.Name) {
		return nil
	}

	p := id.Package()
	d = p.Scope.Lookup(p, nil, Token{Val: n.Name, npos: n.npos})
	if td, ok := d.(*TypeDecl); ok {
		return td
	}

	return nil
}

// UnderlyingType implements Type.
func (n *QualifiedNamedType) UnderlyingType() Type {
	if t := n.alias(); t != n {
		return t.UnderlyingType()
	}

	if td := n.Declaration(); td != nil {
		return td.Type().UnderlyingType()
	}

	return invalidType
}

func (n *QualifiedNamedType) setPos(pos npos) {
	if n.SourceFile() != nil {
		return
	}

	n.npos = pos
}

func (n *QualifiedNamedType) String() string {
	if n.Qualifier == "" {
		return n.Name
	}

	return fmt.Sprintf("%s.%s", n.Qualifier, n.Name)
}

// StructType represents 'struct{...}'.
type StructType struct {
	Fields []FieldDeclaration

	typ
}

func newStructType(pos npos) *StructType {
	n := &StructType{}
	n.npos = pos
	n.outer = n
	return n
}

func (n *StructType) Kind() TypeKind { return Struct }
func (n *StructType) str() string    { return "struct{...}" }

func (n *StructType) check(s *cstack) {
	for i := range n.Fields {
		s.check(n.fld(i), n)
	}
}

func (n *StructType) fld(i int) Type {
	switch x := n.Fields[i].(type) {
	case *FieldDeclarationEmbedded:
		return x.Type
	case *FieldDeclarationNamed:
		return x.Type
	default:
		panic(fmt.Errorf("unexpected type %T", x))
	}
}

func (n *StructType) has(nm string, f *FunctionType) bool {
	for i := range n.Fields {
		x := n.fld(i)
		if x == nil || x == Invalid {
			continue
		}

		if x.has(nm, f) {
			return true
		}
	}
	return false
}

func (n *StructType) comparableWith(t Type) (bool, string) {
	x := t.(*StructType)
	for i := range n.Fields {
		a := n.fld(i)
		if a == nil || a == Invalid {
			continue
		}

		b := x.fld(i)
		if b == nil || b == Invalid {
			continue
		}

		if ok, _ := a.ComparableWith(b); !ok {
			return false, fmt.Sprintf("struct containing %s cannot be compared", t)
		}
	}
	return true, ""
}

func (n *StructType) setPos(pos npos) {
	if n.npos.sf != nil {
		return
	}

	n.npos = pos
}

func (n *StructType) implements(t *InterfaceType) bool {
outer:
	for nm, ft := range t.Methods {
		for _, fld := range n.Fields {
			if efld, ok := fld.(*FieldDeclarationEmbedded); ok {
				if efld.Type.has(nm, ft) {
					continue outer
				}
			}
		}

		return false
	}
	return true
}

func (n *StructType) selector(nm Token) Type {
	for _, v := range n.Fields {
		switch x := v.(type) {
		case *FieldDeclarationEmbedded:
			if x.Type.String() == nm.Val {
				return x.Type
			}
			//TODO deep search
		case *FieldDeclarationNamed:
			if x.Name.Val == nm.Val {
				return x.Type
			}
		default:
			panic(fmt.Sprintf("unexpected type %T", x))
		}
	}
	return invalidType
}

// Two struct types are identical if they have the same sequence of fields, and
// if corresponding fields have the same names, and identical types, and
// identical tags. Non-exported field names from different packages are always
// different.
func (n *StructType) identical(m map[tpair]sentinel, t Type) bool {
	if n == t {
		return true
	}

	x, ok := t.(*StructType)
	if !ok {
		return false
	}

	if len(n.Fields) != len(x.Fields) {
		return false
	}

	for i, v := range n.Fields {
		w := x.Fields[i]
		switch x := v.(type) {
		case *FieldDeclarationEmbedded:
			y, ok := w.(*FieldDeclarationEmbedded)
			if !ok ||
				x.Ptr != y.Ptr ||
				mustUnqote(x.Tag) != mustUnqote(y.Tag) ||
				!x.Type.identicalInner(m, y.Type) {
				return false
			}
		case *FieldDeclarationNamed:
			y, ok := w.(*FieldDeclarationNamed)
			if !ok ||
				x.Name.Val != y.Name.Val ||
				mustUnqote(x.Tag) != mustUnqote(y.Tag) ||
				!x.Type.identicalInner(m, y.Type) {
				return false
			}
		}
	}
	return true
}

func (n *StructType) isGeneric(m map[Type]bool) bool {
	if m[n] {
		return false
	}

	m[n] = true
	for i := range n.Fields {
		if n.fld(i).isGeneric(m) {
			return true
		}
	}
	return false
}

func (n *StructType) String() string {
	var a []string
	for i := range n.Fields {
		a = append(a, n.fld(i).String())
	}
	return fmt.Sprintf("struct{%s}", strings.Join(a, "; "))
}

// TupleType represents '(T1, T2, ...)'.
type TupleType struct {
	List []Type

	typ
}

func newTupleType(pos npos, l []Type) *TupleType {
	n := &TupleType{List: l}
	n.npos = pos
	n.outer = n
	if n.npos.sf == nil && len(l) != 0 {
		n.setPos(l[0].pos())
	}
	return n
}
func (n *TupleType) Kind() TypeKind                       { return Tuple }
func (n *TupleType) comparableWith(t Type) (bool, string) { return false, "" }

func (n *TupleType) check(s *cstack) {
	for _, v := range n.List {
		s.check(v, n)
	}
}

func (n *TupleType) setPos(pos npos) {
	if n.npos.sf != nil {
		return
	}

	n.npos = pos
	for _, v := range n.List {
		v.setPos(pos)
	}
}

func (n *TupleType) identical(m map[tpair]sentinel, t Type) bool {
	if n == t {
		return true
	}

	x, ok := t.(*TupleType)
	if !ok {
		return false
	}

	if len(n.List) != len(x.List) {
		return false
	}

	for i, v := range n.List {
		if !v.identicalInner(m, x.List[i]) {
			return false
		}
	}
	return true
}

func (n *TupleType) isGeneric(m map[Type]bool) bool {
	if n == nil {
		return false
	}

	if m[n] {
		return false
	}

	m[n] = true
	for _, v := range n.List {
		if v.isGeneric(m) {
			return true
		}
	}
	return false
}

func (n *TupleType) String() string {
	var a []string
	for _, v := range n.List {
		a = append(a, v.String())
	}
	return fmt.Sprintf("(%s)", strings.Join(a, ", "))
}
