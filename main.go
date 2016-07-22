package main

import (
	"strings"
	"fmt"
)

func main() {

}

func leftPad(paddingLength int, stringToPad string) string {
	if paddingLength < 0 {
		return fmt.Sprintf("Padding length must no be negative, value entered was %d", paddingLength)
	}
	return strings.Repeat("0",paddingLength)+stringToPad
}

func rightPad(paddingLength int, stringToPad string) string {
	return stringToPad+strings.Repeat("0",paddingLength)
}