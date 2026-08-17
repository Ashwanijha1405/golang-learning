// Returning multiple values from a function

package main

import(
	"fmt"
	"errors"
	"strings"
)

func divide(a, b int) (int, error) { // By convention error is always the last value returned.

	if b==0 {
		return 0, errors.New("Divide by Zero")
	}

	return a / b, nil 
}

func splitName(fullName string) (firstName, lastName string) {

	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]

	return

}

func main() {

	value, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	firstName, lastName := splitName("Ashwani Jha")

	fmt.Println(firstName, lastName)

}


// --------------------------------------- Returning Multiple Values ---------------------------------------
//
// Go functions can return MORE THAN ONE value.
//
// This is very common in Go, especially for:
//     1. Returning a result + an error
//     2. Returning multiple related values
//
// --------------------------------------------------------------------------------
//
// 1. FUNCTION RETURNING VALUE + ERROR
//
// func divide(a, b int) (int, error)
//
// Parameters:
//     a, b int
//
// Return values:
//     int   → result of division
//     error → tells us whether something went wrong
//
// Go commonly follows the convention:
//
//     result, err := someFunction()
//
// The `error` value is usually the LAST return value.
//
// --------------------------------------------------------------------------------
//
// 2. `errors.New()`
//
// errors.New("Divide by Zero")
//
// Creates a new error containing the given message.
//
// Example:
//
//     return 0, errors.New("Divide by Zero")
//
// Means:
//
//     result = 0
//     error  = "Divide by Zero"
//
// We use an error instead of crashing the program.
//
// --------------------------------------------------------------------------------
//
// 3. `nil` FOR NO ERROR
//
// return a / b, nil
//
// `nil` means:
//
//     "There is no error."
//
// So:
//
//     return 5, nil
//
// means:
//
//     result = 5
//     error  = nothing
//
// --------------------------------------------------------------------------------
//
// 4. CHECKING AN ERROR
//
// value, err := divide(10, 0)
//
// The function returns TWO values:
//
//     value → result
//     err   → error
//
// Then:
//
//     if err != nil {
//
// means:
//
//     "Did an error occur?"
//
// If `err` is NOT nil:
//
//     fmt.Println(err)
//
// Otherwise:
//
//     fmt.Println(value)
//
// This pattern is extremely common in Go:
//
//     result, err := function()
//
//     if err != nil {
//         // handle error
//     }
//
// --------------------------------------------------------------------------------
//
// 5. WHY DOES GO RETURN ERRORS LIKE THIS?
//
// Go generally prefers EXPLICIT error handling.
//
// Instead of:
//
//     try {
//         ...
//     } catch {
//         ...
//     }
//
// Go commonly does:
//
//     result, err := function()
//
//     if err != nil {
//         // handle error
//     }
//
// This makes it obvious where errors are being handled.
//
// --------------------------------------------------------------------------------
//
// 6. FUNCTION RETURNING TWO NORMAL VALUES
//
// func splitName(fullName string) (firstName, lastName string)
//
// This function:
//
//     takes:
//         fullName → string
//
//     returns:
//         firstName → string
//         lastName  → string
//
// The return types are:
//
//     (string, string)
//
// We can give the return values names:
//
//     (firstName, lastName string)
//
// This means:
//
//     firstName string
//     lastName  string
//
// --------------------------------------------------------------------------------
//
// 7. `strings.Split()`
//
// parts := strings.Split(fullName, " ")
//
// Splits a string wherever the separator occurs.
//
// Example:
//
//     fullName = "Ashwani Jha"
//
//     strings.Split(fullName, " ")
//
// produces:
//
//     ["Ashwani", "Jha"]
//
// Therefore:
//
//     parts[0] → "Ashwani"
//     parts[1] → "Jha"
//
// --------------------------------------------------------------------------------
//
// 8. NAMED RETURN VALUES
//
// func splitName(fullName string) (firstName, lastName string)
//
// Because we named the return values, we can assign values directly:
//
//     firstName = parts[0]
//     lastName = parts[1]
//
// Then:
//
//     return
//
// A bare `return` means:
//
//     "Return the current values of firstName and lastName."
//
// So:
//
//     return
//
// is equivalent to:
//
//     return firstName, lastName
//
// --------------------------------------------------------------------------------
//
// 9. IMPORTANT: `return` vs `return value`
//
// Normal return:
//
//     func add(a, b int) int {
//         return a + b
//     }
//
// Here we explicitly tell Go what to return.
//
//
//
// Named return:
//
//     func add(a, b int) (result int) {
//         result = a + b
//         return
//     }
//
// Here `result` already represents the return value.
//
// --------------------------------------------------------------------------------
//
// 10. CALLING A MULTI-VALUE FUNCTION
//
// firstName, lastName := splitName("Ashwani Jha")
//
// The function returns:
//
//     "Ashwani", "Jha"
//
// So:
//
//     firstName → "Ashwani"
//     lastName  → "Jha"
//
// Then:
//
//     fmt.Println(firstName, lastName)
//
// prints:
//
//     Ashwani Jha
//
// --------------------------------------------------------------------------------
//
// 11. COMPLETE FLOW
//
// divide(10, 0)
//
//        ↓
//
// b == 0 ?
//
//        ↓ YES
//
// return 0, errors.New("Divide by Zero")
//
//        ↓
//
// value = 0
// err   = error
//
//        ↓
//
// err != nil
//
//        ↓
//
// print error
//
//
// ------------------------------------------------------------
//
// splitName("Ashwani Jha")
//
//        ↓
//
// strings.Split()
//
//        ↓
//
// ["Ashwani", "Jha"]
//
//        ↓
//
// firstName = "Ashwani"
// lastName  = "Jha"
//
//        ↓
//
// return
//
//        ↓
//
// firstName, lastName receive the values
//
// --------------------------------------------------------------------------------
//
// KEY TAKEAWAYS:
//
// `func divide(...) (int, error)`
// → function returns two values.
//
// `result, err := function()`
// → common Go pattern for handling results + errors.
//
// `error`
// → represents something that went wrong.
//
// `errors.New()`
// → creates a new error.
//
// `nil`
// → no error / no value.
//
// `err != nil`
// → an error occurred.
//
// `strings.Split()`
// → splits a string into a slice.
//
// `(firstName, lastName string)`
// → named return values.
//
// `return`
// → returns the named values.
//
// `return firstName, lastName`
// → explicitly returns both values.
//
// Go makes multiple return values especially useful for:
//
//     result + error
//
// which you'll see CONSTANTLY in real Go backend code.