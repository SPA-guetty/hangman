package common

import (
	"fmt"
)

const (
	Reset    = "\033[0m"
	Bold     = "\033[1m"
	FgRed    = "\033[31m"
	FgGreen  = "\033[32m"
	FgYellow = "\033[33m"
	FgBlue   = "\033[34m"
	FgCyan   = "\033[36m"
	FgPink   = "\033[35m"
)

var (
	Title       = FgCyan + Bold
	SubTitle    = FgYellow + Bold
	DisplayJosé = FgGreen
	ErrorText   = FgRed
	WarningText = FgYellow
	CheatCode   = FgPink
)

func UpdateColor(colorName, colorCode string) {
	switch colorName {
	case "title":
		Title = colorCode + Bold
	case "subtitle":
		SubTitle = colorCode + Bold
	case "displayJosé":
		DisplayJosé = colorCode
	case "error":
		ErrorText = colorCode
	case "warning":
		WarningText = colorCode
	default:
		fmt.Println("Unknown color")
	}
}
