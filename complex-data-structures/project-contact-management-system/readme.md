# 📇 Contact Management System

A small Go project built to understand how **structs, slices, maps, pointers, functions, initialization, and lookup patterns** work together in a practical program.

> The goal of this project isn't complexity — it's understanding how Go's core building blocks fit together.

---

## 🚀 Features

- ➕ Add new contacts
- 🚫 Prevent duplicate contacts by name
- 🆔 Automatically generate contact IDs
- 🔎 Find contacts by name
- 📋 List all contacts
- ⚡ Maintain a name → index lookup map
- 👉 Return contacts using pointers
- 🧱 Use structs to model real-world data
- ⚙️ Initialize data structures using `init()`

---

## 🧠 Concepts Practiced

| Concept | Usage |
|---|---|
| `struct` | Represents a contact |
| `slice` | Stores all contacts |
| `map` | Provides fast name-based lookup |
| `pointer` | Returns a reference to an existing contact |
| `function` | Separates different operations |
| `init()` | Initializes global data structures |
| `append()` | Adds contacts to the slice |
| `len()` | Gets number of contacts |
| `value, ok` | Checks whether a map key exists |

---

## 🧱 Data Model

Each contact is represented using a `struct`:

```go
type Contact struct {
    ID    int
    Name  string
    Email string
    Phone string
}
```

A `struct` allows related pieces of information to be grouped together.

Instead of maintaining separate variables:

```go
name  := "Ashwani"
email := "Ashwani@email.com"
phone := "8178050434"
```

we can represent one complete contact:

```go
Contact{
    ID:    1,
    Name:  "Ashwani",
    Email: "Ashwani@email.com",
    Phone: "8178050434",
}
```

This makes the data much easier to organize and work with.

---

# 🗄️ Data Storage

The project uses two main data structures.

## 1. Contact Slice

```go
var contactList []Contact
```

The slice stores the actual contact objects.

Example:

```text
contactList
    ↓
┌───────────────┐
│ Contact #1    │
├───────────────┤
│ Contact #2    │
├───────────────┤
│ Contact #3    │
└───────────────┘
```

A slice is used because the number of contacts can grow dynamically.

---

## 2. Contact Index Map

```go
var contactIndexByName map[string]int
```

This map stores:

```text
Name → Index in contactList
```

For example:

```text
"Ashwani"   → 0
"Kanishka"  → 1
```

So instead of searching the entire slice every time, we can directly find where a contact is stored.

### Why do this?

Without the map:

```text
Find "Kanishka"
        ↓
Check contact 0
        ↓
Check contact 1
        ↓
Found!
```

With the map:

```text
"Kanishka"
     ↓
map lookup
     ↓
index = 1
     ↓
contactList[1]
```

This is a simple example of using a **secondary index** to make lookups faster.

---

# ⚙️ Initialization

The project uses `init()`:

```go
func init() {
    contactList = make([]Contact, 0)
    contactIndexByName = make(map[string]int)
}
```

`init()` is automatically executed by Go before `main()`.

It is useful for preparing data structures or performing setup work before the main program starts.

---

# ➕ Adding a Contact

The function:

```go
func addContact(name, email, phone string)
```

takes three strings:

```text
name
email
phone
```

and creates a new `Contact`.

First, duplicate names are checked:

```go
if _, exists := contactIndexByName[name]; exists {
    return
}
```

This uses Go's **map lookup pattern**:

```go
value, ok := map[key]
```

In this case we don't need the value, so we use `_`:

```go
_, exists := contactIndexByName[name]
```

### Meaning

```text
exists = true
    ↓
Contact already exists

exists = false
    ↓
Contact doesn't exist
```

---

# 🆔 Automatic ID Generation

The project maintains:

```go
var nextID int = 1
```

Whenever a contact is created:

```go
newContact := Contact{
    ID:    nextID,
    Name:  name,
    Email: email,
    Phone: phone,
}

nextID++
```

So the IDs become:

```text
Contact 1 → ID 1
Contact 2 → ID 2
Contact 3 → ID 3
```

This is a simple way to generate unique sequential IDs.

---

# 📦 Adding to the Slice

Once the contact is created:

```go
contactList = append(contactList, newContact)
```

`append()` adds the contact to the slice.

For example:

```text
Before:

contactList = [Ashwani, Kanishka]


After append:

contactList = [Ashwani, Kanishka, Rahul]
```

---

# ⚡ Updating the Index

After adding the contact:

```go
contactIndexByName[name] = len(contactList) - 1
```

Why `len(contactList) - 1`?

Because slice indexes start at `0`.

For example:

```text
Index     Contact

  0       Ashwani
  1       Kanishka
  2       Rahul
```

If:

```go
len(contactList) == 3
```

the last element is:

```go
3 - 1 = 2
```

Therefore:

```go
contactIndexByName["Rahul"] = 2
```

---

# 🔎 Finding a Contact

The function:

```go
func findContact(name string) *Contact
```

returns:

```go
*Contact
```

instead of:

```go
Contact
```

The `*` means the function returns a **pointer to a Contact**.

---

## Why Return a Pointer?

Suppose:

```go
contact := findContact("Kanishka")
```

The returned pointer points to the actual contact stored inside `contactList`.

Conceptually:

```text
contactList
     ↓
┌──────────────┐
│ Ashwani      │
├──────────────┤
│ Kanishka ◄───┼──── pointer
├──────────────┤
│ Rahul        │
└──────────────┘
```

This means we are working with the existing contact rather than creating another copy.

---

# 🧩 Returning `nil`

If the contact doesn't exist:

```go
return nil
```

`nil` means:

> "There is no valid contact to point to."

Therefore the caller can check:

```go
Kanishka := findContact("Kanishka")

if Kanishka == nil {
    fmt.Println("No contact found")
} else {
    fmt.Println("Contact found:", Kanishka.Name)
}
```

The logic becomes:

```text
findContact()
      ↓
 ┌───────────────┐
 │ Contact found?│
 └───────┬───────┘
         │
     YES │ NO
      ↓   ↓
 pointer  nil
```

---

# 📋 Listing Contacts

The project loops through the slice:

```go
for i, contact := range contactList {
    fmt.Printf(
        "%d. ID: %d, Name: %s, Email: %s, Phone: %s\n",
        i+1,
        contact.ID,
        contact.Name,
        contact.Email,
        contact.Phone,
    )
}
```

`range` gives us two values:

```text
index
value
```

So:

```go
for i, contact := range contactList
```

means:

```text
i       → index
contact → actual Contact value
```

---

# 🧠 Why `i + 1`?

Slice indexes start from `0`:

```text
0 → Ashwani
1 → Kanishka
```

But when displaying a list to a user, we usually want:

```text
1. Ashwani
2. Kanishka
```

Therefore:

```go
i + 1
```

is used for display purposes.

---

# 🔑 Important Go Pattern: `value, ok`

One of the most important patterns used in this project is:

```go
value, ok := myMap[key]
```

Example:

```go
index, exists := contactIndexByName[name]
```

The two returned values mean:

```text
index  → value stored for the key
exists → whether the key exists
```

Example:

```go
index, exists := contactIndexByName["Kanishka"]
```

could give:

```text
index  = 1
exists = true
```

If the contact doesn't exist:

```text
index  = 0
exists = false
```

The important part is that `exists` tells us whether the key actually exists.

---

# 🔗 How Everything Connects

The entire system works like this:

```text
                    ADD CONTACT
                         │
                         ▼
                 Check duplicate
                         │
                         ▼
                  Create Contact
                         │
                         ▼
                  Generate ID
                         │
                         ▼
                 Append to slice
                         │
                         ▼
              Update name → index map


                    FIND CONTACT
                         │
                         ▼
                Search name in map
                         │
                         ▼
                  Get slice index
                         │
                         ▼
                contactList[index]
                         │
                         ▼
                  Return *Contact
```

---

# 🧪 Example

The program adds:

```go
addContact(
    "Ashwani",
    "Ashwani@email.com",
    "8178050434",
)

addContact(
    "Kanishka",
    "Kanishka@email.com",
    "8178050334",
)
```

The internal state becomes approximately:

```text
contactList:

[
    {
        ID: 1,
        Name: "Ashwani",
        Email: "Ashwani@email.com",
        Phone: "8178050434"
    },

    {
        ID: 2,
        Name: "Kanishka",
        Email: "Kanishka@email.com",
        Phone: "8178050334"
    }
]
```

And:

```text
contactIndexByName:

"Ashwani"  → 0
"Kanishka" → 1
```

Searching for:

```go
findContact("Kanishka")
```

becomes:

```text
"Kanishka"
     ↓
map lookup
     ↓
1
     ↓
contactList[1]
     ↓
Kanishka Contact
     ↓
*Contact
```

---

# ▶️ Running the Project

Make sure Go is installed:

```bash
go version
```

Run the program:

```bash
go run main.go
```

---

# 📚 Key Takeaways

### `struct`

Groups related data:

```go
type Contact struct {
    ID    int
    Name  string
    Email string
    Phone string
}
```

### `slice`

Stores a dynamic collection:

```go
var contactList []Contact
```

### `map`

Provides key-value storage:

```go
map[string]int
```

### `append`

Adds elements to a slice:

```go
contactList = append(contactList, newContact)
```

### `value, ok`

Checks whether a map key exists:

```go
value, ok := myMap[key]
```

### `pointer`

Stores the address/reference of a value:

```go
*Contact
```

### `nil`

Represents the absence of a valid pointer/value:

```go
return nil
```

### `init()`

Runs automatically before `main()`:

```go
func init() {
    // initialization
}
```

---

# 🎯 Learning Goal

This project is intentionally simple.

The purpose was to understand how these individual Go concepts combine into an actual program:

```text
Structs
   +
Slices
   +
Maps
   +
Pointers
   +
Functions
   +
Initialization
   +
Lookup patterns
        ↓
A working Go program
```

The next step is to evolve this into something more realistic with:

- ✏️ Update contacts
- 🗑️ Delete contacts
- 💾 File/database persistence
- 🖥️ CLI interface
- 🧪 Unit tests
- 🌐 REST API
- 🗄️ Database storage

---

## 🛠️ Built With

- **Go**
- Standard Library
- VS Code
- Linux

---

<div align="center">

### `Learn → Build → Break → Understand → Rebuild`

</div>