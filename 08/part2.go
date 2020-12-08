package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var test = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

var tests []int

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	// content = []byte(test)

	split := strings.Split(string(content), "\n")

	acc := 0
	for i := 0; i < len(split); i++ {
		acc = Run(split)
		if acc != -1 {
			fmt.Printf("Acc: %d\n", acc)
			break
		}
	}
}

func Run(split []string) int {
	counter := []int{}
	i := 0
	acc := 0
	changed := false
	for {
		counter = append(counter, i)

		op := ""
		sign := 0
		value := 0
		fmt.Sscanf(split[i], "%s %c%d", &op, &sign, &value)
		if sign == '-' {
			value = value * -1
		}

		if !changed && !Contains(tests, i) && value != 0 {
			if op == "jmp" {
				changed = true
				op = "nop"
				tests = append(tests, i)
			} else if op == "nop" {
				changed = true
				op = "jmp"
				tests = append(tests, i)
			}
		}

		switch op {
		case "nop":
			i = i + 1
		case "jmp":
			i = i + value
		case "acc":
			acc = acc + value
			i = i + 1
		}

		if Contains(counter, i) {
			break
		}

		if i == len(split) {
			return acc
		}
	}
	return -1
}

func Contains(a []int, i int) bool {
	for _, n := range a {
		if i == n {
			return true
		}
	}
	return false
}
