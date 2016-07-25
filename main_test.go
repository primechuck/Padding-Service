package main

import (
	"testing"
)

func TestLeftPadFunctionReturnspaddedString(t *testing.T) {
	result := leftPad(3, "Howdy")
	expect := "000Howdy"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}

func TestLeftPadFunctionReturnserrorwhennumberislessthanzero(t *testing.T) {
	result := leftPad(-3, "Howdy")
	expect := "Padding length must no be negative, value entered was -3"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}

func TestRightPadFunctionReturnspaddedString(t *testing.T) {
	result := rightPad(3, "Howdy")
	expect := "Howdy000"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}

func TestRightPadFunctionReturnserrorwhennumberislessthanzero(t *testing.T) {
	result := rightPad(-3, "Howdy")
	expect := "Padding length must no be negative, value entered was -3"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}
