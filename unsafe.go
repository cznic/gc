// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

func offsetof(ctx *Context, n *Call) Value {
	args, ddd := n.args()
	if len(args) < 1 {
		todo(n, true)
		return nil
	}

	if len(args) > 1 {
		todo(n, true)
		return nil
	}

	if ddd {
		todo(n, true)
		return nil

	}

	v := args[0]
	if v == nil {
		return nil
	}

	if v.Kind() != RuntimeValue {
		todo(n, true)
		return nil
	}

	f, off := v.StructField()
	if f == nil {
		todo(n, true)
		return nil
	}

	return uintptrConstValueFromUint64(ctx, off)
}

func sizeof(ctx *Context, n *Call) Value {
	args, ddd := n.args()
	if len(args) < 1 {
		todo(n, true)
		return nil
	}

	if len(args) > 1 {
		todo(n, true)
		return nil
	}

	if ddd {
		todo(n, true)
		return nil

	}

	v := args[0]
	if v == nil {
		return nil
	}

	t := v.Type()
	if t == nil {
		return nil
	}

	return uintptrConstValueFromUint64(ctx, t.Size())
}
