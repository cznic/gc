// errorcheck

// Copyright 2009 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foo

const (
	_ = 1 / (len([3]int{}) - 2)
	_ = 1 / (len([3]int{}) - 3) // ERROR "division by zero"
	_ = 1 / (len([3]int{}) - 4)

	_ = 1 / (len([...]int{2: 42, 1: 314}) - 2)
	_ = 1 / (len([...]int{2: 42, 1: 314}) - 3) // ERROR "division by zero"
	_ = 1 / (len([...]int{2: 42, 1: 314}) - 4)
)
