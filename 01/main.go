package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var lines = strings.Split(input, "\n")

	var numbers = []int{}

	for _, line := range lines {
		var first, last = -1, -1
		for i, v := range line {
			if v >= '0' && v <= '9' {
				if first < 0 {
					first = i
				}
				last = i
			}
		}
		var n, _ = strconv.Atoi(fmt.Sprintf("%c%c", line[first], line[last]))
		numbers = append(numbers, n)
	}

	var sum = 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Println(sum)
}
