package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	gamePowerSum := 0
	currGameId := 0
	for scanner.Scan() {
		currGameId++

		rgbSlice := []int{0, 0, 0}

		gameDigits := len(strconv.Itoa(currGameId))
		gameIterations := strings.Split(scanner.Text()[7+gameDigits:], "; ")

		for i := 0; i < len(gameIterations); i++ {
			cubesPulled := strings.Split(gameIterations[i], ", ")
			for j := 0; j < len(cubesPulled); j++ {
				parsedPull := strings.Split(cubesPulled[j], " ")
				num, color := parsedPull[0], parsedPull[1]
				parsedNum, _ := strconv.Atoi(num)

				switch color {
				case "red":
					if parsedNum > rgbSlice[0] {
						rgbSlice[0] = parsedNum
					}
				case "green":
					if parsedNum > rgbSlice[1] {
						rgbSlice[1] = parsedNum
					}
				case "blue":
					if parsedNum > rgbSlice[2] {
						rgbSlice[2] = parsedNum
					}
				}
			}
		}

		gamePowerSum += (rgbSlice[0] * rgbSlice[1] * rgbSlice[2])
	}

	fmt.Println(gamePowerSum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
