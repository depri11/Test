package main

import (
	"fmt"
	"net/http"
	"os"

	palindrome "github.com/depri11/Test/src/palindrom"
)

func main() {

	service := palindrome.NewService() // Create a new palindrome service instance to handle the request
	handler := palindrome.NewHandler(service) // Create a new palindrome handler instance to handle the request

	http.HandleFunc("/palindrome", handler.Palindrome) // Set the handler for the request path

	err := http.ListenAndServe(":3000", nil) // Listen on port 3000 and serve the request
	// print any server-based error messages
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
