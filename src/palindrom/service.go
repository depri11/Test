package palindrome

import (
	"strconv"
	"strings"

	"github.com/depri11/Test/src/helpers"
	"github.com/depri11/Test/src/models"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) CheckPalindrom(input models.Palindrome) *helpers.Response {
	arr := strings.Split(input.Data, " ")

	min, err := strconv.Atoi(arr[0])
	if err != nil {
		return helpers.ResponseAPI(500, "error", err.Error())
	}
	max, err := strconv.Atoi(arr[1])
	if err != nil {
		return helpers.ResponseAPI(500, "error", err.Error())
	}

	if max < min {
		return helpers.ResponseAPI(401, "error", err.Error())
	}

	var result int

	for i := min; i <= max; i++ {
		var res int

		for temp := i; temp > 0; temp /= 10 {
			remainder := i % 10
			res = (res * 10) + remainder
		}

		if i == res {
			result += 1
		}
	}

	return helpers.ResponseAPI(200, "Success", result)
}
