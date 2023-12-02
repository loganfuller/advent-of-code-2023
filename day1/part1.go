package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	calibrationSum := 0
	lineNum := 0

	for scanner.Scan() {
		lineNum++

		calibrationLineSlice := strings.Split(scanner.Text(), "")

		var calibrationValueSlice []string

		for i := 0; i < len(calibrationLineSlice); i++ {
			_, err := strconv.Atoi(calibrationLineSlice[i])
			if err == nil {
				calibrationValueSlice = append(calibrationValueSlice, calibrationLineSlice[i])
			}
		}

		calibrationValue, _ := strconv.Atoi(
			calibrationValueSlice[0] + calibrationValueSlice[len(calibrationValueSlice)-1],
		)

		calibrationSum += calibrationValue
	}

	fmt.Println(calibrationSum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
