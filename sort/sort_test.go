package main

import (
	"testing"
)

func TestAppendToStruct(t *testing.T) {
	input := "3A13"
	actual := AppendToStruct(input)

	if actual[0].CategoryID != 3 {
		t.Errorf("Expected 3, got %v", actual)
	}
}

func TestSortByCategory(t *testing.T) {
	var dataBook = []Book{
		{CategoryID: 3, Name: "A", Height: 13},
		{CategoryID: 5, Name: "X", Height: 19},
		{CategoryID: 9, Name: "Y", Height: 20},
		{CategoryID: 2, Name: "C", Height: 18},
		{CategoryID: 1, Name: "N", Height: 20},
		{CategoryID: 3, Name: "N", Height: 20},
		{CategoryID: 1, Name: "M", Height: 21},
		{CategoryID: 1, Name: "F", Height: 14},
		{CategoryID: 9, Name: "A", Height: 21},
		{CategoryID: 3, Name: "N", Height: 21},
		{CategoryID: 0, Name: "E", Height: 13},
		{CategoryID: 5, Name: "G", Height: 14},
		{CategoryID: 8, Name: "A", Height: 23},
		{CategoryID: 9, Name: "E", Height: 22},
		{CategoryID: 3, Name: "N", Height: 14},
	}

	actual := SortByCategory(dataBook, categories)
	expected := "0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3N14 3A13"

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
