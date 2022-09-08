package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Server unable to start")
	}
}
