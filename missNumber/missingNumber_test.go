package main

import "testing"

func TestFindMissingNumber(t *testing.T) {

	actual := findMissingNumber("23242526272830")
	expected := 29

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFindMissingNumberTwo(t *testing.T) {

	actual := findMissingNumber("101102103104105106107108109111112113")
	expected := 110

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFindMissingNumberThree(t *testing.T) {

	actual := findMissingNumber("12346789")
	expected := 5

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFindMissingNumberFour(t *testing.T) {

	actual := findMissingNumber("79101112")
	expected := 8

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFindMissingNumberFive(t *testing.T) {

	actual := findMissingNumber("7891012")
	expected := 11

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFindMissingNumberSix(t *testing.T) {

	actual := findMissingNumber("9799100101102")
	expected := 98

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFindMissingNumberSeven(t *testing.T) {

	actual := findMissingNumber("100001100002100003100004100006")
	expected := 100005

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
