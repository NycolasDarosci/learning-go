package main

import (
	"fmt"

	// import a packge from the same module
	io "learngo.com/inputOutput"
)

var number = 12

func init() {
	fmt.Println("Initializing init() one")
	prev := number
	number = 5
	fmt.Println("Previous number var: ", prev, " has changed it value to: ", number)
}

func init() {
	fmt.Println("Initializing init() two")
}

func main() {
	// go print("main")
	// go print("other")
	// go print("testing concurrency")
	// go print("go is cool")
	// time.Sleep(time.Second)

	io.PrintData(1, io.Url)
	fmt.Println(io.Math(2, 5.2))
	io.Slices()

	// function return two values
	a, b := io.AddAndSubstract(6, 3)
	fmt.Println("Sum: ", a, " Substract: ", b)

	// the sentece fmt.Println("testing defer")
	// will be executed at the end of the function call
	defer fmt.Println("testing defer")

	// ignoring the second value
	c, _ := io.AddAndSubstract(6, 3)
	fmt.Println("Sum: ", c)

	g, _ := io.AddAndSubstractWithName(6, 3)
	fmt.Println("Sum: ", g)

	// pointer
	// make a change to the reference of the variable
	i := 6
	fmt.Println(i)
	io.Increment(&i)
	fmt.Println(i)
	fmt.Println(&i) // will print the memory address of the i variable 0xc00008e068

	tax, _ := io.CalculateTax(100.0)
	stateTax, cityTax := io.CalculateTaxWithName(100.0)
	fmt.Println(tax, stateTax, cityTax)

	// control instructores

	printControlStructures()

}
