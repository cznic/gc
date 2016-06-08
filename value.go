// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

// Values of type ValueKind.
const (
	RuntimeValue ValueKind = iota // Value known at run time.
	ConstValue                    // Value known at compile time.
)

// Value describes an expression.
type Value interface {
	Kind() ValueKind
	Type() Type
}

// ValueKind is a value category.
type ValueKind int

//TODO
type StringID int

func (s StringID) String() string { return string(dict.S(int(s))) }
