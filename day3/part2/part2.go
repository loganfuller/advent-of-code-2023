package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const NumberString = `0123456789`

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	rNum := regexp.MustCompile(`\d+`)
	rAst := regexp.MustCompile(`\*`)

	gearRatioSum := 0

	scanIdx := 0
	lineSlice := []string{}
	numberSlice := [][][]int{}
	asteriskSlice := [][]int{}
	for scanner.Scan() {
		lineSlice = append(lineSlice, scanner.Text())

		// Pull out numbers as index ranges
		numberSlice = append(numberSlice, rNum.FindAllStringIndex(lineSlice[scanIdx], -1))

		// Pull out asterisks as indices
		asteriskSlice = append(asteriskSlice, []int{})
		asteriskIndices := rAst.FindAllStringIndex(lineSlice[scanIdx], -1)
		for j := 0; j < len(asteriskIndices); j++ {
			asteriskSlice[scanIdx] = append(asteriskSlice[scanIdx], asteriskIndices[j][0])
		}

		scanIdx++
	}

	for i := 0; i < len(lineSlice); i++ {
		for j := 0; j < len(asteriskSlice[i]); j++ {
			asteriskIdx := asteriskSlice[i][j]
			gearNumbers := []int{}

			// Check current line
			for k := 0; k < len(numberSlice[i]); k++ {
				numberIdxStart, numberIdxStop := numberSlice[i][k][0], numberSlice[i][k][1]

				if numberIdxStart == (asteriskIdx+1) || numberIdxStop == asteriskIdx {
					gearNumber, _ := strconv.Atoi(lineSlice[i][numberSlice[i][k][0]:numberSlice[i][k][1]])
					gearNumbers = append(gearNumbers, gearNumber)
				}
			}

			// Check previous line
			if i != 0 {
				for k := 0; k < len(numberSlice[i-1]); k++ {
					numberIdxStart, numberIdxStop := numberSlice[i-1][k][0], numberSlice[i-1][k][1]

					if asteriskIdx >= (numberIdxStart-1) && asteriskIdx <= (numberIdxStop) {
						gearNumber, _ := strconv.Atoi(lineSlice[i-1][numberSlice[i-1][k][0]:numberSlice[i-1][k][1]])
						gearNumbers = append(gearNumbers, gearNumber)
					}
				}
			}

			// Check next line
			if i != (len(lineSlice) - 1) {
				for k := 0; k < len(numberSlice[i+1]); k++ {
					numberIdxStart, numberIdxStop := numberSlice[i+1][k][0], numberSlice[i+1][k][1]

					if asteriskIdx >= (numberIdxStart-1) && asteriskIdx <= (numberIdxStop) {
						gearNumber, _ := strconv.Atoi(lineSlice[i+1][numberSlice[i+1][k][0]:numberSlice[i+1][k][1]])
						gearNumbers = append(gearNumbers, gearNumber)
					}
				}
			}

			if len(gearNumbers) == 2 {
				gearRatioSum += (gearNumbers[0] * gearNumbers[1])
			}
		}
	}

	fmt.Println(gearRatioSum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
