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
		pos1, _ := strconv.Atoi(rules[0])
		pos2, _ := strconv.Atoi(rules[1])
		letter := string(data[1][:1])
		password := data[2]

		if len(password) < pos2 {
			fmt.Printf("pos2 bigger than length of the password %d %s\n", pos2, password)
			break
		}

		if password[pos1-1:pos1] == letter && password[pos2-1:pos2] == letter {
			// invalid, both position contains letter
		} else if password[pos1-1:pos1] != letter && password[pos2-1:pos2] != letter {
			// invalid, neither position contains letter
		} else {
			valid = valid + 1
		}
	}

	fmt.Printf("Valid: %d\n", valid)
}
