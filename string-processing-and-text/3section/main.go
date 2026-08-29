package main

import (
	"fmt"
	"unicode"
)

func main() {

	data := []rune{'律', '平', '等'}

	for _, v := range data {
		fmt.Printf(string(v), unicode.IsLetter(v))
	}

}

// ============================================================
// RUNE + UNICODE — NOTES
// ============================================================
//
// Go strings are UTF-8 encoded sequences of bytes.
// This becomes important when working with non-ASCII characters
// such as Chinese, Japanese, Hindi, emojis, etc.
//
// ------------------------------------------------------------
//
// 1. WHAT IS A RUNE?
//
// In Go:
//
//     rune
//
// is an alias for:
//
//     int32
//
// A rune represents a single Unicode code point.
//
// Example:
//
//     'A'     → one rune
//     '律'    → one rune
//     '平'    → one rune
//     '😀'    → one rune
//
// So:
//
//     []rune{'律', '平', '等'}
//
// creates a slice containing three Unicode code points.
//
// ------------------------------------------------------------
//
// 2. RUNE VS BYTE
//
// A string in Go is a sequence of BYTES.
//
// UTF-8 characters can occupy different numbers of bytes.
//
// ASCII character:
//
//     'A'
//
// takes 1 byte.
//
// A Unicode character such as:
//
//     '律'
//
// takes multiple bytes in UTF-8.
//
// Therefore:
//
//     len("A")      → 1
//
// but:
//
//     len("律")     → 3
//
// because len(string) counts BYTES, not characters.
//
// If we want to work with characters/code points:
//
//     len([]rune("律"))
//
// gives:
//
//     1
//
// ------------------------------------------------------------
//
// 3. CHARACTER LITERAL VS STRING
//
// Single quotes:
//
//     'A'
//
// represent a rune.
//
// Double quotes:
//
//     "A"
//
// represent a string.
//
// Example:
//
//     '律'   → rune
//     "律"   → string
//
// ------------------------------------------------------------
//
// 4. []rune
//
// We can convert a string into a rune slice:
//
//     text := "律平等"
//     data := []rune(text)
//
// Now each Unicode code point can be accessed individually.
//
//     data[0]
//     data[1]
//     data[2]
//
// This is useful when we need character-level operations.
//
// ------------------------------------------------------------
//
// 5. RANGE OVER []rune
//
// Example:
//
//     data := []rune{'律', '平', '等'}
//
//     for _, v := range data {
//         ...
//     }
//
// `v` is a rune (int32).
//
// So conceptually:
//
//     v
//     ↓
//     Unicode code point
//
// ------------------------------------------------------------
//
// 6. CONVERTING RUNE TO STRING
//
// A rune is an integer representing a Unicode code point.
//
// To display the actual character:
//
//     string(v)
//
// Example:
//
//     v = '律'
//
//     string(v)
//
// produces:
//
//     "律"
//
// This is why:
//
//     fmt.Println(string(v))
//
// prints the actual Unicode character.
//
// ------------------------------------------------------------
//
// 7. unicode PACKAGE
//
// Go provides the `unicode` package for working with Unicode
// characters.
//
// Import:
//
//     import "unicode"
//
// It contains useful functions such as:
//
//     unicode.IsLetter(r)
//     unicode.IsDigit(r)
//     unicode.IsNumber(r)
//     unicode.IsSpace(r)
//     unicode.IsUpper(r)
//     unicode.IsLower(r)
//
// These functions operate on runes.
//
// ------------------------------------------------------------
//
// 8. unicode.IsLetter()
//
//     unicode.IsLetter(r)
//
// checks whether a rune represents a Unicode letter.
//
// Example:
//
//     unicode.IsLetter('A')
//     → true
//
//     unicode.IsLetter('律')
//     → true
//
//     unicode.IsLetter('5')
//     → false
//
// This works across many writing systems, not just English.
//
// ------------------------------------------------------------
//
// 9. IMPORTANT: YOUR fmt.Printf() LINE
//
// Your code currently has:
//
//     fmt.Printf(string(v), unicode.IsLetter(v))
//
// This is NOT the correct way to print the character and the
// boolean result.
//
// `fmt.Printf()` expects the first argument to be a FORMAT STRING.
//
// For example:
//
//     fmt.Printf("%s: %t\n", string(v), unicode.IsLetter(v))
//
// Here:
//
//     %s → string(v)
//     %t → unicode.IsLetter(v)
//
// Example output:
//
//     律: true
//     平: true
//     等: true
//
// ------------------------------------------------------------
//
// 10. WHY fmt.Printf(string(v), ...) IS WRONG
//
// Suppose:
//
//     string(v) = "律"
//
// Then Go effectively receives:
//
//     fmt.Printf("律", true)
//
// But "律" contains no formatting verb such as:
//
//     %s
//     %t
//     %d
//
// Therefore the boolean argument has nowhere to be formatted.
//
// `Printf` should always receive a format string when you are
// supplying additional formatting arguments.
//
// ------------------------------------------------------------
//
// 11. BETTER WAY TO WRITE THE EXAMPLE
//
//     package main
//
//     import (
//         "fmt"
//         "unicode"
//     )
//
//     func main() {
//
//         data := []rune{'律', '平', '等'}
//
//         for _, v := range data {
//             fmt.Printf("%s: %t\n",
//                 string(v),
//                 unicode.IsLetter(v),
//             )
//         }
//     }
//
// ------------------------------------------------------------
//
// 12. RANGE OVER A STRING
//
// You don't always need to manually convert a string to []rune.
//
// Go's `range` over a string automatically decodes the UTF-8
// string and gives you runes.
//
// Example:
//
//     text := "律平等"
//
//     for _, r := range text {
//         fmt.Println(string(r))
//     }
//
// Output:
//
//     律
//     平
//     等
//
// This is different from indexing a string.
//
//     text[0]
//
// gives a BYTE, not a rune.
//
// ------------------------------------------------------------
//
// 13. STRING INDEXING VS RANGE
//
// Given:
//
//     text := "律"
//
// `text[0]`
//
// gives the first byte of the UTF-8 encoding.
//
// But:
//
//     for _, r := range text
//
// gives the actual Unicode rune.
//
// Therefore:
//
//     string indexing
//         ↓
//       bytes
//
//     range over string
//         ↓
//       runes
//
// ------------------------------------------------------------
//
// 14. WHEN SHOULD YOU USE []rune?
//
// Convert a string to []rune when you need to perform
// character-level operations such as:
//
//     • accessing characters by position
//     • modifying characters
//     • counting Unicode code points
//     • reversing Unicode text
//     • manipulating characters individually
//
// Example:
//
//     text := "律平等"
//     runes := []rune(text)
//
// ------------------------------------------------------------
//
// 15. IMPORTANT LIMITATION
//
// A rune represents a Unicode CODE POINT, not necessarily what a
// user visually considers one "character".
//
// For example, some visible characters can be made from multiple
// Unicode code points.
//
// So:
//
//     rune
//
// means:
//
//     Unicode code point
//
// not always:
//
//     visual character / grapheme
//
// For most basic Unicode work, however, runes are exactly what
// you need.
//
// ------------------------------------------------------------
//
// QUICK REFERENCE
//
//     rune
//     → alias for int32; represents a Unicode code point.
//
//     []rune
//     → slice of Unicode code points.
//
//     string(r)
//     → converts a rune into a string.
//
//     []rune(s)
//     → converts a string into runes.
//
//     len(s)
//     → number of BYTES.
//
//     len([]rune(s))
//     → number of Unicode code points.
//
//     range over string
//     → iterates through decoded runes.
//
//     unicode.IsLetter(r)
//     → checks whether rune is a Unicode letter.
//
//     unicode.IsDigit(r)
//     → checks whether rune is a Unicode digit.
//
//     unicode.IsSpace(r)
//     → checks whether rune is whitespace.
//
// ------------------------------------------------------------
//
// BIG PICTURE:
//
//             Go String
//                │
//                ↓
//          UTF-8 bytes
//                │
//                ↓
//       ┌─────────────────┐
//       │ Unicode decoding│
//       └────────┬────────┘
//                ↓
//             runes
//                │
//                ↓
//       Unicode code points
//
// Example:
//
//     "律平等"
//         ↓
//     []rune{'律', '平', '等'}
//         ↓
//     each element is a rune (int32)
//
// ============================================================
