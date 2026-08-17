// --------------------------------------- Variadic Functions -------------------------------------

package main

import "fmt"

func sum(numbers ...int) int {

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total

}

func config(numbers ...int) {

	if len(numbers) > 0 {
		first := numbers[0]
		fmt.Println("First number:", first)
	} else {
		fmt.Println("Default number")
	}

}

func main() {

	fmt.Println(sum(1,2,3,4,5))

	config(5)
	config()

}


// --------------------------------------- Variadic Functions -------------------------------------

// A variadic function can accept ANY NUMBER of arguments.
//
// Syntax:
// func functionName(values ...type)
//
// `...int` means:
// → zero or more int arguments.
//
// Example:
// sum()
// sum(1)
// sum(1, 2, 3, 4, 5)

// ------------------------------------------------------------

// `numbers` behaves like a slice inside the function.
//
// So:
//
// func sum(numbers ...int)
//
// is roughly treated inside the function as:
//
// numbers []int

func sum(numbers ...int) int {

	total := 0

	// `range` loops through all the values.
	for _, number := range numbers {
		total += number
	}

	return total
}

// ------------------------------------------------------------

// A variadic function can also be called with ZERO arguments.
//
// `len(numbers)` tells us how many arguments were provided.

func config(numbers ...int) {

	if len(numbers) > 0 {

		// Access the first argument using index 0.
		first := numbers[0]

		fmt.Println("First number:", first)

	} else {

		// No arguments were provided.
		fmt.Println("Default number")
	}
}

// ------------------------------------------------------------

// Calling variadic functions:
//
// sum(1, 2, 3, 4, 5)
// → numbers = []int{1, 2, 3, 4, 5}
//
// config(5)
// → numbers = []int{5}
//
// config()
// → numbers = []int{}

// ------------------------------------------------------------

// IMPORTANT:
//
// You can also pass an existing slice to a variadic function
// using `...`:
//
// nums := []int{1, 2, 3, 4}
//
// sum(nums...)
//
// `...` expands the slice into individual arguments.
//
// Without `...`:
//
// sum(nums)       // ❌ wrong
//
// With `...`:
//
// sum(nums...)    // ✅ correct

// ------------------------------------------------------------
//
// KEY IDEA:
//
// `...int`
// → accepts zero or more integers.
//
// Inside the function:
// → the arguments behave like a slice.
//
// Useful when:
// → you don't know how many arguments the caller will provide.