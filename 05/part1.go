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

	max_seat := 0
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

		if (rows_start*8)+cols_start > max_seat {
			max_seat = (rows_start * 8) + cols_start
		}
	}
	fmt.Printf("Max seat ID: %d\n", max_seat)
}
