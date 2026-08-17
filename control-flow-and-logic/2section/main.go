package main 

import "fmt"

func main() {
    
	fmt.Println("--------------Variation 1----------------")
	tmp := 25 
	if tmp > 30 {
		fmt.Println("greater")
	}else{
		fmt.Println("less")
	}

	fmt.Println("--------------Variation 2----------------")
	score := 85 
	if score >= 80 {
		fmt.Println("Grade: A")
	}else if score >= 70 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else if score < 60 {
		fmt.Println("Failed")
	}

	fmt.Println("--------------Variation 3----------------")
	userAccess := map[string]bool{
		"Ashwani": true,
		"Priyanshu": false,
	}

	if hasAccess, ok := userAccess["Ashwani"]; ok && hasAccess {
		fmt.Println("User can access the system")
	}else{
		fmt.Println("Access Denied")
	}

}

// ============================================================
// IF / ELSE & MAP LOOKUP IN GO
// ============================================================
//
// 1. BASIC IF
//
// if tmp > 30 {
//     fmt.Println("greater")
// } else {
//     fmt.Println("less")
// }
//
// Go does NOT require parentheses around the condition.
//
// ❌ if (tmp > 30)
//
// ✅ if tmp > 30
//
// Curly braces `{}` are mandatory.
//
// ------------------------------------------------------------
//
// 2. ELSE IF
//
// Go supports:
//
// if condition1 {
//     ...
// } else if condition2 {
//     ...
// } else {
//     ...
// }
//
// Conditions are checked from TOP to BOTTOM.
//
// As soon as one condition is true, its block executes
// and the remaining conditions are skipped.
//
// ------------------------------------------------------------
//
// 3. MAP
//
// map[string]bool
//
// Means:
//
//     key   → string
//     value → bool
//
// Example:
//
// userAccess := map[string]bool{
//     "Ashwani":   true,
//     "Priyanshu": false,
// }
//
// Think of a map like a dictionary:
//
// "Ashwani"   → true
// "Priyanshu" → false
//
// ------------------------------------------------------------
//
// 4. NORMAL MAP LOOKUP
//
// hasAccess := userAccess["Ashwani"]
//
// This gives:
//
// hasAccess → true
//
// But there is a problem.
//
// If the key doesn't exist:
//
// userAccess["Rahul"]
//
// Go gives the ZERO VALUE of bool:
//
// false
//
// So you cannot tell whether:
//
// "Rahul exists and hasAccess is false"
//
// OR:
//
// "Rahul doesn't exist"
//
// Both can give:
//
// false
//
// ------------------------------------------------------------
//
// 5. TWO-VALUE MAP LOOKUP
//
// Go solves this using:
//
// value, ok := map[key]
//
// Example:
//
// hasAccess, ok := userAccess["Ashwani"]
//
// `hasAccess` → the value stored in the map
// `ok`        → whether the key exists
//
// For Ashwani:
//
// hasAccess = true
// ok        = true
//
// For Priyanshu:
//
// hasAccess = false
// ok        = true
//
// For Rahul:
//
// hasAccess = false
// ok        = false
//
// IMPORTANT:
//
// `ok` does NOT mean "has access".
//
// `ok` means:
//
//     "Did the key exist in the map?"
//
// ------------------------------------------------------------
//
// 6. `ok && hasAccess`
//
// if hasAccess, ok := userAccess["Ashwani"]; ok && hasAccess {
//
//     ...
//
// }
//
// This means:
//
//     key exists
//         AND
//     access is true
//
// `&&` means AND.
//
// Both conditions must be true.
//
// Example:
//
// ok = true
// hasAccess = true
// → true && true → ACCESS GRANTED
//
// ok = true
// hasAccess = false
// → true && false → ACCESS DENIED
//
// ok = false
// hasAccess = false
// → false && false → ACCESS DENIED
//
// ------------------------------------------------------------
//
// 7. WHY USE `value, ok`?
//
// It is a very common Go pattern.
//
// You will see it frequently with:
//
//     map lookups
//     type assertions
//     channel operations
//
// General idea:
//
// value, ok := something
//
// `value` → result
// `ok`    → whether the operation succeeded / value existed
//
// ------------------------------------------------------------
//
// 8. IMPORTANT DIFFERENCE
//
// `if userAccess["Ashwani"] { ... }`
//
// Only checks the VALUE.
//
// `if hasAccess, ok := userAccess["Ashwani"]; ok && hasAccess { ... }`
//
// Checks BOTH:
//
// 1. Does the user exist?
// 2. Does the user have access?
//
// This is safer when the distinction matters.
//
// ------------------------------------------------------------
//
// 9. VARIABLE SCOPE IN IF
//
// This:
//
// if hasAccess, ok := userAccess["Ashwani"]; ok && hasAccess {
//     ...
// }
//
// declares `hasAccess` and `ok` inside the `if` statement.
//
// They are available inside the entire if/else structure,
// but NOT outside it.
//
// Example:
//
// if value, ok := myMap["key"]; ok {
//     fmt.Println(value) // ✅
// }
//
// fmt.Println(value)     // ❌ not available here
//
// ------------------------------------------------------------
//
// KEY TAKEAWAYS:
//
// `if`
// → executes code when a condition is true.
//
// `else if`
// → checks another condition if the previous one was false.
//
// `else`
// → runs when none of the previous conditions were true.
//
// `map[K]V`
// → map with key type K and value type V.
//
// `value, ok := map[key]`
// → gets the value AND checks whether the key exists.
//
// `ok`
// → means "key exists / operation succeeded",
//    NOT "the value is true".
//
// `&&`
// → logical AND; both conditions must be true.