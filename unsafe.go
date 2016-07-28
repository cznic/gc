// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"strings"
)

func offsetof(ctx *context, n *Call) Value {
	args, _, ddd := n.args()
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

	root, path0, path := v.Selector()
	if root == nil || path == nil {
		todo(n, true)
		return nil
	}

	if _, ok := path[len(path)-1].(*StructField); !ok {
		ctx.err(n.ArgumentList.Argument, "invalid unsafe.Offset expression: argument is a method value")
		return nil
	}

	if len(path0) != len(path) {
		//dbg("==== %s", position(n.Pos()))
		//for _, v := range path0 {
		//	dbg("\t %s", dict.S(v.(*StructField).Name))
		//}
		//dbg("---")
		//for _, v := range path {
		//	dbg("\t %s", dict.S(v.(*StructField).Name))
		//}
		i := 0
		for j, v := range path {
			if path0[i] == v {
				i++
				continue
			}

			f := v.(*StructField)
			if f.Type.Kind() == Ptr && f.Anonymous {
				var a []string
				for _, v := range path[:j+1] {
					a = append(a, string(dict.S(v.(*StructField).Name)))
				}
				ctx.err(
					n.ArgumentList.Argument,
					"invalid unsafe.Offset expression: selector implies indirection of embedded field %s",
					strings.Join(a, "."),
				)
				return nil
			}
		}
	}

	var off uint64
	for i, v := range path {
		f := v.(*StructField)
		if f.Type.Kind() == Ptr && i != len(path)-1 {
			off = 0
			continue
		}

		off += f.Offset
	}
	return uintptrConstValueFromUint64(ctx, off)
}

func sizeof(ctx *context, n *Call) Value {
	args, _, ddd := n.args()
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
