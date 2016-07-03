// errorcheck

// Copyright 2009 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foo

type t struct{}

func (t) m() {}

func f() {}

var (
	_ func()  = nil
	_ func(t) = nil

	_ func()  = t.m // ERROR "TODO"
	_ func(t) = t.m // ok

	_ *func() = &f // ERROR "cannot take address"

	_ func(t) = &t{}.m // ERROR "cannot take address"
	_ func()  = &t{}.m // ERROR "cannot take address"

	_ func(t) = (&t{}).m // ERROR "TODO"
	_ func()  = (&t{}).m // ok

	_ string = 42 // ERROR "TODO"
)
