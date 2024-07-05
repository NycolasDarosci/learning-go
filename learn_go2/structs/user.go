package structs

import "fmt"

type User struct {
	id          int
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
// this method looks like "toString()"
func (u User) String() string {
	return fmt.Sprintf("{id: %d, name: %v, phoneNumber: %v, address: {%v}}", u.id, u.Name, u.PhoneNumber, u.Address.String())
}

//implicit implementation of the Signable interface
// SignUp() bool
// just the assignature of the method
func (c User) SignUp() bool {
	return true
}
