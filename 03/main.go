package main

import (
	_ "embed"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var lines = strings.Split(input, "\n")
	part1(lines)
}

func part1(lines []string) {
	var sum int
	for currentLine, line := range lines {
		var lineLength = len(line)
		var start, end, err = findNumber(line, 0)
		for err == nil {
			fmt.Printf("%d, %d - ", start, end)
			var numbers string
			if end == lineLength-1 {
				numbers = line[start:]
			} else {
				numbers = line[start : end+1]
			}
			fmt.Printf("%s ", numbers)
			if adjacentToSymbol(lines, currentLine, start, end) {
				fmt.Print("added")
				var val, _ = strconv.Atoi(numbers)
				sum += val
			}
			fmt.Println("")
			start, end, err = findNumber(line, end+1)
		}

	}
	fmt.Println(sum)
}

func findNumber(line string, startIndex int) (start, end int, err error) {
	start, end = -1, -1
	for i := startIndex; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			if start < 0 {
				start = i
				end = i
			} else {
				end = i
			}
		} else if start >= 0 {
			break
		}
	}
	if start >= 0 {
		return start, end, nil
	}
	return start, end, errors.New("no number remaining in line")
}

func adjacentToSymbol(lines []string, currentLine, start, end int) bool {
	var xMin = start - 1
	if xMin < 0 {
		xMin = 0
	}
	var xMax = end + 1
	if xMax >= len(lines[currentLine]) {
		xMax = end
	}

	if xMin != start && lines[currentLine][xMin] != '.' {
		return true
	}
	if xMax != end && lines[currentLine][xMax] != '.' {
		return true
	}

	for x := xMin; x <= xMax; x++ {
		if (currentLine-1 >= 0 && lines[currentLine-1][x] != '.') || (currentLine+1 < len(lines) && lines[currentLine+1][x] != '.') {
			return true
		}
	}
	return false
}
