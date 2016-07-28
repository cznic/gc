// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

func builtinCap(ctx *context, call *Call) Value {
	args, lenPoisoned, ddd := call.args()
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
				case lenPoisoned:
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
			if constRe = re.Const().mustConvert(call.ArgumentList.Argument, ctx.float64Type); constRe == nil {
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
			if constIm = im.Const().mustConvert(call.ArgumentList.Argument, ctx.float64Type); constIm == nil {
				return nil
			}
		default:
			todo(call.ArgumentList, true) //
			return nil
		}
	}

	if constRe == nil || constIm == nil {
		todo(call)
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

func builtinLen(ctx *context, call *Call) Value {
	args, lenPoisoned, ddd := call.args()
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
				case lenPoisoned:
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

	for i := 1; i <= 2 && i < len(args); i++ {
		v := args[i]
		if v == nil {
			continue
		}

		switch v.Kind() {
		//TODO check v is nonNegativeInteger()
		default:
			todo(call.ArgumentList.node(i))
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
			todo(call)
		case Map:
			if kt := t.Key(); kt != nil && !kt.Comparable() {
				todo(call, true) // invalid key type
			}

			if len(args) > 2 {
				todo(call.ArgumentList.node(2), true) // too many args
			}
		case Slice:
			todo(call)
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
