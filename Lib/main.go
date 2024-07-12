package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "os":
		osTest()
	case "regex":
		regexTest()
	default:
		fmt.Println("not a right command")
	}
}
