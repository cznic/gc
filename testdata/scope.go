// errorcheck

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
