package data

// just alias
// a way to facilitate the use of the types
// to become more intuitive the use of the types
type integer = int
type json = map[string]string

// NEW TYPE
// semantics
type json2 map[string]string

type distance float64
type distanceKm float64

// Method - special argument, ATTRIBUTE ARGUMENT
// miles it's a special argument, a ATTRIBUTE ARGUMENT
// Method is a function with a type to inject something
func (miles distance) ToKm() distanceKm {
	return distanceKm(1.60934 * miles)
}

func (km distanceKm) ToMiles() distance {
	return distance(km / 1.60934)
}

func test() {
	d := distance(32.4)
	km := d.ToKm()
	miles := km.ToMiles()

	print(d)
	print(km)
	print(miles)

}
