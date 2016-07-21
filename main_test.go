package Padding_Service

import "testing"

func TestLeftPad_Function_Returns_padded_String(t *testing.T) {
	result := leftPad(3, "Howdy")
	expect := "000Howdy"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}

func TestRightPad_Function_StringReturn(t *testing.T) {
	result := rightPad(3, "Howdy")
	expect := "Howdy000"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}
