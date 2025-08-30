package main

import (
	"flag"
	"github.com/kashsuks/termigex/helper"
)

func main() {
	var exampleString int
	flag.IntVar(&exampleString, "number", 0, "team number flag")

	flag.Parse()
}