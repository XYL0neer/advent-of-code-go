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
	var games = strings.Split(input, "\n")

	var sum int

	for _, game := range games {
		var gameSets = strings.Split(game, ":")[1]
		// if gameIsValid(gameSets) {
		// 	sum += i + 1
		// }
		sum += gameScore(gameSets)
	}

	fmt.Println(sum)
}

// function for part 1
func gameIsValid(game string) bool {
	var sets = strings.Split(game, ";")

	for _, set := range sets {
		var draws = strings.Split(set, ",")

		for _, draw := range draws {
			var d = strings.Split(draw, " ")
			var amtStr = d[1]
			var color = d[2]

			amt, err := strconv.Atoi(amtStr)
			if err != nil {
				panic(fmt.Sprintf("%s is NaN", amtStr))
			}

			switch color {
			case "red":
				if amt > 12 {
					return false
				}
			case "green":
				if amt > 13 {
					return false
				}
			case "blue":
				if amt > 14 {
					return false
				}
			}

		}
	}
	return true
}

// function for part 2
func gameScore(game string) int {
	var sets = strings.Split(game, ";")
	var maxRed, maxGreen, maxBlue = 0, 0, 0

	for _, set := range sets {
		var draws = strings.Split(set, ",")

		for _, draw := range draws {
			var d = strings.Split(draw, " ")
			var amtStr = d[1]
			var color = d[2]

			amt, err := strconv.Atoi(amtStr)
			if err != nil {
				panic(fmt.Sprintf("%s is NaN", amtStr))
			}

			switch color {
			case "red":
				if amt > maxRed {
					maxRed = amt
				}
			case "green":
				if amt > maxGreen {
					maxGreen = amt
				}
			case "blue":
				if amt > maxBlue {
					maxBlue = amt
				}
			}

		}
	}
	return maxRed * maxGreen * maxBlue
}
