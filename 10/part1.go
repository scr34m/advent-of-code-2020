package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var test = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	// content = []byte(test)

	list := []int{0}
	split := strings.Split(string(content), "\n")
	for _, l := range split {
		val, _ := strconv.Atoi(l)
		list = append(list, val)
	}
	sort.Ints(list)

	list = append(list, list[len(list)-1]+3)

	joltDiffs := []int{0, 0, 0, 0}
	for i := 1; i < len(list); i++ {
		diff := list[i] - list[i-1]
		joltDiffs[diff] += 1
	}
	fmt.Printf("Result: %d\n", joltDiffs[1]*joltDiffs[3])
}
