package main

import "fmt"

func main() {

	var choice string
	var number1, number2 int

	fmt.Println("-----CALCULATOR-----")
	fmt.Println("Select the operation you want to calculate: (+) (-) (*) (/)")
	// it must send the address of variable choice to change it value
	fmt.Scanf("%s", &choice)

	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter another number: ")
	fmt.Scanf("%d", &number2)

	switch choice {
	case "+":
		fmt.Println(number1 + number2)
	case "-":
		fmt.Println(number1 - number2)
	case "*":
		fmt.Println(number1 * number2)
	case "/":
		fmt.Println(number1 / number2)
	}
}
