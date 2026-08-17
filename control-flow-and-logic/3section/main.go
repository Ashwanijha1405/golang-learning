package main 

import(
	"fmt"
	"time"
)

func main() {

	fmt.Println("------------------- Example 1 --------------------")
	day := "Sunday"
    switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend! No work")

	case "Monday", "Tuesday":
		fmt.Println("Busy day, Lots of meetings!")

	default:
		fmt.Println("Mid-Week")
	}
	
	fmt.Println("------------------- Example 2 --------------------")
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 18:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Night")
	}

	fmt.Println("------------------- Example 3 --------------------")
	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}

	checkType(21)
	checkType("Test")
	checkType(true)
	checkType(312.2)

}

// ============================================================
// SWITCH STATEMENTS IN GO
// ============================================================
//
// Go has three useful ways to use switch:
//
// 1. Value-based switch
// 2. Condition-based switch
// 3. Type switch
//
// ------------------------------------------------------------
//
// 1. VALUE-BASED SWITCH
//
// switch day {
// case "Saturday", "Sunday":
//     ...
// case "Monday", "Tuesday":
//     ...
// default:
//     ...
// }
//
// Go compares the value of `day` with each case.
//
// Multiple values can be handled by one case:
//
// case "Saturday", "Sunday":
//
// This means:
//
//     day == "Saturday" OR day == "Sunday"
//
// Unlike C/C++, Go automatically stops after a matching case.
// You normally do NOT need `break`.
//
// ------------------------------------------------------------
//
// 2. CONDITION-BASED SWITCH
//
// switch {
// case hour < 12:
//     ...
// case hour < 18:
//     ...
// default:
//     ...
// }
//
// When there is no expression after `switch`,
// each `case` is treated as a boolean condition.
//
// This is similar to:
//
// if hour < 12 {
//     ...
// } else if hour < 18 {
//     ...
// } else {
//     ...
// }
//
// So:
//
// switch {
// case hour < 12:
//     ...
// }
//
// basically asks:
//
//     "Is hour < 12?"
//
// ------------------------------------------------------------
//
// 3. IMPORTANT: `interface{}`
//
// This is the most important part of this example.
//
// `interface{}` is the empty interface.
//
// An empty interface can hold a VALUE OF ANY TYPE.
//
// Example:
//
// var x interface{}
//
// x = 10
// x = "Hello"
// x = true
// x = 3.14
//
// All of these are allowed.
//
// In modern Go, `any` is an alias for `interface{}`:
//
// var x any
//
// is equivalent to:
//
// var x interface{}
//
// ------------------------------------------------------------
//
// WHY WOULD WE NEED THIS?
//
// Normally, a function has a fixed parameter type:
//
// func printNumber(n int) {
//     ...
// }
//
// This function only accepts an int.
//
// But suppose we want one function that can receive:
//
//     int
//     string
//     bool
//     float64
//
// We can use:
//
// func checkType(i interface{}) {
//
// }
//
// Now `i` can contain a value of any type.
//
// ------------------------------------------------------------
//
// 4. BUT THERE IS A PROBLEM
//
// If `interface{}` can contain anything,
// how do we know WHAT TYPE is currently inside it?
//
// That's exactly what the TYPE SWITCH solves.
//
// ------------------------------------------------------------
//
// 5. TYPE SWITCH
//
// switch v := i.(type) {
//
// case int:
//     ...
//
// case string:
//     ...
//
// case bool:
//     ...
//
// default:
//     ...
//
// }
//
// This asks:
//
//     "What type of value is currently stored inside `i`?"
//
// It is different from a normal switch.
//
// NORMAL SWITCH:
//
// switch value {
//
// case 10:
// case 20:
// }
//
// → compares VALUES.
//
// TYPE SWITCH:
//
// switch v := i.(type) {
//
// case int:
// case string:
// }
//
// → checks TYPES.
//
// ------------------------------------------------------------
//
// 6. WHAT DOES `i.(type)` MEAN?
//
// This:
//
//     i.(type)
//
// is special syntax that can ONLY be used inside a type switch.
//
// It means:
//
//     "Find out what type is stored inside i."
//
// Example:
//
// i = 21
//
// Then:
//
// case int:
//
// matches.
//
// If:
//
// i = "Test"
//
// Then:
//
// case string:
//
// matches.
//
// ------------------------------------------------------------
//
// 7. WHAT IS `v`?
//
// Look carefully:
//
// switch v := i.(type) {
//
// `v` is NOT a keyword.
//
// It is simply a variable name.
//
// You could write:
//
// switch value := i.(type) {
//
// or:
//
// switch x := i.(type) {
//
// `v` is commonly used because it means "value".
//
// The important part is:
//
//     v contains the actual value
//     after Go determines which type it has.
//
// Example:
//
// i = 21
//
// case int:
//
//     v → 21
//
// Therefore:
//
// fmt.Printf("Integer: %d\n", v)
//
// prints:
//
// Integer: 21
//
// ------------------------------------------------------------
//
// 8. WHY DO WE NEED `v`?
//
// Suppose:
//
// i = 21
//
// We discover:
//
//     i is an int.
//
// But we also want to USE the actual integer.
//
// `v` gives us that value in the correct type.
//
// So:
//
// case int:
//     fmt.Printf("Integer: %d\n", v)
//
// Here `v` is treated as an int.
//
// Similarly:
//
// case string:
//     fmt.Printf("String: %s\n", v)
//
// Here `v` is treated as a string.
//
// ------------------------------------------------------------
//
// 9. UNDERSTANDING YOUR `checkType()`
//
// func checkType(i interface{}) {
//
//     switch v := i.(type) {
//
//     case int:
//         fmt.Printf("Integer: %d\n", v)
//
//     case string:
//         fmt.Printf("String: %s\n", v)
//
//     case bool:
//         fmt.Printf("Boolean: %t\n", v)
//
//     default:
//         fmt.Printf("Unknown type: %T\n", v)
//     }
// }
//
// Think of it as:
//
//     "Give me ANY value."
//              ↓
//        `interface{}`
//              ↓
//     "What type is inside?"
//              ↓
//        type switch
//              ↓
//     use the value appropriately
//
// ------------------------------------------------------------
//
// 10. DRY RUN
//
// checkType(21)
//
// `i` contains:
//
//     21
//
// Type:
//
//     int
//
// Therefore:
//
// case int:
//
// matches.
//
// `v` becomes:
//
//     21
//
// Output:
//
//     Integer: 21
//
// ------------------------------------------------------------
//
// checkType("Test")
//
// `i` contains:
//
//     "Test"
//
// Type:
//
//     string
//
// Therefore:
//
// case string:
//
// matches.
//
// `v` becomes:
//
//     "Test"
//
// Output:
//
//     String: Test
//
// ------------------------------------------------------------
//
// checkType(true)
//
// Type:
//
//     bool
//
// Therefore:
//
// case bool:
//
// matches.
//
// Output:
//
//     Boolean: true
//
// ------------------------------------------------------------
//
// checkType(312.2)
//
// Type:
//
//     float64
//
// There is no:
//
// case float64:
//
// So `default` executes.
//
// `%T` prints the TYPE of a value.
//
// Output:
//
//     Unknown type: float64
//
// ------------------------------------------------------------
//
// 11. `%T` vs `%v`
//
// `%T` → prints the TYPE.
//
// `%v` → prints the VALUE.
//
// Example:
//
// x := 21
//
// fmt.Printf("%T\n", x)
//
// Output:
//
// int
//
// fmt.Printf("%v\n", x)
//
// Output:
//
// 21
//
// ------------------------------------------------------------
//
// 12. TYPE SWITCH vs NORMAL SWITCH
//
// NORMAL SWITCH:
//
// switch day {
//
// case "Sunday":
//     ...
//
// }
//
// Checks:
//
//     VALUE
//
// TYPE SWITCH:
//
// switch v := i.(type) {
//
// case int:
//     ...
//
// case string:
//     ...
//
// }
//
// Checks:
//
//     TYPE
//
// Remember:
//
// Normal switch → "What is the value?"
//
// Type switch → "What is the type?"
//
// ------------------------------------------------------------
//
// 13. `interface{}` vs `any`
//
// These are equivalent:
//
// interface{}
// any
//
// `any` is simply a more readable alias for `interface{}`.
//
// So:
//
// func checkType(i interface{})
//
// and:
//
// func checkType(i any)
//
// mean the same thing.
//
// `any` is often preferred in modern Go code.
//
// ------------------------------------------------------------
//
// KEY TAKEAWAYS:
//
// `switch value`
// → compares values.
//
// `switch {}`
// → evaluates boolean conditions.
//
// `switch v := i.(type)`
// → checks the type stored inside an interface.
//
// `interface{}`
// → can hold a value of any type.
//
// `any`
// → alias for `interface{}`.
//
// `v`
// → normal variable containing the value after the type
//   has been determined.
//
// `%T`
// → prints the type.
//
// `%v`
// → prints the value.
//
// Go automatically breaks after a matching switch case.
// No `break` is normally required.