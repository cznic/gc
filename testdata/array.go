// errorcheck

// Copyright 2009 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foo

var (
	_ = [...]int{1, 2, 3}
	_ = [...]int{2: 42}
	_ = ([...]int){1, 2, 3} // ERROR "syntax error: cannot parenthesize type in composite literal"
	_ = ([...]int){2: 42}   // ERROR "syntax error: cannot parenthesize type in composite literal"
)
