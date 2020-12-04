package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	passport := ""
	valid := 0
	split := strings.Split(string(content)+"\n", "\n")
	for _, l := range split {
		if len(passport) > 0 {
			passport = passport + " "
		}
		passport = passport + l
		if len(l) == 0 {
			missing_field := false
			for _, f := range fields {
				if !strings.Contains(passport, f) {
					missing_field = true
					break
				}
			}

			if !missing_field {
				valid = valid + 1
			}
			passport = ""
		}
	}
	fmt.Printf("valid: %d\n", valid)
}
