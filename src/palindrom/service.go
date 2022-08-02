package palindrome

import (
	"strconv"
	"strings"

	"github.com/depri11/Test/src/helpers"
	"github.com/depri11/Test/src/models"
)

// service struct for handling palindrome
type service struct{}

// NewService returns a new palindrome service
func NewService() *service {
	return &service{}
}

// Logic for checking palindrome
func (s *service) CheckPalindrom(input models.Palindrome) *helpers.Response {
	arr := strings.Split(input.Data, " ") // Split the string into array

	// Convert the string into integer
	min, err := strconv.Atoi(arr[0])
	if err != nil {
		return helpers.ResponseAPI(500, "error", err.Error())
	}
	// Convert the string into integer
	max, err := strconv.Atoi(arr[1])
	if err != nil {
		return helpers.ResponseAPI(500, "error", err.Error())
	}

	// Validation if max is less than min value or not
	if max < min {
		return helpers.ResponseAPI(401, "error", err.Error())
	}

	// Declare variable to store number of palindrom
	var result int

	// Iterate from min to max of palindrom length and check if it's palindrom
	for i := min; i <= max; i++ {
		var reverse int

		// Check number of palindrom or not by reversing the number
		for temp := i; temp > 0; temp /= 10 {
			remainder := i % 10                  // Get the remainder of the number
			reverse = (reverse * 10) + remainder // reverse the remainder and add to the reverse number
		}

		// Compare if the reversed number is equal or not to the original number
		if i == reverse {
			result += 1
		}
	}

	// Return the number of palindrom
	return helpers.ResponseAPI(200, "Success", result)
}
