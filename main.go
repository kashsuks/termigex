package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/kashsuks/termigex/helper"
)

const ERROR_COLOUR = "\033[31m"
const PASS_COLOUR = "\033[32m"
const RESET = "\033[0m"

func main() {
	regex := flag.String("regex", "", "Regex pattern to match")
	str := flag.String("str", "", "Test string to search")
	
	flag.Parse()

	if *regex == "" || *str == "" {
		fmt.Println("Usage: go run main.go -regex='<pattern>' -str='<test string>'") //fix later on to have the proper one with the name of the package
		os.Exit(1)
	}
	
	matches, err := helper.FindMathes(*regex, *str)
	if err != nil {
		fmt.Println(ERROR_COLOUR + "Regex error: " + RESET, err)
		os.Exit(1)
	}

	if len(matches) == 0 {
		fmt.Println("No matches found")
	} else {
		fmt.Println("Matches found: ")
		helper.PrintMatches(matches)
	}
}