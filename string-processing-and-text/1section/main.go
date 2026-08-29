package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "abc"
	s2 := strings.Clone(s1)

	fmt.Println(s2)

	b := strings.Builder{}
	b.WriteString("Here is an example") // also can use b.Write([]byte("Here is an example"))

	fmt.Println(b.String()) 

	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s1))

	s3 := "     test sss      "
	fmt.Println("s3", len(s3))

	s3 = strings.TrimSpace(s3)
	fmt.Println("s3 after trim", len(s3))

	fmt.Println(strings.HasSuffix("test@gmail.com", "gmail.com"))
	fmt.Println(strings.HasPrefix("test@gmail.com", "test"))

	fmt.Println(strings.Replace("test@gmail.com", "test", "john", 1))

	parts := strings.Split("test@gmail.com", "@")
	username, domain := parts[0], parts[1]
	fmt.Println(username, domain)

	parts2 := strings.Fields("jane example.com")
	username2, domain2 := parts2[0], parts2[1]
	fmt.Println(username2, domain2)
}


// ============================================================
// STRING AND STRINGS PACKAGE — NOTES
// ============================================================
//
// Go provides the `strings` package for working with strings.
//
// Import:
//
//     import "strings"
//
// ------------------------------------------------------------
//
// 1. strings.Clone()
//
//     s2 := strings.Clone(s1)
//
// Creates a copy of a string.
//
// Example:
//
//     s1 := "abc"
//     s2 := strings.Clone(s1)
//
// `s2` contains the same text as `s1`.
//
// Useful when you explicitly want an independent copy of the
// string's underlying data.
//
// ------------------------------------------------------------
//
// 2. strings.Builder
//
//     var b strings.Builder
//
// `strings.Builder` is used to efficiently construct strings,
// especially when repeatedly adding pieces of text.
//
// Example:
//
//     b.WriteString("Hello")
//     b.WriteString(" World")
//
//     result := b.String()
//
// Output:
//
//     Hello World
//
// You can also write bytes:
//
//     b.Write([]byte("Hello"))
//
// ------------------------------------------------------------
//
// 3. strings.ToLower()
//
//     strings.ToLower("ABC")
//
// Converts all letters to lowercase.
//
//     "ABC" → "abc"
//
// ------------------------------------------------------------
//
// 4. strings.ToUpper()
//
//     strings.ToUpper("abc")
//
// Converts all letters to uppercase.
//
//     "abc" → "ABC"
//
// ------------------------------------------------------------
//
// 5. strings.TrimSpace()
//
//     strings.TrimSpace(s)
//
// Removes leading and trailing whitespace.
//
// Example:
//
//     "     hello     "
//
// becomes:
//
//     "hello"
//
// It does NOT remove spaces from the middle of the string.
//
// ------------------------------------------------------------
//
// 6. strings.HasSuffix()
//
//     strings.HasSuffix(s, suffix)
//
// Checks whether a string ends with a particular suffix.
//
// Example:
//
//     strings.HasSuffix("test@gmail.com", "gmail.com")
//
//     → true
//
// ------------------------------------------------------------
//
// 7. strings.HasPrefix()
//
//     strings.HasPrefix(s, prefix)
//
// Checks whether a string starts with a particular prefix.
//
// Example:
//
//     strings.HasPrefix("test@gmail.com", "test")
//
//     → true
//
// ------------------------------------------------------------
//
// 8. strings.Replace()
//
//     strings.Replace(s, old, new, n)
//
// Replaces occurrences of one substring with another.
//
// Example:
//
//     strings.Replace(
//         "test@gmail.com",
//         "test",
//         "john",
//         1,
//     )
//
// Result:
//
//     "john@gmail.com"
//
// The final argument controls how many replacements happen.
//
//     1  → replace first occurrence
//     -1 → replace all occurrences
//
// ------------------------------------------------------------
//
// 9. strings.Split()
//
//     strings.Split(s, separator)
//
// Splits a string into a slice of strings.
//
// Example:
//
//     parts := strings.Split("test@gmail.com", "@")
//
// Result:
//
//     ["test", "gmail.com"]
//
// Therefore:
//
//     parts[0] → "test"
//     parts[1] → "gmail.com"
//
// Useful for parsing structured strings.
//
// ------------------------------------------------------------
//
// 10. strings.Fields()
//
//     strings.Fields(s)
//
// Splits a string based on whitespace.
//
// Example:
//
//     strings.Fields("jane example.com")
//
// Result:
//
//     ["jane", "example.com"]
//
// Unlike Split(), Fields automatically handles whitespace and
// ignores extra spaces.
//
// Example:
//
//     strings.Fields("   hello    world   ")
//
// Result:
//
//     ["hello", "world"]
//
// ------------------------------------------------------------
//
// 11. String Length
//
//     len(s)
//
// Returns the number of BYTES in a string.
//
// Example:
//
//     s := "test"
//
//     len(s) → 4
//
// IMPORTANT:
//
// `len()` counts bytes, NOT necessarily characters.
//
// This matters when working with Unicode characters.
//
// ------------------------------------------------------------
//
// QUICK REFERENCE
//
//     strings.Clone(s)
//     → clone a string.
//
//     strings.Builder
//     → efficiently build strings.
//
//     strings.ToLower(s)
//     → convert to lowercase.
//
//     strings.ToUpper(s)
//     → convert to uppercase.
//
//     strings.TrimSpace(s)
//     → remove leading/trailing whitespace.
//
//     strings.HasSuffix(s, suffix)
//     → check ending.
//
//     strings.HasPrefix(s, prefix)
//     → check beginning.
//
//     strings.Replace(s, old, new, n)
//     → replace substrings.
//
//     strings.Split(s, separator)
//     → split into a slice.
//
//     strings.Fields(s)
//     → split using whitespace.
//
//     len(s)
//     → number of bytes in the string.
//
// ============================================================
// KEY IDEA:
//
// The `strings` package provides ready-made tools for common
// string operations, so we don't need to manually implement
// searching, splitting, trimming, replacing, or case conversion.
//
// ============================================================