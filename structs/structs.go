package main

import "fmt"

type contact struct {
	first_name string `default:"first_name"`
	last_name  string `default:"last_name"`
	address    string `default:"address""`
	city       string `default:"city"`
	state      string `default:"state"`
}

func (obj *contact) createContact() {

	if obj.first_name == "" {
		obj.first_name = "first_name"
	}
	if obj.last_name == "" {
		obj.last_name = "last_name"
	}
	if obj.address == "" {
		obj.address = "address"
	}
	if obj.city == "" {
		obj.city = "city"
	}
	if obj.state == "" {
		obj.state = "state"
	}
}

func main() {
	// Using the new keyword
	// Not assigning one value of the struct simply excludes the item from the struct. For example, if I do not assign a value to state, the struct will print out {Kevin Sherrell address city}
	newContact := new(contact)
	newContact.first_name = "Kevin"
	newContact.last_name = "Sherrell"
	newContact.address = "address"
	newContact.city = "city"
	newContact.state = "state"

	// using a struct literal
	// Adding too many arguments results in a error "too many values in struct literal"
	// Adding too few arguments results in an error "too few values in struct literal"
	var contact2 = contact{"Kevin", "Sherrell", "Address", "City", "State"}
	fmt.Println(newContact)
	fmt.Println(contact2)

	// using a struct with default values
	defaultContact := new(contact)
	fmt.Println("Contact with default")
	defaultContact.createContact()
	fmt.Print(defaultContact)

}

// Define a struct with at least four fields, using at least two different data types. Examples might include a person's contact information, a student record, or a more robust account struct than the ones shown here.

// Practice creating and initializing the struct using at least two of the methods shown here.
// What happens if you do not assign a value to every field in the struct? Does the field's data type affect the outcome?
// What happens if you use a struct literal to initialize a struct but the values appear in a different order?
// Can you find a way to use a struct literal to assign values to only some records in the struct?
