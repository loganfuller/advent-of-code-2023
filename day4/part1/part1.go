package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	rCard := regexp.MustCompile(`(?:Card[ ]+\d+:[ ]+)((?:\d+[ ]*)+)(?: +\| +)((?:\d+[ ]*)+)`)
	rSplitNumbers := regexp.MustCompile(`[ ]+`)

	scanner := bufio.NewScanner(f)

	totalPoints := 0

	// Populate scratch cards
	for scanner.Scan() {
		rMatches := rCard.FindAllStringSubmatch(scanner.Text(), -1)

		winningNumbers := rSplitNumbers.Split(rMatches[0][1], -1)
		cardNumbers := rSplitNumbers.Split(rMatches[0][2], -1)

		numWinners := 0
		for _, number := range cardNumbers {
			if slices.Contains(winningNumbers, number) {
				numWinners++
			}
		}

		cardPoints := 0
		for i := 0; i < numWinners; i++ {
			if i == 0 {
				cardPoints = 1
			} else {
				cardPoints *= 2
			}
		}

		totalPoints += cardPoints
	}

	fmt.Println(totalPoints)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
