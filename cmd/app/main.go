package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose output")

	flag.Parse()

	fmt.Println("Hello, world!")

	if *verbose {
		fmt.Println("This is more verbose")
	}
}
