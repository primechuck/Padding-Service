package Padding_Service

import "testing"

func TestLeftPad_Function(t *testing.T) {
	result := leftPad("Howdy")
	expect := "Howdy"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}
