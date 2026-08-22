package main 

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int 
	FirstName string
	LastName string 
	Position string
	Salary int
	IsActive bool 
	JoinedAt time.Time 
}


// A value reciever
func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

 
func (e Employee) Deactivate() {
	e.IsActive = false
}


func deactivate(e * Employee) {
	e.IsActive = false; 
}

func NewEmployee(id int, firstName, lastName, position string, isActive bool) Employee {
	return Employee {
		ID: id, 
		FirstName: firstName, 
		LastName: lastName, 
		Position: position, 
		IsActive: isActive, 
		JoinedAt: time.Now(),
	}
}

func main() {

	ashwani := Employee {
		ID: 1, 
		FirstName: "Ashwani", 
		LastName: "Jha",
		Position: "CEO",
		Salary: 50000000, 
		IsActive: true, 
		JoinedAt: time.Now(),
	}

	fmt.Println(ashwani.FullName())
	//ashwani.Deactivate()
	deactivate(&ashwani)
	fmt.Printf("%+v\n", ashwani)

}


// ============================================================================
//                         METHODS & RECEIVERS IN GO
// ============================================================================
//
// Go does not have traditional classes like Java or C++.
//
// Instead, Go commonly combines:
//
//     struct + methods
//
// to attach behavior to a custom data type.
//
// A METHOD is simply a function that is associated with a particular type.
//
// ============================================================================
// 1. WHAT IS A METHOD?
// ============================================================================
//
// A normal function looks like:
//
//     func FullName(e Employee) string
//
// A method looks like:
//
//     func (e Employee) FullName() string
//
// The important difference is the RECEIVER:
//
//     (e Employee)
//
// The receiver tells Go:
//
//     "This method belongs to the Employee type."
//
// Therefore, instead of calling:
//
//     FullName(ashwani)
//
// we can call:
//
//     ashwani.FullName()
//
// This makes the code feel similar to object-oriented programming while
// still following Go's simpler type system.
//
// ============================================================================
// 2. METHOD RECEIVER
// ============================================================================
//
// The part:
//
//     (e Employee)
//
// is called the RECEIVER.
//
// It consists of:
//
//     e         -> receiver variable name
//     Employee  -> receiver type
//
// `e` is available inside the method just like a normal parameter.
//
// Example:
//
//     func (e Employee) FullName() string {
//         return e.FirstName + " " + e.LastName
//     }
//
// Here:
//
//     e.FirstName
//     e.LastName
//
// access the fields of the Employee on which the method was called.
//
// ============================================================================
// 3. CALLING A METHOD
// ============================================================================
//
// If we have:
//
//     ashwani := Employee{...}
//
// then:
//
//     ashwani.FullName()
//
// calls the FullName method.
//
// Conceptually, you can think of:
//
//     ashwani.FullName()
//
// as being similar to:
//
//     Employee.FullName(ashwani)
//
// The actual Go syntax and method rules are more specific, but this mental
// model is useful for understanding what the receiver is doing.
//
// ============================================================================
// 4. VALUE RECEIVER
// ============================================================================
//
// Our FullName method uses:
//
//     func (e Employee) FullName() string
//
// This is called a VALUE RECEIVER.
//
// `e` receives a COPY of the Employee value.
//
// Therefore, if the method changes `e`, it changes only the copy.
//
// The original Employee is NOT modified.
//
// Example:
//
//     func (e Employee) Deactivate() {
//         e.IsActive = false
//     }
//
// If we call:
//
//     ashwani.Deactivate()
//
// the `e` inside Deactivate is a copy.
//
// So:
//
//     e.IsActive = false
//
// modifies the copy, not the original `ashwani`.
//
// After the method finishes, that copy is discarded.
//
// ============================================================================
// 5. WHY DID `Deactivate()` NOT WORK?
// ============================================================================
//
// This method:
//
//     func (e Employee) Deactivate() {
//         e.IsActive = false
//     }
//
// looks like it should deactivate the employee.
//
// But it doesn't modify the original Employee because `e` is a value
// receiver.
//
// Flow:
//
//     ashwani
//        |
//        | copied
//        v
//     e (copy)
//        |
//        | IsActive = false
//        v
//     copy changes
//
// Original:
//
//     ashwani.IsActive
//
// remains unchanged.
//
// ============================================================================
// 6. POINTER RECEIVER
// ============================================================================
//
// If we want a method to modify the original Employee, we can use a
// POINTER RECEIVER:
//
//     func (e *Employee) Deactivate() {
//         e.IsActive = false
//     }
//
// Now `e` is a pointer to the original Employee.
//
// Therefore:
//
//     e.IsActive = false
//
// modifies the actual Employee.
//
// The pointer receiver is useful when:
//
//     -> the method needs to modify the receiver
//     -> the struct is large and copying it is undesirable
//     -> we want consistent pointer-based method behavior
//
// ============================================================================
// 7. YOUR `deactivate()` FUNCTION VS METHOD
// ============================================================================
//
// You also wrote:
//
//     func deactivate(e *Employee) {
//         e.IsActive = false
//     }
//
// This is a NORMAL FUNCTION, not a method.
//
// It accepts an Employee pointer as an ordinary parameter.
//
// Therefore:
//
//     deactivate(&ashwani)
//
// works because `&ashwani` gives the address of the original Employee.
//
// Compare:
//
//     NORMAL FUNCTION:
//
//     func deactivate(e *Employee)
//
//     call:
//
//     deactivate(&ashwani)
//
//
//
//     METHOD WITH POINTER RECEIVER:
//
//     func (e *Employee) Deactivate()
//
//     call:
//
//     ashwani.Deactivate()
//
// Both can modify the original Employee.
//
// The difference is mainly how the behavior is organized and called.
//
// ============================================================================
// 8. METHOD WITH VALUE RECEIVER
// ============================================================================
//
//     func (e Employee) FullName() string
//
// The receiver is an Employee VALUE.
//
// Conceptually:
//
//     ashwani
//         |
//         | copy
//         v
//     e
//
// The method can READ the Employee and perform calculations without
// modifying the original.
//
// This is perfect for methods such as:
//
//     FullName()
//     Age()
//     DisplayName()
//     CalculateSomething()
//
// when they don't need to modify the receiver.
//
// ============================================================================
// 9. METHOD WITH POINTER RECEIVER
// ============================================================================
//
//     func (e *Employee) Deactivate()
//
// The receiver is a POINTER to Employee.
//
// Conceptually:
//
//     ashwani
//         ^
//         |
//         | points to
//         |
//         e
//
// Changes made through `e` affect the original Employee.
//
// Example:
//
//     func (e *Employee) Deactivate() {
//         e.IsActive = false
//     }
//
// Then:
//
//     ashwani.Deactivate()
//
// changes:
//
//     ashwani.IsActive
//
// to:
//
//     false
//
// ============================================================================
// 10. GO'S AUTOMATIC POINTER HANDLING FOR METHODS
// ============================================================================
//
// Go provides a convenient feature when calling methods.
//
// Suppose:
//
//     func (e *Employee) Deactivate()
//
// and:
//
//     ashwani := Employee{...}
//
// We can write:
//
//     ashwani.Deactivate()
//
// even though the method expects:
//
//     *Employee
//
// Go automatically takes the address when the value is addressable.
//
// Conceptually:
//
//     ashwani.Deactivate()
//
// becomes similar to:
//
//     (&ashwani).Deactivate()
//
// This is why you usually don't need to manually write `&` when calling
// pointer-receiver methods on an addressable variable.
//
// ============================================================================
// 11. VALUE RECEIVER vs POINTER RECEIVER
// ============================================================================
//
// VALUE RECEIVER:
//
//     func (e Employee) FullName() string
//
//     -> receives a copy
//     -> cannot modify the original through `e`
//     -> useful for read-only behavior
//
//
// POINTER RECEIVER:
//
//     func (e *Employee) Deactivate()
//
//     -> receives a pointer
//     -> can modify the original
//     -> avoids copying the entire struct
//
// ============================================================================
// 12. IMPORTANT: POINTER RECEIVER DOES NOT MEAN "POINTER ONLY"
// ============================================================================
//
// A common beginner confusion:
//
//     func (e *Employee) Deactivate()
//
// does NOT mean that you must always have a pointer variable.
//
// If `ashwani` is an addressable Employee value, Go can automatically take
// its address when calling the method:
//
//     ashwani.Deactivate()
//
// So this is usually fine.
//
// You would explicitly use a pointer when needed:
//
//     ptr := &ashwani
//     ptr.Deactivate()
//
// ============================================================================
// 13. WHY RECEIVER NAME IS USUALLY SHORT
// ============================================================================
//
// Go convention commonly uses short receiver names:
//
//     func (e Employee) FullName()
//     func (e *Employee) Deactivate()
//
// Instead of:
//
//     func (employee Employee) FullName()
//
// The short receiver name is idiomatic Go, especially when the type name
// already makes the meaning obvious.
//
// ============================================================================
// 14. METHODS BELONG TO TYPES
// ============================================================================
//
// A method is associated with its receiver type.
//
// For example:
//
//     Employee
//
// can have methods such as:
//
//     FullName()
//     Deactivate()
//     Promote()
//     CalculateSalary()
//
// This allows us to keep data and behavior conceptually together:
//
//     Employee
//       |
//       +-- ID
//       +-- FirstName
//       +-- LastName
//       +-- Salary
//       +-- IsActive
//       |
//       +-- FullName()
//       +-- Deactivate()
//
// This is one of the ways Go supports encapsulation and organization
// without using traditional classes.
//
// ============================================================================
// 15. GENERAL RULE FOR CHOOSING A RECEIVER
// ============================================================================
//
// A useful beginner rule:
//
//     Use a VALUE RECEIVER when:
//         -> the method only needs to read the value
//         -> copying the value is acceptable
//
//     Use a POINTER RECEIVER when:
//         -> the method needs to modify the receiver
//         -> the struct is large
//         -> you want to avoid copying
//
// Example:
//
//     FullName()
//
//     -> value receiver makes sense
//
//     Deactivate()
//
//     -> pointer receiver makes sense because it modifies IsActive
//
// ============================================================================
// 16. ONE IMPORTANT CONSISTENCY RULE
// ============================================================================
//
// If a type has several methods, Go code commonly uses the same receiver
// style consistently when appropriate.
//
// For example, if Employee has many methods that modify its state:
//
//     func (e *Employee) Deactivate()
//     func (e *Employee) Promote()
//     func (e *Employee) UpdateSalary()
//
// using pointer receivers consistently makes the behavior easier to reason
// about.
//
// ============================================================================
// 17. CORE MENTAL MODEL
// ============================================================================
//
// Think of a method as:
//
//     FUNCTION
//        +
//     RECEIVER
//        =
//     METHOD
//
// Example:
//
//     func (e Employee) FullName() string
//          ^      ^
//          |      |
//      receiver  method
//
//
//
// VALUE RECEIVER:
//
//     Employee
//         ↓
//       copy
//         ↓
//     method works on copy
//
//
// POINTER RECEIVER:
//
//     *Employee
//         ↓
//      address
//         ↓
//     method works with original
//
// ============================================================================
// 18. FINAL TAKEAWAYS
// ============================================================================
//
// -> A method is a function associated with a type.
//
// -> The receiver appears before the method name:
//
//        func (e Employee) FullName()
//
// -> `(e Employee)` is a value receiver.
//
// -> `(e *Employee)` is a pointer receiver.
//
// -> Value receivers receive a copy of the receiver.
//
// -> Pointer receivers receive a pointer to the receiver and can modify
//    the original value.
//
// -> Your `FullName()` method is a good example of a value receiver because
//    it only reads the Employee's fields.
//
// -> Your `Deactivate()` method currently uses a value receiver, so it does
//    NOT modify the original Employee.
//
// -> `deactivate(&ashwani)` is a normal function using a pointer parameter,
//    not a method.
//
// -> A pointer-receiver method can usually be called directly on an
//    addressable value because Go automatically takes its address.
//
// -> The key distinction is:
//
//        Value receiver  -> copy
//        Pointer receiver -> original
//
// ============================================================================
//                         ONE-LINE MEMORY TRICK
// ============================================================================
//
//     READ → value receiver is often enough
//     MODIFY → pointer receiver
//
// ============================================================================

