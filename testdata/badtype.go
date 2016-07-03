// errorcheck

// Copyright 2009 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foo

type a int

type t t // ERROR "invalid recursive type t"

type z string
