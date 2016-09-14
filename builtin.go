// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

func builtinAppend(ctx *context, call *Call) Value {
	args, _, ddd := call.args()
	if len(args) < 2 {
		todo(call.ArgumentList, true) // not enough args
		return nil
	}

	sv := args[0]
	if sv == nil {
		return nil
	}

	switch sv.Kind() {
	case TypeValue:
		todo(call, true) // invalid arg
		return nil
	case NilValue:
		todo(call, true) // invalid arg
		return nil
	}

	st := sv.Type()
	if st.Kind() != Slice {
		todo(call, true) // expected slice
		return nil
	}

	et := st.Elem()
	if et == nil {
		return newRuntimeValue(st)
	}

	switch {
	case ddd:
		todo(call)
	default:
		for _, v := range args[1:] {
			if v == nil {
				continue
			}

			if !v.AssignableTo(ctx.Context, et) {
				todo(call, true) // type mismatch
			}
		}
		return newRuntimeValue(st)
	}
	return nil
}

func builtinCap(ctx *context, call *Call) Value {
	args, flags, ddd := call.args()
	if ddd {
		todo(call.ArgumentList, true) // ... invalid
	}
	if len(args) < 1 {
		todo(call.ArgumentList, true) // not enough args
		return nil
	}

	if len(args) > 1 {
		todo(call.ArgumentList, true) // too many args
	}

	v := args[0]
	if v == nil {
		return nil
	}

	switch v.Kind() {
	case RuntimeValue:
		switch t := v.Type(); t.Kind() {
		case Ptr:
			if t.Elem().Kind() != Array {
				todo(call, true) // Invalid
				break
			}

			t = t.Elem()
			fallthrough
		case Array:
			if len := t.Len(); len >= 0 {
				switch {
				case flags.lenPoisoned():
					return newRuntimeValue(ctx.intType)
				default:
					return newConstValue(newIntConst(len, nil, ctx.intType, true))
				}
			}
		default:
			//dbg("", t.Kind())
			todo(call)
		}
	default:
		//dbg("", v.Kind())
		todo(call)
	}
	return nil
}

func builtinComplex(ctx *context, call *Call) Value {
	args, _, ddd := call.args()
	if ddd {
		todo(call.ArgumentList, true) // ... invalid
	}
	if len(args) < 2 {
		todo(call.ArgumentList, true) // not enough args
		return nil
	}

	if len(args) > 2 {
		todo(call.ArgumentList, true) // too many args
	}

	var constRe, constIm Const
	if re := args[0]; re != nil {
		switch re.Kind() {
		case RuntimeValue:
			todo(call)
		case ConstValue:
			if constRe = re.Const().mustConvert(ctx, call.ArgumentList.Argument, ctx.float64Type); constRe == nil {
				return nil
			}
		default:
			todo(call.ArgumentList, true) //
			return nil
		}
	}
	if im := args[1]; im != nil {
		switch im.Kind() {
		case RuntimeValue:
			todo(call)
		case ConstValue:
			if constIm = im.Const().mustConvert(ctx, call.ArgumentList.Argument, ctx.float64Type); constIm == nil {
				return nil
			}
		default:
			todo(call.ArgumentList, true) //
			return nil
		}
	}

	if constRe == nil || constIm == nil {
		return nil
	}

	c := newComplexConst(
		0,
		&bigComplex{constRe.(*floatConst).bigVal, constIm.(*floatConst).bigVal},
		ctx.complex128Type,
		constRe.Untyped() && constIm.Untyped(),
	)
	return newConstValue(c)
}

func builtinImag(ctx *context, call *Call) Value {
	args, _, ddd := call.args()
	if ddd {
		todo(call.ArgumentList, true) // ... invalid
	}
	if len(args) < 1 {
		todo(call.ArgumentList, true) // not enough args
		return nil
	}

	if len(args) > 1 {
		todo(call.ArgumentList, true) // too many args
	}

	v := args[0]
	if v == nil {
		return nil
	}

	switch v.Kind() {
	default:
		//dbg("", v.Kind())
		todo(call)
	}
	return nil
}

func builtinLen(ctx *context, call *Call) Value {
	args, flags, ddd := call.args()
	if ddd {
		todo(call.ArgumentList, true) // ... invalid
	}
	if len(args) < 1 {
		todo(call.ArgumentList, true) // not enough args
		return nil
	}

	if len(args) > 1 {
		todo(call.ArgumentList, true) // too many args
	}

	v := args[0]
	if v == nil {
		return nil
	}

	switch v.Kind() {
	case ConstValue:
		switch c := v.Const(); c.Kind() {
		case StringConst:
			return newConstValue(newIntConst(int64(c.(*stringConst).val.len()), nil, ctx.intType, true))
		default:
			//dbg("", c.Kind())
			todo(call)
		}
	case RuntimeValue:
		switch t := v.Type(); t.Kind() {
		case Ptr:
			if t.Elem().Kind() != Array {
				todo(call, true) // Invalid
				break
			}

			t = t.Elem()
			fallthrough
		case Array:
			if len := t.Len(); len >= 0 {
				switch {
				case flags.lenPoisoned():
					return newRuntimeValue(ctx.intType)
				default:
					return newConstValue(newIntConst(len, nil, ctx.intType, true))
				}
			}
		case Slice:
			return newRuntimeValue(ctx.intType)
		default:
			//dbg("", t.Kind())
			todo(call)
		}
	default:
		//dbg("", v.Kind())
		todo(call)
	}
	return nil
}

func builtinMake(ctx *context, call *Call) Value {
	args, _, ddd := call.args()
	if ddd {
		todo(call.ArgumentList, true) // ... invalid
	}
	if len(args) < 1 {
		todo(call.ArgumentList, true) // not enough args
		return nil
	}

	if len(args) > 3 {
		todo(call.ArgumentList.node(3), true) // too many args
	}

	iarg := [3]int64{1: -1, 2: -1}
	for i := 1; i <= 2 && i < len(args); i++ {
		v := args[i]
		if v == nil {
			continue
		}

		if !v.nonNegativeInteger(ctx) {
			todo(call, true)
			return nil
		}

		switch v.Kind() {
		case ConstValue:
			c := v.Const().Convert(ctx.Context, ctx.intType)
			if c == nil {
				todo(call, true) //TODO ctx.constConversionFail(call.ArgumentList.node(i), ctx.intType, v.Const())
				return nil
			}

			iarg[i] = c.Const().(*intConst).val
		default:
			//dbg("", v.Kind())
			todo(call)
		}
	}

	v := args[0]
	if v == nil {
		return nil
	}

	switch v.Kind() {
	case TypeValue:
		switch t := v.Type(); t.Kind() {
		case Chan:
			if len(args) > 2 {
				todo(call.ArgumentList.node(2), true) // too many args
			}

			return newRuntimeValue(t)
		case Map:
			if kt := t.Key(); kt != nil && !kt.Comparable() {
				todo(call, true) // invalid key type
				break
			}

			if len(args) > 2 {
				todo(call.ArgumentList.node(2), true) // too many args
			}
			return newRuntimeValue(t)
		case Slice:
			if len(args) == 3 {
				todo(call) // check arg2 <= arg3
			}
			return newRuntimeValue(t)
		default:
			todo(call, true) // invalid arg
		}
	default:
		//dbg("", v.Kind())
		todo(call, true) // not a type
	}
	return nil
}

func builtinNew(ctx *context, call *Call) Value {
	args, _, ddd := call.args()
	if ddd {
		todo(call.ArgumentList, true) // ... invalid
	}
	if len(args) < 1 {
		todo(call.ArgumentList, true) // not enough args
		return nil
	}

	if len(args) > 1 {
		todo(call.ArgumentList.node(3), true) // too many args
	}

	v := args[0]
	if v == nil {
		return nil
	}

	switch v.Kind() {
	case TypeValue:
		return newRuntimeValue(newPtrType(ctx, v.Type()))
	default:
		todo(call, true) // not a type
	}
	return nil
}

func builtinPanic(ctx *context, call *Call) Value {
	args, _, ddd := call.args()
	if ddd {
		todo(call.ArgumentList, true) // ... invalid
	}
	if len(args) < 1 {
		todo(call.ArgumentList, true) // not enough args
		return nil
	}

	if len(args) > 1 {
		todo(call.ArgumentList.node(3), true) // too many args
	}
	return newRuntimeValue(ctx.voidType)
}

func builtinPrint(ctx *context, call *Call, nl bool) Value {
	//TODO args, _, ddd := call.args()
	return newRuntimeValue(ctx.voidType)
}

func builtinReal(ctx *context, call *Call) Value {
	args, _, ddd := call.args()
	if ddd {
		todo(call.ArgumentList, true) // ... invalid
	}
	if len(args) < 1 {
		todo(call.ArgumentList, true) // not enough args
		return nil
	}

	if len(args) > 1 {
		todo(call.ArgumentList, true) // too many args
	}

	v := args[0]
	if v == nil {
		return nil
	}

	switch v.Kind() {
	case RuntimeValue:
		switch t := v.Type(); t.Kind() {
		case Complex64:
			return newRuntimeValue(ctx.float32Type)
		case Complex128:
			return newRuntimeValue(ctx.float64Type)
		default:
			todo(call, true) // invalid arg
		}
	default:
		//dbg("", v.Kind())
		todo(call)
	}
	return nil
}
