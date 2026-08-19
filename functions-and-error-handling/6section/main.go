package main

import "fmt"

func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: deferred")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
} // whatever comes deferred will execute last. 

func lifoSimpleDefer() {
	fmt.Println("Function lifosimpleDefer: Start")
	defer fmt.Println("First: deferred")
	defer fmt.Println("Second: deferred")
	fmt.Println("Function lifosimpleDefer: Middle")
}

func main() {

	defer func() {
		fmt.Println("Before the return of main()")
	}()

	//simpleDefer()
	lifoSimpleDefer()

	fmt.Println("Last in main()")
}

// before main returns, defer closes the resources - common practice in go?

// --------------------------------------- Defer in Go ---------------------------------------
//
// `defer` schedules a function call to execute LATER.
//
// The deferred function executes when the surrounding function
// is about to RETURN.
//
// Basic syntax:
//
//     defer functionCall()
//
// Example:
//
//     defer fmt.Println("Done")
//
// The statement is registered immediately, but its execution is
// postponed until the current function is about to finish.
//
// --------------------------------------------------------------------------------
//
// 1. BASIC DEFER
//
//     func simpleDefer() {
//
//         fmt.Println("Start")
//
//         defer fmt.Println("Deferred")
//
//         fmt.Println("Middle")
//     }
//
// Execution:
//
//     Start
//     Middle
//     Deferred
//
// The deferred statement does NOT execute immediately.
//
// It executes when `simpleDefer()` is about to return.
//
// --------------------------------------------------------------------------------
//
// 2. IMPORTANT: DEFER IS ASSOCIATED WITH THE CURRENT FUNCTION
//
// If we have:
//
//     func main() {
//
//         defer fmt.Println("main deferred")
//
//         simpleDefer()
//     }
//
// The defer inside `simpleDefer()` executes when `simpleDefer()`
// returns.
//
// The defer inside `main()` executes when `main()` returns.
//
// Think:
//
//     main()
//       |
//       ├── simpleDefer()
//       |       └── its defer executes when simpleDefer returns
//       |
//       └── main's defer executes when main returns
//
// --------------------------------------------------------------------------------
//
// 3. MULTIPLE DEFER STATEMENTS
//
// Go executes multiple deferred calls in:
//
//     LIFO order
//
// LIFO = Last In, First Out
//
// Example:
//
//     defer fmt.Println("First")
//     defer fmt.Println("Second")
//
// Execution:
//
//     Second
//     First
//
// Why?
//
// The most recently deferred function is executed first.
//
// Think of a STACK:
//
//     defer First
//     defer Second
//
// Stack:
//
//     ┌─────────────┐
//     │   Second    │ ← comes out first
//     ├─────────────┤
//     │   First     │
//     └─────────────┘
//
// --------------------------------------------------------------------------------
//
// 4. YOUR `lifoSimpleDefer()` EXAMPLE
//
//     func lifoSimpleDefer() {
//
//         fmt.Println("Start")
//
//         defer fmt.Println("First: deferred")
//         defer fmt.Println("Second: deferred")
//
//         fmt.Println("Middle")
//     }
//
// Output:
//
//     Start
//     Middle
//     Second: deferred
//     First: deferred
//
// The second defer was registered LAST,
// therefore it executes FIRST.
//
// --------------------------------------------------------------------------------
//
// 5. DEFER EXECUTES BEFORE FUNCTION RETURNS
//
// Example:
//
//     func example() {
//
//         defer fmt.Println("Cleanup")
//
//         fmt.Println("Doing work")
//     }
//
// Execution:
//
//     Doing work
//     Cleanup
//     function returns
//
// So your comment:
//
//     "whatever comes deferred will execute last"
//
// is mostly correct, but more precisely:
//
//     "A deferred call executes when the surrounding function
//      is about to return."
//
// --------------------------------------------------------------------------------
//
// 6. DEFER AND `return`
//
// Example:
//
//     func example() int {
//
//         defer fmt.Println("Deferred")
//
//         return 10
//     }
//
// The deferred function executes before `example()` actually
// finishes returning to its caller.
//
// Conceptually:
//
//     return 10
//         ↓
//     execute deferred calls
//         ↓
//     function returns 10
//
// --------------------------------------------------------------------------------
//
// 7. DEFER IN `main()`
//
// Your code:
//
//     func main() {
//
//         defer func() {
//             fmt.Println("Before the return of main()")
//         }()
//
//         ...
//     }
//
// This anonymous function is deferred.
//
// Therefore it runs when `main()` is about to return.
//
// If the program reaches:
//
//     fmt.Println("Last in main()")
//
// then the deferred function runs AFTER that line.
//
// Output order:
//
//     Last in main()
//     Before the return of main()
//
// --------------------------------------------------------------------------------
//
// 8. DEFER WITH ANONYMOUS FUNCTIONS
//
// You can defer an anonymous function:
//
//     defer func() {
//
//         fmt.Println("Cleanup")
//
//     }()
//
// The `func() { ... }` creates an anonymous function.
//
// The final `()` CALLS that function.
//
// So:
//
//     defer func() {
//         ...
//     }()
//
// means:
//
//     "Create this function and defer its execution."
//
// --------------------------------------------------------------------------------
//
// 9. VERY IMPORTANT: ARGUMENTS ARE EVALUATED IMMEDIATELY
//
// Consider:
//
//     x := 10
//
//     defer fmt.Println(x)
//
//     x = 20
//
// Output:
//
//     10
//
// Why?
//
// The arguments to a deferred function call are evaluated
// when the `defer` statement is executed.
//
// The actual function call happens later.
//
// Think:
//
//     defer fmt.Println(x)
//
// immediately captures:
//
//     x = 10
//
// Then later it executes:
//
//     fmt.Println(10)
//
// --------------------------------------------------------------------------------
//
// 10. DEFER IS COMMONLY USED FOR CLEANUP
//
// This is one of the MOST important real-world uses of defer.
//
// Whenever you acquire a resource:
//
//     open file
//     open database connection
//     acquire mutex
//     network connection
//
// you often defer its cleanup immediately.
//
// Example:
//
//     file, err := os.Open("data.txt")
//
//     if err != nil {
//         return err
//     }
//
//     defer file.Close()
//
// Now no matter where the function returns,
// `file.Close()` will execute.
//
// This makes cleanup safer and easier to remember.
//
// --------------------------------------------------------------------------------
//
// 11. FILE EXAMPLE
//
//     file, err := os.Open("data.txt")
//
//     if err != nil {
//         return err
//     }
//
//     defer file.Close()
//
//     // work with file...
//
// Execution:
//
//     Open file
//        ↓
//     defer Close()
//        ↓
//     work with file
//        ↓
//     function finishes
//        ↓
//     Close() executes
//        ↓
//     function returns
//
// This pattern is extremely common in Go.
//
// --------------------------------------------------------------------------------
//
// 12. WHY DEFER IS USEFUL
//
// Without defer:
//
//     file, err := os.Open("data.txt")
//
//     if err != nil {
//         return err
//     }
//
//     // work
//
//     file.Close()
//
// But if there are multiple return paths:
//
//     if somethingWrong {
//         file.Close()
//         return err
//     }
//
//     if somethingElse {
//         file.Close()
//         return err
//     }
//
//     file.Close()
//
// This can become repetitive and error-prone.
//
// With defer:
//
//     defer file.Close()
//
// we register cleanup ONCE.
//
// --------------------------------------------------------------------------------
//
// 13. DEFER + ERROR HANDLING
//
// Very common Go pattern:
//
//     resource, err := acquire()
//
//     if err != nil {
//         return err
//     }
//
//     defer resource.Close()
//
//     // use resource
//
// Notice:
//
//     defer resource.Close()
//
// comes AFTER successfully acquiring the resource.
//
// We don't defer cleanup if acquisition itself failed.
//
// --------------------------------------------------------------------------------
//
// 14. DEFER + PANIC
//
// Deferred functions also execute when a function is unwinding
// because of a panic.
//
// This is one reason defer is useful for cleanup.
//
// However:
//
//     defer
//
// does NOT mean:
//
//     "catch errors"
//
// Go's normal error handling is still:
//
//     value, err := function()
//
//     if err != nil {
//         ...
//     }
//
// `defer` is mainly about scheduling cleanup/finalization work.
//
// --------------------------------------------------------------------------------
//
// 15. DEFER ORDER
//
// Suppose:
//
//     defer A()
//     defer B()
//     defer C()
//
// Execution:
//
//     C()
//     B()
//     A()
//
// Always remember:
//
//     LIFO
//
//     Last In → First Out
//
// --------------------------------------------------------------------------------
//
// 16. COMPLETE EXECUTION OF YOUR PROGRAM
//
// main()
//   ↓
// defer "Before the return of main()"
//   ↓
// lifoSimpleDefer()
//   ↓
// print "Start"
//   ↓
// defer "First"
//   ↓
// defer "Second"
//   ↓
// print "Middle"
//   ↓
// lifoSimpleDefer() returns
//   ↓
// execute "Second"
//   ↓
// execute "First"
//   ↓
// back to main()
//   ↓
// print "Last in main()"
//   ↓
// main() about to return
//   ↓
// execute "Before the return of main()"
//   ↓
// program ends
//
// --------------------------------------------------------------------------------
//
// KEY TAKEAWAYS:
//
// `defer functionCall()`
// → schedule a function call for later.
//
// Deferred calls execute when the surrounding function is
// about to return.
//
// Multiple defers follow:
//
//     LIFO
//
// Last defer → executes first.
//
// `defer` is heavily used for:
//
//     - file.Close()
//     - database connection cleanup
//     - mutex.Unlock()
//     - network connection cleanup
//     - releasing resources
//
// Arguments to a deferred call are evaluated immediately.
//
// `defer` is NOT an error-handling mechanism.
//
// It is primarily a convenient and reliable cleanup mechanism.
//
// --------------------------------------------------------------------------------
//
// MOST IMPORTANT MENTAL MODEL:
//
//     Acquire resource
//          ↓
//     defer cleanup()
//          ↓
//     do work
//          ↓
//     function is about to return 
//          ↓
//     cleanup executes
//          ↓
//     function actually returns 
//
//
// This is one of the Go patterns you'll use constantly in
// real backend/system programming.