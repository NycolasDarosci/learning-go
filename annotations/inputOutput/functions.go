package inputOutput

// fmt = format
// imports are not per package, they are per file
import "fmt"

// PrintData with uppercase, it is exported externally to the package
// public
func PrintData(n int, str string) {
	fmt.Println("Hello")
	fmt.Println("World")
	fmt.Println(Url)
	fmt.Println(Number)
	fmt.Println(n)
	fmt.Println(str)
	fmt.Println(Pi)
}

func Math(n1 int, n2 float64) float64 {

	n1AsFloat := float64(n1)

	return n1AsFloat * n2
}

func AddAndSubstract(a int, b int) (int, int) {
	return a + b, a - b
}

func AddAndSubstractWithName(a int, b int) (sum int, substract int) {
	return a + b, a - b
}

func CalculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

func CalculateTaxWithName(price float32) (stateTax float32, cityTax float32) {
	stateTax = price * 0.09
	cityTax = price * 0.02
	return
}

// func with a pointer argument
// variable a will change it value
// and whatever package it will be changed it value reference
func Increment(a *int) {
	if *a > 140 {
		// panic function
		// interruption of the flow. Like as exception
		panic("a is greater than 140")
	}
	fmt.Printf("The pointer is %v and the value is %v\n", a, *a)
	*a++
}

// errors desing pattern (error management)

type User struct {
	id   int
	name string
}

func readUser(id int) (user int, err error) {
	// ... proceed with the reading and see a bool ok value

	if 1 > 1 {
		return user, nil
	} else {
		return 1, err
	}
}

func main() {

	user, err := readUser(1)
	fmt.Println(user, err)

}
