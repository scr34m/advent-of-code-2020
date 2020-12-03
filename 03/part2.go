package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("part1-input.txt")

	split := strings.Split(string(content), "\n")
	l := len(split)

	count1 := 0
	x := 0
	for i := 1; i < l; i++ {
		r := strings.Repeat(split[i], l)
		x = x + 1
		if string(r[x]) == "#" {
			count1 = count1 + 1
		}
	}
	fmt.Printf("Trees 1-1: %d\n", count1)

	count2 := 0
	x = 0
	for i := 1; i < l; i++ {
		r := strings.Repeat(split[i], l)
		x = x + 3
		if string(r[x]) == "#" {
			count2 = count2 + 1
		}
	}
	fmt.Printf("Trees 3-1: %d\n", count2)

	count3 := 0
	x = 0
	for i := 1; i < l; i++ {
		r := strings.Repeat(split[i], l)
		x = x + 5
		if string(r[x]) == "#" {
			count3 = count3 + 1
		}
	}
	fmt.Printf("Trees 5-1: %d\n", count3)

	count4 := 0
	x = 0
	for i := 1; i < l; i++ {
		r := strings.Repeat(split[i], l)
		x = x + 7
		if string(r[x]) == "#" {
			count4 = count4 + 1
		}
	}
	fmt.Printf("Trees 7-1: %d\n", count4)

	count5 := 0
	x = 0
	for i := 2; i < l; i += 2 {
		r := strings.Repeat(split[i], l)
		x = x + 1
		if string(r[x]) == "#" {
			count5 = count5 + 1
		}
	}
	fmt.Printf("Trees 1-2: %d\n", count5)

	fmt.Printf("Trees: %d\n", count1*count2*count3*count4*count5)
}
