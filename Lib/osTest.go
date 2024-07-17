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
	//function
	tryFileSystem()

	if len(os.Args) > 1 {
		tryArgs()
	}

	tryEnv()
	tryosExpand()

	tryUserAndGroup()

	tryMem()

	tryProcess()

	tryKernel()

	//type

	tryosReadDir()

	tryFile()
}

func tryFileSystem() {
	fmt.Println(strings.Repeat("-", 8) + "tryFileSystem" + strings.Repeat("-", 8))
	tryChtimes()
	tryReadDir()
	tryDirFS()
	tryChdir()
	tryExecuable()
	tryGetwd()

	fmt.Printf("os.IsPathSeparator('/'): %v\n", os.IsPathSeparator('/'))

	tryLink()
	tryMkdir()
	tryPipe()
	//Rename
	//Truncate
	trySpecialDir()
}

func trySpecialDir() {
	fmt.Println(strings.Repeat("-", 8) + "trySpecialDir" + strings.Repeat("-", 8))
	fmt.Println(os.UserCacheDir())
	fmt.Println(os.UserCacheDir())
	fmt.Println(os.UserHomeDir())
}

func tryPipe() {
	fmt.Println(strings.Repeat("-", 8) + "tryPipe" + strings.Repeat("-", 8))
	buf := make([]byte, 20)
	r, w, err := os.Pipe()
	checkErr(err)
	defer r.Close()
	defer w.Close()
	_, err = w.Write([]byte("beautiful day\n"))
	checkErr(err)
	n, err := r.Read(buf)
	checkErr(err)
	fmt.Printf("buf: %q\n", buf[:n])
}

func tryMkdir() {
	var err error
	defer func() {
		if err != nil {
			err = os.NewSyscallError("os.Mkdir", err)
			log.Fatal(err)
		}
	}()

	fmt.Println(strings.Repeat("-", 8) + "tryMkdir" + strings.Repeat("-", 8))
	err = os.Mkdir("FileSystem/dir", 0777)
	err = os.MkdirAll("FileSystem/dir/all", 0777)
	temp, err := os.MkdirTemp("FileSystem/dir/", "temp")
	defer os.RemoveAll("FileSystem/dir")

	err = os.WriteFile("FileSystem/dir/temp.txt", []byte("NICENICENICE"), 0777)
	content, err := os.ReadFile("FileSystem/dir/temp.txt")
	fmt.Printf("content: %q\n", content)
	fmt.Printf("temp: %v\n", temp)
	fmt.Printf("FileSystem/dir : %v\n", os.DirFS("FileSystem/dir"))
	fmt.Printf("os.TempDir(): %v\n", os.TempDir())
}

func tryLink() {
	fmt.Println(strings.Repeat("-", 8) + "tryLink" + strings.Repeat("-", 8))
	err := os.Link("run.sh", "FileSystem/run.sh")
	checkErr(err)
	defer os.Remove("FileSystem/run.sh")

	err = os.Symlink("run.sh", "FileSystem/runSym.sh")
	checkErr(err)
	defer os.Remove("FileSystem/runSym.sh")
	path, err := os.Readlink("FileSystem/runSym.sh")
	checkErr(err)
	fmt.Printf("path: %v\n", path)
	f, err := os.ReadFile("FileSystem/run.sh")
	checkErr(err)
	fmt.Printf("FileSystem/run.sh: %q\n", f)

	fi1, err := os.Stat("run.sh")
	checkErr(err)
	fi2, err := os.Stat("FileSystem/run.sh")
	checkErr(err)
	fi3, err := os.Stat("FileSystem/runSym.sh")
	checkErr(err)
	fmt.Printf("os.SameFile(fi1, fi2): %v\n", os.SameFile(fi1, fi2))
	fmt.Printf("os.SameFile(fi1, fi3): %v\n", os.SameFile(fi1, fi3))
}

func tryGetwd() {
	fmt.Println(strings.Repeat("-", 8) + "tryGetwd" + strings.Repeat("-", 8))
	fmt.Println(os.Getwd())
}

// Use CreateTemp
func tryCreateTemp() {
	f, err := os.CreateTemp("", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name()) // clean up
	n, err := f.Write([]byte("content"))
	if err != nil {
		log.Fatal(err)
	}
	f.Seek(0, io.SeekStart)
	b := make([]byte, n)
	n, err = f.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f.Name(), "IIIII", string(b[:n]))
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(100 * time.Millisecond)
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
	defer os.Chdir("/BlockChain/Algorithm/Lib")
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
	fmt.Println(strings.Repeat("-", 8) + "tryChtimes" + strings.Repeat("-", 8))
	fileInfo, err := os.Stat("run.sh")
	checkErr(err)
	stat_t := fileInfo.Sys().(*syscall.Stat_t)
	fmt.Println("最后修改时间：", timespecToTime(stat_t.Mtim))
	fmt.Println("最后访问时间：", timespecToTime(stat_t.Atim))
	// 改变文件时间戳为两天前
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("run.sh", lastAccessTime, lastModifyTime)
	defer os.Chtimes("run.sh", time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
	fileInfo, err = os.Stat("run.sh")
	stat_t = fileInfo.Sys().(*syscall.Stat_t)
	if err != nil {
		panic(err)
	}
	fmt.Println("最后修改时间：", timespecToTime(stat_t.Mtim))
	fmt.Println("最后访问时间：", timespecToTime(stat_t.Atim))
}

func tryDirFS() {
	fmt.Println(strings.Repeat("-", 8) + "tryDirFS" + strings.Repeat("-", 8))

	fsys := os.DirFS("/BlockChain/Algorithm")
	fmt.Println(fsys)
}

func tryExecuable() {
	fmt.Println(strings.Repeat("-", 8) + "tryExecuable" + strings.Repeat("-", 8))
	s, err := os.Executable()
	checkErr(err)
	fmt.Println(s)
}

func tryosExpand() {
	fmt.Println(strings.Repeat("-", 8) + "tryosExpand" + strings.Repeat("-", 8))
	mapper := func(placeholderName string) string {
		switch placeholderName {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}
		return ""
	}

	fmt.Println(os.Expand("Good ${DAY_PART}, $NAME!, $NAME", mapper))

	os.Setenv("HSHNAME", "gopher")
	defer os.Unsetenv("HSHNAME")
	os.Setenv("HSHBURROW", "/usr/gopher")
	defer os.Unsetenv("HSHBURROW")

	fmt.Println(os.ExpandEnv("$HSHNAME and $HSHBURROW"))
}

func tryEnv() {
	fmt.Println(strings.Repeat("-", 8) + "tryEnv" + strings.Repeat("-", 8))
	// 拿到所有的环境变量
	kvs := os.Environ()
	fmt.Println(kvs[0])
	// 设置， 取出和删除
	os.Setenv("HSH", "shuai")
	defer os.Unsetenv("HSH")
	s := os.Getenv("HSH")
	fmt.Println("HSH:", s)

	if s, ok := os.LookupEnv("HSH"); ok {
		fmt.Printf("HSH: %v\n", s)
	}
}

func tryUserAndGroup() {
	fmt.Println(strings.Repeat("-", 8) + "tryUserAndGroup" + strings.Repeat("-", 8))
	fmt.Println("euid:", os.Geteuid())
	fmt.Println("egid:", os.Getegid())
	fmt.Printf("os.Geteuid(): %v\n", os.Geteuid())
	groups, err := os.Getgroups()
	checkErr(err)
	fmt.Printf("groups: %v\n", groups)
}

func tryMem() {
	fmt.Println(strings.Repeat("-", 8) + "tryMem" + strings.Repeat("-", 8))
	fmt.Printf("os.Getpagesize(): %v\n", os.Getpagesize())
}

func tryProcess() {
	fmt.Println(strings.Repeat("-", 8) + "tryMem" + strings.Repeat("-", 8))
	fmt.Printf("os.Getpid(): %v\n", os.Getpid())
	fmt.Printf("os.Getppid(): %v\n", os.Getppid())
}

func tryKernel() {
	fmt.Println(strings.Repeat("-", 8) + "tryKernel" + strings.Repeat("-", 8))
	fmt.Println(os.Hostname())
}

func tryosReadDir() {
	fmt.Println(strings.Repeat("-", 8) + "tryosReadDir" + strings.Repeat("-", 8))
	files, err := os.ReadDir("/BlockChain/Algorithm")
	checkErr(err)
	for _, file := range files {
		fmt.Printf("file.Name(): %v\n", file.Name())
	}
}

func tryFile() {
	fmt.Println(strings.Repeat("-", 8) + "tryFile" + strings.Repeat("-", 8))
	tryCreateAndOpen()
	tryReadDir()
	tryRead()
}

func tryCreateAndOpen() {
	fmt.Println(strings.Repeat("-", 8) + "tryCreateAndOpen" + strings.Repeat("-", 8))
	f1, err := os.Create("temp1")
	checkErr(err)
	defer os.Remove(f1.Name())
	fmt.Printf("f1.Name(): %v\n", f1.Name())

	tryCreateTemp()

	f2 := os.NewFile(f1.Fd(), "f2")
	fmt.Printf("f2.Name(): %v\n", f2.Name())

	f3, err := os.Open("FileSystem/test.txt")
	checkErr(err)
	fmt.Printf("f3.Name(): %v\n", f3.Name())
	f3.Close()

	f4, err := os.OpenFile("FileSystem/test.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	checkErr(err)
	f4.Write([]byte("HSH insert"))
	content, err := os.ReadFile(f4.Name())
	checkErr(err)
	fmt.Printf("content: %q\n", content)
	re := regexp.MustCompile(`(?s)(?P<source>.*Job)(HSH.*)`)
	content = re.ReplaceAll(content, []byte("$source"))
	f4.Close()

	f5, err := os.OpenFile("FileSystem/test.txt", os.O_TRUNC|os.O_RDWR, 0644)
	checkErr(err)
	defer f5.Close()
	f5.Write(content)
	contentRestore, err := os.ReadFile(f5.Name())
	checkErr(err)
	fmt.Printf("contentRestore: %q\n", contentRestore)
}

func tryRead() {
	fmt.Println(strings.Repeat("-", 8) + "tryRead" + strings.Repeat("-", 8))
	content := make([]byte, 100)
	f, err := os.OpenFile("FileSystem/test.txt", os.O_RDWR|os.O_CREATE, 0644)
	checkErr(err)
	defer f.Close()
	n, err := f.Read(content)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("content[:n]: %q\n", content[:n])

	f.Seek(0, io.SeekStart)
	n, err = f.ReadAt(content, 4)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("content[:n]: %q\n", content[:n])

	f1, err := os.OpenFile("FileSystem/writer.txt", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	checkErr(err)
	defer f1.Close()

	_, err = f1.ReadFrom(f)
	checkErr(err)

	f1.Seek(0, io.SeekStart)
	n, err = f1.Read(content)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("content[:n]: %q\n", content[:n])
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
	//Readdir

	f.Seek(0, io.SeekStart)
	names, err := f.Readdirnames(-1)
	checkErr(err)
	for _, name := range names {
		fmt.Println(name)
	}
}
