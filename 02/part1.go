package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("part1-input.txt")

	valid := 0
	split := strings.Split(string(content), "\n")
	for _, l := range split {
		data := strings.Split(l, " ")
		rules := strings.Split(data[0], "-")
		min, _ := strconv.Atoi(rules[0])
		max, _ := strconv.Atoi(rules[1])
		letter := data[1][:1]
		password := data[2]

		count := 0
		for _, c := range password {
			if string(c) == letter {
				count = count + 1
			}
		}

		if count >= min && count <= max {
			valid = valid + 1
		}
	}

	fmt.Printf("Valid: %d\n", valid)
}
