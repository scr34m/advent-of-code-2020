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
	//content = []byte(test)
	//scope = 5

	result := 0
	list := []int{}
	split := strings.Split(string(content), "\n")
	for _, l := range split {
		val, _ := strconv.Atoi(l)
		if len(list) < scope {
			list = append(list, val)
			continue
		}
		if !ContainsPair(list, val) {
			result = val
			break
		}

		for i := 0; i < scope-1; i++ {
			list[i] = list[i+1]
		}
		list[scope-1] = val
	}

	start := 0
	for {
		found := false
		min := -1
		max := -1
		sum := 0
		for i := start; i < len(split); i++ {
			val, _ := strconv.Atoi(split[i])
			if val == result {
				break
			}

			if min == -1 || min > val {
				min = val
			}

			if max == -1 || max < val {
				max = val
			}

			sum = sum + val
			if sum > result {
				break
			}

			if sum == result {
				found = true
				break
			}
		}

		if found {
			fmt.Printf("Found min: %d max: %d = %d\n", min, max, min+max)
			break
		}
		start = start + 1
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
