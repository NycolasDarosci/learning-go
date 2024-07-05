package main

import "fmt"

func printControlStructures() {

	a := 1
	b := 2

	ifi(&a, &b)
	four()
}

func ifi(a *int, b *int) {
	if *a > *b {
		fmt.Printf("%v major than %v\n", *a, *b)
	} else if *b > *a {
		fmt.Printf("%v major than %v\n", *b, *a)
	} else {
		fmt.Println("no one is major")
	}

}

func four() {
	var slice [10]int

	for count := 0; count < 1; {
		if count%2 == 0 {
			slice[count] = count
		} else {
			slice[count] = count + 1
		}

		for key, value := range slice {
			fmt.Println("key: ", key, " value: ", value)
		}
		count += 1
	}

}
