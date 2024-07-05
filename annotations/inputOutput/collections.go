package inputOutput

import (
	"fmt"
	"time"
)

var Countries [10]string
var slice [3]int

// map[type key]type value
var ExempleMap map[int]bool

// func init() is a special function
// it is executed typically when app is started
// it can initialize things, just it
func init() {
	time.Sleep(time.Second)
	Countries[0] = "Australia"
	Countries[2] = "Brazil"

	fmt.Println(Countries)
	fmt.Println("Length Countries: ", len(Countries))
}

func Slices() {
	slice[0] = 1
	slice[1] = 2
	slice[2] = 5

	fmt.Println("Slice: ", slice)
	fmt.Println("length of Slice", len(slice))
}
