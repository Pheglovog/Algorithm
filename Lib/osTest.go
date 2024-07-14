package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"syscall"
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
	tryChdir()
	tryChtimes()
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

func tryChdir() {
	fmt.Println(strings.Repeat("-", 8) + "tryChdir" + strings.Repeat("-", 8))
	err := os.Chdir("/BlockChain/Algorithm")
	if err != nil {
		log.Fatal(err)
	}
	f, _ := os.Open(".")
	files, _ := f.ReadDir(-1)
	fmt.Println(files)
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func tryChtimes() {
	fileInfo, err := os.Stat("./run.sh")
	stat_t := fileInfo.Sys().(*syscall.Stat_t)
	if err != nil {
		panic(err)
	}
	fmt.Println("最后修改时间：", timespecToTime(stat_t.Mtim))
	fmt.Println("最后访问时间：", timespecToTime(stat_t.Atim))
	// 改变文件时间戳为两天前
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("./run.sh", lastAccessTime, lastModifyTime)
	if err != nil {
		panic(err)
	}
	fileInfo, err = os.Stat("./run.sh")
	stat_t = fileInfo.Sys().(*syscall.Stat_t)
	if err != nil {
		panic(err)
	}
	fmt.Println("最后修改时间：", timespecToTime(stat_t.Mtim))
	fmt.Println("最后访问时间：", timespecToTime(stat_t.Atim))
}
