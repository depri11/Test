package main

import (
	"fmt"
	"net/http"
	"os"

	palindrome "github.com/depri11/Test/src/palindrom"
)

func main() {

	service := palindrome.NewService()
	handler := palindrome.NewHandler(service)

	http.HandleFunc("/palindrome", handler.Palindrome)

	err := http.ListenAndServe(":3000", nil)
	// print any server-based error messages
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
