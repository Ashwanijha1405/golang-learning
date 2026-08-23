package main 

import "fmt" 

type Person interface {
	GetName() string
}


type BusinessPerson struct {
	ID int
	Name string
}

func (e BusinessPerson) String() string {
	return fmt.Sprintf("Person[ID:%d, Name:%s]", e.ID, e.Name)
}

type ID int 

func (idx ID) String() string {
	return fmt.Sprintf("ID: %d", idx)
}

func main() {

	jane := BusinessPerson{
		ID: 1, 
		Name: "Jane",
	}

	fmt.Println(jane)

	myId := ID(30)
	fmt.Println(myId)

}


// ============================================================
// STRINGER INTERFACE IN GO
// ============================================================
//
// The `Stringer` interface is provided by Go's `fmt` package.
//
// It allows a custom type to control how its value is converted
// into a human-readable string.
//
// The interface is:
//
//     type Stringer interface {
//         String() string
//     }
//
// If a type implements:
//
//     String() string
//
// then `fmt` functions such as `fmt.Println`, `fmt.Printf`,
// etc. can automatically use that method when printing the value.
//
// ============================================================

// ============================================================
// 1. WHY DO WE NEED Stringer?
// ============================================================
//
// Suppose we have:
//
//     type BusinessPerson struct {
//         ID   int
//         Name string
//     }
//
// Without implementing Stringer:
//
//     fmt.Println(jane)
//
// Go prints the struct using its default representation:
//
//     {1 Jane}
//
// This works, but it may not be the format we want.
//
// With Stringer, we can define our own representation:
//
//     Person[ID:1, Name:Jane]
//
// So Stringer is basically a way of telling Go:
//
//     "Whenever someone prints this value,
//      use THIS method to represent it as a string."

// ============================================================
// 2. Stringer INTERFACE
// ============================================================
//
// The Stringer interface is defined inside the `fmt` package:
//
//     type Stringer interface {
//         String() string
//     }
//
// There is only ONE method:
//
//     String() string
//
// Any type that has this method automatically satisfies
// the Stringer interface.
//
// There is no explicit:
//
//     implements Stringer
//
// statement in Go.
//
// This is called implicit interface implementation.

// ============================================================
// 3. BUSINESSPERSON IMPLEMENTING Stringer
// ============================================================
//
// Our type:
//
//     type BusinessPerson struct {
//         ID   int
//         Name string
//     }
//
// implements Stringer because it has:
//
//     func (e BusinessPerson) String() string
//
// The receiver is BusinessPerson and the return type is string.
//
// Therefore:
//
//     BusinessPerson satisfies fmt.Stringer.

// ============================================================
// 4. HOW fmt.Println USES Stringer
// ============================================================
//
// When we write:
//
//     fmt.Println(jane)
//
// `fmt` checks whether the value has a suitable String()
// representation.
//
// Since BusinessPerson implements:
//
//     String() string
//
// fmt can call:
//
//     jane.String()
//
// internally when formatting the value.
//
// Therefore:
//
//     fmt.Println(jane)
//
// effectively gets the string returned by:
//
//     jane.String()
//
// In our example:
//
//     return fmt.Sprintf(
//         "Person[ID:%d, Name:%s]",
//         e.ID,
//         e.Name,
//     )
//
// produces:
//
//     Person[ID:1, Name:Jane]
//
// So the output becomes:
//
//     Person[ID:1, Name:Jane]

// ============================================================
// 5. IMPORTANT: Stringer IS AN INTERFACE
// ============================================================
//
// `fmt.Stringer` is an interface type.
//
// Conceptually:
//
//     fmt.Stringer
//          |
//          ↓
//     String() string
//
// Any type having String() string satisfies it.
//
// For example:
//
//     BusinessPerson
//
// satisfies it.
//
// And:
//
//     ID
//
// also satisfies it.
//
// They are completely different types, but both implement
// the same interface.

// ============================================================
// 6. CUSTOM TYPES CAN ALSO IMPLEMENT Stringer
// ============================================================
//
// In the example:
//
//     type ID int
//
// `ID` is a new named type whose underlying type is int.
//
// We can define methods on it:
//
//     func (idx ID) String() string {
//         return fmt.Sprintf("ID: %d", idx)
//     }
//
// Therefore ID also satisfies Stringer.
//
//
// When we write:
//
//     myId := ID(30)
//
// and then:
//
//     fmt.Println(myId)
//
// fmt can use:
//
//     myId.String()
//
// which returns:
//
//     ID: 30
//
// Output:
//
//     ID: 30

// ============================================================
// 7. WITHOUT Stringer vs WITH Stringer
// ============================================================
//
// WITHOUT Stringer:
//
//     type BusinessPerson struct {
//         ID   int
//         Name string
//     }
//
//     jane := BusinessPerson{
//         ID:   1,
//         Name: "Jane",
//     }
//
//     fmt.Println(jane)
//
// Output:
//
//     {1 Jane}
//
//
//
// WITH Stringer:
//
//     func (e BusinessPerson) String() string {
//         return fmt.Sprintf(
//             "Person[ID:%d, Name:%s]",
//             e.ID,
//             e.Name,
//         )
//     }
//
//     fmt.Println(jane)
//
// Output:
//
//     Person[ID:1, Name:Jane]
//
// Stringer therefore gives us control over the textual
// representation of our custom types.

// ============================================================
// 8. Stringer IS AUTOMATICALLY USED BY fmt
// ============================================================
//
// One of the most useful things about Stringer is that we
// don't normally have to manually call:
//
//     jane.String()
//
// Instead:
//
//     fmt.Println(jane)
//
// is enough.
//
// The fmt package recognizes the Stringer interface and uses
// the String() method when formatting the value.
//
// This makes custom types much cleaner to work with.

// ============================================================
// 9. Stringer METHOD SHOULD RETURN A HUMAN-READABLE STRING
// ============================================================
//
// The purpose of String() is generally to provide a useful,
// human-readable representation of the value.
//
// Good:
//
//     Person[ID:1, Name:Jane]
//
// Less useful:
//
//     1 Jane true 50000000 ...
//
// The String() method is commonly useful for:
//
//     - logging
//     - debugging
//     - displaying objects
//     - error messages
//     - CLI output
//     - formatted output

// ============================================================
// 10. VALUE RECEIVER vs POINTER RECEIVER
// ============================================================
//
// In our example:
//
//     func (e BusinessPerson) String() string
//
// uses a VALUE receiver.
//
// Therefore both:
//
//     BusinessPerson
//
// and, through Go's method-set rules:
//
//     *BusinessPerson
//
// can generally be used where fmt.Stringer is expected.
//
//
//
// If we instead wrote:
//
//     func (e *BusinessPerson) String() string
//
// then only:
//
//     *BusinessPerson
//
// would satisfy the Stringer interface.
//
// This is an important method-set rule to remember:
//
//     Value receiver
//         → method belongs to T and *T
//
//     Pointer receiver
//         → method belongs to *T only

// ============================================================
// 11. Stringer IS JUST AN INTERFACE — NOT MAGIC
// ============================================================
//
// There is nothing special about BusinessPerson itself.
//
// Go simply sees:
//
//     BusinessPerson
//          |
//          ↓
//     String() string
//
// Therefore:
//
//     BusinessPerson implements fmt.Stringer
//
// This is the same interface concept you learned earlier:
//
//     type Person interface {
//         GetName() string
//     }
//
// A type satisfies an interface by having all required methods.
//
// Stringer simply happens to be an interface provided by the
// standard library.

// ============================================================
// 12. KEY TAKEAWAY
// ============================================================
//
// `fmt.Stringer`:
//
//     type Stringer interface {
//         String() string
//     }
//
// A custom type becomes a Stringer when it implements:
//
//     String() string
//
// Once it does, fmt functions can automatically use that method
// to produce a custom textual representation.
//
//
//
// Remember it like this:
//
//     Custom Type
//          |
//          | implements
//          ↓
//     String() string
//          |
//          ↓
//     fmt.Stringer
//          |
//          ↓
//     fmt.Println(value)
//          |
//          ↓
//     value.String()
//          |
//          ↓
//     Human-readable output
//
//
// ============================================================
// ONE-LINE MEMORY TRICK
// ============================================================
//
// Stringer = "I decide how my custom type should look when printed."
//
// ============================================================
