package main 

import "fmt"

func main() {

	names := []string{"Alice", "John", "Mark"} // if we dont put any size then its slice not array
	fmt.Println(names)

	items := make([]int, 3, 5) 
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)
	items = append(items, 4)

	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items)) // capacity increases to original cap as soon as it crossed previous capacity 
}


// ============================================================
// SLICES IN GO
// ============================================================
//
// This section covers:
//   - Slices
//   - Slice literals
//   - make()
//   - len()
//   - cap()
//   - append()
//   - Difference between arrays and slices
//
// ------------------------------------------------------------
//
// 1. ARRAY vs SLICE
//
// ARRAY:
//
//     var numbers [3]int
//
// `[3]int` means:
//     - fixed size
//     - exactly 3 integers
//
// The size is part of the type.
//
//
//
// SLICE:
//
//     names := []string{"Alice", "John", "Mark"}
//
// `[]string` means:
//     - slice of strings
//     - size is NOT fixed
//     - can grow or shrink
//
// Unlike an array, the length of a slice is not part of
// its type.
//
// ------------------------------------------------------------
//
// 2. SLICE LITERAL
//
//     names := []string{"Alice", "John", "Mark"}
//
// Go automatically determines the length.
//
// Therefore:
//
//     len(names) → 3
//
// We can add more elements later using append():
//
//     names = append(names, "Bob")
//
// Now:
//
//     len(names) → 4
//
// This is one of the biggest differences between arrays
// and slices.
//
// ------------------------------------------------------------
//
// 3. `make()`
//
//     items := make([]int, 3, 5)
//
// `make()` is commonly used to create a slice with a
// specified length and capacity.
//
// Syntax:
//
//     make([]type, length, capacity)
//
// So:
//
//     make([]int, 3, 5)
//
// means:
//
//     type     → int
//     length   → 3
//     capacity → 5
//
// Therefore:
//
//     items = [0 0 0]
//     len(items) = 3
//     cap(items) = 5
//
// ------------------------------------------------------------
//
// 4. WHAT IS LENGTH?
//
// `len()` tells us how many elements are currently inside
// the slice.
//
// Example:
//
//     items := make([]int, 3, 5)
//
//     len(items) → 3
//
// Even though the capacity is 5, only 3 elements currently
// belong to the slice.
//
// Think:
//
//     LENGTH = elements currently available
//
// ------------------------------------------------------------
//
// 5. WHAT IS CAPACITY?
//
// `cap()` tells us how much space the slice currently has
// available before Go needs to allocate a larger backing
// array.
//
// Example:
//
//     items := make([]int, 3, 5)
//
//     len(items) → 3
//     cap(items) → 5
//
// Think:
//
//     LENGTH   = how many elements we currently have
//     CAPACITY = how many elements can fit before resizing
//
// ------------------------------------------------------------
//
// 6. WHY IS LENGTH 3 BUT CAPACITY 5?
//
//     items := make([]int, 3, 5)
//
// Initially:
//
//     items
//     ↓
//     [0 0 0]
//      ↑     ↑
//      3 elements
//
// But Go has allocated room for up to 5 elements.
//
// So conceptually:
//
//     [0 0 0 _ _]
//      └─len=3─┘
//      └──cap=5───┘
//
// The last two positions are available capacity, not
// elements currently in the slice.
//
// ------------------------------------------------------------
//
// 7. `append()`
//
// `append()` adds elements to the end of a slice.
//
// Example:
//
//     items = append(items, 1)
//
// Before:
//
//     [0 0 0]
//
// After:
//
//     [0 0 0 1]
//
// Now:
//
//     len(items) → 4
//     cap(items) → 5
//
//
//
// Another:
//
//     items = append(items, 2)
//
// Now:
//
//     [0 0 0 1 2]
//
//     len(items) → 5
//     cap(items) → 5
//
// ------------------------------------------------------------
//
// 8. WHAT HAPPENS WHEN CAPACITY IS FULL?
//
// Suppose:
//
//     len = 5
//     cap = 5
//
// And we do:
//
//     items = append(items, 3)
//
// There isn't enough capacity in the current backing array.
//
// Go will allocate a new backing array with a larger
// capacity and copy the existing elements into it.
//
// Conceptually:
//
//     Old:
//
//     [0 0 0 1 2]
//      len=5 cap=5
//
//             append(3)
//                  ↓
//
//     New:
//
//     [0 0 0 1 2 3 ...]
//      └──── elements ────┘
//
// The exact new capacity is decided by Go's runtime.
// Don't assume it will always exactly double.
//
// ------------------------------------------------------------
//
// 9. IMPORTANT: `append()` RETURNS A SLICE
//
// This is important:
//
//     items = append(items, 1)
//
// NOT:
//
//     append(items, 1)
//
// We normally assign the result back to `items` because
// append() may return a slice backed by a new, larger array.
//
// So:
//
//     items = append(items, 1)
//
// means:
//
//     "Add 1 to the slice and store the resulting slice
//      back in items."
//
// ------------------------------------------------------------
//
// 10. YOUR EXAMPLE
//
//     items := make([]int, 3, 5)
//
// Initially:
//
//     items = [0 0 0]
//     len   = 3
//     cap   = 5
//
//
//     items = append(items, 1)
//
//     items = [0 0 0 1]
//     len   = 4
//     cap   = 5
//
//
//     items = append(items, 2)
//
//     items = [0 0 0 1 2]
//     len   = 5
//     cap   = 5
//
//
//     items = append(items, 3)
//
// Capacity is full.
//
// Go needs a larger backing array.
//
//     items = [0 0 0 1 2 3]
//     len   = 6
//     cap   = larger than 5
//
// ------------------------------------------------------------
//
// 11. VERY IMPORTANT: `make()` DOES NOT CREATE VALUES
//    ONLY WHEN LENGTH IS ZERO
//
// Compare:
//
//     make([]int, 3, 5)
//
// gives:
//
//     [0 0 0]
//
// because length = 3.
//
// But:
//
//     make([]int, 0, 5)
//
// gives:
//
//     []
//
// length = 0
// capacity = 5
//
// We can then append:
//
//     items = append(items, 10)
//
// Now:
//
//     [10]
//
//     len = 1
//     cap = 5
//
// This pattern is very common in real Go programs.
//
// ------------------------------------------------------------
//
// 12. SLICE INDEXING
//
// Slices use zero-based indexing just like arrays.
//
//     names := []string{"Alice", "John", "Mark"}
//
//     names[0] → "Alice"
//     names[1] → "John"
//     names[2] → "Mark"
//
// We can modify an existing element:
//
//     names[1] = "Bob"
//
// Now:
//
//     ["Alice", "Bob", "Mark"]
//
// ------------------------------------------------------------
//
// 13. KEY DIFFERENCE
//
// ARRAY:
//
//     [3]int
//
//     Fixed size.
//     Size is part of the type.
//
//
//
// SLICE:
//
//     []int
//
//     Dynamic size.
//     Can grow using append().
//
//     Internally, a slice is a small descriptor that refers
//     to an underlying (backing) array.
//
// ------------------------------------------------------------
//
// 14. QUICK REVISION
//
//     []int
//     → slice of integers
//
//     make([]int, 3, 5)
//     → slice with length 3 and capacity 5
//
//     len(slice)
//     → number of elements currently in the slice
//
//     cap(slice)
//     → capacity of the slice
//
//     append(slice, value)
//     → adds value to the end
//
//     slice[index]
//     → accesses an element
//
//     slice[index] = value
//     → modifies an element
//
// ------------------------------------------------------------
//
// MENTAL MODEL:
//
//     ARRAY
//
//     [ fixed size ]
//
//
//
//     SLICE
//
//     Slice
//       ↓
//     [ backing array ]
//       ↓
//     can grow using append()
//
//
//
// The most important thing to remember:
//
//     len = "How many elements do I have?"
//
//     cap = "How much room do I currently have?"
//
//     append = "Add another element."
//
// ============================================================