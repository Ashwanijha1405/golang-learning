package main 

import "fmt" 

func main() {

	studentGrades := map[string]int{
		"Alice": 90, 
		"James": 85, 
		"Dan": 60,
	}

	fmt.Printf("%+v\n", studentGrades)
	studentGrades["Alice"] = 95
	fmt.Printf("%+v\n", studentGrades)

	alice, ok := studentGrades["Alice"]
	if ok {
		fmt.Printf("%+v\n", alice)
	}

	key := "Dan"
	if value, ok := studentGrades[key]; ok {
		fmt.Printf("%s: %+v\n", key, value)
	}

	delete(studentGrades, "Alice")
	fmt.Printf("%+v\n", studentGrades)

	configs := make(map[string]int)
	fmt.Printf("%+v %T\n", configs, configs)

	if configs == nil {
		fmt.Printf("Config is nil")
	}
}

// ============================================================
// WORKING WITH MAPS IN GO
// ============================================================
//
// A MAP stores data in KEY-VALUE pairs.
//
// Example:
//
//     studentGrades := map[string]int{
//         "Alice": 90,
//         "James": 85,
//         "Dan":   60,
//     }
//
// Think of it as:
//
//     Student Name → Grade
//
//     Alice → 90
//     James → 85
//     Dan   → 60
//
// ------------------------------------------------------------
//
// 1. CREATING A MAP
//
//     map[string]int
//
// means:
//
//     key   → string
//     value → int
//
// General syntax:
//
//     map[keyType]valueType
//
// Example:
//
//     map[string]int
//     map[int]string
//     map[string]float64
//
// ------------------------------------------------------------
//
// 2. ACCESSING A VALUE
//
// We can access a map value using its KEY:
//
//     studentGrades["Alice"]
//
// This gives:
//
//     90
//
// Similar to an array/slice:
//
//     array[index]
//
// But maps use KEYS instead of numeric indexes.
//
//
//
// ARRAY / SLICE:
//
//     numbers[0]
//
//     → uses an index
//
//
//
// MAP:
//
//     studentGrades["Alice"]
//
//     → uses a key
//
// ------------------------------------------------------------
//
// 3. UPDATING A VALUE
//
// Maps are mutable.
//
// We can change the value associated with an existing key:
//
//     studentGrades["Alice"] = 95
//
// Before:
//
//     Alice → 90
//
// After:
//
//     Alice → 95
//
// ------------------------------------------------------------
//
// 4. ADDING A NEW KEY-VALUE PAIR
//
// We can also use the same syntax to add a completely
// new key:
//
//     studentGrades["Bob"] = 88
//
// Now:
//
//     Bob → 88
//
// So:
//
//     map[key] = value
//
// can either:
//
//     - update an existing key
//     - add a new key
//
// ------------------------------------------------------------
//
// 5. CHECKING WHETHER A KEY EXISTS
//
// One of the MOST IMPORTANT patterns when working with
// maps is:
//
//     value, ok := studentGrades["Alice"]
//
// Go gives us TWO values:
//
//     value → value stored for the key
//     ok    → whether the key actually exists
//
// If "Alice" exists:
//
//     value = 95
//     ok    = true
//
// If "Bob" does NOT exist:
//
//     value = 0
//     ok    = false
//
// The `0` comes from the ZERO VALUE of int.
//
// ------------------------------------------------------------
//
// 6. `ok` IS NOT A KEYWORD
//
// `ok` is just a normal variable name.
//
// You could write:
//
//     value, ok := studentGrades["Alice"]
//
// Or:
//
//     value, found := studentGrades["Alice"]
//
// Or:
//
//     value, exists := studentGrades["Alice"]
//
// All of them work.
//
// The important thing is the TWO-VALUE MAP LOOKUP:
//
//     value, boolean := map[key]
//
// Usually Go programmers use:
//
//     value, ok
//
// because it is a very common convention.
//
// ------------------------------------------------------------
//
// 7. WHY DO WE NEED `ok`?
//
// Suppose we write:
//
//     grade := studentGrades["Unknown"]
//
// If "Unknown" doesn't exist, Go gives us:
//
//     0
//
// But now we don't know:
//
//     Does the student actually have grade 0?
//
// OR:
//
//     Does the student not exist?
//
// The `ok` value solves this:
//
//     grade, ok := studentGrades["Unknown"]
//
// If:
//
//     ok == true
//
// the key exists.
//
// If:
//
//     ok == false
//
// the key does not exist.
//
// ------------------------------------------------------------
//
// 8. MAP LOOKUP DIRECTLY INSIDE IF
//
// Go allows us to combine the lookup and condition:
//
//     key := "Dan"
//
//     if value, ok := studentGrades[key]; ok {
//         fmt.Printf("%s: %d\n", key, value)
//     }
//
// This means:
//
//     1. Look for `key` in the map.
//     2. Store the value in `value`.
//     3. Store whether it exists in `ok`.
//     4. If `ok` is true, execute the if block.
//
// So:
//
//     if value, ok := map[key]; ok {
//
// means:
//
//     "If this key exists, give me its value and execute
//      this block."
//
// ------------------------------------------------------------
//
// 9. DELETING FROM A MAP
//
// Go provides a built-in `delete()` function:
//
//     delete(studentGrades, "Alice")
//
// This removes the key:
//
//     "Alice"
//
// and its associated value.
//
// Before:
//
//     Alice → 95
//     James → 85
//     Dan   → 60
//
// After:
//
//     James → 85
//     Dan   → 60
//
// Syntax:
//
//     delete(mapName, key)
//
// ------------------------------------------------------------
//
// 10. CREATING AN EMPTY MAP WITH `make()`
//
// We can create an empty map using:
//
//     configs := make(map[string]int)
//
// This creates an initialized map with:
//
//     key   → string
//     value → int
//
// Currently it contains no elements.
//
//     {}
//
// But it is READY TO USE.
//
// For example:
//
//     configs["timeout"] = 30
//
// works perfectly.
//
// ------------------------------------------------------------
//
// 11. NIL MAP
//
// We can also declare a map without initializing it:
//
//     var configs map[string]int
//
// This creates a NIL MAP.
//
// We can check:
//
//     if configs == nil {
//
//         fmt.Println("Config is nil")
//
//     }
//
// The result is:
//
//     true
//
// ------------------------------------------------------------
//
// 12. NIL MAP vs EMPTY MAP
//
// These two can look similar but are different:
//
//     configs := make(map[string]int)
//
// and:
//
//     var configs map[string]int
//
//
//
// `make()`:
//
//     configs := make(map[string]int)
//
//     → initialized
//     → empty
//     → can write to it
//
//
//
// `var`:
//
//     var configs map[string]int
//
//     → nil
//     → no entries
//     → CANNOT write to it
//
// ------------------------------------------------------------
//
// IMPORTANT:
//
// Reading from a nil map is allowed:
//
//     value := configs["test"]
//
// But writing to a nil map causes a RUNTIME PANIC:
//
//     configs["test"] = 10 // ❌ panic
//
// Therefore, if you want to write to a map, initialize it:
//
//     configs := make(map[string]int)
//
// ------------------------------------------------------------
//
// 13. `%T`
//
// In:
//
//     fmt.Printf("%+v %T\n", configs, configs)
//
// `%T` prints the TYPE of a value.
//
// Example:
//
//     configs := make(map[string]int)
//
//     fmt.Printf("%T\n", configs)
//
// Output:
//
//     map[string]int
//
// So:
//
//     %v  → value
//     %T  → type
//
// ------------------------------------------------------------
//
// 14. MAPS ARE UNORDERED
//
// Maps do NOT guarantee a particular order when their
// contents are printed or iterated.
//
// For example:
//
//     studentGrades := map[string]int{
//         "Alice": 90,
//         "James": 85,
//         "Dan":   60,
//     }
//
// You might see:
//
//     map[Alice:90 James:85 Dan:60]
//
// But you should NOT depend on that order.
//
// Another execution could display the entries in a
// different order.
//
// ------------------------------------------------------------
//
// 15. MAP KEYS
//
// Map keys must be of a type that Go can compare.
//
// Common examples:
//
//     map[string]int
//     map[int]string
//     map[bool]string
//
// A slice cannot be used directly as a map key:
//
//     map[[]int]string // ❌
//
// ------------------------------------------------------------
//
// 16. MAP vs ARRAY vs SLICE
//
// ARRAY:
//
//     [3]int
//
//     Uses indexes.
//     Fixed size.
//
//
//
// SLICE:
//
//     []int
//
//     Uses indexes.
//     Dynamic size.
//
//
//
// MAP:
//
//     map[string]int
//
//     Uses keys.
//     Stores key-value pairs.
//
// Think:
//
//     ARRAY / SLICE
//
//     index → value
//
//     0 → Alice
//     1 → James
//     2 → Dan
//
//
//
// MAP
//
//     key → value
//
//     "Alice" → 90
//     "James" → 85
//     "Dan"   → 60
//
// ------------------------------------------------------------
//
// 17. COMPLETE MAP FLOW
//
//     Create Map
//          ↓
//     Add key-value pairs
//          ↓
//     Access using map[key]
//          ↓
//     Update using map[key] = value
//          ↓
//     Check existence using value, ok
//          ↓
//     Delete using delete(map, key)
//
// ------------------------------------------------------------
//
// QUICK REVISION:
//
//     map[string]int
//     → string keys with int values
//
//     map[key]
//     → access a value
//
//     map[key] = value
//     → add or update a value
//
//     value, ok := map[key]
//     → get value + check whether key exists
//
//     ok
//     → normal variable, commonly used by convention
//
//     delete(map, key)
//     → remove an entry
//
//     make(map[string]int)
//     → create an initialized empty map
//
//     var m map[string]int
//     → creates a nil map
//
//     len(map)
//     → number of key-value pairs
//
//     %T
//     → prints the type
//
// MOST IMPORTANT:
//
//     value, ok := myMap[key]
//
//     value → what was stored
//     ok    → whether the key existed
//
// ============================================================
// END OF MAP NOTES
// ============================================================