package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var test = `abc

a
b
c

ab
ac

a
a
a
a

b
`

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	// content = []byte(test)

	result := 0
	group := ""
	split := strings.Split(string(content)+"\n", "\n")
	for _, l := range split {
		if len(l) == 0 {
			result = result + CountUnique(group)
			group = ""
		} else {
			group = group + l
		}
	}
	fmt.Printf("Result %d\n", result)
}

func CountUnique(s string) int {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] = m[s[i]] + 1
		} else {
			m[s[i]] = 1
		}
	}
	return len(m)
}
