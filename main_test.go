package main

import (
	"testing"
)

func TestLeftPad_Function_Returns_padded_String(t *testing.T) {
	result := leftPad(3, "Howdy")
	expect := "000Howdy"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}

func TestLeftPad_Function_Returns_error_when_number_is_less_than_zero(t *testing.T) {
	result := leftPad(-3, "Howdy")
	expect := "Padding length must no be negative, value entered was -3"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}

func TestRightPad_Function_Returns_padded_String(t *testing.T) {
	result := rightPad(3, "Howdy")
	expect := "Howdy000"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}

func TestRightPad_Function_Returns_error_when_number_is_less_than_zero(t *testing.T) {
	result := rightPad(-3, "Howdy")
	expect := "Padding length must no be negative, value entered was -3"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}
