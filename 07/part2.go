package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var test = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`

type SmallBag struct {
	Name     string
	Quantity int
}

type BigBag struct {
	Bags []SmallBag
}

var bags map[string]BigBag

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	//content = []byte(test)

	bags = map[string]BigBag{}

	split := strings.Split(string(content), "\n")
	for _, l := range split {
		if len(l) == 0 {
			continue
		}

		b := strings.Split(strings.Trim(l, "."), " bags contain ")

		bag := BigBag{}
		if b[1] != "no other bags" {
			for _, sb := range strings.Split(b[1], ", ") {
				sf := strings.Index(sb, " ")
				sl := strings.LastIndex(sb, " ")
				q, _ := strconv.Atoi(string(sb[0:sf]))
				sbag := SmallBag{Name: string(sb[sf+1 : sl]), Quantity: q}
				bag.Bags = append(bag.Bags, sbag)
			}
		}

		bags[b[0]] = bag
	}

	result := CountNested(bags["shiny gold"].Bags)
	fmt.Printf("Result: %d\n", result)
}

func CountNested(smallBags []SmallBag) int {
	found := 0
	for _, bag := range smallBags {
		found = found + bag.Quantity + (bag.Quantity * CountNested(bags[bag.Name].Bags))
	}
	return found
}
