package main

import "fmt"

type ConfigItem struct {
	Key   string
	Value interface{} // a.k.a `any`
	isSet bool
}

/*
%v
%+v
%#v
%T
%s
%d
%f (%.2f)
%t
%q
%%
*/

func (c ConfigItem) String() string {
	return fmt.Sprintf("Key: %s, Value: %s, isSet: %t", c.Key, c.Value, c.isSet)
}

func main() {
	appName := "EnvParser"
	version := 1.2
	port := 8000
	isEnabled := true

	status := fmt.Sprintf("Application: %s (Version: %.1f) running on port %d. Enabled: %t", appName, version, port, isEnabled)
	fmt.Println(status)

	item1 := ConfigItem{Key: "API_URL", Value: "http://localhost:3000/api", isSet: true}
	item2 := ConfigItem{Key: "TIMEOUT_MS", Value: 5000, isSet: true}
	item3 := ConfigItem{Key: "DEBUG_MODE", Value: false, isSet: false}

	fmt.Printf("Item 1 (%%V): %v\n", item1)
	fmt.Printf("Item 2 (%%+v): %+v\n", item2)
	fmt.Printf("Item 3 (%%#V): %#v\n", item3)

	err := errors.New("test")

	fmt.Errorf("here is the error on the port %d: %w", port, err)
}


// ============================================================
// fmt FORMATTING + interface{} / any — NOTES
// ============================================================
//
// This example demonstrates two useful Go concepts:
//
//     1. fmt formatting verbs
//     2. interface{} (also written as any)
//
// ------------------------------------------------------------
//
// 1. interface{} / any
//
//     Value interface{}
//
// `interface{}` is the empty interface.
//
// It can hold a value of ANY type:
//
//     string
//     int
//     float64
//     bool
//     struct
//     etc.
//
// In modern Go, `any` is simply an alias for `interface{}`:
//
//     interface{} == any
//
// So this:
//
//     Value interface{}
//
// can also be written as:
//
//     Value any
//
// Example:
//
//     Value: "hello"
//     Value: 5000
//     Value: false
//
// All of these are valid because `Value` can store any type.
//
// ------------------------------------------------------------
//
// 2. fmt.Sprintf()
//
//     fmt.Sprintf(format, values...)
//
// Formats values according to the supplied format string and
// RETURNS the resulting string.
//
// Example:
//
//     fmt.Sprintf("Port: %d", 8000)
//
// Result:
//
//     "Port: 8000"
//
// Unlike fmt.Printf(), Sprintf does not print directly.
//
//     Printf  → formats + prints
//     Sprintf → formats + returns
//
// ------------------------------------------------------------
//
// 3. %v — Default Value
//
//     %v
//
// Prints the value in its default format.
//
// Example:
//
//     fmt.Printf("%v", 100)
//
// Output:
//
//     100
//
// It is the general-purpose formatting verb.
//
// ------------------------------------------------------------
//
// 4. %+v — Value with Field Names
//
//     %+v
//
// Mainly useful with structs.
//
// Example:
//
//     fmt.Printf("%+v", item)
//
// It includes the struct's field names.
//
// This is especially useful while debugging structs.
//
// ------------------------------------------------------------
//
// 5. %#v — Go-Syntax Representation
//
//     %#v
//
// Prints a Go-syntax representation of the value.
//
// It tries to show the value in a form that resembles valid Go
// code.
//
// Very useful for debugging because it gives more structural
// information about the value.
//
// ------------------------------------------------------------
//
// 6. %T — Type
//
//     %T
//
// Prints the type of the value.
//
// Example:
//
//     fmt.Printf("%T", 5000)
//
// Output:
//
//     int
//
// Example:
//
//     fmt.Printf("%T", 1.2)
//
// Output:
//
//     float64
//
// Very useful when working with interface{} / any.
//
// ------------------------------------------------------------
//
// 7. %s — String
//
//     %s
//
// Used for strings.
//
// Example:
//
//     fmt.Printf("%s", "hello")
//
// Output:
//
//     hello
//
// `%s` expects a string-like value.
//
// ------------------------------------------------------------
//
// 8. %d — Integer
//
//     %d
//
// Formats an integer using base 10.
//
// Example:
//
//     fmt.Printf("%d", 8000)
//
// Output:
//
//     8000
//
// Commonly used with:
//
//     int
//     int8
//     int16
//     int32
//     int64
//
// ------------------------------------------------------------
//
// 9. %f — Floating Point
//
//     %f
//
// Used for floating-point values.
//
// Example:
//
//     fmt.Printf("%f", 12.3456)
//
// Output:
//
//     12.345600
//
// You can control decimal precision:
//
//     %.2f
//
// Example:
//
//     fmt.Printf("%.2f", 12.3456)
//
// Output:
//
//     12.35
//
// ------------------------------------------------------------
//
// 10. %t — Boolean
//
//     %t
//
// Formats a boolean value.
//
// Example:
//
//     fmt.Printf("%t", true)
//
// Output:
//
//     true
//
// ------------------------------------------------------------
//
// 11. %q — Quoted String
//
//     %q
//
// Prints a string with quotation marks.
//
// Example:
//
//     fmt.Printf("%q", "hello")
//
// Output:
//
//     "hello"
//
// This is useful when you want whitespace or special characters
// to be visually obvious.
//
// ------------------------------------------------------------
//
// 12. %% — Literal %
//
//     %%
//
// Prints an actual `%` character.
//
// Why is this needed?
//
// Because `%` normally begins a formatting verb.
//
// Example:
//
//     fmt.Printf("Progress: 100%%")
//
// Output:
//
//     Progress: 100%
//
// ------------------------------------------------------------
//
// 13. CUSTOM String() METHOD
//
// ConfigItem defines:
//
//     func (c ConfigItem) String() string
//
// This implements the `fmt.Stringer` interface.
//
// Stringer requires:
//
//     String() string
//
// Once a type implements String(), fmt functions can use that
// method when formatting the value as a string.
//
// Example:
//
//     fmt.Println(item1)
//
// can use ConfigItem.String() to determine how the value should
// be displayed.
//
// ------------------------------------------------------------
//
// 14. IMPORTANT DIFFERENCE: %v vs %+v vs %#v
//
// For a struct:
//
//     %v
//     → default representation
//
//     %+v
//     → includes field names
//
//     %#v
//     → Go-syntax representation
//
// These are particularly useful for debugging and inspecting
// structs.
//
// ------------------------------------------------------------
//
// 15. Formatting a Complete Message
//
// Example:
//
//     fmt.Sprintf(
//         "Application: %s (Version: %.1f) running on port %d. Enabled: %t",
//         appName,
//         version,
//         port,
//         isEnabled,
//     )
//
// Each formatting verb corresponds to the value:
//
//     %s  → appName
//     %.1f → version
//     %d  → port
//     %t  → isEnabled
//
// The number and order of formatting arguments should match the
// formatting verbs.
//
// ------------------------------------------------------------
//
// 16. DEBUGGING WITH %T
//
// This becomes especially useful with:
//
//     interface{}
//     any
//
// Example:
//
//     item := ConfigItem{
//         Key: "TIMEOUT_MS",
//         Value: 5000,
//     }
//
// Since Value can hold different types, `%T` can tell us what
// concrete type is currently stored.
//
// Concept:
//
//     Value
//       ↓
//     interface{} / any
//       ↓
//     actual value
//       ↓
//     %T → tells us its concrete type
//
// ------------------------------------------------------------
//
// QUICK REFERENCE
//
//     %v
//     → default value
//
//     %+v
//     → value with struct field names
//
//     %#v
//     → Go-syntax representation
//
//     %T
//     → type
//
//     %s
//     → string
//
//     %d
//     → integer
//
//     %f
//     → floating point
//
//     %.2f
//     → floating point with 2 decimal places
//
//     %t
//     → boolean
//
//     %q
//     → quoted string
//
//     %%
//     → literal `%`
//
// ------------------------------------------------------------
//
// IMPORTANT CORRECTION IN THIS EXAMPLE:
//
// If `Value` is declared as:
//
//     Value interface{}
//
// then this String() implementation:
//
//     fmt.Sprintf("Key: %s, Value: %s, isSet: %t",
//         c.Key, c.Value, c.isSet)
//
// is problematic because `%s` expects a string, while `c.Value`
// can contain ANY type.
//
// A safer general-purpose formatter is:
//
//     %v
//
// Example:
//
//     fmt.Sprintf("Key: %s, Value: %v, isSet: %t",
//         c.Key, c.Value, c.isSet)
//
// This works for strings, integers, booleans, floats, structs,
// etc.
//
// ------------------------------------------------------------
//
// ALSO NOTE:
//
// The example uses:
//
//     errors.New(...)
//
// so the `errors` package must be imported:
//
//     import "errors"
//
// If you call `errors.New()` without importing `errors`, Go will
// report:
//
//     undefined: errors
//
// ------------------------------------------------------------
//
// BIG PICTURE:
//
//     fmt
//      │
//      ├── Printf()
//      ├── Sprintf()
//      └── formatting verbs
//              │
//              ├── %v
//              ├── %+v
//              ├── %#v
//              ├── %T
//              ├── %s
//              ├── %d
//              ├── %f
//              ├── %t
//              ├── %q
//              └── %%
//
//     interface{} / any
//              │
//              ↓
//       Can hold any type
//              │
//              ↓
//          %T → inspect type
//          %v → inspect value
//
// ============================================================
