package structs

import "fmt"

type Address struct {
	Id           int
	Number       int32
	Street       string
	Neighborhood string
	City         string
	State        string
	Country      string
}

// design pattern to create a factory
// start with 'New..
func NewAddress(
	number int32,
	street string,
	neighborhood string,
	city string,
	state string,
	country string,
) Address {
	return Address{
		Number:       number,
		Street:       street,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
		Country:      country,
	}
}

func (a Address) PrintConstructor() string {
	return fmt.Sprintf("id: %d, number: %v, street: %v, neighborhood: %v, city: %v, state: %v, country: %v",
		a.Id,
		a.Number,
		a.Street,
		a.Neighborhood,
		a.City,
		a.State,
		a.Country,
	)
}
