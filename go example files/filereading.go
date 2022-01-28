package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func splitLines(text string) []string {
	return strings.Split(text, "\n")
}

func main() {
	content, _ := ioutil.ReadFile("../other/helloworld.txt")
	text := string(content)
	lines := splitLines(text)
	fmt.Printf("%v\n", lines)
}
