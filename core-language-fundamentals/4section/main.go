package main

import "fmt"

const(
	Sunday=iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	//Go doesnt have enum keyword as java or c# but it use iota identifier.
	
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}

// ENUM-LIKE CONSTANTS WITH `iota`
//
// Go does not have an `enum` keyword like Java or C#.
// Instead, Go commonly uses `const` + `iota` to create
// enum-like sets of constants.
//
// 1. `iota`
//
// `iota` is a special identifier used inside a `const` block.
// It starts at 0 and automatically increases by 1 for each
// constant declaration.
//
// const (
//     Sunday = iota  // 0
//     Monday          // 1
//     Tuesday         // 2
//     Wednesday       // 3
//     Thursday        // 4
//     Friday          // 5
//     Saturday        // 6
// )
//
// Go automatically repeats the previous expression (`iota`)
// for the following lines.
//
// 2. WHY USE `iota`?
//
// Without iota:
//
// const (
//     Sunday    = 0
//     Monday    = 1
//     Tuesday   = 2
//     Wednesday = 3
// )
//
// With iota:
//
// const (
//     Sunday = iota
//     Monday
//     Tuesday
//     Wednesday
// )
//
// This avoids manually maintaining sequential numbers and
// reduces the chance of numbering mistakes.
//
// 3. IMPORTANT
//
// `iota` resets to 0 at the beginning of every `const` block.
//
// const (
//     A = iota // 0
//     B        // 1
// )
//
// const (
//     C = iota // 0 again
//     D        // 1
// )
//
// 4. `iota` IS NOT AN ENUM
//
// `iota` simply generates successive integer constants.
// The combination of `const` + `iota` is commonly used to
// create enum-like values in Go.
//
// 5. COMMON USE CASES
//
// `iota` is useful for:
// - Days of the week
// - Months
// - Status values
// - Permissions / flags
// - Categories represented by numbers
//
// Example:
//
// const (
//     Pending = iota
//     Approved
//     Rejected
// )
//
// Pending  → 0
// Approved → 1
// Rejected → 2