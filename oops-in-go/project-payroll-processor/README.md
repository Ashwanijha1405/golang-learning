# 💰 Go Payroll Processor

A small payroll processing system built in Go to practice **interfaces, interface composition, methods, receivers, generics, variadic concepts, and polymorphism**.

The project models different types of employees and calculates their monthly salary using a common `Payable` interface.

---

## 📌 Project Overview

This project demonstrates how Go can model different objects that behave differently while exposing a common interface.

We have three employee types:

- 👨‍💼 `SalariedEmployee`
- 👷 `HourlyEmployee`
- 💼 `CommissionEmployee`

Although their salary calculations are different, all of them satisfy the same `Payable` interface.

This allows the payroll processor to work with all employee types through one common abstraction.

---

## 🧠 Concepts Practiced

This project combines several important Go concepts:

- Structs
- Custom types
- Methods
- Value receivers
- Interfaces
- Interface composition
- `fmt.Stringer`
- Polymorphism
- Slices of interfaces
- Multiple concrete types implementing one interface
- Generic functions
- Type constraints
- Variadic functions
- `fmt.Printf`
- `Sprintf`
- Floating-point calculations

---

# 🏗️ High-Level Architecture

The overall design looks like this:

    ┌─────────────────────────────────────────┐
    │              Payable Interface          │
    │                                         │
    │  fmt.Stringer                           │
    │  CalculatePay() float64                 │
    └───────────────────┬─────────────────────┘
                        │
          ┌─────────────┼─────────────┐
          │             │             │
          ▼             ▼             ▼
    ┌───────────┐ ┌───────────┐ ┌──────────────────┐
    │ Salaried  │ │  Hourly   │ │   Commission     │
    │ Employee  │ │ Employee  │ │    Employee      │
    └───────────┘ └───────────┘ └──────────────────┘
          │             │             │
          ▼             ▼             ▼
      Annual/12     Rate × Hours    Base + Commission
          │             │             │
          └─────────────┼─────────────┘
                        ▼
                 ProcessPayroll()
                        │
                        ▼
                Total Monthly Payroll


---

# 1. 👨‍💼 Salaried Employee

A salaried employee receives an annual salary.

The monthly salary is calculated as:

    Monthly Pay = Annual Salary / 12

Example:

    Annual Salary = $72,000

    Monthly Pay = 72,000 / 12
                = $6,000


The type contains:

    Name
    AnnualSalary


And implements:

    CalculatePay() float64

and:

    String() string


---

# 2. 👷 Hourly Employee

An hourly employee is paid based on:

    Hourly Rate × Hours Worked

Example:

    Hourly Rate  = $25
    Hours Worked = 160

    Monthly Pay = 25 × 160
                = $4,000


The type contains:

    Name
    HourlyRate
    HoursWorked


It also implements:

    CalculatePay() float64

and:

    String() string


---

# 3. 💼 Commission Employee

A commission employee has:

- A monthly base salary
- A commission rate
- A sales amount

The calculation is:

    Monthly Pay =
        Base Salary + (Commission Rate × Sales Amount)

Example:

    Base Salary     = $2,000
    Commission Rate = 10%
    Sales Amount    = $15,000

    Commission = 0.10 × 15,000
               = $1,500

    Monthly Pay = 2,000 + 1,500
                = $3,500


The commission rate is stored as:

    0.10

which represents:

    10%


---

# 4. 🔌 Payable Interface

The central abstraction of the project is:

    type Payable interface {
        fmt.Stringer
        CalculatePay() float64
    }


This means any type that wants to be considered `Payable` must provide:

    CalculatePay() float64

AND must satisfy:

    fmt.Stringer


Since `fmt.Stringer` requires:

    String() string

every `Payable` type must effectively implement:

    String() string
    CalculatePay() float64


Visual representation:

    Payable
       │
       ├── String() string
       │
       └── CalculatePay() float64


---

# 5. 🧩 Interface Composition

One interesting part of the project is:

    type Payable interface {
        fmt.Stringer
        CalculatePay() float64
    }


Here, `fmt.Stringer` is embedded inside another interface.

This is called:

    Interface Composition


Instead of writing:

    type Payable interface {
        String() string
        CalculatePay() float64
    }


we reuse the existing `fmt.Stringer` interface.

Since `fmt.Stringer` already defines:

    String() string


`Payable` automatically requires that method as well.


Visual:

    fmt.Stringer
         │
         │ String() string
         ▼
    ┌──────────────────┐
    │     Payable      │
    │                  │
    │ String() string  │
    │ CalculatePay()   │
    └──────────────────┘


This is a very common and useful Go pattern.

---

# 6. 🖨️ Why `fmt.Stringer`?

`fmt.Stringer` is an interface from the standard `fmt` package.

Conceptually:

    type Stringer interface {
        String() string
    }


If a type implements `String() string`, Go's formatting functions can use that method when displaying the value.

For example:

    fmt.Println(employee)


can use:

    employee.String()


instead of printing the default struct representation.

This allows us to control how an employee appears in logs and terminal output.


---

# 7. 🎨 Custom String Representations

### Salaried Employee

The `String()` method returns something similar to:

    Salaried: Alice Wonderland (Annual: $72000.00)


### Hourly Employee

    Hourly: Bob the Builder (Rate: $25.00/hr, Hours: 160.0)


### Commission Employee

    Commission: Charlie Chaplin (Base: $2000.00, CommRate: 10.00%, Sales: $15000.00)


This makes the output much more readable than the default struct representation.


---

# 8. 🔄 Methods and Receivers

Each employee type defines methods using a value receiver.

For example:

    func (se SalariedEmployee) CalculatePay() float64

Here:

    se SalariedEmployee

is the receiver.

It means:

    CalculatePay()

belongs to the `SalariedEmployee` type.


Similarly:

    func (he HourlyEmployee) CalculatePay() float64

belongs to:

    HourlyEmployee


And:

    func (ce CommissionEmployee) CalculatePay() float64

belongs to:

    CommissionEmployee


---

# 9. 🎯 Same Method Name, Different Behaviour

All employee types have:

    CalculatePay() float64


But the implementation is different.

    SalariedEmployee
          │
          ▼
    AnnualSalary / 12


    HourlyEmployee
          │
          ▼
    HourlyRate × HoursWorked


    CommissionEmployee
          │
          ▼
    BaseSalary + Commission


This is one of the core ideas behind interfaces and polymorphism.

---

# 10. 🧬 Polymorphism in Go

The payroll processor does NOT need to know the concrete employee type.

It only knows:

    Payable


So all of these can be stored together:

    SalariedEmployee
    HourlyEmployee
    CommissionEmployee


Inside:

    []Payable


Visual:

    []Payable
       │
       ├── SalariedEmployee
       │
       ├── HourlyEmployee
       │
       ├── CommissionEmployee
       │
       └── HourlyEmployee


The payroll processor can then call:

    emp.CalculatePay()


without knowing which concrete employee it is dealing with.


---

# 11. 🔥 Why This Is Powerful

Without interfaces, we might have to write separate payroll logic for every employee type.

Something like:

    processSalariedEmployee()
    processHourlyEmployee()
    processCommissionEmployee()


That becomes difficult to maintain as the system grows.

With the `Payable` interface:

    ProcessPayroll()

only needs to understand:

    CalculatePay()


Therefore, adding another employee type becomes much easier.

For example:

    type ContractEmployee struct {
        Name string
        MonthlyPay float64
    }


As long as it implements:

    CalculatePay() float64

and:

    String() string

it can become a `Payable` employee.

No changes to the core payroll processing logic are required.


---

# 12. 🔁 ProcessPayroll Flow

The payroll processing pipeline looks like this:

    main()
      │
      ▼
    Create employees
      │
      ▼
    Store employees in []Payable
      │
      ▼
    ProcessPayroll(payrollList)
      │
      ▼
    Iterate through employees
      │
      ▼
    PrintEmployeeSummary(emp)
      │
      ▼
    emp.CalculatePay()
      │
      ▼
    Add result to totalPayroll
      │
      ▼
    Print total monthly payroll


---

# 13. 📦 `[]Payable`

The following idea is very important:

    payrollList := []Payable{
        salEmp,
        hrEmp,
        comEmp,
        HourlyEmployee{...},
    }


This is a slice whose elements are of interface type `Payable`.

It can hold different concrete types as long as they satisfy the interface.

So:

    SalariedEmployee
    HourlyEmployee
    CommissionEmployee


can all exist inside the same slice.


This is the Go equivalent of achieving polymorphic behaviour without traditional class inheritance.


---

# 14. 🧠 Implicit Interface Implementation

Go does NOT require:

    implements Payable


There is no explicit declaration saying:

    SalariedEmployee implements Payable


Instead, Go checks whether the type has all the required methods.

For `Payable`:

    String() string
    CalculatePay() float64


`SalariedEmployee` has both.

Therefore:

    SalariedEmployee satisfies Payable


Same for:

    HourlyEmployee

and:

    CommissionEmployee


Visual:

    SalariedEmployee
          │
          ├── String()
          └── CalculatePay()
                 │
                 ▼
              Payable ✓


This is one of the defining characteristics of Go interfaces.


---

# 15. 🧬 Generic `PrintEmployeeSummary`

The project also contains:

    func PrintEmployeeSummary[P fmt.Stringer](employee P)


This is a generic function.

`P` is a type parameter.

The constraint:

    fmt.Stringer


means:

    P must implement String() string


Therefore the function can accept any type satisfying `fmt.Stringer`.

Inside the function:

    fmt.Printf(" - Processing: %s\n", employee)


the formatting relies on the type's `String()` method.


---

# 16. 🔍 Why Use a Generic Here?

The generic function demonstrates that generics can work with interfaces as constraints.

The mental model is:

    PrintEmployeeSummary[P fmt.Stringer]

means:

    "Give me some type P,
     but P must satisfy fmt.Stringer."


For example:

    SalariedEmployee
    HourlyEmployee
    CommissionEmployee


all satisfy `fmt.Stringer`.

Therefore they can be passed to the generic function.


Note:

The generic function is intentionally kept in the project as a learning example to demonstrate Go generics and interface constraints.


---

# 17. 🔗 Generics + Interfaces

This project combines two concepts:

    Generics
       +
    Interfaces


The generic constraint:

    [P fmt.Stringer]


restricts what types can be passed to:

    PrintEmployeeSummary()


Meanwhile:

    Payable


is used to model the behaviour required by the payroll system.


Visual:

    Generic Function
          │
          ▼
    P must satisfy
          │
          ▼
    fmt.Stringer
          │
          ▼
    String() string


and separately:

    Payroll System
          │
          ▼
       Payable
          │
          ├── String()
          └── CalculatePay()


---

# 18. 💰 Payroll Calculation Flow

For every employee:

    Employee
       │
       ▼
    CalculatePay()
       │
       ├───────────────┐
       │               │
       ▼               ▼
    Salary          Hourly
    Annual/12       Rate × Hours
       │               │
       └───────┬───────┘
               │
               ▼
          Monthly Pay
               │
               ▼
         totalPayroll


For commission employees:

    Base Salary
         │
         +
    Commission Rate × Sales
         │
         ▼
    Monthly Pay


---

# 19. 🧮 Example Payroll

The project creates these employees:

    Alice
    Annual Salary = $72,000

    Monthly Pay = $72,000 / 12
                = $6,000


    Bob
    Hourly Rate = $25
    Hours = 160

    Monthly Pay = $25 × 160
                = $4,000


    Charlie
    Base Salary = $2,000
    Commission = 10%
    Sales = $15,000

    Commission = 0.10 × $15,000
               = $1,500

    Monthly Pay = $2,000 + $1,500
                = $3,500


    Diana
    Hourly Rate = $30
    Hours = 150

    Monthly Pay = $30 × 150
                = $4,500


Total:

    $6,000
    + $4,000
    + $3,500
    + $4,500
    ----------------
    $18,000


Therefore:

    Total Monthly Payroll: $18000.00


---

# 20. 🧱 Project Structure

The current project is intentionally kept simple:

    payroll-processor/
    │
    ├── main.go
    └── README.md


`main.go` contains:

    Employee definitions
    Payroll interface
    Salary calculations
    String representations
    Generic helper
    Payroll processing logic
    Program entry point


This structure is appropriate for a small learning project.


---

# 21. 🚀 How to Run

Make sure Go is installed.

Check:

    go version


Run the program directly:

    go run main.go


Or initialize a Go module:

    go mod init payroll-processor


Then run:

    go run .


---

# 22. 📤 Expected Output

The output will look approximately like:

    Welcome to the Payroll Processor!

    ---- Processing Payroll -----

     - Processing: Salaried: Alice Wonderland (Annual: $72000.00)
    Monthly Pay: $6000.00

     - Processing: Hourly: Bob the Builder (Rate: $25.00/hr, Hours: 160.0)
    Monthly Pay: $4000.00

     - Processing: Commission: Charlie Chaplin (Base: $2000.00, CommRate: 10.00%, Sales: $15000.00)
    Monthly Pay: $3500.00

     - Processing: Hourly: Diana Price (Rate: $30.00/hr, Hours: 150.0)
    Monthly Pay: $4500.00

    Total Monthly Payroll: $18000.00
    ---------------- END ------------------


---

# 23. 🧠 Important Go Lessons From This Project

## Interfaces

    Interface = behaviour contract


A type satisfies an interface by implementing all required methods.


## Interface Composition

    type Payable interface {
        fmt.Stringer
        CalculatePay() float64
    }


An interface can embed another interface.


## Stringer

    String() string


Allows custom types to control their formatted string representation.


## Methods

    func (e Employee) CalculatePay() float64


Methods associate behaviour with a type.


## Receivers

    (e Employee)


The receiver determines the type the method belongs to.


## Polymorphism

Different concrete employee types can be handled through:

    Payable


## Generics

    func PrintEmployeeSummary[P fmt.Stringer](employee P)


Allows a function to work with multiple types while enforcing a compile-time constraint.


## Slices of Interfaces

    []Payable


Allows different concrete types to be stored and processed together.


---

# 24. 🎯 Core Design Principle

The most important idea in this project is:

    Program against behaviour, not concrete types.


The payroll processor doesn't care whether an employee is:

    SalariedEmployee
    HourlyEmployee
    CommissionEmployee


It only cares that the employee can:

    CalculatePay()
    String()


That's exactly what the `Payable` interface expresses.


---

# 25. 🔥 Complete Mental Model

Think of the entire project like this:

    ┌─────────────────────────────┐
    │          Employee           │
    │        Concrete Types       │
    └──────────────┬──────────────┘
                   │
          ┌────────┼─────────┐
          │        │         │
          ▼        ▼         ▼
      Salaried   Hourly   Commission
          │        │         │
          └────────┼─────────┘
                   │
                   ▼
              Payable
                   │
          ┌────────┴────────┐
          │                 │
          ▼                 ▼
     String()          CalculatePay()
          │                 │
          ▼                 ▼
      Display          Monthly Salary
          │                 │
          └────────┬────────┘
                   ▼
             ProcessPayroll()
                   │
                   ▼
          Total Monthly Payroll


---

# 📝 Key Takeaways

1. Go interfaces are implemented implicitly.

2. `Payable` combines `fmt.Stringer` with `CalculatePay()`.

3. `fmt.Stringer` allows custom types to control their printed representation.

4. Different employee types can implement the same interface differently.

5. `[]Payable` allows different employee types to be processed together.

6. `CalculatePay()` demonstrates polymorphism because the same method call produces different behaviour depending on the concrete employee type.

7. Interface composition allows existing interfaces to be reused.

8. Generic functions can use interfaces as type constraints.

9. Generics and interfaces solve different problems:
   
       Interfaces → common behaviour
       Generics   → reusable type-safe algorithms

10. The payroll processor depends on the `Payable` abstraction rather than individual employee implementations.

---

# 🛠️ Possible Future Improvements

This project can be extended into a more realistic payroll system.

Possible improvements:

- Add employee IDs
- Add departments
- Add tax calculation
- Add bonuses
- Add overtime pay
- Add deductions
- Add monthly payslips
- Add employee database persistence
- Add JSON/CSV export
- Add unit tests
- Add error handling
- Add employee search
- Add payroll history
- Add configuration
- Add REST API
- Split the project into packages
- Add PostgreSQL persistence

A natural next step would be to separate:

    Employee Models
          ↓
    Payroll Services
          ↓
    Repository Layer
          ↓
    Database


---

# 📚 Concepts Demonstrated

    Go
    ├── Structs
    ├── Methods
    ├── Value Receivers
    ├── Interfaces
    ├── Interface Composition
    ├── fmt.Stringer
    ├── Polymorphism
    ├── Slices of Interfaces
    ├── Generics
    ├── Type Constraints
    ├── Variadic Functions
    └── Formatted Output


---

# 🐹 Final Thought

This project is less about building a production payroll system and more about understanding how Go models behaviour.

The key design is:

    Concrete Types
          ↓
       Methods
          ↓
      Interfaces
          ↓
     Polymorphism
          ↓
     Payroll Processor


Go doesn't need inheritance to achieve this.

Define the behaviour you need, make types satisfy that behaviour, and let the compiler connect everything together.

That's the Go way.