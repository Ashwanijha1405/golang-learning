package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {

	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already exists: ", name)
		return
	}

	newContact := Contact{
		ID:    nextID,
		Name:  name, 
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1
	fmt.Printf("Contact added: %v\n", name)

}

func findContact(name string) *Contact {

	index, exists := contactIndexByName[name]
	if exists {
		return &contactList[index]
	}
	return nil

}

func listContacts() {

	fmt.Println("---------------- Listing Contacts -------------------")
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	
	for i, contact := range contactList {
		fmt.Printf("%d. ID: %d, Name: %s, Email: %s, Phone: %s\n", i+1, contact.ID, contact.Name, contact.Email, contact.Phone)
	}

	fmt.Println("")

}

func main() {

	addContact("Ashwani", "Ashwani@email.com", "8178050434")
	addContact("Kanishka", "Kanishka@email.com", "8178050334")
	listContacts()

	Kanishka := findContact("Kanishka")
	if Kanishka == nil {
		fmt.Println("No contact found")
	} else {
		fmt.Println("Contact found:", Kanishka.Name)
	}

}