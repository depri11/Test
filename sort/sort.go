package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Category struct {
	ID   int
	Name string
}

type Book struct {
	CategoryID int
	Name       string
	Height     int
	Category   Category
}

// Category struct for storing category data
var categories = []Category{
	{6, "Applied Sciences (600)"},
	{7, "Arts (700)"},
	{0, "General (000)"},
	{9, "Geography, History (900)"},
	{4, "Language (400)"},
	{8, "Literature (800)"},
	{1, "Philosophy, Psychology (100)"},
	{2, "Religion (200)"},
	{5, "Science, Math (500)"},
	{3, "Social Sciences (300)"},
}

// Function to sort book by category
func SortByCategory(input string) string {
	var temp []string                // Declare slice to store temporary data
	arr := strings.Split(input, " ") // Convert input string to slice

	var books []Book // Declare array of books to store book data
	// Iterate through the input string and convert it to book data
	for _, v := range arr {
		categoryID, _ := strconv.Atoi(v[0:1])
		bookName := v[1:2]
		bookHeight, _ := strconv.Atoi(v[2:4])

		// Iterate through the category slice to get the category name
		for _, v := range categories {
			if v.ID == categoryID {
				book := Book{
					CategoryID: categoryID,
					Name:       bookName,
					Height:     bookHeight,
					Category:   v,
				}
				books = append(books, book) // Append the book data to the books slice
			}
		}
	}

	// Sort the books slice by category with books slice name
	for i := 0; i < len(categories); i++ {
		var arrNum []int // Declare array of arrNum to store the number of books in each category
		for j := 0; j < len(books); j++ {
			if categories[i].ID == books[j].CategoryID {
				if len(arrNum) < 3 { // A maximum of 2 books for the same category id and name
					arrNum = append(arrNum, books[j].Height) // Append the book data to the books slice by category with books slice height
				}
			}
		}

		// Sort the arrNum slice by height
		sort.Slice(arrNum, func(a, b int) bool {
			return arrNum[a] > arrNum[b]
		})

		// Iterate through the arrNum slice to get the number of books in each category
		for _, v := range arrNum {
			for k := 0; k < len(books); k++ {
				// Check if the book height is equal or not to the arrNum slice height
				if categories[i].ID == books[k].CategoryID {
					if v == books[k].Height {
						str := strconv.Itoa(books[k].CategoryID) + books[k].Name + strconv.Itoa(books[k].Height)
						temp = append(temp, str) // Append the book data to the temp slice
					}
				}
			}
		}
	}

	result := strings.Join(temp, " ") // Convert temp slice to string
	return result                     // Return the result
}

func main() {
	res := SortByCategory("3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14")
	fmt.Println(res)
}
