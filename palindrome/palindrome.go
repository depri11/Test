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
	arr := strings.Split(input, " ")

	min, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0
	}
	max, err := strconv.Atoi(arr[1])
	if err != nil {
		return 0
	}

	if max < min {
		return 0
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

	return result
}
