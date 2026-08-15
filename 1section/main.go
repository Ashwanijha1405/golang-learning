package main

import "fmt"

func main() {

	var t any

	fmt.Println("Hello " + "World\n")

	fmt.Println(1 + 1)

	fmt.Println(3.14)

	fmt.Println(true, false)

	fmt.Println("%+v\n", []int{1, 2, 3})

	fmt.Println("%+v\n", t)
}


// VALUES IN GO
//
// Go supports different kinds of values:
//
// 1. Strings
//    "Hello World"
//
// 2. Integers
//    1, 42, -10
//
// 3. Floating-point numbers
//    3.14, 10.5
//
// 4. Booleans
//    true, false
//
// 5. Arrays / Slices
//    []int{1, 2, 3}
//
// %v is a general-purpose format verb.
// It prints the value in its default format.
//
// Examples:
//
// fmt.Println("Hello " + "World") // String concatenation
// fmt.Println(1 + 1)              // Arithmetic
// fmt.Println(3.14)               // Float
// fmt.Println(true, false)        // Boolean
// fmt.Printf("%v\n", []int{1, 2, 3}) // Slice
//
// `any` can store a value of any type.
// Example:
// var t any
// t = 10
// t = "hello"
//
// NOTE:
// `any` is an alias for `interface{}`.
// It is useful when you don't know the type beforehand,
// but avoid using it unnecessarily because Go is statically typed.

