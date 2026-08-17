package main

import "fmt"

func modifyValue(val int) {

	val = val * 10
	fmt.Printf("modifyValue: %+v\n", val)
}

func modifyPointer(val *int) {
	if val == nil {
		fmt.Printf("val is nil")
		return 
	}
	*val = *val * 10 // derefencing
	fmt.Printf("modifyPointer: %+v\n", val)
}

func main() {

	num := 10
	modifyValue(num)
	fmt.Println(num)

	modifyPointer(&num)
	fmt.Println(num)

	grade := 50
	gradePtr := &grade
	fmt.Printf("gradePtr grade: %+v\n", gradePtr)
	fmt.Printf("gradePtr: %+v\n", *(&gradePtr))
}


// ============================================================
// POINTERS IN GO
// ============================================================
//
// A pointer is a variable that stores the MEMORY ADDRESS
// of another variable.
//
// The two most important operators are:
//
//     &  → gives the address of a variable
//     *  → dereferences a pointer and gives the value stored
//          at that address
//
// Think:
//
//     &  → "WHERE is the value?"
//     *  → "WHAT is at that address?"
//
// ------------------------------------------------------------
//
// 1. NORMAL VARIABLE
//
//     num := 10
//
// `num` stores the actual value:
//
//     num → 10
//
// Somewhere in memory, Go stores this value.
//
// Every variable has:
//     - a value
//     - a memory address
//
// ------------------------------------------------------------
//
// 2. `&` — ADDRESS-OF OPERATOR
//
//     &num
//
// means:
//
//     "Give me the memory address of num."
//
// For example, conceptually:
//
//     num
//      ↓
//     Address: 1000
//     Value:   10
//
// Therefore:
//
//     &num
//
// gives:
//
//     1000
//
// In an actual Go program, the address may look like:
//
//     0xc0000120c0
//
// ------------------------------------------------------------
//
// 3. WHAT IS A POINTER?
//
// A pointer is a variable that stores a MEMORY ADDRESS.
//
// Example:
//
//     num := 10
//     ptr := &num
//
// Now:
//
//     num
//      ↓
//     10
//
//     ptr
//      ↓
//     address of num
//
// So:
//
//     num     → actual value
//     ptr     → address of that value
//
// ------------------------------------------------------------
//
// 4. DECLARING A POINTER
//
//     var ptr *int
//
// `*int` means:
//
//     "pointer to an int"
//
// So:
//
//     int
//     → actual integer value
//
//     *int
//     → pointer to an integer
//
// Example:
//
//     num := 10
//     var ptr *int
//     ptr = &num
//
// Now ptr points to num.
//
// ------------------------------------------------------------
//
// 5. `*` HAS TWO MEANINGS IN GO
//
// This is VERY important.
//
// When used during a declaration:
//
//     var ptr *int
//
// `*int` means:
//
//     "ptr is a pointer to an int."
//
//
//
// When used with an existing pointer:
//
//     *ptr
//
// it means:
//
//     "Go to the address stored in ptr and give me the
//      value stored there."
//
// This is called DEREFERENCING.
//
// ------------------------------------------------------------
//
// 6. DEREFERENCING
//
// Suppose:
//
//     num := 10
//     ptr := &num
//
// Then:
//
//     ptr
//
// contains the address of num.
//
// But:
//
//     *ptr
//
// gives the actual value stored at that address:
//
//     *ptr → 10
//
// So:
//
//     num
//     → 10
//
//     &num
//     → address of num
//
//     ptr
//     → address of num
//
//     *ptr
//     → 10
//
// ------------------------------------------------------------
//
// 7. PASSING A NORMAL VALUE TO A FUNCTION
//
//     func modifyValue(val int) {
//
//         val = val * 10
//
//     }
//
// Here:
//
//     val int
//
// means the function receives an integer VALUE.
//
// Example:
//
//     num := 10
//     modifyValue(num)
//
// Go passes the value to the function.
//
// Conceptually:
//
//     main:
//
//     num → 10
//
//             ↓ copy
//
//     modifyValue:
//
//     val → 10
//
// `val` is a separate copy.
//
// Therefore:
//
//     val = val * 10
//
// changes:
//
//     val → 100
//
// but the original:
//
//     num → 10
//
// remains unchanged.
//
// ------------------------------------------------------------
//
// 8. PASS-BY-VALUE
//
// Go commonly passes values by VALUE.
//
// So:
//
//     modifyValue(num)
//
// does NOT give the function direct access to `num`.
//
// It gives the function a copy of its value.
//
// Therefore:
//
//     modifyValue: 100
//     main:        10
//
//
//
// IMPORTANT:
//
// Go does not have a special "pass by reference" parameter
// syntax like some languages.
//
// Instead, we can pass a POINTER when we want a function
// to modify the original value.
//
// ------------------------------------------------------------
//
// 9. POINTER FUNCTION
//
//     func modifyPointer(val *int)
//
// `val *int` means:
//
//     val is a pointer to an int.
//
// Therefore, the function receives an ADDRESS rather than
// just a copied integer value.
//
// ------------------------------------------------------------
//
// 10. PASSING THE ADDRESS
//
//     num := 10
//
//     modifyPointer(&num)
//
// `&num` means:
//
//     "Give the function the address of num."
//
// So inside the function:
//
//     val
//
// contains the address of num.
//
// Conceptually:
//
//     num
//      ↓
//     Address: 1000
//     Value:   10
//
//     val
//      ↓
//     1000
//
// Therefore:
//
//     val → address of num
//
// ------------------------------------------------------------
//
// 11. DEREFERENCING INSIDE THE FUNCTION
//
//     *val = *val * 10
//
// Break it down:
//
//     *val
//
// means:
//
//     "Go to the address stored in val and get the value."
//
// If:
//
//     val → address of num
//
// and:
//
//     num = 10
//
// then:
//
//     *val → 10
//
// Therefore:
//
//     *val = *val * 10
//
// becomes:
//
//     10 = 10 * 10
//
// resulting in:
//
//     num = 100
//
// Because `val` points directly to the original `num`.
//
// ------------------------------------------------------------
//
// 12. WHY DID `num` CHANGE?
//
// Normal value:
//
//     modifyValue(num)
//
//     num → 10
//     val → 10 (copy)
//
// Changing val does NOT change num.
//
//
//
// Pointer:
//
//     modifyPointer(&num)
//
//     num → 10
//           ↑
//           │
//     val ──┘
//
// `val` points to the original num.
//
// Therefore:
//
//     *val = 100
//
// changes the original:
//
//     num → 100
//
// ------------------------------------------------------------
//
// 13. `nil` POINTER
//
// A pointer can contain no address.
//
// Example:
//
//     var ptr *int
//
// Here:
//
//     ptr → nil
//
// `nil` basically means:
//
//     "This pointer is not pointing anywhere."
//
// We can check:
//
//     if ptr == nil {
//
//         fmt.Println("Pointer is nil")
//
//     }
//
// ------------------------------------------------------------
//
// 14. WHY CHECK FOR `nil`?
//
// Dereferencing a nil pointer:
//
//     *ptr
//
// is invalid.
//
// It will cause a RUNTIME PANIC.
//
// Therefore, when a pointer might be nil, we can check:
//
//     if val == nil {
//         return
//     }
//
// before doing:
//
//     *val
//
// This protects the program from dereferencing an invalid
// pointer.
//
// ------------------------------------------------------------
//
// 15. POINTER EXAMPLE
//
//     grade := 50
//     gradePtr := &grade
//
// Now:
//
//     grade
//       ↓
//      50
//
//     gradePtr
//       ↓
//     address of grade
//
// Therefore:
//
//     grade       → 50
//     &grade      → address of grade
//     gradePtr    → address of grade
//     *gradePtr   → 50
//
// ------------------------------------------------------------
//
// 16. `*gradePtr`
//
// If:
//
//     gradePtr = &grade
//
// then:
//
//     *gradePtr
//
// means:
//
//     "Go to the address stored inside gradePtr and retrieve
//      the value there."
//
// Therefore:
//
//     *gradePtr → 50
//
// We can also modify the original value:
//
//     *gradePtr = 100
//
// Now:
//
//     grade → 100
//
// because gradePtr points to grade.
//
// ------------------------------------------------------------
//
// 17. UNDERSTANDING `*(&gradePtr)`
//
// This expression can look confusing:
//
//     *(&gradePtr)
//
// Break it from inside out.
//
// First:
//
//     &gradePtr
//
// means:
//
//     "Give me the address of gradePtr."
//
// Then:
//
//     *(&gradePtr)
//
// means:
//
//     "Go to that address and retrieve the value stored there."
//
// The value stored there is simply `gradePtr`.
//
// Therefore:
//
//     *(&gradePtr)
//
// is effectively the same as:
//
//     gradePtr
//
// This demonstrates that `&` and `*` can conceptually
// cancel each other when used together correctly.
//
// ------------------------------------------------------------
//
// 18. MEMORY MODEL
//
// Suppose:
//
//     grade := 50
//     gradePtr := &grade
//
// Conceptually:
//
//                    MEMORY
//
//              ┌─────────────────┐
//              │      grade      │
//              │                 │
//              │       50        │
//              └─────────────────┘
//                     ▲
//                     │
//                     │ address
//                     │
//              ┌──────┴──────────┐
//              │    gradePtr     │
//              │                 │
//              │   0xc000...     │
//              └─────────────────┘
//
// gradePtr stores the address where grade lives.
//
// ------------------------------------------------------------
//
// 19. THE FOUR THINGS TO MEMORIZE
//
//     num := 10
//
// Creates a normal variable.
//
//
//
//     &num
//
// Gives the ADDRESS of num.
//
//
//
//     ptr := &num
//
// Creates a pointer storing the address of num.
//
//
//
//     *ptr
//
// DEREFERENCES ptr and gives the VALUE stored at that address.
//
// ------------------------------------------------------------
//
// 20. SIMPLE MENTAL MODEL
//
// Think of a house.
//
// The VALUE is what is inside the house:
//
//     House → 50
//
// The ADDRESS tells you where the house is:
//
//     House → "123 Main Street"
//
// A POINTER is like writing that address on a piece of paper:
//
//     pointer → "123 Main Street"
//
// Then:
//
//     *pointer
//
// means:
//
//     "Go to that address and see what's inside."
//
// And:
//
//     &variable
//
// means:
//
//     "Tell me the address of this variable."
//
// ------------------------------------------------------------
//
// QUICK REVISION:
//
//     &variable
//     → address of the variable
//
//     *pointer
//     → value stored at the pointer's address
//
//     *int
//     → pointer to an int
//
//     ptr := &num
//     → ptr stores the address of num
//
//     *ptr = value
//     → modifies the value at that address
//
//     nil
//     → pointer is not pointing anywhere
//
//     val *int
//     → function receives a pointer to an integer
//
// ------------------------------------------------------------
//
// MOST IMPORTANT RELATIONSHIP:
//
//             &
//             ↓
//
//     VALUE ───────→ ADDRESS
//                       ↑
//                       │
//                    POINTER
//                       │
//                       ↓
//             * ─────── VALUE
//
// Remember:
//
//     & → "WHERE is it?"
//
//     * → "WHAT is there?"
//
// ============================================================
// END OF POINTER NOTES
// ============================================================