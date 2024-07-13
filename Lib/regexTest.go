package main

import (
	"fmt"
	"regexp"
	"strings"
)

func regexTest() {
	tryGroup()
	tryMatch()
	tryMatchReader()
	tryMatchString()
	tryQuoteMeta()
	tryms()
	tryRegexp()
}

func tryGroup() {
	fmt.Println(strings.Repeat("-", 8) + "tryGroup" + strings.Repeat("-", 8))
	str := "John Doe, jane@example.com"
	re := regexp.MustCompile(`(\w+)\s(\w+),\s(\w+@\w+\.\w+)`)

	match := re.FindStringSubmatch(str)
	fmt.Println(len(match))
	fmt.Println(match)                       // [John Doe, John Doe, jane@example.com]
	fmt.Println("Name:", match[1], match[2]) // Name: John Doe
	fmt.Println("Email:", match[3])          // Email: jane@example.com
	fmt.Println("All: ", match[0])
}

func tryMatch() {
	fmt.Println(strings.Repeat("-", 8) + "tryMatch" + strings.Repeat("-", 8))
	matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
	fmt.Println(matched, err)
	matched, err = regexp.Match(`bar.*`, []byte(`seafood`))
	fmt.Println(matched, err)
}

func tryMatchReader() {
	fmt.Println(strings.Repeat("-", 8) + "tryMatchReader" + strings.Repeat("-", 8))
	r := strings.NewReader("abcd12Am.com")
	b, err := regexp.MatchReader(`\w+\.\w+`, r)
	checkErr(err)
	fmt.Println("T or F", b)
}

func tryMatchString() {
	fmt.Println(strings.Repeat("-", 8) + "tryMatchString" + strings.Repeat("-", 8))
	matched, err := regexp.MatchString(`foo.*`, "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString(`bar.*`, "seafood")
	fmt.Println(matched, err)
}

func tryQuoteMeta() {
	fmt.Println(strings.Repeat("-", 8) + "tryQuoteMeta" + strings.Repeat("-", 8))
	fmt.Println(regexp.QuoteMeta(`Escaping symbols like: .+*?()|[]{}^$`))
}

func tryms() {
	fmt.Println(strings.Repeat("-", 8) + "try ?m ?s" + strings.Repeat("-", 8))
	res := []*regexp.Regexp{regexp.MustCompile(`^[A-Z]{3}`),
		regexp.MustCompile(`(?m)^[A-Z]{3}`), regexp.MustCompile(`(?ms)^[A-Z]{3}.`)}
	for _, re := range res {
		matchs := re.FindAllStringSubmatch("ABC DEF\nGHI\n", -1)
		fmt.Println(matchs)
	}
}

func tryRegexp() {
	//	找共同前缀
	tryLiteralPrefix()
	// 设置最长匹配
	tryLongest()
	//关于正则表达式
	tryMarshalText()
	//判断是否匹配
	tryReMatch()
	//分割
	trySplit()
	//查找匹配项
	tryFind()
	// 查找匹配项和Submatch
	trySubmatch()
	tryNumSubexp()
	trySubexp()
	// 用于将 匹配项 of src->dst, 并且更改格式
	tryExpand()
	tryExpandString()
	// 更改本身的匹配项
	tryReplace()
}

func tryExpand() {
	fmt.Println(strings.Repeat("-", 8) + "tryExpand" + strings.Repeat("-", 8))
	content := []byte(`
	# comment line
	# option1: value1
	option2: value2

	# another comment line
	option3: value3
	`)
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	template := []byte("$key=$value\n")
	result := []byte{}

	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		// Apply the captured submatches to the template and append the output
		// to the result.
		result = pattern.Expand(result, template, content, submatches)
	}
	fmt.Println(string(result))
}

func tryExpandString() {
	fmt.Println(strings.Repeat("-", 8) + "tryExpandString" + strings.Repeat("-", 8))
	content := `
	# comment line
	# option1: value1
	option2: value2

	# another comment line
	option3: value3
	`
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?<value>\w+)$`)
	templateString := "$key=$2\n"
	result := []byte{}

	for _, submatches := range pattern.FindAllStringSubmatchIndex(content, -1) {
		result = pattern.ExpandString(result, templateString, content, submatches)
	}
	fmt.Println(string(result))
}

func tryFind() {
	fmt.Println(strings.Repeat("-", 8) + "tryFind" + strings.Repeat("-", 8))
	re1 := regexp.MustCompile("foo.?")
	data1 := `seafood fool`

	f1 := re1.Find([]byte(data1))
	fmt.Println(string(f1))

	fmt.Println(strings.Repeat("-", 8) + "tryFindIndex" + strings.Repeat("-", 8))
	f6 := re1.FindIndex([]byte(data1))
	fmt.Println(f6)

	fmt.Println(strings.Repeat("-", 8) + "tryFindAll" + strings.Repeat("-", 8))
	f2 := re1.FindAll([]byte(data1), -1)
	fmt.Printf("%q\n", f2)

	fmt.Println(strings.Repeat("-", 8) + "tryFindAllIndex" + strings.Repeat("-", 8))
	f3 := re1.FindAllIndex([]byte(data1), -1)
	fmt.Println(f3)

	fmt.Println(strings.Repeat("-", 8) + "tryFindString" + strings.Repeat("-", 8))
	f8 := re1.FindString(data1)
	fmt.Println(f8)

	fmt.Println(strings.Repeat("-", 8) + "tryFindStringIndex" + strings.Repeat("-", 8))
	f9 := re1.FindStringIndex(data1)
	fmt.Println(f9)

	fmt.Println(strings.Repeat("-", 8) + "tryFindAllString" + strings.Repeat("-", 8))
	f4 := re1.FindAllString(data1, -1)
	fmt.Println(f4)

	fmt.Println(strings.Repeat("-", 8) + "tryFindAllStringIndex" + strings.Repeat("-", 8))
	f5 := re1.FindAllStringIndex(data1, -1)
	fmt.Println(f5)

	fmt.Println(strings.Repeat("-", 8) + "tryFindReaderIndex" + strings.Repeat("-", 8))
	f7 := re1.FindReaderIndex(strings.NewReader(data1))
	fmt.Println(f7)
}

func trySubmatch() {
	re2 := regexp.MustCompile(`a(x*)b`)
	data2 := "-axxb-ab-"

	fmt.Println(strings.Repeat("-", 8) + "tryFindSubmatch" + strings.Repeat("-", 8))
	f13 := re2.FindSubmatch([]byte(data2))
	fmt.Printf("%q\n", f13)

	fmt.Println(strings.Repeat("-", 8) + "tryFindSubmatchIndex" + strings.Repeat("-", 8))
	f14 := re2.FindSubmatchIndex([]byte(data2))
	fmt.Println(f14)

	fmt.Println(strings.Repeat("-", 8) + "tryFindStringSubmatch" + strings.Repeat("-", 8))
	f11 := re2.FindStringSubmatch(data2)
	fmt.Println(f11)

	fmt.Println(strings.Repeat("-", 8) + "tryFindStringSubmatchIndex" + strings.Repeat("-", 8))
	f12 := re2.FindStringSubmatchIndex(data2)
	fmt.Println(f12)

	fmt.Println(strings.Repeat("-", 8) + "tryFindAllSubmatch" + strings.Repeat("-", 8))
	f8 := re2.FindAllSubmatch([]byte(data2), -1)
	fmt.Printf("%q\n", f8)

	fmt.Println(strings.Repeat("-", 8) + "tryFindAllSubmatchIndex" + strings.Repeat("-", 8))
	f9 := re2.FindAllSubmatchIndex([]byte(data2), -1)
	fmt.Println(f9)

	fmt.Println(strings.Repeat("-", 8) + "tryFindAllStringSubmatch" + strings.Repeat("-", 8))
	f6 := re2.FindAllStringSubmatch(data2, -1)
	fmt.Println(f6, len(f6))

	fmt.Println(strings.Repeat("-", 8) + "tryFindAllStringSubmatchIndex" + strings.Repeat("-", 8))
	f7 := re2.FindAllStringSubmatchIndex(data2, -1)
	fmt.Println(f7)

	fmt.Println(strings.Repeat("-", 8) + "tryFindReaderSubmatchIndex" + strings.Repeat("-", 8))
	f10 := re2.FindReaderSubmatchIndex(strings.NewReader(data2))
	fmt.Println(f10)
}

func tryNumSubexp() {
	fmt.Println(strings.Repeat("-", 8) + "tryNumSubexp" + strings.Repeat("-", 8))
	re := regexp.MustCompile(`(.*)((a)b)(.*)a`)
	fmt.Println(re.NumSubexp())
	fmt.Println(re.FindAllStringSubmatch("xxxabjjjja", -1))
}

func trySubexp() {
	fmt.Println(strings.Repeat("-", 8) + "trySubexp" + strings.Repeat("-", 8))
	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	fmt.Println(re.MatchString("Alan Turing"))
	matches := re.FindStringSubmatch("Alan Turing")
	lastIndex := re.SubexpIndex("last")
	fmt.Printf("last => %d\n", lastIndex)
	fmt.Println(matches[lastIndex])

	fmt.Printf("%q\n", re.SubexpNames())
	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	fmt.Println(reversed)
	fmt.Println(re.ReplaceAllString("Alan Turing", reversed))
}

func tryLiteralPrefix() {
	// 就是看看所有匹配的string所共有的前缀
	fmt.Println(strings.Repeat("-", 8) + "tryLiteralPrefix" + strings.Repeat("-", 8))
	re := regexp.MustCompile(`_a`)
	s, b := re.LiteralPrefix()
	fmt.Println(s, b)
}

func tryLongest() {
	fmt.Println(strings.Repeat("-", 8) + "tryLongest" + strings.Repeat("-", 8))
	re := regexp.MustCompile(`a(|b)`)
	fmt.Println(re.FindString("ab"))
	re.Longest()
	fmt.Println(re.FindString("ab"))
}

func tryMarshalText() {
	fmt.Println(strings.Repeat("-", 8) + "tryMarshalText" + strings.Repeat("-", 8))
	re := regexp.MustCompile(`(\w+)\s+(\w+)\.(\w+)`)
	s, err := re.MarshalText()
	checkErr(err)
	fmt.Printf("%q\n", s)
	fmt.Println(re.String())

	err = re.UnmarshalText([]byte(`ab*`))
	checkErr(err)

	s, err = re.MarshalText()
	checkErr(err)
	fmt.Printf("%q\n", s)

	fs := re.FindString("ab*\nadfsa")
	fmt.Println(fs)

	fmt.Println(re.String())
}

func tryReMatch() {
	fmt.Println(strings.Repeat("-", 8) + "tryReMatch" + strings.Repeat("-", 8))
	re := regexp.MustCompile(`\w+@\w+`)
	data := "abc@com"
	fmt.Println(re.Match([]byte(data)))
	fmt.Println(re.MatchReader(strings.NewReader(data)))
	fmt.Println(re.MatchString(data))
}

func tryReplace() {
	fmt.Println(strings.Repeat("-", 8) + "tryReplace" + strings.Repeat("-", 8))
	re2 := regexp.MustCompile(`a(?P<1W>x*)b`)
	fmt.Printf("%s\n", re2.ReplaceAll([]byte("-ab-axxb-"), []byte("$1W")))
	fmt.Printf("%s\n", re2.ReplaceAll([]byte("-ab-axxb-"), []byte("${1}W")))

	fmt.Printf("%s\n", re2.ReplaceAllFunc([]byte("-ab-axxb-"), func(match []byte) []byte {
		var b []byte
		b = append(b, "^^"...)
		b = append(b, match...)
		b = append(b, "^^"...)
		return b
	}))

	fmt.Printf("%s\n", re2.ReplaceAllLiteral([]byte("-ab-axxb-"), []byte("${1}W")))

	fmt.Printf("%s\n", re2.ReplaceAllLiteralString("-ab-axxb-", "${1}W"))

	fmt.Printf("%s\n", re2.ReplaceAllString("-ab-axxb-", "${1}W"))

	fmt.Printf("%s\n", re2.ReplaceAllStringFunc("-ab-axxb-", func(match string) string {
		return "@@" + match + "@@"
	}))
}

func trySplit() {
	fmt.Println(strings.Repeat("-", 8) + "trySplit" + strings.Repeat("-", 8))
	re := regexp.MustCompile(`w`)
	s := re.Split("wawwaaaw", -1)
	fmt.Println(s, len(s))
}
