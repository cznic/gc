// errorcheck

// Copyright 2009 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test forward reference to a definition in the same local scope fails.
package foo

const a = 42

func bar() {
	const b = 24
	var c = a
	var d = b
	var e = f // ERROR "undefined"
	const f = 314
	var g = f
	var h = i // ERROR "undefined"
	var j = 1
}
