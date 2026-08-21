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