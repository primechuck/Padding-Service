package Padding_Service

import (
	"strings"
)

func leftPad(paddingLength int, stringToPad string) string {
	return strings.Repeat("0",paddingLength)+stringToPad
}

func rightPad(paddingLength int, stringToPad string) string {
	return stringToPad+strings.Repeat("0",paddingLength)
}