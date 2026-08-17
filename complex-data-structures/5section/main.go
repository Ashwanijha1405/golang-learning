package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Println("------------------ Advanced Slicing Operations -------------------")

	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf(
		"Original: %v, len: %d, cap: %d\n",
		original,
		len(original),
		cap(original),
	)

	s1 := original[2:5]
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))

	s2 := original[:4]
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	s3 := original[6:]
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))

	s4 := original[:]
	fmt.Printf("s4: %v, len: %d, cap: %d\n", s4, len(s4))

	// Check whether 8 exists in s4
	if slices.Contains(s4, 8) {
		fmt.Println("8 exists in s4")
	}

	// Insert 1000 at index 3
	s4 = slices.Insert(s4, 3, 1000)

	fmt.Printf(
		"s4 modified: %v, len: %d, cap: %d\n",
		s4,
		len(s4),
		cap(s4),
	)
}

// ============================================================
// ADVANCED SLICING OPERATIONS IN GO
// ============================================================
//
// A slice is a flexible view over an underlying array.
//
// Important:
//
//     A slice does NOT necessarily own its own data.
//
// Multiple slices can refer to the SAME underlying array.
//
// This becomes very important when working with:
//
//     - slicing expressions
//     - len()
//     - cap()
//     - append()
//     - slices.Contains()
//     - slices.Insert()
//
// ------------------------------------------------------------
//
// 1. ORIGINAL SLICE
//
//     original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//
// This creates a slice:
//
//     index:   0 1 2 3 4 5 6 7 8 9
//     value:   0 1 2 3 4 5 6 7 8 9
//
// len(original):
//
//     10
//
// cap(original):
//
//     10
//
// Since the slice starts at index 0 and currently has access
// to the entire underlying array:
//
//     len = 10
//     cap = 10
//
// ------------------------------------------------------------
//
// 2. BASIC SLICE EXPRESSION
//
//     original[2:5]
//
// General syntax:
//
//     slice[start:end]
//
// IMPORTANT:
//
//     start → INCLUDED
//     end   → EXCLUDED
//
// Therefore:
//
//     original[2:5]
//
// gives:
//
//     [2, 3, 4]
//
// NOT:
//
//     [2, 3, 4, 5]
//
// Think:
//
//     2 ≤ index < 5
//
// ------------------------------------------------------------
//
// 3. S1
//
//     s1 := original[2:5]
//
// Original:
//
//     index:  0 1 2 3 4 5 6 7 8 9
//     value:  0 1 2 3 4 5 6 7 8 9
//
// s1:
//
//     [2 3 4]
//
// len(s1):
//
//     3
//
//
//
// BUT:
//
//     cap(s1) = 8
//
// Why?
//
// Because the slice starts at index 2 and can potentially
// access the rest of the underlying array:
//
//     index 2 → 3 → 4 → 5 → 6 → 7 → 8 → 9
//
// That's 8 elements.
//
// So:
//
//     len = number of currently visible elements
//     cap = number of elements available from the starting
//            position to the end of the underlying array
//
// ------------------------------------------------------------
//
// 4. `len()` vs `cap()`
//
// `len()`:
//
//     How many elements are currently inside the slice?
//
// `cap()`:
//
//     How much capacity does the slice have from its starting
//     position before reaching the end of its backing array?
//
// Example:
//
//     s1 := original[2:5]
//
//     s1 = [2 3 4]
//
//     len(s1) = 3
//     cap(s1) = 8
//
// Visual:
//
//     original:
//
//     [0 1 | 2 3 4 | 5 6 7 8 9]
//           ↑       ↑
//         start    len ends
//
//     s1:
//
//     [2 3 4]
//
//     But its underlying array still has:
//
//     [2 3 4 5 6 7 8 9]
//
//     Therefore:
//
//     len = 3
//     cap = 8
//
// ------------------------------------------------------------
//
// 5. S2 — OMITTING THE START
//
//     s2 := original[:4]
//
// If we don't specify the start:
//
//     original[:4]
//
// means:
//
//     start from index 0
//
// So:
//
//     original[:4]
//
// is equivalent to:
//
//     original[0:4]
//
// Result:
//
//     [0 1 2 3]
//
// Therefore:
//
//     len(s2) = 4
//     cap(s2) = 10
//
// ------------------------------------------------------------
//
// 6. S3 — OMITTING THE END
//
//     s3 := original[6:]
//
// If we don't specify the end:
//
//     original[6:]
//
// means:
//
//     start at index 6
//     continue until the end
//
// Result:
//
//     [6 7 8 9]
//
// Therefore:
//
//     len(s3) = 4
//     cap(s3) = 4
//
// Why is capacity 4?
//
// Because the slice starts at index 6:
//
//     6 → 7 → 8 → 9
//
// There are only 4 elements available until the end.
//
// ------------------------------------------------------------
//
// 7. COPYING THE ENTIRE SLICE
//
//     s4 := original[:]
//
// This means:
//
//     start = 0
//     end = end of slice
//
// Equivalent to:
//
//     s4 := original[0:]
//
// Result:
//
//     [0 1 2 3 4 5 6 7 8 9]
//
// Therefore:
//
//     len(s4) = 10
//     cap(s4) = 10
//
// ------------------------------------------------------------
//
// 8. IMPORTANT: SLICES SHARE THE UNDERLYING ARRAY
//
// This is one of the MOST IMPORTANT concepts.
//
// When we write:
//
//     s1 := original[2:5]
//
// Go does NOT necessarily create a brand-new array containing
// 2, 3 and 4.
//
// Instead:
//
//     s1
//      ↓
//     points into the same underlying array as original.
//
// Conceptually:
//
//     original
//
//     [0 1 2 3 4 5 6 7 8 9]
//         ↑
//         │
//         └── s1 starts here
//
// Therefore, modifying elements through one slice can affect
// another slice that shares the same underlying array.
//
// ------------------------------------------------------------
//
// 9. THE `slices` PACKAGE
//
// Modern Go provides the `slices` package for common
// slice operations.
//
// Import it using:
//
//     import "slices"
//
// It provides useful functions such as:
//
//     slices.Contains()
//     slices.Insert()
//     slices.Delete()
//     slices.Sort()
//     slices.Reverse()
//     slices.Equal()
//
// ------------------------------------------------------------
//
// 10. `slices.Contains()`
//
//     slices.Contains(s4, 8)
//
// Checks whether the slice contains a particular value.
//
// It returns a boolean:
//
//     true
//     or
//     false
//
// Example:
//
//     if slices.Contains(s4, 8) {
//         fmt.Println("8 exists")
//     }
//
// If:
//
//     s4 = [0 1 2 3 4 5 6 7 8 9]
//
// then:
//
//     slices.Contains(s4, 8)
//     → true
//
// While:
//
//     slices.Contains(s4, 100)
//     → false
//
// ------------------------------------------------------------
//
// 11. `slices.Insert()`
//
//     slices.Insert(s4, 3, 1000)
//
// General syntax:
//
//     slices.Insert(slice, index, value)
//
// Meaning:
//
//     Insert `value` at `index`.
//
// Example:
//
//     s4 = [0 1 2 3 4 5]
//
//     slices.Insert(s4, 3, 1000)
//
// Result:
//
//     [0 1 2 1000 3 4 5]
//
// The existing elements from index 3 onward are shifted
// to the right.
//
// ------------------------------------------------------------
//
// 12. IMPORTANT: `slices.Insert()` RETURNS A SLICE
//
// Therefore, normally write:
//
//     s4 = slices.Insert(s4, 3, 1000)
//
// instead of simply:
//
//     slices.Insert(s4, 3, 1000)
//
// Why?
//
// Because the returned slice may have a different underlying
// array or different length/capacity.
//
// So the safe/common pattern is:
//
//     slice = slices.Insert(slice, index, value)
//
// ------------------------------------------------------------
//
// 13. `slices.Insert()` DOES NOT USE `i:` OR `v:`
//
// This is NOT valid Go:
//
//     slices.Insert(s4, i:1000)
//
// Go does not use named arguments like that.
//
// Correct:
//
//     slices.Insert(s4, 3, 1000)
//
// Meaning:
//
//     slice = s4
//     index = 3
//     value = 1000
//
// ------------------------------------------------------------
//
// 14. IMPORTANT SLICING PATTERNS
//
//     original[2:5]
//
//     → from index 2 up to, but NOT including, 5
//
//
//
//     original[:5]
//
//     → from beginning up to, but NOT including, 5
//
//
//
//     original[2:]
//
//     → from index 2 until the end
//
//
//
//     original[:]
//
//     → entire slice
//
// ------------------------------------------------------------
//
// 15. QUICK EXAMPLE
//
//     numbers := []int{10, 20, 30, 40, 50}
//
//     a := numbers[1:4]
//
// Result:
//
//     a = [20 30 40]
//
// Because:
//
//     index:   0   1   2   3   4
//     value:  10  20  30  40  50
//                  └───────┘
//
//     len(a) = 3
//
//     cap(a) = 4
//
// Why?
//
//     a starts at index 1.
//
//     Available elements from index 1:
//
//     20 30 40 50
//
//     → 4 elements
//
// ------------------------------------------------------------
//
// 16. SLICE MENTAL MODEL
//
// Think of a slice as a VIEW into an underlying array.
//
// Example:
//
//     original:
//
//     [0 1 2 3 4 5 6 7 8 9]
//           ↑
//           │
//           s1 := original[2:5]
//
//     s1:
//
//     [2 3 4]
//
// `s1` is a view over part of `original`.
//
// This is why slices are lightweight and useful.
//
// ------------------------------------------------------------
//
// 17. COMMON MISTAKE
//
// DON'T think:
//
//     original[2:5]
//
// means:
//
//     "Create a completely independent array containing
//      2, 3 and 4."
//
// Instead, initially think:
//
//     "Give me a slice/view of the original data from
//      index 2 up to index 5."
//
// ------------------------------------------------------------
//
// QUICK REVISION:
//
//     slice[start:end]
//     → start included, end excluded
//
//     slice[:end]
//     → starts from 0
//
//     slice[start:]
//     → continues to the end
//
//     slice[:]
//     → entire slice
//
//     len(slice)
//     → number of visible elements
//
//     cap(slice)
//     → available capacity from the slice's starting position
//
//     slices.Contains(slice, value)
//     → checks whether value exists
//
//     slices.Insert(slice, index, value)
//     → inserts value at index
//
//     slice = slices.Insert(...)
//     → store the returned slice
//
// ------------------------------------------------------------
//
// MOST IMPORTANT:
//
//     len = "How many elements do I currently have?"
//
//     cap = "How much room do I have available from my
//            starting position?"
//
//     [start:end]
//
//     start → INCLUDED
//     end   → EXCLUDED
//
// ============================================================
// END OF ADVANCED SLICING NOTES
// ============================================================