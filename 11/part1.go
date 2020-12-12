package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"strings"
)

var test = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	//content = []byte(test)

	grid := [][]byte{}

	split := strings.Split(string(content), "\n")
	for _, l := range split {
		line := []byte{}
		for i := 0; i < len(l); i++ {
			line = append(line, l[i])
		}
		grid = append(grid, line)
	}

	hash := ""
	result := 0
	for {
		grid2 := Rule1(grid)
		grid3 := Rule2(grid2)

		result = 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				grid[i][j] = grid3[i][j]
				if grid[i][j] == '#' {
					result += 1
				}
			}
		}

		new_hash := CalculateChecksum(grid)
		if new_hash == hash {
			break
		} else {
			hash = new_hash
		}
	}

	fmt.Printf("Result: %d\n", result)
}

func CalculateChecksum(grid [][]byte) string {
	algorithm := md5.New()
	for i := range grid {
		algorithm.Write(grid[i])
	}
	return fmt.Sprintf("%x", algorithm.Sum(nil))
}

func Rule1(grid [][]byte) [][]byte {
	grid_new := make([][]byte, len(grid))
	for i := range grid {
		grid_new[i] = make([]byte, len(grid[i]))
		copy(grid_new[i], grid[i])
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid_new[i][j] = grid[i][j]

			if grid[i][j] == 'L' && SeatsAdjacent(grid, i, j) == 0 {
				grid_new[i][j] = '#'
			}
		}
	}

	return grid_new
}

func Rule2(grid [][]byte) [][]byte {
	grid_new := make([][]byte, len(grid))
	for i := range grid {
		grid_new[i] = make([]byte, len(grid[i]))
		copy(grid_new[i], grid[i])
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid_new[i][j] = grid[i][j]

			if grid[i][j] == '#' && SeatsAdjacent(grid, i, j) >= 4 {
				grid_new[i][j] = 'L'
			}
		}
	}

	return grid_new
}

func SeatsAdjacent(grid [][]byte, i, j int) int {
	found := 0
	for r := -1; r <= 1; r++ {

		if i+r < 0 || i+r == len(grid[i]) {
			continue
		}

		if j > 0 {
			if grid[i+r][j-1] == '#' {
				found += 1
			}
		}
		if r != 0 {
			if grid[i+r][j] == '#' {
				found += 1
			}
		}
		if j < len(grid[i+r])-1 {
			if grid[i+r][j+1] == '#' {
				found += 1
			}
		}

	}

	return found
}

func Draw(grid [][]byte) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Printf("\n")
	}
}
