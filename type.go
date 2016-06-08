// Copyright 2016 The GC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Portions of this source are a derived work. The original code is at
//
//	https://github.com/golang/go/blob/e805bf39458915365924228dc53969ce04e32813/src/reflect/type.go
//
// and contains the following copyright notice.
//
// ---------------------------------------------------
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// ---------------------------------------------------
//
// The LICENSE file mentioned just above is available as GO-LICENSE in this
// repository.

package gc

// Type is the representation of a Go type.
//
// Not all methods apply to all kinds of types.  Restrictions, if any, are
// noted in the documentation for each method.  Use the Kind method to find out
// the kind of type before calling kind-specific methods.  Calling a method
// inappropriate to the kind of type causes a run-time panic.
type Type interface {
}
