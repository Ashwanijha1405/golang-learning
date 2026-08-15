package main

import "fmt"

func main() {

	var greeting string 
	// variable is assigned but not initialised manually. But in go it is initiliased with "zero-value" which in this case is ""(empty string)

	greeting = "Hello, World!" //memory is allocated as its initialised

	fmt.Println(greeting)


	var count int 
	count = 10
	fmt.Println(count)

	var isRunning bool
	isRunning = true
	fmt.Println(isRunning)

	var firstName, lastName string
	firstName = "Ashwani"
	lastName = "jha"
	fmt.Println(firstName, lastName)

	//Now the shortcut way to do same things as above

	email := "test@test.com" 
	/* ":=" is used for short variable declaration and initialisation at same time.
	    there is no worry to write datatype explicitly because go compiler can internally infer the datatype
		by figuring out whats stored in the variable, which in this case is a string. */
	fmt.Println(email)

	age := 24
	fmt.Println(age)
}


// VARIABLES IN GO
//
// A variable stores a value in memory.
//
// 1. DECLARING A VARIABLE
//
// var greeting string
//
// `var` declares a variable.
// `string` specifies its type.
//
// If no value is explicitly assigned, Go gives the variable
// its type's ZERO VALUE.
//
// Examples:
// string → ""
// int    → 0
// bool   → false
//
// 2. ASSIGNING A VALUE
//
// greeting = "Hello, World!"
//
// The variable was already declared, so we can assign a value
// to it later.
//
// 3. DECLARING MULTIPLE VARIABLES
//
// var firstName, lastName string
//
// Multiple variables of the same type can be declared together.
//
// 4. SHORT VARIABLE DECLARATION
//
// email := "test@test.com"
//
// `:=` declares AND initializes a variable at the same time.
//
// Go automatically INFERS the type from the assigned value.
//
// email := "test@test.com" → string
// age := 24               → int
//
// `:=` can only be used INSIDE functions.
//
// 5. `var` vs `:=`
//
// var age int
// age = 24
//
// OR:
//
// age := 24
//
// Use `:=` when declaring and initializing a new variable.
// Use `var` when you want to explicitly declare a variable,
// especially when you want to rely on its zero value.
//
// IMPORTANT:
// A variable declared with `var` can be reassigned:
//
// age = 25
//
// But `:=` is NOT an assignment operator.
// It is a SHORT VARIABLE DECLARATION.
//
// 6. GO IS STATISTICALLY TYPED
//
// Once a variable has a type, it cannot hold a value of
// another incompatible type.
//
// age := 24      // int
// age = "hello"  // ❌ compile-time error