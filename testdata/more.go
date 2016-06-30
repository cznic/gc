// errorcheck

package foo

import 42 // ERROR "import statement not a string"

func init(int) { // ERROR "func init must have no arguments and no return values"
}

func init() int { // ERROR "func init must have no arguments and no return values"
}

type i interface {
	m()
	n()
	m() // ERROR "duplicate method m"
}

var _ = stringID(0)

type stringID int

var forcegcperiod int64 = 2 * 60 * 1e9
