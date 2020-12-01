package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./part1-input.txt")
	if err != nil {
		panic(err)
	}

	var numbers []int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		numbers = append(numbers, i)
	}

	for _, n1 := range numbers {
		for _, n2 := range numbers {
			if n1+n2 == 2020 {
				fmt.Println(n1 * n2)
			}
		}
	}

}
