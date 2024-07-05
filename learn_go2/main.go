package main

import (
	"fmt"
	"os"

	"example.go/files/data"
	fu "example.go/files/fileUtils"
	u "example.go/files/structs"
)

// Alias to facilitate the use of u.User
type user = u.User

func main() {

	readWrite()
	structs()
}

func readWrite() {
	// reading a file
	dir, _ := os.Getwd() // get the rooted path name corresponding to the current directory

	filePath := dir + "/data/text.txt"
	content, err := fu.ReadTextFile(filePath)

	// error desing pattern
	if err == nil {
		data := fmt.Sprintf("Original: %v \nAdd two sentences: %v%v", content, content, content)
		err := fu.WriteToFile(filePath, data)
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}

	// type alias
	type distance float32
	var v int = 32
	a := distance(v)
	fmt.Println(a)
}

func structs() {
	// Structs..

	u2 := user{ // does not need to add the property
		Name:        "Carlos",
		PhoneNumber: "9956432345",
	}

	john := u.NewUser(
		"John",
		"5643322765",
		u.NewAddress(
			456,
			"Calendar Street",
			"St Joaquim",
			"Los Angeles",
			"California",
			"USA",
		),
	)

	car := u.NewCar(
		"Volkswagen",
		"Alexandro",
		u.NewAddress(
			456,
			"Calendar Street",
			"St Joaquim",
			"Los Angeles",
			"California",
			"USA",
		),
	)

	// method String()
	// the standard output for string
	fmt.Printf("%v\n", u2)
	fmt.Printf("%v\n", john)

	// Car does not have String()
	// this will not print the name of the car
	// but will still print the other types, the embbeded types
	// so you are embbeding those methods from the types you embbeded too
	fmt.Printf("%v\n", car)

	// both User (john) and Car (car)
	// are implicit implemented the interface Signable
	var signables [2]data.Signable
	signables[0] = john
	signables[1] = car

	for _, sign := range signables {
		fmt.Printf("sign: %v", sign)
		fmt.Printf(" is sign? %v\n", sign.SignUp())
	}
}
