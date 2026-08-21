# 🧮 Go Math Library

A small Go project built to practice **error handling, custom errors, multiple return values, variadic functions, and `defer`**.

The project implements integer division with a custom `MathError` type and a variadic `Sum` function.

---

## 🎯 Learning Goals

- Go's `error` interface
- Multiple return values
- Explicit error handling with `if err != nil`
- Creating errors with `errors.New`
- Custom error types
- Implementing the `Error() string` method
- Variadic functions
- `defer`
- Structs and constants
- `fmt.Sprintf`
- `strings.Join`

---

## 🚀 Features

- ➕ Sum any number of integers
- ➗ Safely divide two integers
- ❌ Detect division by zero
- 🧱 Represent errors using a custom `MathError` struct
- 📝 Generate descriptive error messages
- ⏳ Demonstrate `defer` with the `Sum` function

---

## 📁 Project Structure

```text
math-library/
├── main.go
└── README.md
```

The current implementation intentionally lives in one `main.go` file because this is a learning project.

---

## 🛠️ Requirements

Check that Go is installed:

```bash
go version
```

Run the project:

```bash
go run .
```

Or:

```bash
go run main.go
```

---

# 🧱 Core Concepts

## 1. `MathError` Struct

The project defines a custom error structure:

```go
type MathError struct {
    Operation string
    InputA    int
    InputB    int
    Message   string
}
```

A struct groups related data together.

| Field | Purpose |
|---|---|
| `Operation` | Mathematical operation that failed |
| `InputA` | First input |
| `InputB` | Second input |
| `Message` | Description of the error |

Example:

```go
MathError{
    Operation: "Division",
    InputA:    10,
    InputB:    0,
    Message:   "division by zero is not allowed",
}
```

---

# ⚠️ Go Error Handling

Go does **not** use traditional exception handling like Java, Python, or JavaScript.

Instead, functions commonly return an error as one of their return values.

```go
func safeDivision(a, b int) (int, error)
```

The function returns:

```text
result + error
```

Successful operation:

```text
value = 5
err   = nil
```

Failed operation:

```text
value = 0
err   = some error
```

The standard pattern is:

```go
if err != nil {
    // handle error
}
```

---

# 🔢 Multiple Return Values

Go allows a function to return multiple values.

```go
func safeDivision(a, b int) (int, error)
```

The caller receives both:

```go
value, err := safeDivision(10, 2)
```

This pattern is extremely common in Go.

---

# ❌ Division by Zero

`safeDivision` checks whether the divisor is zero:

```go
if b == 0 {
    return 0, &MathError{
        Operation: division,
        InputA:    a,
        InputB:    b,
        Message:   divisionErrMsg,
    }
}
```

Instead of continuing with an invalid operation, the function returns:

```text
0 + error
```

The caller decides how to handle that error.

---

# 🔌 The `error` Interface

Go's `error` is an interface.

A type satisfies the `error` interface by implementing:

```go
Error() string
```

Our custom type does this:

```go
func (e MathError) Error() string {
    // ...
}
```

Therefore `MathError` can be returned where an `error` is expected.

Conceptually:

```text
MathError
    │
    │ implements Error() string
    ↓
error interface
```

That is why this works:

```go
return 0, &MathError{...}
```

---

# 📝 Custom `Error()` Method

The project defines:

```go
func (e MathError) Error() string {
    var inputs []string

    if e.Operation == "Division" {
        inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
        inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
    }

    return fmt.Sprintf(
        "Math error in %s (%s): %s",
        e.Operation,
        strings.Join(inputs, ","),
        e.Message,
    )
}
```

This converts the structured error into a readable string.

For example:

```text
Math error in Division (a=10,b=0): division by zero is not allowed
```

When:

```go
fmt.Println(err)
```

is used, Go can call the error's `Error()` method to obtain its string representation.

---

# 🧩 `fmt.Sprintf`

`fmt.Sprintf` formats a string and **returns the formatted string**.

```go
fmt.Sprintf("a=%d", e.InputA)
```

If `e.InputA` is `10`, the result is:

```text
a=10
```

Difference:

```text
fmt.Printf
    ↓
formats + prints

fmt.Sprintf
    ↓
formats + returns string
```

---

# 🧵 Building the Error Message

The code starts with an empty slice:

```go
var inputs []string
```

Then:

```go
inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
```

The slice becomes conceptually:

```text
["a=10", "b=0"]
```

Then:

```go
strings.Join(inputs, ",")
```

produces:

```text
a=10,b=0
```

That value is inserted into the final error message.

---

# 🧱 Constants

The project uses constants:

```go
const (
    division       = "Division"
    divisionErrMsg = "division by zero is not allowed"
)
```

Constants are useful for values that should not be reassigned during program execution.

Instead of repeatedly writing:

```go
"Division"
```

we can use:

```go
division
```

---

# ➕ Variadic `Sum` Function

The project contains:

```go
func Sum(numbers ...int) int
```

The `...int` syntax makes the function **variadic**.

It can accept any number of integers:

```go
Sum()
Sum(1)
Sum(1, 2)
Sum(1, 2, 3, 4, 5)
```

Inside the function, the variadic parameter behaves like a slice:

```go
for _, n := range numbers {
    total += n
}
```

Conceptually:

```text
Sum(1, 2, 3, 4)

        ↓

numbers = []int{1, 2, 3, 4}
```

---

# ⏳ `defer`

The `Sum` function contains:

```go
defer fmt.Println("Sum finished")
```

`defer` schedules a function call to execute when the surrounding function is about to return.

The lifecycle is:

```text
Enter Sum()
      ↓
Register deferred call
      ↓
Execute normal code
      ↓
Calculate total
      ↓
Function is about to return
      ↓
Deferred call executes
      ↓
Sum() returns
```

The precise idea is:

> A deferred function executes during the function's return process, before control leaves that function.

It is **not** executed immediately when the `defer` statement is encountered.

---

# 🔄 `defer` and LIFO

Multiple deferred calls execute in **Last In, First Out** order.

```go
defer fmt.Println("First")
defer fmt.Println("Second")
defer fmt.Println("Third")
```

Output:

```text
Third
Second
First
```

Think of it as a stack:

```text
┌─────────┐
│ Third   │ ← executes first
├─────────┤
│ Second  │
├─────────┤
│ First   │ ← executes last
└─────────┘
```

---

# 🧹 Why `defer` Is Useful

`defer` is commonly used for cleanup:

```text
Acquire resource
      ↓
defer cleanup()
      ↓
Do work
      ↓
Function returns
      ↓
cleanup() executes
```

Common examples:

```go
defer file.Close()
```

```go
defer mutex.Unlock()
```

```go
defer response.Body.Close()
```

This makes cleanup easier to guarantee across different return paths.

---

# 🔍 `safeDivision` Flow

The function:

```go
func safeDivision(a, b int) (int, error)
```

works like this:

```text
safeDivision(a, b)
        ↓
Is b == 0?
   ┌────┴────┐
  YES        NO
   ↓          ↓
Return      Calculate
MathError   a / b
   ↓          ↓
0 + error   result + nil
```

For:

```go
safeDivision(10, 2)
```

we get:

```text
5
nil
```

For:

```go
safeDivision(10, 0)
```

we get:

```text
0
MathError
```

---

# 🔍 Error Checking in `main`

The caller uses:

```go
value, err := safeDivision(10, 2)

if err != nil {
    fmt.Println(err)
}
```

The standard Go pattern is:

```text
Call function
     ↓
Receive value + error
     ↓
Is error nil?
   ↙       ↘
 YES       NO
  ↓         ↓
Use       Handle
value     error
```

`nil` means there is no error.

---

# 🧠 Important Concepts Learned

| Concept | Meaning |
|---|---|
| `struct` | Groups related data |
| `error` | Standard Go interface for errors |
| `Error() string` | Method used to satisfy the `error` interface |
| `errors.New()` | Creates a basic error |
| Multiple returns | Function can return several values |
| `err != nil` | Standard error check |
| `...int` | Variadic function parameter |
| `defer` | Schedules a function call for the return process |
| `const` | Declares a constant |
| `fmt.Sprintf` | Formats and returns a string |
| `strings.Join` | Combines slice elements into one string |
| `&MathError{}` | Creates a pointer to a `MathError` |

---

# 🔥 Complete Execution Flow

When the program starts:

```text
main()
  │
  ├── Sum(1, 2, 3)
  │      │
  │      ├── register defer
  │      ├── total = 0
  │      ├── total = 1
  │      ├── total = 3
  │      ├── total = 6
  │      ├── deferred print executes
  │      └── return 6
  │
  └── safeDivision(10, 2)
         │
         ├── b == 0?
         │      ↓
         │     NO
         │
         ├── calculate 10 / 2
         │
         └── return 5, nil
```

---

# 📌 Expected Output

For the current `main()`:

```text
Sum finished
6
5
```

`Sum finished` appears when `Sum()` is completing because of the deferred statement.

---

# 💡 What This Project Taught Me

This small project introduces patterns that appear frequently in real Go applications:

```text
Functions
    ↓
Multiple return values
    ↓
Explicit error handling
    ↓
Custom error types
    ↓
Interfaces
    ↓
Deferred cleanup
    ↓
More reliable backend code
```

The key mindset is:

> Go generally makes errors explicit instead of hiding them behind traditional exceptions.

---

# 🚧 Possible Future Improvements

The project can be extended with:

- `Add()`
- `Subtract()`
- `Multiply()`
- Floating-point operations
- More mathematical validation
- More custom error types
- Error wrapping with `fmt.Errorf`
- `errors.Is`
- `errors.As`
- Unit tests
- Benchmarks
- Separate Go packages
- CLI interface
- Input validation
- Table-driven tests

A possible future structure:

```text
math-library/
├── cmd/
│   └── math/
│       └── main.go
├── math/
│   ├── operations.go
│   └── errors.go
├── tests/
└── README.md
```

---

## 🧭 Go Learning Progression

This project fits into a broader Go learning path:

```text
Go Syntax
    ↓
Variables & Types
    ↓
Arrays & Slices
    ↓
Maps
    ↓
Structs
    ↓
Pointers
    ↓
Functions
    ↓
Multiple Return Values
    ↓
Error Handling
    ↓
Custom Errors
    ↓
Defer
    ↓
Packages
    ↓
Testing
    ↓
Concurrency
    ↓
HTTP / REST APIs
    ↓
Backend Engineering
```

---

## 📜 License

This project is created for learning and experimentation.
