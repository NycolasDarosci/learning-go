package structs

import "fmt"

type User struct {
	Id          int
	Name        string
	PhoneNumber string
	Address     Address
}

// design pattern to create a factory
// start with 'New..'
func NewUser(name string, phoneNumber string, address Address) User {
	return User{Name: name, PhoneNumber: phoneNumber, Address: address}
}

// Structs with methods
// Structs are functions attached to a type declared outside of the structure
// method of the "class"
func (u User) PrintConstructor() string {
	return fmt.Sprintf("{id: %d, name: %v, phoneNumber: %v, address: {%v}}", u.Id, u.Name, u.PhoneNumber, u.Address.PrintConstructor())
}
