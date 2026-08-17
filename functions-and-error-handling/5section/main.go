package main 

import(
	"fmt"
	"errors"
	"time"
)

func divide(a, b int) (int, error) { // By convention error is always the last value returned.

	if b==0 {
		return 0, errors.New("Divide by Zero")
	}

	if a > 1000 {
		return 0, errors.New("a is too large")
	}

	return a / b, nil 
}

var ErrDivisionByZero = errors.New("division by zero")
var ErrNumberTooLarge = errors.New("number too large")

type OpError struct {
	Op      string
	Code    int 
	Message string
	Time    time.Time
}

func (op OpError) Error() string {
	return op.Message
}

func NewOpError(op string, code int, message string, t time.Time) *OpError {

	return &OpError{
		Op:      op, 
		Code:    code, 
		Message: message, 
		Time:    t,
	}

}

func doSomething() error {

	return NewOpError("doSomething", 100, "failed", time.Now())
}

func main() {

	value, err := divide(10, 5)
	if err != nil {

		if errors.Is(err, ErrDivisionByZero) {
			fmt.Println("divide by zero")
		} else if errors.Is(err, ErrNumberTooLarge) {
			fmt.Println("number too large")
		}else{
			fmt.Println(err)
		}
		
		return
	}

	fmt.Println(value)

}



// --------------------------------------- Advanced Error Handling ---------------------------------------
//
// This example introduces more advanced Go error-handling concepts:
//
//     - Sentinel errors
//     - errors.Is()
//     - Custom error types
//     - Implementing the error interface
//     - Error constructors
//     - Returning custom errors
//     - time.Time
//
// --------------------------------------------------------------------------------
//
// 1. BASIC ERROR RETURNING
//
// func divide(a, b int) (int, error)
//
// Go commonly returns:
//
//     result, error
//
// Example:
//
//     value, err := divide(10, 5)
//
// If everything succeeds:
//
//     value = 2
//     err   = nil
//
// If something goes wrong:
//
//     value = 0
//     err   = some error
//
// --------------------------------------------------------------------------------
//
// 2. CREATING AN ERROR WITH errors.New()
//
// errors.New("division by zero")
//
// creates an error value.
//
// Example:
//
//     return 0, errors.New("division by zero")
//
// This is useful when we simply want to report what went wrong.
//
// However, there is an important problem.
//
// Every time we call:
//
//     errors.New("division by zero")
//
// Go creates a NEW error value.
//
// Even if the message is identical, the error values are different.
//
// --------------------------------------------------------------------------------
//
// 3. SENTINEL ERRORS
//
// These are predefined error values that we can compare against.
//
//     var ErrDivisionByZero = errors.New("division by zero")
//     var ErrNumberTooLarge = errors.New("number too large")
//
// These are called SENTINEL ERRORS.
//
// Think of them as predefined error categories:
//
//     ErrDivisionByZero
//          ↓
//     "This operation failed because of division by zero."
//
//
//     ErrNumberTooLarge
//          ↓
//     "This operation failed because the number was too large."
//
// Instead of creating a new error every time, we return these predefined
// errors.
//
// --------------------------------------------------------------------------------
//
// 4. WHY SENTINEL ERRORS?
//
// Instead of:
//
//     return 0, errors.New("division by zero")
//
// we use:
//
//     return 0, ErrDivisionByZero
//
// Then the caller can identify the specific error:
//
//     if errors.Is(err, ErrDivisionByZero) {
//
//         fmt.Println("divide by zero")
//
//     }
//
// This is much better than checking strings manually.
//
// DON'T do:
//
//     if err.Error() == "division by zero" {
//
//         ...
//
//     }
//
// Prefer:
//
//     errors.Is(err, ErrDivisionByZero)
//
// --------------------------------------------------------------------------------
//
// 5. errors.Is()
//
// errors.Is() checks whether an error matches a particular error.
//
// Example:
//
//     errors.Is(err, ErrDivisionByZero)
//
// returns:
//
//     true
//
// if `err` represents ErrDivisionByZero.
//
// Otherwise:
//
//     false
//
// This becomes especially useful when errors are WRAPPED.
//
// --------------------------------------------------------------------------------
//
// 6. IMPORTANT DIFFERENCE
//
// These are NOT the same error value:
//
//     errors.New("division by zero")
//
//     errors.New("division by zero")
//
// Even though their messages are identical, they were created separately.
//
// Think:
//
//     Error A → errors.New(...)
//     Error B → errors.New(...)
//
//     A != B
//
// But:
//
//     var ErrDivisionByZero = errors.New("division by zero")
//
// gives us one reusable error value.
//
// We can return that same error whenever the condition occurs.
//
// --------------------------------------------------------------------------------
//
// 7. CORRECT divide() FUNCTION
//
//     func divide(a, b int) (int, error) {
//
//         if b == 0 {
//             return 0, ErrDivisionByZero
//         }
//
//         if a > 1000 {
//             return 0, ErrNumberTooLarge
//         }
//
//         return a / b, nil
//     }
//
// Possible outcomes:
//
//     divide(10, 5)
//         → 2, nil
//
//     divide(10, 0)
//         → 0, ErrDivisionByZero
//
//     divide(2000, 5)
//         → 0, ErrNumberTooLarge
//
// --------------------------------------------------------------------------------
//
// 8. ERROR CHECKING
//
// A common Go pattern:
//
//     value, err := divide(10, 5)
//
//     if err != nil {
//
//         // handle error
//
//     }
//
//     // use value
//
// `err != nil` means:
//
//     "Something went wrong."
//
// `err == nil` means:
//
//     "No error occurred."
//
// --------------------------------------------------------------------------------
//
// 9. MULTIPLE ERROR TYPES
//
// We can handle different errors differently:
//
//     if errors.Is(err, ErrDivisionByZero) {
//
//         fmt.Println("divide by zero")
//
//     } else if errors.Is(err, ErrNumberTooLarge) {
//
//         fmt.Println("number too large")
//
//     } else {
//
//         fmt.Println(err)
//
//     }
//
// This lets the caller decide what to do based on the type/category
// of error that occurred.
//
// --------------------------------------------------------------------------------
//
// 10. CUSTOM ERROR TYPES
//
// Go allows us to create our own error types.
//
// Example:
//
//     type OpError struct {
//
//         Op      string
//         Code    int
//         Message string
//         Time    time.Time
//
//     }
//
// This is a struct containing additional information about an error.
//
// Instead of only having:
//
//     "failed"
//
// we can store:
//
//     Operation → doSomething
//     Code      → 100
//     Message   → failed
//     Time      → when the error occurred
//
// --------------------------------------------------------------------------------
//
// 11. WHY CREATE A CUSTOM ERROR?
//
// A normal error:
//
//     errors.New("failed")
//
// only gives us a message.
//
// A custom error can contain structured information:
//
//     OpError{
//         Op:      "doSomething",
//         Code:    100,
//         Message: "failed",
//         Time:    time.Now(),
//     }
//
// This can be useful for:
//
//     - APIs
//     - backend services
//     - logging
//     - debugging
//     - monitoring
//     - HTTP errors
//     - database errors
//
// --------------------------------------------------------------------------------
//
// 12. MAKING OpError AN ERROR
//
// We have:
//
//     type OpError struct {
//         ...
//     }
//
// But a struct is NOT automatically an error.
//
// To make it an error, it must implement the `error` interface.
//
// The error interface is essentially:
//
//     type error interface {
//         Error() string
//     }
//
// So we create:
//
//     func (op OpError) Error() string {
//
//         return op.Message
//
//     }
//
// Now OpError satisfies the error interface.
//
// --------------------------------------------------------------------------------
//
// 13. WHAT DOES "IMPLEMENTS error" MEAN?
//
// In Go, interfaces are satisfied IMPLICITLY.
//
// We don't write:
//
//     implements error
//
// Instead, Go checks whether the type has the required method.
//
// `error` requires:
//
//     Error() string
//
// Our OpError has:
//
//     func (op OpError) Error() string
//
// Therefore:
//
//     OpError satisfies error
//
// Conceptually:
//
//     error interface
//          ↑
//          │
//     Error() string
//          ↑
//          │
//     OpError
//
// --------------------------------------------------------------------------------
//
// 14. THE RECEIVER
//
//     func (op OpError) Error() string
//
// `(op OpError)` is the RECEIVER.
//
// It means this method belongs to the OpError type.
//
// `op` represents the current OpError value.
//
// Therefore:
//
//     op.Message
//
// accesses the Message field of that particular error.
//
// --------------------------------------------------------------------------------
//
// 15. ERROR CONSTRUCTOR
//
// We create:
//
//     func NewOpError(
//         op string,
//         code int,
//         message string,
//         t time.Time,
//     ) *OpError
//
// This is a constructor-style function.
//
// Go doesn't have constructors like C++ or Java.
//
// Instead, developers commonly create functions such as:
//
//     NewOpError(...)
//
// to construct and initialize values.
//
// --------------------------------------------------------------------------------
//
// 16. WHY RETURN *OpError?
//
// The return type:
//
//     *OpError
//
// means:
//
//     "Return a pointer to an OpError."
//
// The function returns:
//
//     &OpError{
//         Op:      op,
//         Code:    code,
//         Message: message,
//         Time:    t,
//     }
//
// `&` gives us the address of the newly-created OpError.
//
// So:
//
//     NewOpError(...)
//
// returns:
//
//     *OpError
//
// --------------------------------------------------------------------------------
//
// 17. time.Time
//
// The struct contains:
//
//     Time time.Time
//
// `time.Time` is Go's standard type for representing a point in time.
//
// We need to import:
//
//     "time"
//
// Then:
//
//     time.Now()
//
// returns the current time.
//
// Example:
//
//     time.Now()
//
// could represent:
//
//     2026-08-18 01:20:00 +0530 IST
//
// --------------------------------------------------------------------------------
//
// 18. doSomething()
//
//     func doSomething() error {
//
//         return NewOpError(
//             "doSomething",
//             100,
//             "failed",
//             time.Now(),
//         )
//
//     }
//
// Notice that the function returns:
//
//     error
//
// but NewOpError returns:
//
//     *OpError
//
// This works because:
//
//     *OpError implements error
//
// Therefore a *OpError can be returned wherever an error is expected.
//
// This is one of the most important ideas here.
//
// --------------------------------------------------------------------------------
//
// 19. TYPE RELATIONSHIP
//
// Think of it like:
//
//     OpError
//        │
//        │ has method
//        ▼
//     Error() string
//        │
//        ▼
//     satisfies
//        │
//        ▼
//     error interface
//
// Therefore:
//
//     func doSomething() error
//
// can return:
//
//     *OpError
//
// --------------------------------------------------------------------------------
//
// 20. CURRENT main()
//
//     value, err := divide(10, 5)
//
// Since:
//
//     10 / 5 = 2
//
// the function returns:
//
//     value = 2
//     err   = nil
//
// Therefore:
//
//     if err != nil
//
// is false.
//
// So the program prints:
//
//     2
//
// --------------------------------------------------------------------------------
//
// 21. IF WE USED divide(10, 0)
//
//     value, err := divide(10, 0)
//
// returns:
//
//     value = 0
//     err   = ErrDivisionByZero
//
// Then:
//
//     errors.Is(err, ErrDivisionByZero)
//
// becomes:
//
//     true
//
// and:
//
//     fmt.Println("divide by zero")
//
// executes.
//
// --------------------------------------------------------------------------------
//
// 22. IF WE USED divide(2000, 5)
//
// The condition:
//
//     a > 1000
//
// becomes true.
//
// Therefore:
//
//     return 0, ErrNumberTooLarge
//
// Then:
//
//     errors.Is(err, ErrNumberTooLarge)
//
// returns true.
//
// --------------------------------------------------------------------------------
//
// 23. SENTINEL ERROR vs CUSTOM ERROR
//
// SENTINEL ERROR:
//
//     var ErrDivisionByZero = errors.New("division by zero")
//
// Best when:
//
//     You only need to identify a known category of error.
//
// Example:
//
//     errors.Is(err, ErrDivisionByZero)
//
//
//
// CUSTOM ERROR:
//
//     type OpError struct {
//         Op      string
//         Code    int
//         Message string
//         Time    time.Time
//     }
//
// Best when:
//
//     You need additional structured information.
//
// Example:
//
//     err.Code
//     err.Op
//     err.Message
//     err.Time
//
// --------------------------------------------------------------------------------
//
// 24. BIG PICTURE
//
// Basic error:
//
//     errors.New("something failed")
//
//             ↓
//
// Sentinel error:
//
//     ErrDivisionByZero
//
//             ↓
//
// Check known error:
//
//     errors.Is(err, ErrDivisionByZero)
//
//             ↓
//
// Custom error:
//
//     OpError{
//         Op,
//         Code,
//         Message,
//         Time,
//     }
//
//             ↓
//
// Implement error interface:
//
//     Error() string
//
//             ↓
//
// Rich, structured error handling
//
// --------------------------------------------------------------------------------
//
// KEY TAKEAWAYS:
//
// `errors.New()`
// → creates an error.
//
// `ErrSomething`
// → commonly used naming convention for sentinel errors.
//
// `errors.Is()`
// → checks whether an error matches another error.
//
// `error`
// → Go's standard error interface.
//
// `Error() string`
// → method required to satisfy the error interface.
//
// `type OpError struct`
// → creates a custom structured error type.
//
// `NewOpError()`
// → constructor-style function.
//
// `*OpError`
// → pointer to an OpError.
//
// `time.Time`
// → represents a point in time.
//
// `time.Now()`
// → gets the current time.
//
// `nil`
// → means no error when used with an error value.
//
// --------------------------------------------------------------------------------
//
// MOST IMPORTANT MENTAL MODEL:
//
//
//
//     Something goes wrong
//             ↓
//        create/return error
//             ↓
//     caller receives `err`
//             ↓
//        err != nil?
//          /       \
//        YES        NO
//         ↓          ↓
//    inspect it    continue
//         ↓
//   errors.Is()
//         ↓
// identify specific error
//
//
// And when simple errors aren't enough:
//
//     Custom Error Type
//            ↓
//     Op + Code + Message + Time
//            ↓
//     implements Error() string
//            ↓
//     becomes an `error`
//