package main 

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	text1 := "Hello World! Welcome to Go"

	regGo, err := regexp.Compile(`Go`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Text '%s, matches 'Go': %t\n", text1, regGo.MatchString(text1))

	text2 := "Products codes: p123, X342, P789"
	rProductP := regexp.MustCompile(`P\d+`)
	firstProduct := rProductP.FindString(text2)
	fmt.Println(firstProduct)

}


// ============================================================
// REGULAR EXPRESSIONS (REGEX) IN GO — NOTES
// ============================================================
//
// Go provides the `regexp` package for matching patterns inside
// strings.
//
// Import:
//
//     import "regexp"
//
// Regex is useful when we need to search, validate, extract, or
// replace text based on a PATTERN rather than an exact string.
//
// ------------------------------------------------------------
//
// 1. COMPILE A REGEX
//
//     reg, err := regexp.Compile(`Go`)
//
// `regexp.Compile()` takes a regex pattern and compiles it into
// a *regexp.Regexp object.
//
// It returns TWO values:
//
//     *regexp.Regexp
//     error
//
// This is because the pattern might be invalid.
//
// Example:
//
//     regGo, err := regexp.Compile(`Go`)
//
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
// Once compiled, `regGo` can be used for matching operations.
//
// ------------------------------------------------------------
//
// 2. RAW STRING LITERALS
//
// Regex patterns are commonly written using BACKTICKS:
//
//     `Go`
//
// This is a raw string literal in Go.
//
// Raw strings are convenient for regex because backslashes do not
// need to be escaped by Go's string parser.
//
// Example:
//
//     `\d+`
//
// is easier to write than:
//
//     "\\d+"
//
// This becomes especially useful with complex regex patterns.
//
// ------------------------------------------------------------
//
// 3. MatchString()
//
//     reg.MatchString(text)
//
// Checks whether the regex pattern matches ANYWHERE inside the
// supplied string.
//
// Example:
//
//     pattern := `Go`
//     text := "Hello World! Welcome to Go"
//
//     reg.MatchString(text)
//
//     → true
//
// It returns:
//
//     true  → pattern was found
//     false → pattern was not found
//
// IMPORTANT:
// MatchString() checks for a match somewhere in the string.
// It does NOT automatically mean the entire string matches.
//
// ------------------------------------------------------------
//
// 4. MUST COMPILE
//
//     regexp.MustCompile(`P\d+`)
//
// `MustCompile()` is a shortcut when you KNOW the regex pattern
// is valid.
//
// It returns:
//
//     *regexp.Regexp
//
// Unlike Compile(), it does NOT return an error.
//
// If the pattern is invalid, MustCompile() PANICS.
//
// Therefore:
//
//     regexp.Compile()
//     → returns error
//
//     regexp.MustCompile()
//     → panics if pattern is invalid
//
// Use MustCompile when the regex is a fixed pattern known at
// development time.
//
// ------------------------------------------------------------
//
// 5. BASIC REGEX SYMBOLS
//
// Some important regex metacharacters:
//
//     .       → any character
//
//     \d      → digit
//
//     \w      → word character
//
//     \s      → whitespace
//
//     +       → one or more occurrences
//
//     *       → zero or more occurrences
//
//     ?       → zero or one occurrence
//
//     ^       → beginning of string
//
//     $       → end of string
//
//     [abc]   → one character from a, b, or c
//
//     [0-9]   → one digit from 0 to 9
//
//     [A-Z]   → uppercase letter
//
//     [a-z]   → lowercase letter
//
//     |       → OR
//
//     (...)   → capturing/grouping expression
//
// ------------------------------------------------------------
//
// 6. UNDERSTANDING `P\d+`
//
// Your example:
//
//     rProductP := regexp.MustCompile(`P\d+`)
//
// Break it down:
//
//     P
//     → must contain the letter P
//
//     \d
//     → one digit
//
//     +
//     → one or more digits
//
// Therefore:
//
//     P\d+
//
// matches:
//
//     P123
//     P5
//     P789
//
// but does NOT match:
//
//     p123       // lowercase p
//     P           // no digit
//     PX123      // X is not a digit
//
// ------------------------------------------------------------
//
// 7. FindString()
//
//     r.FindString(text)
//
// Finds the FIRST substring that matches the regex.
//
// Example:
//
//     text := "Products codes: p123, X342, P789"
//
//     pattern := `P\d+`
//
//     r.FindString(text)
//
// returns:
//
//     P789
//
// Notice:
//
//     p123
//
// does not match because regex matching is case-sensitive by
// default.
//
// ------------------------------------------------------------
//
// 8. FINDING ALL MATCHES
//
//     FindAllString()
//
// can be used when you want ALL matching substrings.
//
// Concept:
//
//     r.FindAllString(text, -1)
//
// The `-1` means:
//
//     return all matches
//
// Whereas a positive number can limit the number of matches.
//
// ------------------------------------------------------------
//
// 9. CASE SENSITIVITY
//
// Regex matching is case-sensitive by default.
//
// Example:
//
//     `P\d+`
//
// matches:
//
//     P123
//
// but not:
//
//     p123
//
// You can use the `(?i)` flag for case-insensitive matching:
//
//     `(?i)p\d+`
//
// Now both:
//
//     P123
//     p123
//
// can match.
//
// ------------------------------------------------------------
//
// 10. REGEX IS PATTERN MATCHING
//
// Normal string search:
//
//     strings.Contains(text, "Go")
//
// asks:
//
//     "Does this exact text exist?"
//
// Regex:
//
//     regexp.MatchString()
//
// asks:
//
//     "Does this text follow the pattern?"
//
// Example:
//
//     `P\d+`
//
// can match:
//
//     P123
//     P456
//     P999
//
// without knowing the exact number beforehand.
//
// ------------------------------------------------------------
//
// 11. COMMON REGEX USE CASES
//
// Regex is commonly used for:
//
//     • Email pattern checking
//     • Phone number validation
//     • Extracting product codes
//     • Log parsing
//     • Finding numbers in text
//     • Extracting URLs
//     • Input validation
//     • Searching structured text
//     • Parsing simple text formats
//
// ------------------------------------------------------------
//
// 12. COMPILE ONCE, USE MANY TIMES
//
// If the same regex is used repeatedly, compile it once:
//
//     re := regexp.MustCompile(`P\d+`)
//
// Then:
//
//     re.MatchString(text1)
//     re.MatchString(text2)
//     re.FindString(text3)
//
// This avoids repeatedly compiling the same pattern.
//
// ------------------------------------------------------------
//
// 13. IMPORTANT: REGEX IS NOT ALWAYS THE BEST TOOL
//
// Don't automatically use regex for every string problem.
//
// For simple operations, the `strings` package is usually clearer:
//
//     strings.Contains()
//     strings.HasPrefix()
//     strings.HasSuffix()
//     strings.Split()
//     strings.Replace()
//
// Use regex when the matching logic is actually PATTERN-BASED.
//
// Example:
//
//     strings.HasPrefix(email, "admin")
//
// is simpler than creating a regex just to check a prefix.
//
// ------------------------------------------------------------
//
// QUICK REFERENCE
//
//     regexp.Compile(pattern)
//     → compile regex and return error if invalid.
//
//     regexp.MustCompile(pattern)
//     → compile regex and panic if invalid.
//
//     re.MatchString(text)
//     → check whether a match exists.
//
//     re.FindString(text)
//     → return first matching substring.
//
//     re.FindAllString(text, -1)
//     → return all matching substrings.
//
//     \d
//     → digit.
//
//     +
//     → one or more.
//
//     *
//     → zero or more.
//
//     ?
//     → zero or one.
//
//     ^
//     → beginning.
//
//     $
//     → end.
//
//     [abc]
//     → one of the listed characters.
//
//     (?i)
//     → case-insensitive matching.
//
// ------------------------------------------------------------
//
// BIG PICTURE:
//
//                 REGEX PATTERN
//                       │
//                       ↓
//              regexp.Compile()
//                       │
//                       ↓
//                *regexp.Regexp
//                       │
//             ┌─────────┴─────────┐
//             ↓                   ↓
//       MatchString()       FindString()
//             │                   │
//             ↓                   ↓
//        true / false       matched text
//
// ============================================================