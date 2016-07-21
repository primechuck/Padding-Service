package Padding_Service

import (
	"strings"
)

func leftPad(paddingLength int, stringToPad string) string {
	return strings.Repeat("0",paddingLength)+stringToPad
}