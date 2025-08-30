package helper

import "fmt"

const ERROR_COLOUR = "\033[31m"
const PASS_COLOUR = "\033[32m"
const RESET = "\033[0m"

func PrintMatches(matches []string) {
	for _, m := range matches {
		fmt.Println(PASS_COLOUR + m + RESET)
	}
}