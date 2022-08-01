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

func SortByCategory(input string) string {
	var temp []string
	arr := strings.Split(input, " ")

	var books []Book
	for _, v := range arr {
		categoryID, _ := strconv.Atoi(v[0:1])
		bookName := v[1:2]
		bookHeight, _ := strconv.Atoi(v[2:4])

		for _, v := range categories {
			if v.ID == categoryID {
				book := Book{
					CategoryID: categoryID,
					Name:       bookName,
					Height:     bookHeight,
					Category:   v,
				}
				books = append(books, book)
			}
		}
	}

	for i := 0; i < len(categories); i++ {
		var arrNum []int
		for j := 0; j < len(books); j++ {
			if categories[i].ID == books[j].CategoryID {
				if len(arrNum) < 3 {
					arrNum = append(arrNum, books[j].Height)
				}
			}
		}

		sort.Slice(arrNum, func(a, b int) bool {
			return arrNum[a] > arrNum[b]
		})

		for _, v := range arrNum {
			for k := 0; k < len(books); k++ {
				if categories[i].ID == books[k].CategoryID {
					if v == books[k].Height {
						str := strconv.Itoa(books[k].CategoryID) + books[k].Name + strconv.Itoa(books[k].Height)
						temp = append(temp, str)
					}
				}
			}
		}
	}

	result := strings.Join(temp, " ")

	return result
}

func main() {
	res := SortByCategory("3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14")
	fmt.Println(res)
}
