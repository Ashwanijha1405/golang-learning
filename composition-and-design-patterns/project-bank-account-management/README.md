# 🏦 Bank Account System — Go

A small Go project demonstrating how to build a simple bank account system using **structs, methods, pointers, struct embedding, method promotion, method overriding, error handling, and encapsulation-like behaviour through methods**.

The project implements a base `Account` type and builds specialized account types such as `SavingsAccount` and `OverdraftAccount` on top of it.

---

## 📌 Concepts Covered

- Structs
- Methods
- Pointer receivers
- Struct embedding
- Anonymous fields
- Promoted fields
- Promoted methods
- Method overriding
- Custom `String()` method
- `fmt.Stringer`
- Error handling
- `errors.New()`
- `fmt.Errorf()`
- Encapsulation through methods
- Floating-point calculations
- Composition in Go

---

# 🧠 Project Structure

The project has three main types:

    Account
       │
       ├── SavingsAccount
       │      └── InterestRate
       │
       └── OverdraftAccount
              └── OverdraftLimit

`SavingsAccount` and `OverdraftAccount` both embed the base `Account` struct.

This allows them to reuse the fields and methods of `Account` while also providing their own specialized behaviour.

---

# 1. Base `Account` Struct

The base account is defined as:

    type Account struct {
        AccountNumber string
        Balance       float64
        OwnerName     string
    }

It contains the common information that every bank account needs:

| Field | Type | Purpose |
|---|---|---|
| `AccountNumber` | `string` | Unique account identifier |
| `Balance` | `float64` | Current account balance |
| `OwnerName` | `string` | Name of the account owner |

The idea is that specialized account types can reuse this common functionality instead of duplicating it.

---

# 2. `Deposit()` Method

The `Account` type has a `Deposit()` method:

    func (acc *Account) Deposit(amount float64) error

Notice the pointer receiver:

    *Account

This is important because depositing money modifies the account's balance.

The method performs validation:

    if amount <= 0 {
        return errors.New("deposit amount must be positive")
    }

If the amount is valid:

    acc.Balance += amount

The balance is modified directly because `acc` points to the original `Account`.

---

# 3. `Withdraw()` Method

The normal withdrawal logic is:

    func (acc *Account) Withdraw(amount float64) error

It performs two validations.

### Invalid amount

    if amount <= 0 {
        return errors.New("Withdrawal amount must be positive")
    }

### Insufficient balance

    if acc.Balance < amount {
        return fmt.Errorf(...)
    }

Only after validation does the withdrawal happen:

    acc.Balance -= amount

This prevents the normal account from going below zero.

---

# 4. Why Pointer Receivers?

Both `Deposit()` and `Withdraw()` use:

    *Account

instead of:

    Account

because these methods modify the account.

For example:

    acc.Balance += amount

We want the original account's balance to change.

With a pointer receiver, the method works with the original object.

Conceptually:

    Account
       │
       │ pointer
       ↓
    Modify original Balance

This is different from a value receiver, where the method receives a copy.

---

# 5. The `String()` Method

The `Account` type implements:

    func (acc *Account) String() string

It returns a formatted representation of the account:

    Account [SAV101] Owner: Alice Saver, Balance: $1170.00

This is useful for displaying account information consistently.

The method uses:

    fmt.Sprintf()

Remember:

    fmt.Printf()
        → formats AND prints

    fmt.Sprintf()
        → formats AND returns a string

---

# 6. `SavingsAccount`

The savings account is defined as:

    type SavingsAccount struct {
        Account
        InterestRate float64
    }

The interesting part is:

    Account

There is no field name such as:

    Account Account

Instead, `Account` is embedded as an anonymous field.

This is called **struct embedding**.

---

# 7. What Does Embedding Give Us?

Because `Account` is embedded, its fields are promoted.

For example, instead of writing:

    savAcc.Account.Balance

we can write:

    savAcc.Balance

Similarly:

    savAcc.Account.AccountNumber

can become:

    savAcc.AccountNumber

And methods can also be promoted.

For example:

    savAcc.Deposit(200)

works because `Deposit()` belongs to the embedded `Account`.

Conceptually:

    SavingsAccount
          │
          ├── Account
          │     ├── AccountNumber
          │     ├── Balance
          │     ├── OwnerName
          │     ├── Deposit()
          │     ├── Withdraw()
          │     └── String()
          │
          └── InterestRate

---

# 8. `AddInterest()`

Savings accounts have additional behaviour:

    func (sa *SavingsAccount) AddInterest()

The interest is calculated using:

    interest := sa.Balance * sa.InterestRate

Because `Balance` is promoted from the embedded `Account`, we can access it directly.

For example:

    Balance = $1000
    InterestRate = 0.02

Therefore:

    Interest = 1000 × 0.02
             = $20

Then the interest is deposited:

    sa.Deposit(interest)

`Deposit()` is a method of `Account`, but because `Account` is embedded, the method is promoted and can be called directly on `SavingsAccount`.

---

# 9. Method Promotion

This is an important consequence of embedding.

`Account` has:

    Deposit()
    Withdraw()
    String()

Because `Account` is embedded inside `SavingsAccount`, those methods become available through `SavingsAccount`.

For example:

    savAcc.Deposit(200)

is effectively accessing the embedded account's method.

Conceptually:

    savAcc.Deposit()
          │
          ↓
    embedded Account
          │
          ↓
    Account.Deposit()

The method still belongs to `Account`.

Embedding simply makes it conveniently accessible.

---

# 10. `OverdraftAccount`

The overdraft account is:

    type OverdraftAccount struct {
        Account
        OverdraftLimit float64
    }

It also embeds `Account`.

But it introduces additional behaviour:

    OverdraftLimit

This represents how far the account is allowed to go below zero.

For example:

    Balance = $100
    OverdraftLimit = $200

The maximum amount available for withdrawal is:

    $100 + $200
    = $300

---

# 11. Overriding `Withdraw()`

This is one of the most interesting parts of the project.

`Account` already has:

    Withdraw()

But `OverdraftAccount` defines its own:

    func (oa *OverdraftAccount) Withdraw(amount float64) error

This method has the same name but different behaviour.

The overdraft account allows the balance to become negative.

For example:

    Balance = $100
    OverdraftLimit = $200

A withdrawal of:

    $200

is allowed.

The resulting balance becomes:

    $100 - $200
    = -$100

This would not be allowed by the normal `Account.Withdraw()` method.

---

# 12. Overdraft Calculation

The important condition is:

    if (oa.Balance + oa.OverdraftLimit) < amount

This calculates the maximum amount currently available.

Example:

    Balance = 100
    OverdraftLimit = 200

Therefore:

    Available = 100 + 200
              = 300

A withdrawal of:

    250

is allowed.

A withdrawal of:

    350

is rejected.

---

# 13. Method Shadowing / Overriding Behaviour

Because `OverdraftAccount` has its own:

    Withdraw()

it takes precedence when calling:

    ovdAcc.Withdraw(...)

The specialized method is used instead of the promoted `Account.Withdraw()`.

Conceptually:

    OverdraftAccount
          │
          ├── Account
          │      └── Withdraw()
          │
          └── Withdraw()
                 ↑
          specialized method

The outer method shadows the promoted method.

---

# 14. Composition Through Embedding

The project demonstrates Go's composition-oriented design.

Instead of inheritance:

    SavingsAccount IS-A Account
    OverdraftAccount IS-A Account

we can think of the implementation as:

    SavingsAccount HAS-A Account
    OverdraftAccount HAS-A Account

The specialized structs are built by composing the base `Account` with additional fields and behaviour.

---

# 15. Error Handling

Go does not use traditional exception-based error handling.

Functions commonly return:

    value, error

For example:

    func (acc *Account) Deposit(amount float64) error

If everything succeeds:

    return nil

If something goes wrong:

    return errors.New("deposit amount must be positive")

The caller checks:

    err := acc.Deposit(200)

    if err != nil {
        fmt.Println("Error:", err)
    }

---

# 16. `errors.New()`

For simple errors:

    errors.New("deposit amount must be positive")

creates a new error containing that message.

It is useful when no dynamic information needs to be included.

---

# 17. `fmt.Errorf()`

When the error needs dynamic information, `fmt.Errorf()` is useful.

For example:

    return fmt.Errorf(
        "insufficient funds in %s. Balance: $%.2f, Tried to Withdraw: $%.2f",
        acc.AccountNumber,
        acc.Balance,
        amount,
    )

This allows runtime values to be included in the error message.

Conceptually:

    fmt.Errorf()
         │
         ├── static message
         └── runtime values

---

# 18. Account Lifecycle

A normal deposit looks like:

    Deposit(amount)
          ↓
    Is amount <= 0?
       /       \
     YES       NO
      ↓         ↓
    Error    Add amount
                ↓
          Update Balance
                ↓
             return nil

A normal withdrawal:

    Withdraw(amount)
          ↓
    Is amount <= 0?
       /       \
     YES       NO
      ↓         ↓
    Error    Check Balance
                ↓
        Balance < amount?
             /       \
           YES       NO
            ↓         ↓
          Error    Subtract
                       ↓
                  Update Balance
                       ↓
                    return nil

---

# 19. Savings Account Flow

    SavingsAccount
          ↓
    Existing Account
          +
    InterestRate
          ↓
    AddInterest()
          ↓
    Balance × InterestRate
          ↓
      Interest
          ↓
    Deposit(interest)
          ↓
    Updated Balance

Example:

    Balance = $1000
    InterestRate = 2%

    Interest = 1000 × 0.02
             = $20

    New Balance = $1020

---

# 20. Overdraft Account Flow

    OverdraftAccount
          ↓
    Existing Account
          +
    OverdraftLimit
          ↓
    Withdraw(amount)
          ↓
    Balance + OverdraftLimit
          ↓
    Is amount within available limit?
          /                    \
        NO                      YES
        ↓                        ↓
      Error                Subtract amount
                                ↓
                         Balance may become
                            negative

---

# 21. Important Go Concepts Demonstrated

### Struct

Groups related data together.

    type Account struct {
        AccountNumber string
        Balance       float64
        OwnerName     string
    }

### Method

Associates behaviour with a type.

    func (acc *Account) Deposit(amount float64) error

### Pointer Receiver

Allows a method to modify the original value.

    func (acc *Account) Deposit(...)

### Struct Embedding

Embeds one struct inside another without explicitly naming the field.

    type SavingsAccount struct {
        Account
        InterestRate float64
    }

### Promoted Fields

Embedded fields can be accessed directly.

    savAcc.Balance

instead of:

    savAcc.Account.Balance

### Promoted Methods

Embedded methods can also be accessed directly.

    savAcc.Deposit(200)

### Specialized Method

A type can define its own method with the same name as an embedded method.

    func (oa *OverdraftAccount) Withdraw(...)

This gives the specialized type different behaviour.

---

# 22. Project Execution

The program creates two types of accounts.

### Savings Account

    Account Number: SAV101
    Owner: Alice Saver
    Initial Balance: $1000
    Interest Rate: 2%

Operations:

    Deposit $200
        ↓
    Balance = $1200

    Add 2% Interest
        ↓
    Interest = $24
        ↓
    Balance = $1224

    Withdraw $50
        ↓
    Balance = $1174

### Overdraft Account

    Account Number: 0VD002
    Owner: Bob Spender
    Initial Balance: $100
    Overdraft Limit: $200

Operations:

    Deposit $50
        ↓
    Balance = $150

    Withdraw $200
        ↓
    Balance = -$50

The negative balance is allowed because the overdraft limit covers the withdrawal.

---

# 23. Project Architecture

    ┌───────────────────────────────┐
    │           Account             │
    ├───────────────────────────────┤
    │ AccountNumber                 │
    │ Balance                       │
    │ OwnerName                     │
    ├───────────────────────────────┤
    │ Deposit()                     │
    │ Withdraw()                    │
    │ String()                      │
    └───────────────┬───────────────┘
                    │
             ┌──────┴──────┐
             │             │
             ▼             ▼
    ┌────────────────┐ ┌──────────────────┐
    │ SavingsAccount │ │ OverdraftAccount │
    ├────────────────┤ ├──────────────────┤
    │ Account        │ │ Account          │
    │ InterestRate   │ │ OverdraftLimit   │
    ├────────────────┤ ├──────────────────┤
    │ AddInterest()  │ │ Withdraw()       │
    └────────────────┘ └──────────────────┘

The important design principle is:

    Common functionality
            ↓
        Account
            ↓
    Specialized behaviour
       ↙            ↘
    Savings      Overdraft

---

# 🎯 Key Takeaways

- `Account` contains common banking functionality.
- `SavingsAccount` and `OverdraftAccount` embed `Account`.
- Embedding provides promoted fields and methods.
- Pointer receivers are used when methods need to modify state.
- `SavingsAccount` adds interest-related behaviour.
- `OverdraftAccount` provides specialized withdrawal behaviour.
- A specialized method can shadow an embedded method.
- `errors.New()` is useful for simple errors.
- `fmt.Errorf()` is useful when errors need dynamic information.
- Go encourages composition instead of traditional inheritance.
- Struct embedding is a powerful way to reuse data and behaviour.

---

# 🚀 What This Project Teaches

This project is a practical example of how several Go features work together:

    Structs
       ↓
    Methods
       ↓
    Pointer Receivers
       ↓
    Composition
       ↓
    Struct Embedding
       ↓
    Method Promotion
       ↓
    Specialized Behaviour
       ↓
    Error Handling

The main idea is not just learning individual Go features, but understanding how they can be combined to design reusable and maintainable code.

---

## 🛠️ Technologies

- **Language:** Go
- **Standard Library:** `fmt`, `errors`

---

## ▶️ Run the Project

Make sure Go is installed, then run:

    go run main.go

To format the code:

    gofmt -w main.go

To build the project:

    go build

---

## 📚 Learning Focus

This project is part of my Go learning journey, focused on understanding the language through small practical systems rather than isolated syntax examples.

Current concepts explored include:

    Structs
    Methods
    Pointer Receivers
    Composition
    Struct Embedding
    Interfaces
    Stringer
    Generics
    Error Handling
    Defer
    Panic & Recovery
    Maps
    Slices
    Arrays
    Variadic Functions
    Multiple Return Values

---

## ⭐ Summary

This Bank Account System demonstrates how Go can model real-world entities using **structs + methods + composition**.

The core design is:

    Account
       ↓
    Shared state + behaviour
       ↓
    ┌───────────────────┐
    │                   │
    ▼                   ▼
    SavingsAccount     OverdraftAccount
    │                   │
    Interest            Overdraft
    behaviour           behaviour

Rather than relying on traditional inheritance, Go allows us to build specialized types by **composing reusable components and adding focused behaviour**.