package main

import (
	"fmt"
	"regexp"
)

func regexTest() {
	tryGroup()
}

func tryGroup() {
	str := "John Doe, jane@example.com"
	re := regexp.MustCompile(`(\w+)\s(\w+),\s(\w+@\w+\.\w+)`)

	match := re.FindStringSubmatch(str)
	fmt.Println(len(match))
	fmt.Println(match)                       // [John Doe, John Doe, jane@example.com]
	fmt.Println("Name:", match[1], match[2]) // Name: John Doe
	fmt.Println("Email:", match[3])          // Email: jane@example.com
	fmt.Println("All: ", match[0])
}
