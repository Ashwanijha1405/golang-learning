// ---------------------------------------- Working with Arrays ----------------------------------------- \\

package main 

import "fmt"

func main() {

	var numbers [2]int 
	fmt.Printf("%+v\n", numbers)

	numbers[0] = 1
	numbers[1] = 2
	fmt.Printf("%+v\n", numbers)

	primes := [4]int{2,3,5,7}
	fmt.Printf("%+v\n", primes)
	primes[3] = 11
	fmt.Printf("%+v\n", primes)

	for i := 0; i < len(primes); i++ {
		fmt.Printf("%d\n", primes[i])
	}

	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3

	fmt.Printf("%+v\n", matrix)
}

// ============================================================
// WORKING WITH ARRAYS IN GO
// ============================================================
//
// An ARRAY is a fixed-size collection of values of the SAME type.
//
// Example:
//
//     var numbers [2]int
//
// means:
//
//     [2]  → array has exactly 2 elements
//     int  → every element must be an int
//
// So the array looks like:
//
//     [0 0]
//
// ------------------------------------------------------------
//
// 1. DECLARING AN ARRAY
//
// var numbers [2]int
//
// This creates an array containing 2 integers.
//
// Go automatically gives each element its ZERO VALUE.
//
// For int:
//
//     zero value = 0
//
// Therefore:
//
//     var numbers [2]int
//
// produces:
//
//     [0 0]
//
// ------------------------------------------------------------
//
// 2. ACCESSING ARRAY ELEMENTS
//
// Array indexing starts from 0.
//
//     numbers[0] = 1
//     numbers[1] = 2
//
// The array becomes:
//
//     [1 2]
//
// IMPORTANT:
//
// For an array of size 2:
//
//     index 0 → first element
//     index 1 → second element
//
// There is NO index 2.
//
// Trying:
//
//     numbers[2] = 10
//
// causes a runtime error because index 2 doesn't exist.
//
// ------------------------------------------------------------
//
// 3. ARRAY LITERAL
//
// Instead of declaring an empty array and filling it manually:
//
//     var primes [4]int
//
//     primes[0] = 2
//     primes[1] = 3
//     primes[2] = 5
//     primes[3] = 7
//
// We can directly initialize it:
//
//     primes := [4]int{2, 3, 5, 7}
//
// This creates:
//
//     [2 3 5 7]
//
// ------------------------------------------------------------
//
// 4. ARRAYS HAVE FIXED SIZE
//
// This is one of the MOST IMPORTANT things about arrays in Go.
//
//     primes := [4]int{2, 3, 5, 7}
//
// means the array can contain EXACTLY 4 elements.
//
// You can change the values:
//
//     primes[3] = 11
//
// Result:
//
//     [2 3 5 11]
//
// But you CANNOT add another element:
//
//     primes = append(primes, 13) // ❌
//
// `append()` is used with SLICES, not arrays.
//
//
//
// ARRAY:
//
//     [4]int
//
//     Fixed size
//
//
//
// SLICE:
//
//     []int
//
//     Dynamic size
//
// We will learn slices separately.
//
// ------------------------------------------------------------
//
// 5. len() WITH ARRAYS
//
//     len(primes)
//
// returns the number of elements in the array.
//
// Example:
//
//     primes := [4]int{2, 3, 5, 7}
//
//     len(primes)
//
// returns:
//
//     4
//
// ------------------------------------------------------------
//
// 6. LOOPING THROUGH AN ARRAY
//
// We can use a normal for loop:
//
//     for i := 0; i < len(primes); i++ {
//
//         fmt.Printf("%+d\n", primes[i])
//
//     }
//
// Let's understand it:
//
//     i := 0
//
// Start from index 0.
//
//     i < len(primes)
//
// Continue while i is smaller than the array length.
//
//     i++
//
// Increase i by 1 after every iteration.
//
//     primes[i]
//
// Access the element at the current index.
//
// For:
//
//     primes := [4]int{2, 3, 5, 7}
//
// The loop accesses:
//
//     primes[0] → 2
//     primes[1] → 3
//     primes[2] → 5
//     primes[3] → 7
//
// ------------------------------------------------------------
//
// 7. MULTIDIMENSIONAL ARRAYS
//
// Go also supports arrays inside arrays.
//
// Example:
//
//     var matrix [2][3]int
//
// This means:
//
//     2 rows
//     3 columns
//
// Visually:
//
//     [ [0 0 0]
//       [0 0 0] ]
//
// So:
//
//     matrix[0][0]
//
// means:
//
//     row 0, column 0
//
// ------------------------------------------------------------
//
// 8. MODIFYING A MATRIX
//
//     matrix[0][0] = 1
//     matrix[0][1] = 2
//     matrix[0][2] = 3
//
// Now the matrix becomes:
//
//     [ [1 2 3]
//       [0 0 0] ]
//
// Remember:
//
//     matrix[row][column]
//
// ------------------------------------------------------------
//
// 9. PRINTING VALUES WITH fmt.Printf()
//
// `%v`
//
// Prints a value using its default format.
//
// Example:
//
//     fmt.Printf("%v\n", primes)
//
// Output:
//
//     [2 3 5 7]
//
//
// `%+v`
//
// Prints the value with additional formatting information.
//
// For simple arrays, `%+v` and `%v` will usually look the SAME.
//
// Example:
//
//     fmt.Printf("%+v\n", primes)
//
// Output:
//
//     [2 3 5 7]
//
// `%d`
//
// Used for integers.
//
// Example:
//
//     fmt.Printf("%d\n", primes[0])
//
// Output:
//
//     2
//
// ------------------------------------------------------------
//
// 10. IMPORTANT DIFFERENCE: Printf vs Println
//
// `fmt.Printf()` understands FORMAT VERBS:
//
//     fmt.Printf("%v\n", numbers)
//
// But `fmt.Println()` does NOT interpret `%v`, `%d`, etc.
//
// ❌ WRONG:
//
//     fmt.Println("%+v\n", matrix)
//
// This will literally print:
//
//     %+v
//
// along with the matrix as another argument.
//
//
//
// ✅ CORRECT:
//
//     fmt.Printf("%+v\n", matrix)
//
// OR simply:
//
//     fmt.Println(matrix)
//
// ------------------------------------------------------------
//
// 11. ARRAY LENGTH vs ARRAY INDEX
//
// Suppose:
//
//     numbers := [5]int{10, 20, 30, 40, 50}
//
// Length:
//
//     len(numbers) → 5
//
// Valid indexes:
//
//     0
//     1
//     2
//     3
//     4
//
// Notice:
//
//     length = 5
//     last index = 4
//
// Therefore:
//
//     last index = len(array) - 1
//
// ------------------------------------------------------------
//
// 12. IMPORTANT: ARRAYS ARE VALUES IN GO
//
// Arrays in Go are value types.
//
// Example:
//
//     a := [3]int{1, 2, 3}
//     b := a
//
// Now `b` gets a COPY of `a`.
//
// If:
//
//     b[0] = 100
//
// Then:
//
//     a → [1 2 3]
//     b → [100 2 3]
//
// Changing `b` does NOT change `a`.
//
// This is different from how slices behave.
//
// ------------------------------------------------------------
//
// QUICK REVISION:
//
//     [2]int
//     → array of 2 integers
//
//     [4]string
//     → array of 4 strings
//
//     [3]float64
//     → array of 3 floats
//
//     len(array)
//     → number of elements
//
//     array[index]
//     → access an element
//
//     array[index] = value
//     → modify an element
//
//     [4]int{1, 2, 3, 4}
//     → array literal
//
//     [2][3]int
//     → 2D array with 2 rows and 3 columns
//
//     fmt.Printf("%v\n", array)
//     → formatted printing
//
//     fmt.Println(array)
//     → simple printing
//
// MOST IMPORTANT:
//
//     ARRAY = FIXED SIZE
//     SLICE = DYNAMIC SIZE
//
// ------------------------------------------------------------
// END OF ARRAY NOTES
// ============================================================