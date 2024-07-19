package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func jsonTest() {
	tryMarshal()
	tryCompact()
	tryHTMLEscape()
	tryIndent()
	tryMarshalIndent()
	tryValid()
	tryDecoder()
	tryEncoder()
}

func tryMarshal() {
	fmt.Println(strings.Repeat("-", 8) + "tryMarshal" + strings.Repeat("-", 8))
	type Message struct {
		Name string `json:"message_name"`
		Body string `json:"-"`
		Time int64  `json:"message_time,omitempty"`
	}
	m := Message{"Alice", "Hello", 1294706395881547000}
	// m := Message{
	// 	Name: "Alice",
	// 	Body: "Hello",
	// }
	b, err := json.Marshal(m)
	checkErr(err)
	fmt.Printf("b: %s\n", b)

	var m1 Message
	err = json.Unmarshal(b, &m1)
	checkErr(err)
	fmt.Printf("m1: %v\n", m1)

	b = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f any
	err = json.Unmarshal(b, &f)
	checkErr(err)
	m2 := f.(map[string]any)
	fmt.Printf("m2: %v\n", f)
	for k, v := range m2 {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []any:
			fmt.Println(k, "is an array")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println("an unexpected type")
		}
	}
}

func tryCompact() {
	fmt.Println(strings.Repeat("-", 8) + "tryCompact" + strings.Repeat("-", 8))

	dst := bytes.NewBuffer([]byte{})

	const src = `{
			  "release_date": "2004-11-09",
			  "status": "retired",
			  "engine": "Gecko",
			  "engine_version": "1.7"
			}`
	if err := json.Compact(dst, []byte(src)); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dst: %v\n", dst)
}

func tryHTMLEscape() {
	fmt.Println(strings.Repeat("-", 8) + "tryHTMLEscape" + strings.Repeat("-", 8))
	var out bytes.Buffer
	json.HTMLEscape(&out, []byte(`{"Name":"<b>HTML content</b>"}`))
	out.WriteTo(os.Stdout)
	fmt.Println()
}

func tryIndent() {
	fmt.Println(strings.Repeat("-", 8) + "tryIndent" + strings.Repeat("-", 8))
	type Road struct {
		Name   string
		Number int
	}
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("b: %s\n", b)
	var out bytes.Buffer
	json.Indent(&out, b, "=", "\t")
	out.WriteTo(os.Stdout)
}

func tryMarshalIndent() {
	fmt.Println(strings.Repeat("-", 8) + "tryMarshalIndent" + strings.Repeat("-", 8))
	data := map[string]int{
		"a": 1,
		"b": 2,
	}

	b, err := json.MarshalIndent(data, "<prefix>", "<indent>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

func tryValid() {
	fmt.Println(strings.Repeat("-", 8) + "tryValid" + strings.Repeat("-", 8))
	goodJSON := `{"example": 1}`
	badJSON := `{"example":2:]}}`

	fmt.Println(json.Valid([]byte(goodJSON)), json.Valid([]byte(badJSON)))
}

func tryDecoder() {
	fmt.Println(strings.Repeat("-", 8) + "tryDecoder" + strings.Repeat("-", 8))
	const jsonStream = `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v----%v\n", m.Name, m.Text, dec.InputOffset())
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

}

func tryEncoder() {
	fmt.Println(strings.Repeat("-", 8) + "tryEncoder" + strings.Repeat("-", 8))
	type Student struct {
		Name  string
		Class string
		Id    int
	}
	s := Student{
		Name:  "<HSH>",
		Class: "cs1",
		Id:    1234,
	}
	buf := bytes.NewBuffer([]byte{})
	en := json.NewEncoder(buf)
	en.SetEscapeHTML(true)
	en.SetIndent("<pre>", "<idt>")
	en.Encode(s)
	fmt.Printf("buf: %s\n", buf)
}
