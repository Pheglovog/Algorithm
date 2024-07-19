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
	case "fs":
		fsTest()
	case "io":
		ioTest()
	case "time":
		timeTest()
	case "json":
		jsonTest()
	default:
		fmt.Println("not a right command")
	}
}
