package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strings"
)

func fsTest() {
	tryWalkDir()
	tryFormatFileInfor()
	tryGlob()
	tryReadFile()
	tryValidPath()
	tryFormatDirEntry()

	tryDirEntry()
	tryFS()
}

func tryFormatDirEntry() {
	fmt.Println(strings.Repeat("-", 8) + "tryFormatDirEntry" + strings.Repeat("-", 8))

	filesys := os.DirFS("/BlockChain/Algorithm")
	fileinfo, err := fs.Stat(filesys, "Lib")
	checkErr(err)
	direntry := fs.FileInfoToDirEntry(fileinfo)
	fmt.Println(fs.FormatDirEntry(direntry))
}

func tryFormatFileInfor() {
	fmt.Println(strings.Repeat("-", 8) + "tryFormatFileInfor" + strings.Repeat("-", 8))

	root := "/BlockChain/Algorithm"
	fileSystem := os.DirFS(root)
	fileInfo, err := fs.Stat(fileSystem, "Lib/run.sh")

	checkErr(err)
	fmt.Println(fs.FormatFileInfo(fileInfo))
	fmt.Println(fileInfo)
}

func tryGlob() {
	fmt.Println(strings.Repeat("-", 8) + "tryGlob" + strings.Repeat("-", 8))
	fileSystem := os.DirFS("/BlockChain/Algorithm")
	matchs, err := fs.Glob(fileSystem, `L*/*`)
	checkErr(err)
	fmt.Printf("%q\n", matchs)
	fmt.Printf("%s\n", matchs)
}

func tryReadFile() {
	fmt.Println(strings.Repeat("-", 8) + "tryReadFile" + strings.Repeat("-", 8))
	fileSystem := os.DirFS("/BlockChain/Algorithm")
	names, err := fs.Glob(fileSystem, `Lib/*.sh`)
	checkErr(err)
	for _, name := range names {
		content, err := fs.ReadFile(fileSystem, name)
		checkErr(err)
		fmt.Printf("%s\n", content)
	}
}

func tryValidPath() {
	//注意不能有"/"开头，不能有".." 或者 ".“
	fmt.Println(strings.Repeat("-", 8) + "tryValidPath" + strings.Repeat("-", 8))
	fmt.Println(fs.ValidPath("BlockChain/Algorithm/Lib/run.sh"))
	fmt.Println(fs.ValidPath("/BlockChain/Algorithm/Lib/run.sh"))
}

func tryWalkDir() {
	fmt.Println(strings.Repeat("-", 8) + "tryWalkDir" + strings.Repeat("-", 8))
	root := "/BlockChain/Algorithm"
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if ok, _ := regexp.MatchString(`L.*`, path); ok {
			fmt.Println("path:", path)
			fmt.Println("DirEntry: ", fs.FormatDirEntry(d))
		}
		return nil
	})
}

func tryDirEntry() {
	fmt.Println(strings.Repeat("-", 8) + "tryDirEntry" + strings.Repeat("-", 8))
	tryFileInfoToDirEntry()
	tryfsReadDir()
}

func tryFileInfoToDirEntry() {
	fmt.Println(strings.Repeat("-", 8) + "tryFileInfoToDirEntry" + strings.Repeat("-", 8))
	fsys := os.DirFS("/BlockChain/Algorithm")
	finfo, err := fs.Stat(fsys, "Lib/run.sh")
	checkErr(err)
	dEntry := fs.FileInfoToDirEntry(finfo)
	i1, err := dEntry.Info()
	checkErr(err)
	fmt.Println(dEntry.Name(), dEntry.IsDir(), dEntry.Type().String())
	fmt.Println(fs.FormatFileInfo(i1))
}

func tryfsReadDir() {
	fmt.Println(strings.Repeat("-", 8) + "tryfsReadDir" + strings.Repeat("-", 8))
	fsys := os.DirFS("/BlockChain/Algorithm")
	dirEntrys, err := fs.ReadDir(fsys, "Lib")
	checkErr(err)
	for _, de := range dirEntrys {
		fmt.Println(de)
	}
}

func tryFS() {
	fmt.Println(strings.Repeat("-", 8) + "tryFS" + strings.Repeat("-", 8))
	fsys := os.DirFS("/BlockChain/Algorithm")
	file, err := fsys.Open("Lib/run.sh")
	checkErr(err)
	defer file.Close()

	buf := make([]byte, 20)
	n, err := file.Read(buf)
	checkErr(err)
	fmt.Printf("%q\n", buf[:n])

	subfsys, err := fs.Sub(fsys, "Lib")
	checkErr(err)
	subfile, err := subfsys.Open("run.sh")
	checkErr(err)
	defer subfile.Close()

	n, err = subfile.Read(buf)
	checkErr(err)
	fmt.Printf("%q\n", buf[:n])
}
