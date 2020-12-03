package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("part1-input.txt")

	split := strings.Split(string(content), "\n")
	x := 0
	count := 0
	for i, l := range split {
		l = strings.Repeat(l, len(split))
		if i == 0 {
			continue
		}
		x = x + 3
		if string(l[x]) == "#" {
			count = count + 1
		}
	}
	fmt.Printf("Trees: %d\n", count)
}
