package main

import (
	"fmt"
	"strings"
)

func main() {

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
