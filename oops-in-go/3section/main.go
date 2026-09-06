package main 

import "fmt" 

type Person interface {
	GetName() string
}

type Employee struct {
	ID  int 
	Name string
}

type BusinessPerson struct {
	ID int
	Name string
}

func (e BusinessPerson) GetName() string {
	return e.Name
}

func (e Employee) GetName() string {
	return e.Name
}

func displayPerson(p Person) {
	fmt.Println(p.GetName())
}

func main() {

	joe := Employee {
		ID:  1, 
		Name: "Joe",
	}

	jane := BusinessPerson{
		ID: 1, 
		Name: "Jane",
	}

	displayPerson(jane)
	displayPerson(joe)

}


// ============================================================================
//                              INTERFACES IN GO
// ============================================================================
//
// An interface defines a SET OF BEHAVIORS that a type must provide.
//
// Instead of saying:
//
//     "This function only accepts Employee"
//
// an interface allows us to say:
//
//     "This function accepts anything that can provide GetName()."
//
// This is one of the most important ideas in Go:
//
//     PROGRAM TO BEHAVIOR, NOT TO A CONCRETE TYPE.
//
// ============================================================================
// 1. DEFINING AN INTERFACE
// ============================================================================
//
//     type Person interface {
//         GetName() string
//     }
//
// `Person` is an interface type.
//
// It says:
//
//     Any type that has a method:
//
//         GetName() string
//
//     satisfies the Person interface.
//
// Notice that the interface does NOT contain any implementation.
//
// It only describes WHAT behavior is required.
//
// ============================================================================
// 2. INTERFACE = CONTRACT
// ============================================================================
//
// Think of an interface as a contract.
//
//     Person
//        |
//        +---- GetName() string
//
// Any type that wants to be treated as a `Person` must satisfy this contract.
//
// For example:
//
//     Employee
//         |
//         +---- GetName() string
//
//     BusinessPerson
//         |
//         +---- GetName() string
//
// Both satisfy the Person interface.
//
// ============================================================================
// 3. EMPLOYEE IMPLEMENTS PERSON
// ============================================================================
//
// Employee has:
//
//     func (e Employee) GetName() string {
//         return e.Name
//     }
//
// The Person interface requires:
//
//     GetName() string
//
// Employee provides exactly that method.
//
// Therefore:
//
//     Employee satisfies Person.
//
// There is NO explicit declaration like:
//
//     implements Person
//
// This is very important.
//
// Go uses IMPLICIT INTERFACE IMPLEMENTATION.
//
// ============================================================================
// 4. BUSINESSPERSON ALSO IMPLEMENTS PERSON
// ============================================================================
//
// BusinessPerson also defines:
//
//     func (e BusinessPerson) GetName() string {
//         return e.Name
//     }
//
// Therefore BusinessPerson also satisfies:
//
//     Person
//
// So both types can be used wherever a Person is expected.
//
// ============================================================================
// 5. NO `implements` KEYWORD IN GO
// ============================================================================
//
// Languages such as Java commonly use:
//
//     class Employee implements Person
//
// Go does NOT do this.
//
// You simply define the required methods.
//
// If a type has all the methods required by an interface:
//
//     Type -> automatically satisfies interface
//
// This is called IMPLICIT INTERFACE SATISFACTION.
//
// ============================================================================
// 6. THE `displayPerson` FUNCTION
// ============================================================================
//
//     func displayPerson(p Person) {
//         fmt.Println(p.GetName())
//     }
//
// This function does NOT care whether `p` is:
//
//     Employee
//     BusinessPerson
//     Customer
//     Manager
//     Student
//
// It only cares that `p` satisfies:
//
//     Person
//
// In other words, the function depends on the BEHAVIOR:
//
//     GetName()
//
// rather than the concrete type.
//
// ============================================================================
// 7. WHAT HAPPENS WHEN `displayPerson(jane)` IS CALLED?
// ============================================================================
//
// `jane` is a BusinessPerson.
//
// We call:
//
//     displayPerson(jane)
//
// The parameter expects:
//
//     Person
//
// BusinessPerson satisfies Person because it has:
//
//     GetName() string
//
// Therefore Go allows the conversion/assignment of the concrete value to
// the interface value.
//
// Inside the function:
//
//     p.GetName()
//
// calls the BusinessPerson implementation of GetName().
//
// Therefore the output is:
//
//     Jane
//
// ============================================================================
// 8. WHAT HAPPENS WHEN `displayPerson(joe)` IS CALLED?
// ============================================================================
//
// `joe` is an Employee.
//
// We call:
//
//     displayPerson(joe)
//
// Employee satisfies Person because it has:
//
//     GetName() string
//
// Therefore `joe` can be passed as a Person.
//
// Inside:
//
//     p.GetName()
//
// calls Employee's GetName() method.
//
// Therefore the output is:
//
//     Joe
//
// ============================================================================
// 9. POLYMORPHISM
// ============================================================================
//
// This is an example of POLYMORPHISM.
//
// The same function:
//
//     displayPerson()
//
// can work with multiple concrete types:
//
//     Employee
//     BusinessPerson
//
// because both satisfy the same interface.
//
// The function doesn't need separate versions such as:
//
//     displayEmployee()
//     displayBusinessPerson()
//
// Instead:
//
//     displayPerson(Person)
//
// handles both.
//
// This is one of the biggest practical benefits of interfaces.
//
// ============================================================================
// 10. THE IMPORTANT IDEA: BEHAVIOR OVER TYPE
// ============================================================================
//
// Without an interface, we might write:
//
//     func displayEmployee(e Employee)
//
// This function is tightly coupled to Employee.
//
// It cannot directly accept BusinessPerson.
//
// With an interface:
//
//     func displayPerson(p Person)
//
// the function only depends on:
//
//     GetName()
//
// Therefore any future type that provides GetName() can also be used.
//
// Example:
//
//     type Student struct {
//         Name string
//     }
//
//     func (s Student) GetName() string {
//         return s.Name
//     }
//
// Student automatically satisfies Person.
//
// We do NOT need to modify displayPerson.
//
// This makes code easier to extend.
//
// ============================================================================
// 11. INTERFACE DOES NOT STORE "ANYTHING"
// ============================================================================
//
// A beginner-friendly way to think about an interface variable:
//
//     p Person
//
// can hold a value of any concrete type that satisfies Person.
//
// For example:
//
//     p = joe
//
// or:
//
//     p = jane
//
// But the interface only exposes the methods defined by the interface.
//
// Since Person defines:
//
//     GetName()
//
// we can safely call:
//
//     p.GetName()
//
// But we cannot directly assume that `p` has some Employee-specific field
// such as:
//
//     p.ID
//
// because ID is not part of the Person interface.
//
// ============================================================================
// 12. INTERFACE HIDES IMPLEMENTATION DETAILS
// ============================================================================
//
// The caller knows:
//
//     p.GetName()
//
// The caller does NOT need to know:
//
//     How Employee.GetName() works
//
// or:
//
//     How BusinessPerson.GetName() works
//
// Each concrete type handles its own implementation.
//
// This provides a useful separation:
//
//     INTERFACE
//          ↓
//     defines WHAT
//          ↓
//     CONCRETE TYPE
//          ↓
//     defines HOW
//
// ============================================================================
// 13. INTERFACE SATISFACTION IS ALL OR NOTHING
// ============================================================================
//
// Suppose we have:
//
//     type Person interface {
//         GetName() string
//         GetID() int
//     }
//
// Now a type must provide BOTH methods:
//
//     GetName() string
//     GetID() int
//
// Having only GetName() is not enough.
//
// Therefore:
//
//     Required methods
//            ↓
//     type must provide ALL of them
//            ↓
//     interface satisfied
//
// ============================================================================
// 14. METHOD SIGNATURE MUST MATCH
// ============================================================================
//
// The method must match the interface exactly.
//
// Interface:
//
//     GetName() string
//
// This satisfies it:
//
//     GetName() string
//
// But this does NOT:
//
//     GetName() int
//
// And this does NOT:
//
//     GetName(name string) string
//
// The method name, parameters and return values must match the interface
// method signature.
//
// ============================================================================
// 15. VALUE RECEIVER AND INTERFACE SATISFACTION
// ============================================================================
//
// Your code uses VALUE RECEIVERS:
//
//     func (e Employee) GetName() string
//
//     func (e BusinessPerson) GetName() string
//
// This means BOTH:
//
//     Employee
//
// and:
//
//     *Employee
//
// can generally satisfy the Person interface because the method belongs to
// the value receiver method set.
//
// This becomes more important when using POINTER RECEIVERS.
//
// ============================================================================
// 16. POINTER RECEIVER AND INTERFACES
// ============================================================================
//
// Suppose we instead define:
//
//     func (e *Employee) GetName() string
//
// Now the method belongs to `*Employee` rather than `Employee`.
//
// Therefore:
//
//     *Employee
//
// satisfies the interface.
//
// But:
//
//     Employee
//
// does not satisfy the interface in the same way because its method set does
// not include pointer-receiver methods.
//
// This is an important connection between:
//
//     METHODS
//          +
//     RECEIVERS
//          +
//     INTERFACES
//
// ============================================================================
// 17. INTERFACES ENABLE LOOSE COUPLING
// ============================================================================
//
// `displayPerson()` is loosely coupled to concrete types.
//
// It doesn't know about:
//
//     Employee
//     BusinessPerson
//
// It only knows:
//
//     Person
//
// This means we can introduce new types without changing the function.
//
// For example:
//
//     Customer
//     Student
//     Manager
//
// As long as they implement:
//
//     GetName() string
//
// they can be passed to:
//
//     displayPerson()
//
// ============================================================================
// 18. INTERFACE AS A COMMON ABSTRACTION
// ============================================================================
//
// Suppose a system has:
//
//     Employee
//     Customer
//     Supplier
//     Manager
//
// They may all be completely different structs.
//
// But perhaps all of them have:
//
//     GetName() string
//
// Instead of writing separate logic for every type, we can define:
//
//     Person
//
// and operate on the common behavior.
//
// This gives us a common abstraction over otherwise unrelated types.
//
// ============================================================================
// 19. INTERFACES ARE VERY IMPORTANT IN BACKEND ENGINEERING
// ============================================================================
//
// Interfaces are heavily used in real Go applications.
//
// Common examples include:
//
//     Repository interfaces
//     Database interfaces
//     HTTP clients
//     Storage systems
//     Logging abstractions
//     Service layers
//     Testing/mocking
//
// For example:
//
//     type Repository interface {
//         Create()
//         Get()
//         Update()
//         Delete()
//     }
//
// Then we could have:
//
//     SQLiteRepository
//     PostgresRepository
//     MockRepository
//
// All of them can satisfy the same Repository interface.
//
// The rest of the application can depend on:
//
//     Repository
//
// rather than a specific database implementation.
//
// ============================================================================
// 20. THE CORE MENTAL MODEL
// ============================================================================
//
// Think about an interface like this:
//
//     INTERFACE
//          |
//          | defines required behavior
//          ↓
//     +----------------+
//     | GetName()      |
//     +----------------+
//          ↑
//          |
//     +----+-----------+
//     |                |
//     |                |
// Employee      BusinessPerson
//     |                |
// GetName()         GetName()
//
// Both types satisfy the same contract.
//
// Therefore:
//
//     displayPerson()
//
// can work with both.
//
// ============================================================================
// 21. FINAL TAKEAWAYS
// ============================================================================
//
// -> An interface defines behavior, not data.
//
// -> An interface contains method signatures.
//
// -> A type satisfies an interface by implementing ALL required methods.
//
// -> Go uses IMPLICIT interface implementation.
//
// -> There is no `implements` keyword in Go.
//
// -> Interfaces allow functions to depend on behavior instead of concrete
//    types.
//
// -> `displayPerson(p Person)` can accept Employee and BusinessPerson because
//    both provide GetName() string.
//
// -> Interfaces enable polymorphism.
//
// -> Interfaces reduce coupling between components.
//
// -> A value-receiver method affects interface satisfaction differently from
//    a pointer-receiver method.
//
// -> Interfaces become especially powerful when designing larger systems,
//    such as repositories, services, storage layers and APIs.
//
// ============================================================================
//                         ONE-LINE MEMORY TRICK
// ============================================================================
//
//     INTERFACE = WHAT CAN YOU DO?
//
//     CONCRETE TYPE = HOW DO YOU DO IT?
//
//     Employee       -> "I can GetName()"
//     BusinessPerson -> "I can GetName()"
//                          ↓
//                       Person
//
// ============================================================================
