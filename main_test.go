package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRightPadAPIReturns(t *testing.T) {
	expect := `{String: "rightPadAPI", Length: "3"}`
	b := strings.NewReader(expect)

	r, err := http.NewRequest("POST", "/rightpad", b)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	rightPadAPI(responseRecorder, r)

	if responseRecorder.Body.String() != expect {
		t.Errorf("Handler returned unexpected body: got %v wanted %v",
			responseRecorder.Body.String(), expect)
	}
}

func TestLeftPadAPIReturns(t *testing.T) {

	expect := `{String: "leftPadAPI", Length: "3"}`
	b := strings.NewReader(expect)

	r, err := http.NewRequest("POST", "/leftpad", b)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	leftPadAPI(responseRecorder, r)

	if responseRecorder.Body.String() != expect {
		t.Errorf("Handler returned unexpected body: got %v wanted %v",
			responseRecorder.Body.String(), expect)
	}
}

func TestLeftPadFunctionReturnsPaddedString(t *testing.T) {
	result := leftPad(3, "Howdy")
	expect := "000Howdy"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}

func TestLeftPadFunctionReturnsErrorWhenNumberIsLessThanZero(t *testing.T) {
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

func TestRightPadFunctionReturnsErrorWhenNumberIsLessThanZero(t *testing.T) {
	result := rightPad(-3, "Howdy")
	expect := "Padding length must no be negative, value entered was -3"

	if result != expect {
		t.Errorf("Test Failed, expected: %s but got %s.", expect, result)
	}
}
