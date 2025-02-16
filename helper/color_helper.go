package helper

import "fmt"

const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
)

func PrintError(text string) {
	fmt.Println(ColorRed + text + ColorReset)
}

func PrintInfo(text string) {
	fmt.Println(ColorYellow + text + ColorReset)
}

func PrintSuccess(text string) {
	fmt.Println(ColorGreen + text + ColorReset)
}
