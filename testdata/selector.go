// errorcheck

// Modified content from the Go Language Specification.

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foo

type T0 struct {
	x int
}

func (*T0) M0() int { return 0 }

type T1 struct {
	y int
}

func (T1) M1() int { return 0 }

type T2 struct {
	z int
	T1
	*T0
}

func (*T2) M2() int { return 0 }

type Q *T2

var t T2  // with t.T0 != nil
var p *T2 // with p != nil and (*p).T0 != nil
var q Q = p

var (
	_ = t.z // t.z
	_ = t.y // t.T1.y
	_ = t.T1.y
	_ = t.x    // (*t.T0).x
	_ = t.T0.x // (*t.T0).x

	_ = p.z // (*p).z
	_ = p.y // (*p).T1.y
	_ = p.x // (*(*p).T0).x

	_ = q.x // (*(*q).T0).x        (*q).x is a valid field selector

	_ = p.M0() // ((*p).T0).M0()      M0 expects *T0 receiver
	_ = p.M1() // ((*p).T1).M1()      M1 expects T1 receiver
	_ = p.M2() // p.M2()              M2 expects *T2 receiver
	_ = t.M2() // (&t).M2()           M2 expects *T2 receiver, see section on Calls

	_ = (*q).M0 // valid but not a field selector
	_ = (*q).M0()
	_ = q.M0() // ERROR "M0 undefined .type foo.Q has no field or method M0."

)
