# 🧾 Simple Sales Order Processor

A small Go project built while learning the fundamentals of **Go programming**.

This project simulates a basic sales order system that looks up product prices, handles `_SALE` items, applies discounts, and calculates the final subtotal.

> **Learn → Build → Break → Debug → Understand**

---

## 🚀 What This Project Does

The program maintains a product-price map:

```go
var productPrices = map[string]float64{
    "TSHIRT": 20.00,
    "MUG":    12.50,
    "HAT":    18.00,
    "BOOK":   25.99,
}
```

It then processes an order:

```go
orderItems := []string{
    "TSHIRT",
    "MUG_SALE",
    "HAT",
    "BOOK",
}
```

The program:

- 🔎 Looks up products in the price map
- 🏷️ Detects `_SALE` products
- ✂️ Removes the `_SALE` suffix
- 💰 Applies a 10% discount
- ❌ Handles products that don't exist
- 🧮 Calculates the final subtotal

---

## 📌 Example

```text
TSHIRT       → $20.00
MUG_SALE     → $11.25
HAT          → $18.00
BOOK         → $25.99
```

### Final Subtotal

```text
$75.24
```

Because:

```text
20.00 + 11.25 + 18.00 + 25.99 = 75.24
```

---

# 🧠 Go Concepts Learned

This project combines several Go fundamentals into one small program.

---

## 1. 🗺️ Maps

A map stores data in **key-value pairs**.

```go
map[string]float64
```

means:

```text
key   → string
value → float64
```

Our map:

```go
var productPrices = map[string]float64{
    "TSHIRT": 20.00,
    "MUG":    12.50,
    "HAT":    18.00,
    "BOOK":   25.99,
}
```

Think of it as:

```text
Product Code → Price

TSHIRT → 20.00
MUG    → 12.50
HAT    → 18.00
BOOK   → 25.99
```

To access a value:

```go
productPrices["MUG"]
```

returns:

```text
12.50
```

---

## 2. 🔄 Functions With Multiple Return Values

Our function:

```go
func calculateItemPrice(itemCode string) (float64, bool)
```

takes:

```text
string
```

and returns:

```text
float64
bool
```

So we can do:

```go
price, found := calculateItemPrice("MUG")
```

Where:

```text
price → item's price
found → whether the item was successfully processed
```

Go allows functions to return multiple values, and this pattern is extremely common in Go.

---

## 3. 🔍 Map Lookup — `value, ok` Pattern

One of the most important Go patterns in this project:

```go
basePrice, found := productPrices[itemCode]
```

Go gives us **two values**:

```text
basePrice → value stored in the map
found     → whether the key exists
```

For:

```go
productPrices["MUG"]
```

we get:

```text
basePrice = 12.50
found     = true
```

But:

```go
productPrices["PHONE"]
```

gives:

```text
basePrice = 0
found     = false
```

### ⚠️ Important

`found` does **NOT** mean:

> "Is the price correct?"

It means:

> "Does this key exist in the map?"

---

## 4. ❓ Why Do We Need `found`?

Suppose we write:

```go
price := productPrices["PHONE"]
```

If `"PHONE"` doesn't exist, Go gives the zero value for `float64`:

```text
0
```

But now we don't know whether:

```text
PHONE exists and costs $0
```

or:

```text
PHONE doesn't exist
```

That's why Go provides:

```go
value, ok := map[key]
```

Now we can distinguish between the two cases.

---

## 5. 🔤 `strings.HasSuffix()`

We use:

```go
strings.HasSuffix(itemCode, "_SALE")
```

to check whether a string **ends with** `_SALE`.

Example:

```go
strings.HasSuffix("MUG_SALE", "_SALE")
```

returns:

```text
true
```

While:

```go
strings.HasSuffix("MUG", "_SALE")
```

returns:

```text
false
```

The `strings` package must be imported:

```go
import "strings"
```

---

## 6. ✂️ `strings.TrimSuffix()`

Once we know an item is a sale item:

```go
originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
```

For:

```text
MUG_SALE
```

we get:

```text
MUG
```

Now we can search for the original product:

```go
productPrices["MUG"]
```

---

## 7. 💸 Sale Price Calculation

Sale items receive a 10% discount.

```go
salesPrice := basePrice * 0.90
```

Why `0.90`?

Because after a 10% discount, the customer pays **90%** of the original price.

Example:

```text
Original price = $12.50

10% discount
      ↓
Customer pays 90%
      ↓
12.50 × 0.90
      ↓
$11.25
```

---

## 8. 🚪 `return` Immediately Ends the Function

Inside the sale logic:

```go
if found {
    salesPrice := basePrice * 0.90

    return salesPrice, true
}
```

Once Go executes:

```go
return salesPrice, true
```

the function **ends immediately**.

It does NOT continue executing the remaining code.

So:

```go
return salesPrice, true
```

means:

> "The sale item was successfully processed. Here's its price."

---

## 9. 🆚 Normal Item vs Sale Item

### Normal Item

Example:

```text
TSHIRT
```

Flow:

```text
TSHIRT
   ↓
Search map
   ↓
Found
   ↓
$20.00
   ↓
return 20.00, true
```

### Sale Item

Example:

```text
MUG_SALE
```

Flow:

```text
MUG_SALE
    ↓
Search map
    ↓
Not found
    ↓
Ends with "_SALE"
    ↓
Remove "_SALE"
    ↓
MUG
    ↓
Search map
    ↓
Found → $12.50
    ↓
Apply 10% discount
    ↓
$12.50 × 0.90
    ↓
$11.25
    ↓
return 11.25, true
```

---

## 10. ❌ Invalid Item

Suppose we process:

```go
calculateItemPrice("PHONE_SALE")
```

First:

```text
PHONE_SALE
    ↓
Not found in map
```

Then we check:

```text
Does it end with "_SALE"?
```

Yes.

Remove the suffix:

```text
PHONE_SALE
     ↓
PHONE
```

Search:

```go
productPrices["PHONE"]
```

Still not found.

Therefore:

```go
return 0.0, false
```

This tells the caller:

```text
Price   = 0
Success = false
```

---

## 11. 🔁 `for _, item := range`

The order is stored in a slice:

```go
orderItems := []string{
    "TSHIRT",
    "MUG_SALE",
    "HAT",
    "BOOK",
}
```

We process every item using:

```go
for _, item := range orderItems {
    ...
}
```

`range` normally gives us **two values**:

```text
index
value
```

For example:

```go
for index, item := range orderItems
```

But we don't need the index.

So we use:

```go
_
```

The `_` is called the **blank identifier**.

It means:

> "I don't need this value."

Therefore:

```go
for _, item := range orderItems
```

means:

> "Give me every item, but ignore its index."

---

## 12. 🧮 Subtotal

We declare:

```go
var subTotal float64
```

Go automatically gives it the zero value:

```text
0.0
```

Every successfully processed item's price is added:

```go
subTotal += price
```

This is shorthand for:

```go
subTotal = subTotal + price
```

---

## 13. 💵 `%.2f`

We use:

```go
fmt.Printf("Subtotal Price: %.2f\n", subTotal)
```

`%.2f` means:

> Display a floating-point number with exactly 2 digits after the decimal point.

Example:

```text
75.2
```

becomes:

```text
75.20
```

This is useful when displaying prices.

---

# 🔄 Complete Program Flow

```text
                  ORDER ITEMS
                       │
                       ▼
             calculateItemPrice()
                       │
                       ▼
                Search the map
                       │
              ┌────────┴────────┐
              │                 │
            FOUND           NOT FOUND
              │                 │
              ▼                 ▼
        Return price       Check "_SALE"
          + true                │
                               ▼
                        Remove "_SALE"
                               │
                               ▼
                       Find original item
                               │
                     ┌─────────┴─────────┐
                     │                   │
                   FOUND             NOT FOUND
                     │                   │
                     ▼                   ▼
               Apply 10% discount    Return false
                     │
                     ▼
              Return sale price
                     │
                     ▼
              Add to subtotal
```

---

# 🧩 Key Takeaways

| Go Concept | What It Does |
|---|---|
| `map[string]float64` | String keys with float64 values |
| `value, ok := map[key]` | Gets value + checks if key exists |
| `func (...) (float64, bool)` | Function returning multiple values |
| `strings.HasSuffix()` | Checks whether a string ends with something |
| `strings.TrimSuffix()` | Removes a suffix from a string |
| `for _, item := range` | Loops through values while ignoring index |
| `_` | Blank identifier; intentionally ignores a value |
| `+=` | Adds and assigns |
| `%.2f` | Formats float to 2 decimal places |
| `return value, true` | Successful return |
| `return 0.0, false` | Failed/unsuccessful return |

---

# 🧪 Example Output

```text
--------------------- Simple Sales Order Processor -----------------------------
--------------- Processing Order Items:
 - Item MUG (Sale! Original: $12.50, Sale Price: $11.25)
Subtotal Price: 75.24
```

---

# 📂 Project Structure

```text
sales-order-processor/
│
├── main.go
└── README.md
```

---

# 🎯 Why I Built This

This isn't meant to be a complex production application.

The goal was to take individual Go concepts and combine them into something that behaves like a small real-world program.

Instead of learning:

```text
Maps
Functions
Strings
Loops
```

as completely isolated topics, this project connects them together.

### Learning Philosophy

```text
Learn a concept
      ↓
Write code
      ↓
Break the code
      ↓
Debug it
      ↓
Understand WHY it works
      ↓
Build something with it
```

> **Don't just learn the syntax. Understand what the syntax is doing.**

---

# 🛠️ Future Improvements

As I learn more Go, I can extend this project with:

- [ ] Add quantities to order items
- [ ] Add tax calculation
- [ ] Add customer information
- [ ] Add order IDs
- [ ] Move product data into a database
- [ ] Split code into multiple packages
- [ ] Add unit tests
- [ ] Add proper error handling
- [ ] Convert it into a REST API

---

# 📚 Part of My Go Learning Journey

This project is part of my **Go learning repository**, where I'm documenting my progress through small programs and progressively larger projects.

The goal isn't just to finish projects.

The goal is to understand **why the code works**.

> **Build it → Break it → Understand it → Ship it.**