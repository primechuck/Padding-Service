package Padding_Service

import "testing"

func TestLeftPad_Function_StringReturn(t *testing.T) {
	result := leftPad(3, "Howdy")
	expect := "000Howdy"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}
