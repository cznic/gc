// errorcheck

package foo

type a int

type t t // ERROR "invalid recursive type t"

type z string
