// Panic and recovery is not exceptional handling. If coming from java, js, python then understand that 
// Go doesnt have exceptional handling, it has panic and recovery. 

package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("something went wrong in mightPanic")
	}

	fmt.Println("This function get executed without a panic")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from Panic:", r)
		}
	}()

	mightPanic(true)

}

func main() {

	recoverable()

}


// ============================================================
// PANIC & RECOVERY IN GO
// ============================================================
//
// Go does NOT use try/catch/finally as its normal error-handling
// mechanism.
//
// Normal, expected failures are usually handled using:
//
//     value, err := someFunction()
//     if err != nil {
//         // handle error
//     }
//
// `panic` and `recover` are different.
//
//     panic()   → interrupts normal execution and starts
//                 panic propagation / stack unwinding.
//
//     recover() → retrieves the currently active panic and
//                 stops that panic from propagating further.
//
// `recover()` is a built-in Go function.
//
// ------------------------------------------------------------
//
// 1. PANIC
//
// `panic()` is used when something has gone seriously wrong and
// normal execution cannot continue.
//
// Example:
//
//     panic("something went wrong")
//
// Once panic executes:
//
//     - normal execution of the current function stops
//     - the panic becomes active
//     - Go starts panic propagation
//     - stack unwinding begins
//     - deferred functions are executed
//
// ------------------------------------------------------------
//
// 2. WHAT HAPPENS DURING PANIC?
//
// Consider:
//
//     func mightPanic() {
//         panic("something went wrong")
//     }
//
// When `panic()` executes, the function does NOT continue normally.
//
// For example:
//
//     func mightPanic() {
//
//         panic("oops")
//
//         fmt.Println("Hello")
//     }
//
// `Hello` will NEVER execute.
//
// Why?
//
// Because panic immediately interrupts normal execution.
//
// ------------------------------------------------------------
//
// 3. FUNCTION CALL STACK
//
// Suppose:
//
//     func main() {
//         recoverable()
//     }
//
//     func recoverable() {
//         mightPanic()
//     }
//
//     func mightPanic() {
//         panic("oops")
//     }
//
// The conceptual call stack looks like:
//
//     ┌─────────────────────┐
//     │    mightPanic()     │  ← currently executing
//     ├─────────────────────┤
//     │    recoverable()    │
//     ├─────────────────────┤
//     │       main()        │
//     └─────────────────────┘
//
// Each function call creates an execution context that allows
// the runtime to keep track of the current function and where
// control should return.
//
// ------------------------------------------------------------
//
// 4. WHAT IS STACK UNWINDING?
//
// When a panic occurs, Go cannot continue normal execution.
//
// Therefore, the runtime starts UNWINDING the call stack.
//
// Simple meaning:
//
//     "Move back through the active function calls while
//      executing their deferred functions."
//
// Example:
//
//     main()
//       ↓
//     recoverable()
//       ↓
//     mightPanic()
//       ↓
//     panic()
//
// After panic:
//
//     mightPanic()
//          ↓
//     panic
//          ↓
//     stack unwinding starts
//          ↓
//     deferred functions execute
//
// ------------------------------------------------------------
//
// 5. DEFER + PANIC
//
// `defer` is extremely important during panic handling.
//
// Example:
//
//     func recoverable() {
//
//         defer func() {
//             // recovery logic
//         }()
//
//         mightPanic()
//     }
//
// The deferred function is NOT executed immediately.
//
// It is REGISTERED first.
//
// Normal execution continues:
//
//     defer registered
//          ↓
//     mightPanic()
//          ↓
//     panic()
//          ↓
//     stack unwinding
//          ↓
//     deferred function executes
//
// ------------------------------------------------------------
//
// 6. IMPORTANT: DEFER DOES NOT "WATCH" FOR PANIC
//
// This is a common misunderstanding.
//
// `defer` does NOT sit there waiting and detecting:
//
//     "Oh, panic happened!"
//
// Instead:
//
//     defer function
//
// means:
//
//     "When this surrounding function is about to finish,
//      execute this function."
//
// During normal return:
//
//     function work
//          ↓
//     function about to return
//          ↓
//     deferred function executes
//          ↓
//     function returns
//
// During panic:
//
//     panic
//          ↓
//     stack unwinding
//          ↓
//     deferred function executes
//
// ------------------------------------------------------------
//
// 7. RECOVER
//
// `recover()` is a BUILT-IN Go function.
//
// Its job is to retrieve the currently active panic.
//
// Example:
//
//     defer func() {
//
//         if r := recover(); r != nil {
//             fmt.Println("Recovered:", r)
//         }
//
//     }()
//
// If the panic was:
//
//     panic("something went wrong")
//
// then:
//
//     r := recover()
//
// can retrieve:
//
//     "something went wrong"
//
// ------------------------------------------------------------
//
// 8. RECOVER DOES NOT DETECT THE PANIC
//
// This distinction is VERY important.
//
// `recover()` does NOT:
//
//     "detect that stack unwinding has started"
//
// The runtime already knows that a panic occurred.
//
// The actual sequence is:
//
//     panic()
//        ↓
//     runtime knows panic occurred
//        ↓
//     panic propagation begins
//        ↓
//     stack unwinding begins
//        ↓
//     deferred function executes
//        ↓
//     recover()
//        ↓
//     retrieves the active panic
//
// Therefore:
//
//     panic()  → creates the panic
//     runtime  → performs unwinding
//     recover() → retrieves/recover the active panic
//
// ------------------------------------------------------------
//
// 9. WHY MUST RECOVER BE INSIDE DEFER?
//
// `recover()` is useful for recovering a panic only when called
// from a deferred function during panic handling.
//
// Correct:
//
//     defer func() {
//
//         if r := recover(); r != nil {
//             fmt.Println("Recovered:", r)
//         }
//
//     }()
//
// Incorrect mental model:
//
//     recover()
//
// sitting somewhere before the panic happens.
//
// `recover()` does not continuously monitor the program.
//
// ------------------------------------------------------------
//
// 10. COMPLETE PANIC → RECOVER FLOW
//
//     main()
//       ↓
//     recoverable()
//       ↓
//     defer registered
//       ↓
//     mightPanic(true)
//       ↓
//     panic() executes
//       ↓
//     normal execution stops
//       ↓
//     panic propagation begins
//       ↓
//     stack unwinding starts
//       ↓
//     deferred function executes
//       ↓
//     recover() retrieves active panic
//       ↓
//     panic propagation stops
//       ↓
//     recoverable() finishes
//       ↓
//     main() continues
//
// ------------------------------------------------------------
//
// 11. YOUR EXAMPLE
//
//     func recoverable() {
//
//         defer func() {
//
//             if r := recover(); r != nil {
//                 fmt.Println("Recovered from Panic:", r)
//             }
//
//         }()
//
//         mightPanic(true)
//     }
//
// Execution:
//
//     recoverable()
//          ↓
//     defer function REGISTERED
//          ↓
//     mightPanic(true)
//          ↓
//     panic("something went wrong")
//          ↓
//     mightPanic() normal execution stops
//          ↓
//     stack unwinding begins
//          ↓
//     deferred function executes
//          ↓
//     recover()
//          ↓
//     panic value retrieved
//          ↓
//     panic propagation stops
//
// ------------------------------------------------------------
//
// 12. PANIC VALUE
//
// `panic()` can receive a value:
//
//     panic("something went wrong")
//
// The value passed to panic becomes the panic value.
//
// `recover()` retrieves that value:
//
//     r := recover()
//
// Therefore:
//
//     panic("oops")
//
// gives:
//
//     r == "oops"
//
// The panic value does not have to be a string; panic can be
// called with a value of any type.
//
// ------------------------------------------------------------
//
// 13. WHAT IF WE DON'T RECOVER?
//
// Suppose:
//
//     func mightPanic() {
//         panic("oops")
//     }
//
// And nobody recovers the panic.
//
// Then:
//
//     panic()
//       ↓
//     stack unwinding
//       ↓
//     deferred functions execute
//       ↓
//     panic continues propagating
//       ↓
//     eventually program terminates
//
// So `recover()` provides a controlled point where a panic can
// be recovered instead of continuing to propagate.
//
// ------------------------------------------------------------
//
// 14. PANIC vs ERROR
//
// `error`:
//
//     Used for expected or normally handleable failures.
//
// Example:
//
//     value, err := divide(10, 0)
//
//     if err != nil {
//         fmt.Println(err)
//     }
//
// Think:
//
//     "This operation failed, but the program knows how to
//      handle the failure."
//
//
//
// `panic`:
//
//     Used when normal execution cannot reasonably continue.
//
// Example:
//
//     panic("invalid internal state")
//
// Think:
//
//     "Something has gone seriously wrong."
//
//
//
// `recover`:
//
//     Used at a controlled boundary when we deliberately want
//     to recover from a panic.
//
// ------------------------------------------------------------
//
// 15. PANIC IS NOT A REPLACEMENT FOR ERROR
//
// DON'T normally do this:
//
//     func divide(a, b int) int {
//
//         if b == 0 {
//             panic("division by zero")
//         }
//
//         return a / b
//     }
//
// for an ordinary input validation problem.
//
// Prefer:
//
//     func divide(a, b int) (int, error) {
//
//         if b == 0 {
//             return 0, errors.New("division by zero")
//         }
//
//         return a / b, nil
//     }
//
// Normal Go philosophy:
//
//     EXPECTED FAILURE → error
//
//     ABNORMAL / SERIOUS FAILURE → panic
//
//     CONTROLLED RECOVERY FROM PANIC → recover
//
// ------------------------------------------------------------
//
// 16. DEFER + RECOVER PATTERN
//
// The common pattern is:
//
//     defer func() {
//
//         if r := recover(); r != nil {
//             // handle recovered panic
//         }
//
//     }()
//
// This creates a recovery boundary.
//
// It is especially useful when you want to prevent a panic from
// bringing down a larger part of the application.
//
// ------------------------------------------------------------
//
// 17. VERY IMPORTANT MENTAL MODEL
//
// Think of the whole mechanism as:
//
//
//     panic()
//        │
//        ▼
//     "Something went seriously wrong."
//        │
//        ▼
//     Normal execution stops
//        │
//        ▼
//     Stack unwinding begins
//        │
//        ▼
//     Deferred functions execute
//        │
//        ▼
//     recover()
//        │
//        ▼
//     Retrieve active panic
//        │
//        ▼
//     Stop panic propagation
//
// ------------------------------------------------------------
//
// 18. THREE ROLES TO REMEMBER
//
//     panic()
//     ↓
//     STARTS the panic.
//
//     runtime
//     ↓
//     HANDLES the panic propagation and stack unwinding.
//
//     recover()
//     ↓
//     RETRIEVES the active panic and RECOVERS from it.
//
// ------------------------------------------------------------
//
// 19. ONE-LINE DEFINITIONS
//
//     panic()
//     → Stops normal execution and starts panic propagation.
//
//     defer
//     → Schedules a function to execute when the surrounding
//       function is about to return or while panic unwinding
//       passes through it.
//
//     stack unwinding
//     → The runtime moving back through active function calls
//       during panic propagation, executing deferred functions.
//
//     recover()
//     → Retrieves the currently active panic from a deferred
//       function and stops that panic from propagating further.
//
// ------------------------------------------------------------
//
// FINAL MENTAL MODEL:
//
//     FUNCTION EXECUTION
//            ↓
//        panic()
//            ↓
//     normal execution stops
//            ↓
//     stack unwinding
//            ↓
//     deferred functions
//            ↓
//        recover()
//            ↓
//     panic recovered
//            ↓
//     surrounding execution can continue/finish normally
//
// ============================================================
//
// IMPORTANT:
//
// Don't memorize this as:
//
//     "defer catches panic"
//
// Better mental model:
//
//     "panic starts unwinding,
//      unwinding executes deferred functions,
//      and recover inside a deferred function can recover
//      the active panic."
//
// ============================================================