package main

import (
	"fmt"
	"math"
)

func main() {
	result := findMissingNumber("456789101213")
	fmt.Println(result)
}

// Function for getting the digit/integer at position i with the length of main
func getDigit(s string, i int, main int) int {
	if i+main > len(s) {
		return -1
	}

	// Find value at index i with the length of main
	value := 0
	for l := 0; l < main; l++ {
		temp := int(s[i+l]) - '0'
		if temp < 0 || temp > 9 {
			return -1
		}
		value = value*10 + temp
	}

	return value
}

// Logical function, search the missing number
func findMissingNumber(s string) int {

	maxDigits := 6
	// Iterate maxDigits as main
	for main := 1; main <= maxDigits; main++ {

		// Get the value of first number with current length
		n := getDigit(s, 0, main)
		if n == -1 { // If n returns -1, break the loop. Consider it an error
			break
		}

		tempMissingNumber := -1 // Store the missing number
		isError := false        // To indicate whether the iteration failed or no

		// Iterate the subsequent numbers after the first digit/previous number, we indicate it as variable n
		for i := main; i != len(s); i += 1 + int(math.Log10(float64(n))) {

			//If we haven't yet found the missing number for the current iteration. Next number will be n + 2
			if tempMissingNumber == -1 && getDigit(s, i, 1+int(math.Log10(float64(n+2)))) == n+2 { // If returned data from getDigit equal or the same as n + 2, it means we finally found the missing number
				tempMissingNumber = n + 1
				n += 2
			} else if getDigit(s, i, 1+int(math.Log10(float64(n+1)))) == n+1 { // If next value is equal or the same as n + 1 then do increment to n and do another cycle of iteration
				n++
			} else { // If there's nothing to iterate, break the loop and indicate it as an error
				isError = true
				break
			}
		}

		if !isError {
			return tempMissingNumber
		}
	}
	return -1 // Not found (error) or there's no missing number
}
