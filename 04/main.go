package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const winningStartIndex = 10
const winningEnd = 39
const drawnStartIndex = 42

func main() {
	var solution int
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		var winningNumbers, drawnNumbers = parseGame(line)
		var count = countWins(winningNumbers, drawnNumbers)
		fmt.Println(count)
		if count > 0 {
			solution += 1 << (count - 1)
		}
	}
	fmt.Println(solution)
}

func parseGame(game string) (winningNumbers, drawnNumbers []int) {
	var winning = game[winningStartIndex:winningEnd]
	var drawn = game[drawnStartIndex:]
	winningNumbers = splitAndCastToInt(winning)
	drawnNumbers = splitAndCastToInt(drawn)
	return winningNumbers, drawnNumbers
}

func splitAndCastToInt(s string) []int {
	var numbers = []int{}
	var trimmed = strings.Join(strings.Fields(s), " ")
	var nStr = strings.Split(trimmed, " ")
	for _, n := range nStr {
		var num, _ = strconv.Atoi(n)
		numbers = append(numbers, num)
	}
	return numbers
}

func countWins(wins, draws []int) int {
	var count int
	for _, w := range wins {
		if slices.Contains(draws, w) {
			count++
		}
	}
	return count
}
