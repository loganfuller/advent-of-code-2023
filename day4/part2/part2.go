package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

func countCards(gameCards [][2][]string, cardIdx int) int {
	gameCard := gameCards[cardIdx]
	winningNumbers := gameCard[0]
	cardNumbers := gameCard[1]

	numWinners := 0
	numCards := 1

	for _, number := range cardNumbers {
		if slices.Contains(winningNumbers, number) {
			numWinners++
		}
	}

	for i := 1; i <= numWinners; i++ {
		nextIdx := cardIdx + i
		if nextIdx < len(gameCards) {
			numCards += countCards(gameCards, nextIdx)
		} else {
			break
		}
	}

	return numCards
}

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	rCard := regexp.MustCompile(`(?:Card[ ]+\d+:[ ]+)((?:\d+[ ]*)+)(?: +\| +)((?:\d+[ ]*)+)`)
	rSplitNumbers := regexp.MustCompile(`[ ]+`)

	scanner := bufio.NewScanner(f)

	totalCards := 0

	// Populate scratch cards
	gameCards := [][2][]string{}
	for scanner.Scan() {
		rMatches := rCard.FindAllStringSubmatch(scanner.Text(), -1)

		winningNumbers := rSplitNumbers.Split(rMatches[0][1], -1)
		cardNumbers := rSplitNumbers.Split(rMatches[0][2], -1)

		gameCards = append(gameCards, [2][]string{winningNumbers, cardNumbers})
	}

	// Iterate over scratch cards for scoring
	for i := 0; i < len(gameCards); i++ {
		totalCards += countCards(gameCards, i)
	}

	fmt.Println(totalCards)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
