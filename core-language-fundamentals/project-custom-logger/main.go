package main

import "fmt" 

type LogLevel int // what's type and what's alias?

const(
	LevelTrace  LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

var levelNames = []string{"Trace", "Debug", "Info", "Warning", "Error"}

func (l LogLevel) String() string {
	
	if l < LevelTrace || l > LevelError {
		return "Unknown"
	}

	return levelNames[l]

}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log Level: %d %s\n", level, level.String())
}

func main() {

	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)
	printLogLevel(20)

}


// --------------------------------------------------------- NOTES TO REVISE ------------------------------------------------------------------------

// TYPES, TYPE ALIASES & FUNCTIONS IN GO
//
// 1. CREATING A NEW TYPE
//
// type LogLevel int
//
// This creates a NEW type called `LogLevel`.
// Its underlying type is `int`.
//
// LogLevel and int are related, but they are NOT the same type.
//
// Why create a new type?
//
// It gives meaning to the values.
//
// `int` could represent anything:
//     10, 20, 30...
//
// `LogLevel` tells us:
//     "This integer represents a logging level."
//
// This improves type safety and makes code easier to understand.
//
// ------------------------------------------------------------
//
// 2. TYPE VS TYPE ALIAS
//
// NEW TYPE:
//
// type LogLevel int
//
// Creates a completely new type.
//
// TYPE ALIAS:
//
// type LogLevel = int
//
// This does NOT create a new type.
// `LogLevel` is simply another name for `int`.
//
// Think:
//
// type LogLevel int
// → "Create a new type based on int."
//
// type LogLevel = int
// → "Give int another name."
//
// ------------------------------------------------------------
//
// 3. ENUM-LIKE VALUES WITH `iota`
//
// Go doesn't have an `enum` keyword like C++/Java/C#.
//
// Instead, `const` + `iota` is commonly used:
//
// const (
//     LevelTrace LogLevel = iota // 0
//     LevelDebug                  // 1
//     LevelInfo                   // 2
//     LevelWarning                // 3
//     LevelError                  // 4
// )
//
// `iota` automatically generates sequential integer constants.
//
// Because `LogLevel` is specified on the first constant,
// the following constants use the same type.
//
// ------------------------------------------------------------
//
// 4. FUNCTION DECLARATION IN GO
//
// func printLogLevel(level LogLevel) {
//     ...
// }
//
// General syntax:
//
// func functionName(parameter parameterType) returnType {
//     // body
// }
//
// Example:
//
// func add(a int, b int) int {
//     return a + b
// }
//
// Go can also combine parameters of the same type:
//
// func add(a, b int) int {
//     return a + b
// }
//
// IMPORTANT:
// Go puts the TYPE AFTER the variable name.
//
// C++:
//
// int add(int a, int b) {
//     return a + b;
// }
//
// Go:
//
// func add(a int, b int) int {
//     return a + b
// }
//
// ------------------------------------------------------------
//
// 5. METHODS IN GO
//
// func (l LogLevel) String() string
//
// This is a METHOD, not just a normal function.
//
// `(l LogLevel)` is called the RECEIVER.
//
// It means:
//
// "String() belongs to the LogLevel type."
//
// In C++, this would be similar to:
//
// class LogLevel {
// public:
//     string String() {
//         ...
//     }
// };
//
// In Go, we don't put the method inside the type declaration.
// We attach it using a receiver:
//
// func (l LogLevel) String() string
//
// ------------------------------------------------------------
//
// 6. `l` IS THE RECEIVER
//
// func (l LogLevel) String() string
//
// `l` represents the current LogLevel value.
//
// For example:
//
// level := LevelError
// level.String()
//
// Inside String(), `l` is LevelError.
//
// ------------------------------------------------------------
//
// 7. RANGE CHECK
//
// if l < LevelTrace || l > LevelError {
//     return "Unknown"
// }
//
// Valid values:
//
// LevelTrace  → 0
// LevelDebug  → 1
// LevelInfo   → 2
// LevelWarning → 3
// LevelError  → 4
//
// Anything outside this range is considered unknown.
//
// Example:
//
// printLogLevel(20)
//
// 20 is outside 0–4,
// so String() returns "Unknown".
//
// ------------------------------------------------------------
//
// 8. `fmt.Print` vs `fmt.Println` vs `fmt.Printf`
//
// fmt.Print("Hello")
//
// Prints exactly what you give it.
// DOES NOT automatically add a newline.
//
// Output:
// Hello
//
// ------------------------------------------------------------
//
// fmt.Println("Hello")
//
// Prints the values and automatically adds a newline.
//
// Output:
// Hello
//
// ------------------------------------------------------------
//
// fmt.Printf("Age: %d\n", age)
//
// `Printf` means PRINT FORMATTED.
//
// It understands format verbs such as:
//
// %d → integer
// %f → floating-point number
// %s → string
// %v → general/default value
// %t → boolean
// %T → type of a value
// \n → newline
//
// Example:
//
// fmt.Printf("Name: %s Age: %d\n", name, age)
//
// ------------------------------------------------------------
//
// IMPORTANT:
//
// Println:
//
// fmt.Println("Age: %d", 20)
//
// ❌ Does NOT replace %d with 20.
//
// It will print the arguments normally.
//
// Use Printf when you need format specifiers:
//
// fmt.Printf("Age: %d\n", 20)
//
// ------------------------------------------------------------
//
// 9. WHY USE A CUSTOM TYPE HERE?
//
// Without LogLevel:
//
// func printLogLevel(level int)
//
// This accepts ANY integer:
//
// printLogLevel(10)
// printLogLevel(999)
// printLogLevel(-50)
//
// The compiler can't tell that these numbers are supposed
// to represent log levels.
//
// With:
//
// type LogLevel int
//
// func printLogLevel(level LogLevel)
//
// We clearly communicate:
//
// "This function expects a LogLevel."
//
// This gives our code more meaning and better type safety.
//
// ------------------------------------------------------------
//
// 10. WHY USE A TYPE ALIAS?
//
// Example:
//
// type ID = int
//
// An alias is useful when you want another name for an
// existing type without creating a new type.
//
// Example:
//
// type UserID = int
//
// UserID is still exactly the same type as int.
//
// Unlike:
//
// type UserID int
//
// which creates a NEW type.
//
// In simple terms:
//
// `type X int`
// → NEW type
//
// `type X = int`
// → ALIAS / another name
//
// ------------------------------------------------------------
//
// KEY TAKEAWAYS:
//
// `type LogLevel int`
// → creates a new type.
//
// `type LogLevel = int`
// → creates an alias.
//
// `iota`
// → generates sequential constant values.
//
// `func`
// → declares a function.
//
// `(l LogLevel)`
// → receiver; makes String() a method of LogLevel.
//
// `fmt.Print()`
// → print without automatic newline.
//
// `fmt.Println()`
// → print with automatic newline.
//
// `fmt.Printf()`
// → formatted printing using %d, %s, %v, etc.
//
// Custom types
// → give meaning to values and provide stronger type safety.