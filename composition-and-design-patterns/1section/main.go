package main

import (
	"fmt"
)

// Composition --> Has-A relationship
// A Customer HAS-A BillingAddress and HAS-A ShippingAddress.
// A Customer is composed of Address objects.

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" {
		return "No address provided"
	}

	return fmt.Sprintf("%s, %s, %s %s",
		a.Street,
		a.City,
		a.State,
		a.ZipCode,
	)
}

type Customer struct {
	CustomerID      int
	Name            string
	Email           string
	BillingAddress  Address
	ShippingAddress Address
}

func (c Customer) PrintDetails() {
	fmt.Printf("Customer ID: %d\n", c.CustomerID)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Email: %s\n", c.Email)
	fmt.Printf("Billing Address: %s\n", c.BillingAddress.FullAddress())
	fmt.Printf("Shipping Address: %s\n", c.ShippingAddress.FullAddress())
}

func main() {

	fmt.Printf("-------------- Composition ---------------\n")

	cust1 := Customer{
		CustomerID: 1001,
		Name:       "Gadget Corp",
		Email:      "sales@gadgetcorp.com",

		BillingAddress: Address{
			Street:  "MG Road",
			City:    "Gurgaon",
			State:   "Haryana",
			ZipCode: "110063",
		},

		ShippingAddress: Address{
			Street:  "Peepal Road",
			City:    "New Delhi",
			State:   "Delhi",
			ZipCode: "110059",
		},
	}

	cust1.PrintDetails()
}

// ============================================================
// COMPOSITION IN GO — NOTES
// ============================================================
//
// Composition means building a larger type by including other
// types inside it.
//
// The key idea:
//
//     Composition = HAS-A relationship
//
// Example:
//
//     Customer HAS-A BillingAddress
//     Customer HAS-A ShippingAddress
//
// This is different from inheritance:
//
//     Customer IS-A Address   ❌
//
// We are not saying Customer is an Address.
// We are saying Customer contains Address objects.
//
// ------------------------------------------------------------
//
// 1. ADDRESS TYPE
//
// type Address struct {
//     Street  string
//     City    string
//     State   string
//     ZipCode string
// }
//
// Address represents one complete address.
//
// It has its own data:
//
//     Street
//     City
//     State
//     ZipCode
//
// It can also have its own behaviour, such as:
//
//     FullAddress()
//
// ------------------------------------------------------------
//
// 2. CUSTOMER COMPOSES ADDRESS
//
// type Customer struct {
//     CustomerID      int
//     Name            string
//     Email           string
//     BillingAddress  Address
//     ShippingAddress Address
// }
//
// Here Customer contains TWO Address values.
//
// Therefore:
//
//     Customer HAS-A BillingAddress
//     Customer HAS-A ShippingAddress
//
// Visually:
//
//     Customer
//        │
//        ├── CustomerID
//        ├── Name
//        ├── Email
//        │
//        ├── BillingAddress
//        │        └── Address
//        │             ├── Street
//        │             ├── City
//        │             ├── State
//        │             └── ZipCode
//        │
//        └── ShippingAddress
//                 └── Address
//                      ├── Street
//                      ├── City
//                      ├── State
//                      └── ZipCode
//
// ------------------------------------------------------------
//
// 3. COMPOSITION REUSES EXISTING TYPES
//
// Instead of putting address fields directly inside Customer:
//
//     type Customer struct {
//         Name          string
//         BillingStreet string
//         BillingCity   string
//         ...
//     }
//
// We create a reusable Address type:
//
//     type Address struct {
//         Street
//         City
//         State
//         ZipCode
//     }
//
// Then Customer can reuse it:
//
//     BillingAddress  Address
//     ShippingAddress Address
//
// This avoids duplicating the same group of fields.
//
// ------------------------------------------------------------
//
// 4. COMPOSITION ALSO REUSES BEHAVIOUR
//
// Address has its own method:
//
//     func (a Address) FullAddress() string
//
// Because Customer contains Address values, Customer can access
// the Address behaviour through those values.
//
// Example:
//
//     c.BillingAddress.FullAddress()
//
// This means:
//
//     Customer
//        ↓
//     BillingAddress
//        ↓
//     FullAddress()
//
// Similarly:
//
//     c.ShippingAddress.FullAddress()
//
// ------------------------------------------------------------
//
// 5. METHOD CALL THROUGH COMPOSITION
//
// When we write:
//
//     c.BillingAddress.FullAddress()
//
// Go does NOT magically make FullAddress() a Customer method.
//
// FullAddress() still belongs to Address.
//
// We are simply accessing the Address stored inside Customer
// and calling its method.
//
// Think:
//
//     Customer
//         │
//         └── BillingAddress → Address
//                                │
//                                └── FullAddress()
//
// ------------------------------------------------------------
//
// 6. VALUE RECEIVER
//
// The Address method:
//
//     func (a Address) FullAddress() string
//
// uses a value receiver.
//
// `a` is a copy of the Address value on which the method is called.
//
// The method only needs to READ the address and create a string,
// so a value receiver is perfectly appropriate here.
//
// ------------------------------------------------------------
//
// 7. EMPTY ADDRESS CHECK
//
// The method checks:
//
//     if a.Street == "" && a.City == ""
//
// `&&` means logical AND.
//
// Therefore both conditions must be true:
//
//     Street is empty
//            AND
//     City is empty
//
// Only then do we return:
//
//     "No address provided"
//
// IMPORTANT:
//
//     &   → bitwise AND
//     &&  → logical AND
//
// ------------------------------------------------------------
//
// 8. `fmt.Sprintf()`
//
// The method returns:
//
//     fmt.Sprintf("%s, %s, %s %s",
//         a.Street,
//         a.City,
//         a.State,
//         a.ZipCode,
//     )
//
// `Sprintf` formats the values and RETURNS the resulting string.
//
// Difference:
//
//     fmt.Printf()
//         → formats and prints
//
//     fmt.Sprintf()
//         → formats and returns a string
//
// Example:
//
//     fmt.Sprintf("%s, %s", "Delhi", "India")
//
// produces a string such as:
//
//     Delhi, India
//
// ------------------------------------------------------------
//
// 9. CUSTOMER PRINT DETAILS
//
// Customer has its own method:
//
//     func (c Customer) PrintDetails()
//
// This method is responsible for displaying customer information.
//
// It accesses its own fields:
//
//     c.CustomerID
//     c.Name
//     c.Email
//
// And it accesses behaviour from its composed Address values:
//
//     c.BillingAddress.FullAddress()
//     c.ShippingAddress.FullAddress()
//
// ------------------------------------------------------------
//
// 10. DATA FLOW
//
// When we create:
//
//     cust1 := Customer{
//         ...
//         BillingAddress: Address{
//             ...
//         },
//         ShippingAddress: Address{
//             ...
//         },
//     }
//
// we are constructing one Customer containing TWO Address
// objects.
//
// Visual:
//
//     cust1
//       │
//       ├── Customer information
//       │
//       ├── BillingAddress
//       │       └── Address object
//       │
//       └── ShippingAddress
//               └── Address object
//
// ------------------------------------------------------------
//
// 11. WHY COMPOSITION IS USEFUL
//
// Imagine a bigger application:
//
//     Customer
//     Employee
//     Supplier
//     Company
//
// Multiple types may need an Address.
//
// Instead of repeating:
//
//     Street
//     City
//     State
//     ZipCode
//
// everywhere, we define Address once.
//
// Then:
//
//     Customer HAS-A Address
//     Employee HAS-A Address
//     Supplier HAS-A Address
//
// This improves:
//
//     - Code reuse
//     - Maintainability
//     - Organization
//     - Separation of responsibility
//
// ------------------------------------------------------------
//
// 12. COMPOSITION vs INHERITANCE
//
// Traditional OOP often uses inheritance:
//
//     Car IS-A Vehicle
//
// Go does not have traditional class inheritance.
//
// Instead, Go heavily encourages composition.
//
// Example:
//
//     Car HAS-A Engine
//     Car HAS-A Transmission
//     Customer HAS-A Address
//     Order HAS-A Customer
//
// This allows types to be assembled from smaller reusable
// components.
//
// ------------------------------------------------------------
//
// 13. IMPORTANT MENTAL MODEL
//
// Don't think:
//
//     Customer "inherits" Address.
//
// Think:
//
//     Customer contains Address.
//
// Therefore:
//
//     Customer
//        │
//        └── Address
//
// This is the HAS-A relationship.
//
// ------------------------------------------------------------
//
// 14. EMBEDDING vs COMPOSITION
//
// Normal composition:
//
//     type Customer struct {
//         BillingAddress Address
//     }
//
// Access:
//
//     customer.BillingAddress.City
//
// The Address field has an explicit name.
//
// Go also supports embedding:
//
//     type Customer struct {
//         Address
//     }
//
// With embedding, fields and methods can be promoted so that
// they can sometimes be accessed directly through Customer.
//
// Example:
//
//     customer.City
//     customer.FullAddress()
//
// This is a different Go feature called embedding.
//
// The current project uses NORMAL COMPOSITION because Address
// is explicitly stored as:
//
//     BillingAddress
//     ShippingAddress
//
// ------------------------------------------------------------
//
// 15. BIG PICTURE
//
//     Address
//        │
//        ├── Data
//        │    ├── Street
//        │    ├── City
//        │    ├── State
//        │    └── ZipCode
//        │
//        └── Behaviour
//             └── FullAddress()
//
//                ↓
//
//     Customer
//        │
//        ├── Customer information
//        │
//        ├── BillingAddress
//        │       └── Address
//        │
//        └── ShippingAddress
//                └── Address
//
//                ↓
//
//       PrintDetails()
//
//                ↓
//
//     BillingAddress.FullAddress()
//     ShippingAddress.FullAddress()
//
// ------------------------------------------------------------
//
// KEY TAKEAWAYS:
//
//     Composition
//     → Building a type using other types.
//
//     HAS-A relationship
//     → One type contains another type.
//
//     Address
//     → Reusable type representing address information.
//
//     Customer
//     → Composes two Address values.
//
//     c.BillingAddress.FullAddress()
//     → Access the composed Address and call its method.
//
//     fmt.Sprintf()
//     → Formats values and returns a string.
//
//     &&
//     → Logical AND.
//
//     &
//     → Bitwise AND.
//
//     Value receiver
//     → Method receives a copy of the value.
//
//     Composition ≠ inheritance
//     → Go generally favours combining types rather than
//       traditional class inheritance.
//
// ------------------------------------------------------------
//
// FINAL IDEA:
//
//     Small reusable types
//              ↓
//        Compose together
//              ↓
//       Build larger types
//              ↓
//       Reuse data + behaviour
//
// This is one of the fundamental design ideas in Go:
//
//     "Build complex things by composing simple things."
//
// ============================================================