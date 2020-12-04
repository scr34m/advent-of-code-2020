package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")

	passport := map[string]string{}

	valid := 0
	split := strings.Split(string(content)+"\n", "\n")
	for _, l := range split {
		r := strings.Split(l, " ")
		for _, kv := range r {
			var key, value string
			fmt.Sscanf(kv, "%3s:%s", &key, &value)
			if key != "cid" {
				passport[key] = value
			}
		}

		if len(l) == 0 {
			invalid_field := false
			if val, ok := passport["byr"]; ok {
				d, _ := strconv.Atoi(val)
				if d < 1920 || d > 2002 {
					invalid_field = true
				}
			} else {
				invalid_field = true
			}

			if val, ok := passport["iyr"]; ok {
				d, _ := strconv.Atoi(val)
				if d < 2010 || d > 2020 {
					invalid_field = true
				}
			} else {
				invalid_field = true
			}

			if val, ok := passport["eyr"]; ok {
				d, _ := strconv.Atoi(val)
				if d < 2020 || d > 2030 {
					invalid_field = true
				}
			} else {
				invalid_field = true
			}

			if val, ok := passport["hgt"]; ok {
				var d int
				var t string
				_, err := fmt.Sscanf(val, "%d%s", &d, &t)
				if err == nil {
					if t == "cm" && d >= 150 && d <= 193 {
						// OK
					} else if t == "in" && d >= 59 && d <= 76 {
						// OK
					} else {
						invalid_field = true
					}
				} else {
					invalid_field = true
				}
			} else {
				invalid_field = true
			}

			if val, ok := passport["hcl"]; ok {
				var r, g, b string
				_, err := fmt.Sscanf(val, "#%02x%02x%02x", &r, &g, &b)
				if err != nil {
					invalid_field = true
				}
			} else {
				invalid_field = true
			}

			if val, ok := passport["ecl"]; ok {
				if val != "amb" && val != "blu" && val != "brn" && val != "gry" && val != "grn" && val != "hzl" && val != "oth" {
					invalid_field = true
				}
			} else {
				invalid_field = true
			}

			if val, ok := passport["pid"]; ok {
				if len(val) != 9 {
					invalid_field = true
				} else {
					_, err := strconv.Atoi(val)
					if err != nil {
						invalid_field = true
					}
				}
			} else {
				invalid_field = true
			}

			if !invalid_field {
				valid = valid + 1
			}

			passport = map[string]string{}
		}
	}
	fmt.Printf("valid: %d\n", valid)
}
