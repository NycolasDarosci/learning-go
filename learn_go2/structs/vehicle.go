package structs

import "time"

// embedding
// all properties from User are now part of the Car
type Car struct {
	User // embedding all the properties of User
	Name string
	Time time.Time
}

//implicit implementation of the Signable interface
// SignUp() bool
// just the assignature of the method
func (c Car) SignUp() bool {
	return true
}

// factory
// when you have similar embbeded properties name
// work with embbeded properties
func NewCar(name string, nameUser string, address Address) Car {
	c := Car{}
	c.Name = name
	c.User.Name = nameUser // working with embbeding properties
	c.Address = address
	return c
}
