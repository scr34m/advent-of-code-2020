package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var test = `FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	// content = []byte(test)

	seats := []int{}
	split := strings.Split(string(content), "\n")
	for _, l := range split {

		rows_start := 0
		rows_end := 127

		cols_start := 0
		cols_end := 7

		for i := 0; i < 10; i++ {
			switch l[i] {
			case 'F':
				rows_start = rows_start
				rows_end = rows_end - (((rows_end - rows_start) / 2) + 1) // +1 for ceiling fix
			case 'B':
				rows_start = rows_end - ((rows_end - rows_start) / 2)
				rows_end = rows_end
			case 'R':
				cols_start = cols_end - ((cols_end - cols_start) / 2)
				cols_end = cols_end
			case 'L':
				cols_end = cols_end - (((cols_end - cols_start) / 2) + 1) // +1 for ceiling fix
				cols_start = cols_start
			}
		}

		seats = append(seats, (rows_start*8)+cols_start)
	}

	for i := 1; i < ((127*8)+5)-1; i++ {
		if !Contains(seats, i) && Contains(seats, i-1) && Contains(seats, i+1) {
			fmt.Printf("Found: %d\n", i)
		}
	}
}

func Contains(a []int, i int) bool {
	for _, n := range a {
		if i == n {
			return true
		}
	}
	return false
}
