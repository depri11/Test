package main

import (
	"testing"
)

func TestNumOfPalindrome(t *testing.T) {
	var data = "1 10"

	actual := NumOfPalindrom(data)

	if actual != 9 {
		t.Errorf("Expected 9, got %v", actual)
	}
}

func TestNumOfPalindrome_WithSecondNumberGreaterThanFirstNumber(t *testing.T) {
	var data = "21 10"

	actual := NumOfPalindrom(data)
	expected := 0

	if actual != expected {
		t.Errorf("Expected 0, got %v", actual)
	}
}

func TestNumOfPalindrome_WithInputAlphabet(t *testing.T) {
	var data = "asdsad"

	actual := NumOfPalindrom(data)
	expected := 0

	if actual != expected {
		t.Errorf("Expected 0, got %v", actual)
	}
}
