package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello from new go server!"))
		})
	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Printf("something happened with server: %v", err.Error())
	}
	fmt.Println("Server initialized..")
}
