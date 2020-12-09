package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var test = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

var scope = 25

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	// content = []byte(test)
	// scope = 5

	list := []int{}
	split := strings.Split(string(content), "\n")
	for _, l := range split {
		val, _ := strconv.Atoi(l)
		if len(list) < scope {
			list = append(list, val)
			continue
		}
		if !ContainsPair(list, val) {
			fmt.Printf("%d\n", val)
			break
		}

		for i := 0; i < scope-1; i++ {
			list[i] = list[i+1]
		}
		list[scope-1] = val

	}
}

func ContainsPair(a []int, n int) bool {
	for i := 0; i < scope-1; i++ {
		for j := 1; j < scope; j++ {
			if a[i]+a[j] == n {
				return true
			}
		}
	}
	return false
}
