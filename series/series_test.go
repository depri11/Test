package main

import "testing"

func TestDeretAngka(t *testing.T) {
	input := "12346789"
	actual := deretAngka(input)
	expected := 5

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDeretAngka_WithOneDigitToTwoDigit(t *testing.T) {
	input := "7891012"
	actual := deretAngka(input)
	expected := 11

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDeretAngka_WithNumberFromBigToSmall(t *testing.T) {
	input := "98765321"
	actual := deretAngka(input)
	expected := 0

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDeretAngka_WithInputAlphabet(t *testing.T) {
	input := "Test"
	actual := deretAngka(input)
	expected := 0

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
