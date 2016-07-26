package main

import (
	"fmt"
	"net/http"
	"strings"
	"io"
	"io/ioutil"
)

func main() {
}

func rightPadAPI(w http.ResponseWriter, r *http.Request) {
	request, _ := ioutil.ReadAll(r.Body)
	io.WriteString(w, string(request))
}

func leftPadAPI(w http.ResponseWriter, r *http.Request) {
	request, _ := ioutil.ReadAll(r.Body)
	io.WriteString(w, string(request))
}

func leftPad(paddingLength int, stringToPad string) string {
	if paddingLength < 0 {
		return fmt.Sprintf("Padding length must no be negative, value entered was %d", paddingLength)
	}
	return strings.Repeat("0", paddingLength) + stringToPad
}

func rightPad(paddingLength int, stringToPad string) string {
	if paddingLength < 0 {
		return fmt.Sprintf("Padding length must no be negative, value entered was %d", paddingLength)
	}
	return stringToPad + strings.Repeat("0", paddingLength)
}
