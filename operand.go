// Copyright 2018 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"math/big"
)

var (
	_ Operand = UntypedFalse{}
	_ Operand = UntypedTrue{}
	_ Operand = UntypedIntOperand{}

	invalidValue        Value = InvalidValue{}
	untypedBooleanFalse Value = UntypedBooleanValue{V: false}
	untypedBooleanTrue  Value = UntypedBooleanValue{V: true}
)

// Value represents operand's value.
type Value interface {
	isValue()
	equal(Value) bool
}

type value struct{}

func (value) isValue()         {}
func (value) equal(Value) bool { return true } //TODO

// InvalidValue represents an invalidValue.
type InvalidValue struct {
	value
}

// UntypedBooleanValue represents a true/false value.
type UntypedBooleanValue struct {
	value
	V bool
}

// UntypedIntValue represents an untyped integer.
type UntypedIntValue struct {
	value
	V *big.Int
}

// Operand describes operation's operand.
type Operand interface {
	Type() Type
	Value() Value
}

// UntypedFalse represents a boolean false constant.
type UntypedFalse struct{}

// Type implements Operand.
func (UntypedFalse) Type() Type { return UntypedBool }

// Value implements Operand.
func (UntypedFalse) Value() Value { return untypedBooleanFalse }

// UntypedTrue represents a bool true constant.
type UntypedTrue struct{}

// Type implements Operand.
func (UntypedTrue) Type() Type { return UntypedBool }

// Value implements Operand.
func (UntypedTrue) Value() Value { return untypedBooleanTrue }

// UntypedIntOperand represents an integer operand.
type UntypedIntOperand struct {
	v UntypedIntValue
}

// Type implements Operand.
func (UntypedIntOperand) Type() Type { return UntypedInt }

// Value implements Operand.
func (o UntypedIntOperand) Value() Value { return o.v }
