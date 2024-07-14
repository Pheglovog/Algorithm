package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ioTest() {
	tryCopy()
}

func tryCopy() {
	fmt.Println(strings.Repeat("-", 8) + "tryCopy" + strings.Repeat("-", 8))
	r := strings.NewReader("abd")
	//可以将一个类型转化成一个interface
	ir := io.ReadSeeker(r)
	buf := make([]byte, 10)
	n, err := ir.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]), ": buf")

	//这个seek非常重要
	ir.Seek(0, io.SeekStart)
	//copy也输入到标准输出中，但是为什么不见了
	_, err = io.Copy(os.Stdout, ir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" : copy")
}
