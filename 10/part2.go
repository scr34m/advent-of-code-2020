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

var level = 0

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	//content = []byte(test)

	list := []int{0}
	split := strings.Split(string(content), "\n")
	for _, l := range split {
		val, _ := strconv.Atoi(l)
		list = append(list, val)
	}
	sort.Ints(list)

	list = append(list, list[len(list)-1]+3)

	fmt.Printf("%d\n", Try(list)+1)
}

func Try(list []int) int {
	c := 0
	s := len(list) - 1
	for i := 1; i < s; i++ {
		jmax := s
		if i+3 < jmax {
			jmax = i + 3
		}
		for j := i + 1; j < jmax; j++ {
			p := j - i
			if list[i+p]-list[i-1] <= 3 {
				c += 1 + Try(list[i+p:])
			}
		}
	}
	return c
}
