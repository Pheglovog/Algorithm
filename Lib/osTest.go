package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func osTest() {
	tryCreateTemp()
	tryReadDir()
	if len(os.Args) > 1 {
		tryArgs()
	}
}

// Use CreateTemp
func tryCreateTemp() {
	fmt.Println(strings.Repeat("-", 8) + "tryCreateTemp" + strings.Repeat("-", 8))
	f, err := os.CreateTemp("", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name()) // clean up
	n, err := f.Write([]byte("content"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	f.Seek(0, io.SeekStart)
	b := make([]byte, 7)
	n, err = f.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	fmt.Println(f.Name(), "IIIII", string(b))
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(100 * time.Millisecond)
}

// Use ReadDir
func tryReadDir() {
	fmt.Println(strings.Repeat("-", 8) + "tryReadDir" + strings.Repeat("-", 8))
	f, err := os.Open("..")
	checkErr(err)
	defer f.Close()
	files, err := f.ReadDir(-1)
	checkErr(err)
	re := regexp.MustCompile(`(\.+)[A-z]`)
	for _, v := range files {
		if v.IsDir() {
			if !re.Match([]byte(v.Name())) {
				fmt.Println("Dir:", v.Name())
			}
		} else {
			fmt.Println("File", v.Name())
			nf, err := os.Open("../" + v.Name())
			checkErr(err)
			b, err := io.ReadAll(nf)
			checkErr(err)
			fmt.Println(string(b))
		}
	}
}

// Use Args
func tryArgs() {
	fmt.Println(strings.Repeat("-", 8) + "tryArgs" + strings.Repeat("-", 8))
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "+"
	}
	fmt.Println(s)
}
