package main

import (
	"fmt"
	"sync"

	"examble.com/crypto/api"
)

// main goroutine
func main() {
	// waitgroup = it will wait for goroutines to finish
	var wg sync.WaitGroup

	currencies := []string{"BTC", "ETH", "SOL"}

	for _, currency := range currencies {
		// every time goroutine is open, say 1, add 1
		wg.Add(1)

		// new goroutine with a function that will synchronously execute, get data, and when it will finish call done
		// closures =  nested function that allows us to access variables of the outer function even after the outer function is closed.
		go func() {
			getRate(currency)
			// every time goroutine end, should call done
			wg.Done()
		}()

	}

	// done() wiil be executed immediately after call a new goroutine(getCurrency) without waiting
	// Wait blocks until the WaitGroup counter is zero.
	wg.Wait()
}

// new goroutine
func getRate(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is $%.2f\n", rate.Currency, rate.Price)
	}
}
