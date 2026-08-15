package main 

import "fmt"

const HOST string = "localhost"

func main() {

	fmt.Println(HOST)

	const appName string = "GoApp"
	fmt.Println(appName)

	const pi float64 = 3.1415926
	fmt.Println(pi)

	const rate float32 = 5.2
	fmt.Println(rate)
}

// CONSTANTS IN GO
//
// A constant is a value that CANNOT be changed after it is declared.
//
// Use `const` instead of `var` when a value should remain fixed.
//
// 1. DECLARING A CONSTANT
//
// const HOST string = "localhost"
//
// `const` → declares a constant
// `HOST`  → constant name
// `string` → explicit type
// `"localhost"` → value
//
// 2. CONSTANTS CAN BE DECLARED OUTSIDE FUNCTIONS
//
// const HOST string = "localhost"
//
// This is declared at package level, so it can be used
// throughout the package.
//
// Constants declared inside `main()` have local scope.
//
// 3. CONSTANTS CAN HAVE DIFFERENT TYPES
//
// const appName string = "GoApp"
// const pi float64 = 3.1415926
// const rate float32 = 5.2
//
// Just like variables, constants can have types such as:
// string, int, float32, float64, bool, etc.
//
// 4. CONSTANTS CANNOT BE REASSIGNED
//
// const pi float64 = 3.14
// pi = 4.0 // ❌ compile-time error
//
// Unlike variables, a constant's value cannot change.
//
// 5. `const` vs `var`
//
// var age int = 24
// age = 25          // ✅ allowed
//
// const pi float64 = 3.14
// pi = 4.0          // ❌ not allowed
//
// Use `var` for values that may change.
// Use `const` for values that should remain fixed.
//
// 6. CONSTANTS ARE EVALUATED AT COMPILE TIME
//
// Constants represent values known at compile time.
// They are useful for fixed configuration values, limits,
// mathematical values, and other values that should not change.
//
// NOTE:
// Constants in Go are more powerful than variables because
// Go constants can be untyped until they are given a specific
// type or used in a context that requires one.

