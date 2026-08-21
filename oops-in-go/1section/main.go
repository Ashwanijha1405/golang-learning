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

	fmt.Printf("%+v\n", ashwani)

	kanishka := NewEmployee(2, "Kanishka", "Bhardwaj", "Senior Staff", true)

	fmt.Println(kanishka.ID)
	fmt.Println(kanishka.FirstName)
	fmt.Println(kanishka.LastName)
	fmt.Println(kanishka.IsActive)
	fmt.Println(kanishka.Position)

	kanishka.Salary = 10000000
	fmt.Println(kanishka.Salary)

	//accessing values by pointer
	kanishkaPtr := &kanishka
	fmt.Println(*(kanishkaPtr))
	fmt.Println(kanishkaPtr) // even without using `*` go does the derefencing by its own. 
	fmt.Println(kanishkaPtr.FirstName)
	fmt.Println(kanishkaPtr.Salary)

}





// ============================================================================
// CUSTOM TYPES WITH STRUCTS
// ============================================================================
//
// A struct is a custom data type used to group related fields together.
//
// Example:
// An employee has multiple pieces of information:
// ID, FirstName, LastName, Position, Salary, IsActive, JoinedAt.
//
// Instead of keeping these values separately, we can create ONE custom type:
// Employee.
//
// This makes the code easier to organize, understand and maintain.
//
// ============================================================================
// 1. DEFINING A CUSTOM STRUCT TYPE
// ============================================================================
//
// type Employee struct { ... } creates a new custom type named Employee.
//
// Each field inside the struct has:
// FieldName + DataType
//
// Example:
//
// ID int
// FirstName string
// Salary int
// IsActive bool
//
// A struct can contain fields of different data types.
//
// ============================================================================
// 2. CREATING A STRUCT VALUE
// ============================================================================
//
// We can create an Employee value using a struct literal:
//
// Employee{
// ID: 1,
// FirstName: "Ashwani",
// LastName: "Jha",
// }
//
// Named fields are recommended because they make the code clear and
// prevent mistakes when a struct has many fields.
//
// Fields that are not explicitly provided receive their zero value.
//
// Examples of zero values:
// int -> 0
// string -> ""
// bool -> false
// pointer -> nil
//
// ============================================================================
// 3. ACCESSING STRUCT FIELDS
// ============================================================================
//
// Use the . operator to access a field:
//
// employee.FirstName
// employee.Salary
// employee.IsActive
//
// We can also modify fields:
//
// employee.Salary = 10000000
//
// Struct values are mutable, so changing a field changes that particular
// struct value.
//
// ============================================================================
// 4. STRUCTS ARE VALUES
// ============================================================================
//
// In Go, a struct variable normally contains the actual struct value.
//
// For example:
//
// kanishka := Employee{...}
//
// kanishka itself is an Employee value.
//
// If we assign it to another variable:
//
// employee := kanishka
//
// employee receives a COPY of the struct.
//
// Therefore:
//
// employee.Salary = 500
//
// does NOT change:
//
// kanishka.Salary
//
// because they are separate struct values.
//
// ============================================================================
// 5. CONSTRUCTOR-LIKE FUNCTIONS
// ============================================================================
//
// Go does NOT have constructors like Java or C++.
//
// Instead, Go commonly uses a normal function to create and initialize
// a struct.
//
// Example:
//
// func NewEmployee(...) Employee
//
// The New prefix is a common Go naming convention for functions that
// create and return initialized values.
//
// It is NOT a special keyword.
//
// NewEmployee() can provide default values automatically.
//
// For example, JoinedAt can be initialized using:
//
// time.Now()
//
// This avoids repeatedly writing the same initialization logic.
//
// ============================================================================
// 6. WHY USE A CREATION FUNCTION?
// ============================================================================
//
// Without a creation function, every caller may have to manually initialize
// every field:
//
// Employee{
// ID: 2,
// FirstName: "Kanishka",
// LastName: "Bhardwaj",
// Position: "Senior Staff",
// IsActive: true,
// JoinedAt: time.Now(),
// }
//
// With NewEmployee():
//
// NewEmployee(2, "Kanishka", "Bhardwaj", "Senior Staff", true)
//
// The function centralizes the initialization logic.
//
// This becomes especially useful when a struct has many fields or when
// initialization requires validation/default values/calculations.
//
// ============================================================================
// 7. POINTERS TO STRUCTS
// ============================================================================
//
// We can store the address of a struct using &:
//
// kanishkaPtr := &kanishka
//
// kanishkaPtr is now a pointer to an Employee.
//
// In other words:
//
// kanishkaPtr
// |
// v
// +------------------+
// | Employee |
// | ID: 2 |
// | FirstName: ... |
// +------------------+
//
// The pointer itself stores the memory address of kanishka.
//
// ============================================================================
// 8. DEREFERENCING A STRUCT POINTER
// ============================================================================
//
// * can be used to dereference a pointer:
//
// (kanishkaPtr)
//
// This means:
//
// "Give me the actual Employee value stored at this address."
//
// Therefore:
//
// fmt.Println((kanishkaPtr))
//
// prints the Employee value itself.
//
// Whereas:
//
// fmt.Println(kanishkaPtr)
//
// prints the pointer/address representation.
//
// ============================================================================
// 9. GO AUTOMATICALLY DEREFERENCES STRUCT POINTERS FOR FIELD ACCESS
// ============================================================================
//
// This is an important Go convenience.
//
// If kanishkaPtr is a pointer to Employee, we can write:
//
// kanishkaPtr.FirstName
//
// instead of manually writing:
//
// (*kanishkaPtr).FirstName
//
// Go automatically performs the dereference when accessing a struct field.
//
// So these are effectively equivalent:
//
// kanishkaPtr.FirstName
//
// (*kanishkaPtr).FirstName
//
// This automatic dereferencing applies to STRUCT FIELD ACCESS.
//
// It does NOT mean Go automatically dereferences pointers everywhere.
//
// ============================================================================
// 10. POINTERS ALLOW US TO MODIFY THE ORIGINAL STRUCT
// ============================================================================
//
// Suppose:
//
// employeePtr := &employee
//
// Then:
//
// employeePtr.Salary = 500000
//
// modifies the original employee because the pointer refers to the
// same struct in memory.
//
// Again, Go automatically dereferences the pointer for field access.
//
// Therefore:
//
// employeePtr.Salary = 500000
//
// is effectively:
//
// (*employeePtr).Salary = 500000
//
// ============================================================================
// 11. STRUCT POINTERS AND FUNCTIONS
// ============================================================================
//
// Struct pointers become especially important when passing structs to
// functions.
//
// Passing a struct value:
//
// func updateEmployee(e Employee)
//
// gives the function a COPY of the Employee.
//
// Changes made to e won't affect the original struct.
//
// Passing a pointer:
//
// func updateEmployee(e *Employee)
//
// gives the function access to the original Employee.
//
// Therefore, changes made through e can modify the original value.
//
// ============================================================================
// 12. VALUE vs POINTER
// ============================================================================
//
// VALUE:
//
// employee := Employee{...}
//
// employee contains the actual Employee value.
//
// POINTER:
//
// employeePtr := &employee
//
// employeePtr contains the address of the Employee.
//
// Think:
//
// Employee value -> actual data
// *Employee -> pointer to that data
//
// ============================================================================
// 13. STRUCT + POINTER MENTAL MODEL
// ============================================================================
//
// If:
//
// employee := Employee{...}
//
// then:
//
// employee
// |
// v
// [ Employee data ]
//
// If:
//
// employeePtr := &employee
//
// then:
//
// employeePtr
// |
// | address
// v
// [ Employee data ]
//
// *employeePtr gives the actual Employee value.
//
// ============================================================================
// 14. IMPORTANT TAKEAWAYS
// ============================================================================
//
// -> struct creates a custom type containing related fields.
//
// -> Structs can contain fields of different data types.
//
// -> Use . to access or modify struct fields.
//
// -> Structs are values and are normally copied when assigned/passed by value.
//
// -> Go does not have traditional constructors.
//
// -> NewEmployee() is simply a normal function following the common
// Go New... naming convention.
//
// -> &employee gives the address of a struct.
//
// -> *employeePtr dereferences the pointer and gives the actual struct.
//
// -> Go automatically dereferences a struct pointer when accessing its fields:
//
// employeePtr.FirstName
//
// is equivalent to:
//
// (*employeePtr).FirstName
//
// -> Pointers are useful when we want to work with or modify the original
// struct instead of a copy.
//
// ============================================================================
// CORE MENTAL MODEL
// ============================================================================
//
// struct
// ↓
// groups related data
// ↓
// Employee
// ↓
// Employee value
// ↓
// &Employee
// ↓
// pointer to Employee
// ↓
// *pointer
// ↓
// actual Employee value
//
// ============================================================================ 
