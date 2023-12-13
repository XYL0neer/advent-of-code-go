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
	var blocks = strings.Split(input, "\n\n")
	var nextSources = readSeeds(blocks[0])
	for _, block := range blocks[1:] {
		fmt.Println(nextSources)
		var mappings = parseBlock(block)
		nextSources = determineMapping(mappings, nextSources)
	}
	fmt.Println(nextSources)
	fmt.Println(findSmallest(nextSources))
}
func findSmallest(in []int) int {
	smallest := in[0]

	for _, num := range in[1:] {
		if num < smallest {
			smallest = num
		}
	}

	return smallest
}

func determineMapping(mappings []string, sources []int) []int {
	var mapped = make(map[int]int)

	for _, s := range sources {
		mapped[s] = s
	}

	for _, mapping := range mappings {
		var m = strings.Split(mapping, " ")
		var destination, _ = strconv.Atoi(m[0])
		var source, _ = strconv.Atoi(m[1])
		var length, _ = strconv.Atoi(m[2])

		for _, s := range sources {
			if s >= source && s < source+length {
				var diff = s - source
				mapped[s] = destination + diff
			}
		}
	}

	var out = []int{}
	for _, v := range mapped {
		out = append(out, v)
	}

	return out
}

func parseBlock(blockStr string) []string {
	return strings.Split(blockStr, "\n")[1:]
}

func readSeeds(seedsStr string) []int {
	var seeds = strings.Split(seedsStr, " ")[1:]
	var s = make([]int, len(seeds))
	for i, str := range seeds {
		var v, _ = strconv.Atoi(str)
		s[i] = v
	}
	return s
}
