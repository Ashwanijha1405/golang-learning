package main

import "fmt"

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

type ContactInfo struct {
	Email string
	Phone string
}

func (ci ContactInfo) DisplayContact() string {
	return fmt.Sprintf("Email: %s, Phone: %s", ci.Email, ci.Phone)
}

type Company struct {
	Name    string
	Address // Here is the interesting part, we are just embedding other struct not using any name to it.
	ContactInfo
	BusinessType string
}

func (c Company) GetProfile() {
	fmt.Printf("Company Name: %s\n", c.Name)

	fmt.Printf("Location: %s\n", c.FullAddress())
	fmt.Printf("Street (promoted): %s\n", c.Street)

	fmt.Printf("Email (promoted): %s\n", c.Email)
	fmt.Printf("Business Type: %s\n", c.BusinessType)
}

type CompanyWithOwnEmail struct {
	Name string
	Address
	ContactInfo
	Email string // This Email field shadows the one in ContactInfo.
}

func main() {

	fmt.Println("---------------- Struct Embeddings -------------------")

	comp := Company{
		Name: "NK Securities",
		Address: Address{
			Street:  "Peepal Road",
			City:    "New Delhi",
			State:   "Delhi",
			ZipCode: "110059",
		},
		ContactInfo: ContactInfo{
			Email: "contact@email.com",
			Phone: "555-0100",
		},
		BusinessType: "Technology",
	}

	comp.GetProfile()
}


// ============================================================
// STRUCT EMBEDDING IN GO — SHORT NOTES
// ============================================================
//
// Struct embedding allows one struct to be embedded inside
// another struct WITHOUT giving the embedded field a name.
//
// Example:
//
//     type Company struct {
//         Name string
//         Address
//         ContactInfo
//     }
//
// Here Address and ContactInfo are embedded structs.
//
// ------------------------------------------------------------
//
// 1. PROMOTED FIELDS
//
// Because Address is embedded, its fields are promoted to Company.
//
// So instead of:
//
//     comp.Address.Street
//
// we can write:
//
//     comp.Street
//
// Similarly, ContactInfo's fields are promoted:
//
//     comp.Email
//     comp.Phone
//
// This makes accessing commonly used fields more convenient.
//
// ------------------------------------------------------------
//
// 2. PROMOTED METHODS
//
// Address has:
//
//     FullAddress()
//
// Since Address is embedded inside Company, its method is also
// promoted.
//
// Therefore we can call:
//
//     comp.FullAddress()
//
// instead of:
//
//     comp.Address.FullAddress()
//
// The method still belongs to Address — it is simply promoted
// through the embedding.
//
// ------------------------------------------------------------
//
// 3. EMBEDDING vs NORMAL COMPOSITION
//
// Normal composition:
//
//     type Company struct {
//         Address Address
//     }
//
// Access:
//
//     comp.Address.Street
//
// Embedding:
//
//     type Company struct {
//         Address
//     }
//
// Access:
//
//     comp.Street
//
// Embedding is therefore a convenient form of composition with
// field and method promotion.
//
// ------------------------------------------------------------
//
// 4. FIELD SHADOWING
//
// Suppose Company has its own Email:
//
//     type CompanyWithOwnEmail struct {
//         ContactInfo
//         Email string
//     }
//
// Now there are two possible Email fields:
//
//     comp.Email
//     comp.ContactInfo.Email
//
// `comp.Email` refers to Company's own Email field.
//
// To explicitly access the embedded field:
//
//     comp.ContactInfo.Email
//
// The outer field shadows the promoted embedded field.
//
// ------------------------------------------------------------
//
// 5. IMPORTANT
//
// Embedding does NOT mean inheritance.
//
// Company does NOT become an Address.
//
// It simply contains an embedded Address and gets convenient
// access to its fields and methods through promotion.
//
// Think:
//
//     Company
//        │
//        ├── Name
//        ├── Address
//        │     ├── Street
//        │     ├── City
//        │     └── FullAddress()
//        │
//        └── ContactInfo
//              ├── Email
//              ├── Phone
//              └── DisplayContact()
//
// ------------------------------------------------------------
//
// KEY TAKEAWAYS:
//
//     Embedding
//     → Include another type without naming the field.
//
//     Promoted fields
//     → Embedded fields can be accessed directly.
//
//     Promoted methods
//     → Embedded methods can be called directly.
//
//     Shadowing
//     → An outer field takes precedence over a promoted field.
//
//     Embedding ≠ inheritance
//     → It is a Go mechanism built around composition.
//
// ============================================================