package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbersSlice = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

const numbersString = "0123456789"

func findFirstNumberInString(str string) (string, bool) {
	foundNumber := -1

	for i := 0; i < len(numbersSlice); i++ {
		if idx := strings.Index(str, numbersSlice[i]); idx == 0 {
			foundNumber = i
		}
	}

	if foundNumber != -1 {
		return strconv.Itoa(foundNumber), true
	} else {
		return "", false
	}
}

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	calibrationSum := 0

	for scanner.Scan() {
		line := scanner.Text()

		foundNumbers := []string{}

		for i := 0; i < len(line); i++ {
			char := string(line[i])

			// Try to find integer string. If not present, see if current char matches the start
			// of a written number string
			if strings.Contains(numbersString, char) {
				foundNumbers = append(foundNumbers, char)
			} else {
				foundNumber, ok := findFirstNumberInString(line[i:])
				if ok {
					foundNumbers = append(foundNumbers, foundNumber)
				}
			}
		}

		calibrationValue, _ := strconv.Atoi(foundNumbers[0] + foundNumbers[len(foundNumbers)-1])

		calibrationSum += calibrationValue
	}

	fmt.Println(calibrationSum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
