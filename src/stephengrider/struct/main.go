package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
	// contact   contactInfo
}

func main() {
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	/** EMBEDDING STRUCT */
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jim.updateName("Jimbo")
	// jim.print()
	/** WITH POJNTER DEFAULT */
	// jimPointer := &jim
	// jimPointer.updateName("Jimbo")
	/** WITH POJNTER SHORTCUT */
	jim.updateName("Jimbo")
	jim.print()

}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
