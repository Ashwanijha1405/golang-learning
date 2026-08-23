package main

import "fmt" 

/*func Sum[T int | float64 | float32](numbers ...T) T {  
	var total T 
	for _, number := range numbers {
		total += number
	}
	return total
}*/

type Number interface {
	int | float32 | float64
}

func Sum[T Number](numbers ...T) T {  
	var total T 
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {

	fmt.Println(Sum(1,2,3,4,5))
	fmt.Println(Sum(30.2, 1.9))

}


// ============================================================================
//                              GENERICS IN GO
// ============================================================================
//
// Generics allow us to write functions and data structures that work with
// MULTIPLE TYPES while still maintaining compile-time type safety.
//
// Without generics, we might need to write separate functions:
//
//     SumInt()
//     SumFloat32()
//     SumFloat64()
//
// With generics, we can write ONE:
//
//     Sum[T Number](numbers ...T)
//
// and use it with different numeric types.
//
// ============================================================================
// 1. THE PROBLEM GENERICS SOLVE
// ============================================================================
//
// Suppose we want to calculate the sum of integers:
//
//     func Sum(numbers ...int) int
//
// This works for int.
//
// But what if we want:
//
//     float32
//     float64
//
// We would traditionally need separate functions.
//
// Generics allow us to say:
//
//     "This function can work with multiple types,
//      as long as those types satisfy certain requirements."
//
// ============================================================================
// 2. TYPE PARAMETERS
// ============================================================================
//
// In:
//
//     func Sum[T Number](numbers ...T) T
//
// `T` is a TYPE PARAMETER.
//
// It is similar to a placeholder for a type.
//
// Think:
//
//     T = int
//
// or:
//
//     T = float64
//
// depending on how the function is called.
//
// So:
//
//     Sum(1, 2, 3)
//
// allows Go to infer:
//
//     T = int
//
// While:
//
//     Sum(30.2, 1.9)
//
// allows Go to infer:
//
//     T = float64
//
// ============================================================================
// 3. UNDERSTANDING THE SYNTAX
// ============================================================================
//
//     func Sum[T Number](numbers ...T) T
//              ↑      ↑          ↑
//              |      |          |
//         type param constraint  return type
//
// `[T Number]` means:
//
//     T is a type parameter
//     T must satisfy the Number constraint
//
// `numbers ...T` means:
//
//     the function accepts any number of arguments of type T
//
// final `T` means:
//
//     the function returns a value of type T
//
// ============================================================================
// 4. TYPE CONSTRAINT
// ============================================================================
//
// We created:
//
//     type Number interface {
//         int | float32 | float64
//     }
//
// This is a GENERIC TYPE CONSTRAINT.
//
// It tells Go:
//
//     T can only be one of these types:
//
//         int
//         float32
//         float64
//
// Therefore:
//
//     Sum(1, 2, 3)
//
// is valid.
//
//     Sum(1.2, 3.4)
//
// is valid if the inferred type is allowed.
//
// But something like:
//
//     Sum("hello", "world")
//
// is NOT valid because string is not included in Number.
//
// ============================================================================
// 5. `|` IN A GENERIC CONSTRAINT
// ============================================================================
//
// Inside a type constraint:
//
//     int | float32 | float64
//
// means:
//
//     T may be int OR float32 OR float64.
//
// This is called a TYPE UNION.
//
// It does NOT mean that a variable can contain all three types at once.
//
// Instead, it defines the set of types that are allowed for T.
//
// ============================================================================
// 6. `Number` IS A CONSTRAINT, NOT A NORMAL DATA INTERFACE
// ============================================================================
//
// This is an important distinction.
//
// We wrote:
//
//     type Number interface {
//         int | float32 | float64
//     }
//
// This interface is being used as a GENERIC CONSTRAINT.
//
// It tells the compiler which types are allowed as type arguments.
//
// This is different from a normal runtime-style interface such as:
//
//     type Person interface {
//         GetName() string
//     }
//
// `Person` describes required METHODS.
//
// `Number` describes allowed TYPES.
//
// ============================================================================
// 7. WHY DOES `total += number` WORK?
// ============================================================================
//
// Inside:
//
//     var total T
//
// `total` has type T.
//
// `number` also has type T.
//
// Since Number restricts T to numeric types:
//
//     int
//     float32
//     float64
//
// Go knows that addition is valid for these types.
//
// Therefore:
//
//     total += number
//
// is allowed.
//
// ============================================================================
// 8. ZERO VALUE OF A GENERIC TYPE
// ============================================================================
//
// We write:
//
//     var total T
//
// Go initializes `total` with the ZERO VALUE of whatever T becomes.
//
// If:
//
//     T = int
//
// then:
//
//     total = 0
//
// If:
//
//     T = float64
//
// then:
//
//     total = 0.0
//
// This is a very useful pattern when working with generic algorithms.
//
// ============================================================================
// 9. TYPE INFERENCE
// ============================================================================
//
// We call:
//
//     Sum(1, 2, 3, 4, 5)
//
// We don't explicitly tell Go:
//
//     T = int
//
// Go looks at the arguments and infers the type.
//
// Therefore it understands:
//
//     T = int
//
// Similarly:
//
//     Sum(30.2, 1.9)
//
// allows Go to infer:
//
//     T = float64
//
// This is called TYPE INFERENCE.
//
// ============================================================================
// 10. EXPLICIT TYPE ARGUMENTS
// ============================================================================
//
// We can also explicitly provide the type parameter.
//
// Instead of relying on inference:
//
//     Sum(1, 2, 3)
//
// we could write:
//
//     Sum[int](1, 2, 3)
//
// Here:
//
//     [int]
//
// explicitly tells Go:
//
//     T = int
//
// Usually type inference makes this unnecessary when the compiler can
// determine the type from the arguments.
//
// ============================================================================
// 11. VARIADIC + GENERICS
// ============================================================================
//
// Our function contains:
//
//     numbers ...T
//
// `...` makes the function VARIADIC.
//
// That means it can accept zero or more arguments.
//
// Examples:
//
//     Sum()
//     Sum(1)
//     Sum(1, 2, 3)
//     Sum(1, 2, 3, 4, 5)
//
// The important part is that every argument must have the same inferred T.
//
// For example:
//
//     Sum(1, 2, 3)
//
// means:
//
//     numbers = []int{1, 2, 3}
//
// while:
//
//     Sum(30.2, 1.9)
//
// means the values are handled using the inferred floating-point type.
//
// ============================================================================
// 12. GENERICS VS `interface{}` / `any`
// ============================================================================
//
// Before generics, developers often used:
//
//     interface{}
//
// or:
//
//     any
//
// to accept values of arbitrary types.
//
// But this loses useful compile-time type information.
//
// For example, a function accepting `any` may need type assertions or
// type switches to determine what it received.
//
// Generics allow us to say:
//
//     "I support multiple types,
//      but ONLY these types."
//
// Therefore generics provide both:
//
//     flexibility
//     +
//     type safety
//
// ============================================================================
// 13. YOUR ORIGINAL VERSION
// ============================================================================
//
// You initially had:
//
//     func Sum[T int | float64 | float32](numbers ...T) T
//
// This is valid.
//
// It directly places the allowed types in the function's constraint.
//
// However, creating:
//
//     type Number interface {
//         int | float32 | float64
//     }
//
// gives the constraint a reusable name.
//
// Then we can write:
//
//     func Sum[T Number](numbers ...T) T
//
// This is cleaner and becomes more useful when multiple functions need
// the same constraint.
//
// ============================================================================
// 14. REUSABLE CONSTRAINTS
// ============================================================================
//
// Suppose we later create:
//
//     Average()
//     Min()
//     Max()
//     Multiply()
//
// and all of them should work with:
//
//     int
//     float32
//     float64
//
// We can reuse:
//
//     Number
//
// instead of repeatedly writing:
//
//     int | float32 | float64
//
// This is one reason named constraints are useful.
//
// ============================================================================
// 15. GENERICS ARE COMPILE-TIME
// ============================================================================
//
// Generics are primarily a compile-time feature.
//
// The compiler checks whether the supplied type satisfies the constraint.
//
// For example:
//
//     Sum(1, 2, 3)
//
// is checked against:
//
//     Number
//
// and int is allowed.
//
// But:
//
//     Sum("A", "B")
//
// fails at compile time because string does not satisfy Number.
//
// This is much safer than accepting arbitrary values and discovering
// the problem at runtime.
//
// ============================================================================
// 16. GENERIC FUNCTION MENTAL MODEL
// ============================================================================
//
// Think of:
//
//     func Sum[T Number](numbers ...T) T
//
// as:
//
//     "Give me a type T."
//
//          ↓
//
//     "T must be allowed by Number."
//
//          ↓
//
//     "Give me many values of type T."
//
//          ↓
//
//     "I'll return a value of that same T."
//
// So:
//
//     Sum(1, 2, 3)
//
// becomes conceptually:
//
//     Sum[int](1, 2, 3)
//
// and:
//
//     Sum(30.2, 1.9)
//
// becomes conceptually:
//
//     Sum[float64](30.2, 1.9)
//
// ============================================================================
// 17. GENERICS + INTERFACES
// ============================================================================
//
// This is an important connection with the interfaces you just learned.
//
// A generic constraint can use an interface to define what types are
// allowed.
//
// There are two broad ideas:
//
//     NORMAL INTERFACE
//         ↓
//     describes behavior through methods
//
//     GENERIC CONSTRAINT
//         ↓
//     restricts which types can be used as type arguments
//
// Your `Number` constraint is an example of the second concept.
//
// ============================================================================
// 18. WHEN SHOULD YOU USE GENERICS?
// ============================================================================
//
// Generics are useful when:
//
//     -> the same algorithm works for multiple types
//     -> the logic is type-independent
//     -> you want compile-time type safety
//     -> duplicating the same function for different types would be
//        unnecessary
//
// Common examples:
//
//     Sum()
//     Min()
//     Max()
//     Contains()
//     Map/filter-style utilities
//     Generic data structures
//     Reusable algorithms
//
// ============================================================================
// 19. WHEN NOT TO USE GENERICS
// ============================================================================
//
// Don't use generics simply because they exist.
//
// If a function only needs to work with one type:
//
//     func calculateSalary(salary int) int
//
// there is no reason to make it generic.
//
// Generics are useful when multiple types genuinely share the same logic.
//
// A good rule:
//
//     If the algorithm is the same,
//     but only the type changes,
//     generics may be a good fit.
//
// ============================================================================
// 20. FINAL TAKEAWAYS
// ============================================================================
//
// -> Generics allow code to work with multiple types while maintaining
//    compile-time type safety.
//
// -> `T` is a type parameter.
//
// -> `[T Number]` means T must satisfy the Number constraint.
//
// -> `Number` defines the allowed types:
//
//        int
//        float32
//        float64
//
// -> `|` creates a type union inside a generic constraint.
//
// -> `...T` means a variadic list of values of type T.
//
// -> The final `T` means the function returns the same type T.
//
// -> Go can often infer T automatically from the arguments.
//
// -> `Sum[int](...)` can explicitly specify the type parameter.
//
// -> `var total T` gets the zero value of whatever T becomes.
//
// -> Named constraints like `Number` can be reused across multiple
//    generic functions.
//
// ============================================================================
//                         ONE-LINE MEMORY TRICK
// ============================================================================
//
//     GENERICS = "Same logic, different types, still type-safe."
//
//     T       → type placeholder
//     Number  → allowed types
//     [T Number] → T must satisfy Number
//
// ============================================================================
