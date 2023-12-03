package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NumRedCubes   = 12
	NumGreenCubes = 13
	NumBlueCubes  = 14
)

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	gameIdSum := 0
	currGameId := 0
	for scanner.Scan() {
		currGameId++

		gameDigits := len(strconv.Itoa(currGameId))
		gameIterations := strings.Split(scanner.Text()[7+gameDigits:], "; ")

		gamePossible := true
		for i := 0; i < len(gameIterations); i++ {
			cubesPulled := strings.Split(gameIterations[i], ", ")
			for j := 0; j < len(cubesPulled); j++ {
				parsedPull := strings.Split(cubesPulled[j], " ")
				num, color := parsedPull[0], parsedPull[1]

				parsedNum, _ := strconv.Atoi(num)

				switch color {
				case "red":
					if parsedNum > NumRedCubes {
						gamePossible = false
					}
				case "green":
					if parsedNum > NumGreenCubes {
						gamePossible = false
					}
				case "blue":
					if parsedNum > NumBlueCubes {
						gamePossible = false
					}
				}

				if !gamePossible {
					break
				}
			}
		}

		if gamePossible {
			gameIdSum += currGameId
		}
	}

	fmt.Println(gameIdSum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
