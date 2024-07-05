package main

import (
	"fmt"
	"os"

	fu "example.go/files/fileUtils"
	u "example.go/files/structs"
)

// Alias to facilitate the use of u.User
type user = u.User

func main() {
	// reading a file
	dir, _ := os.Getwd() // get the rooted path name corresponding to the current directory

	filePath := dir + "/data/text.txt"
	content, err := fu.ReadTextFile(filePath)

	// error desing pattern
	if err == nil {
		fmt.Println(content)

		//data := fmt.Sprintf("Original: %v \nAdd two sentences: %v%v", content, content, content)
		//err := fu.WriteToFile(filePath, data)
		//if err != nil {
		//	panic(err)
		//}
	} else {
		panic(err)
	}

	type distance float32
	var v int = 32

	a := distance(v)

	fmt.Println(a)

	u2 := user{ // does not need to add the property
		Id:          1,
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

	fmt.Println(u.User.PrintConstructor(u2))
	fmt.Println(john.PrintConstructor())
}
