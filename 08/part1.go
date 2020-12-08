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
acc +6
`

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	//content = []byte(test)

	split := strings.Split(string(content), "\n")
	counter := []int{}
	i := 0
	acc := 0
	for {
		counter = append(counter, i)

		op := ""
		sign := 0
		value := 0
		fmt.Sscanf(split[i], "%s %c%d", &op, &sign, &value)
		if sign == '-' {
			value = value * -1
		}

		fmt.Printf("%s %c %d\n", op, sign, value)

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
	}

	fmt.Printf("Acc: %d\n", acc)
}

func Contains(a []int, i int) bool {
	for _, n := range a {
		if i == n {
			return true
		}
	}
	return false
}
