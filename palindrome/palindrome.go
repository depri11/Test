package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := NumOfPalindrom("1 10")
	fmt.Println(res)
}

func NumOfPalindrom(input string) int {
	// Convert input string to slice
	arr := strings.Split(input, " ")

	// Convert string to int
	min, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0
	}

	// Convert string to int
	max, err := strconv.Atoi(arr[1])
	if err != nil {
		return 0
	}

	// Validation if max is less than min value or not
	if max < min {
		return 0
	}

	// Declare variable to store number of palindrom
	var result int

	// Iterate from min to max of palindrom length and check if it's palindrom
	for i := min; i <= max; i++ {
		var reverse int

		// Check number of palindrom or not by reversing the number
		for temp := i; temp > 0; temp /= 10 {
			remainder := i % 10 // Get the remainder of the number
			reverse = (reverse * 10) + remainder // reverse the remainder and add to the reverse number
		}

		// Compare if the reversed number is equal or not to the original number
		if i == reverse {
			result += 1
		}
	}

	// Return the number of palindrom
	return result
}
