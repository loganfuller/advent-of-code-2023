package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const SymbolString = `~!@#$%^&*()_-+={[}]|\:;"'<,>?/`

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	rNum := regexp.MustCompile(`\d+`)
	rSym := regexp.MustCompile(`[^\d|\.]+`)

	lineSlice := []string{}
	for scanner.Scan() {
		lineSlice = append(lineSlice, scanner.Text())
	}

	partNumberSum := 0
	for i := 0; i < len(lineSlice); i++ {
		numbers := rNum.FindAllStringIndex(lineSlice[i], -1)

		for j := 0; j < len(numbers); j++ {
			idxStart, idxStop := numbers[j][0], numbers[j][1]
			altIdxStart, altIdxStop := 0, 0
			isPartNumber := false

			// Check current line
			prevCharIsSymbol := idxStart > 0 && strings.Contains(SymbolString, lineSlice[i][idxStart-1:idxStart])
			nextCharIsSymbol := idxStop != len(lineSlice[i]) && strings.Contains(SymbolString, lineSlice[i][idxStop:idxStop+1])
			if prevCharIsSymbol || nextCharIsSymbol {
				isPartNumber = true
				goto addSum
			}

			// Substring indices for prev / next lines
			altIdxStart = idxStart - 1
			if altIdxStart < 0 {
				altIdxStart = 0
			}
			altIdxStop = idxStop + 1
			if altIdxStop > len(lineSlice[i]) {
				altIdxStop = len(lineSlice[i])
			}

			// Check previous line
			if i != 0 {
				substringToCheck := lineSlice[i-1][altIdxStart:altIdxStop]

				if rSym.MatchString(substringToCheck) {
					isPartNumber = true
					goto addSum
				}
			}

			// Check next line
			if i != (len(lineSlice) - 1) {
				substringToCheck := lineSlice[i+1][altIdxStart:altIdxStop]

				if rSym.MatchString(substringToCheck) {
					isPartNumber = true
					goto addSum
				}
			}

		addSum:
			if isPartNumber {
				partNumber, _ := strconv.Atoi(lineSlice[i][idxStart:idxStop])
				partNumberSum += partNumber
			}
		}
	}

	fmt.Println(partNumberSum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
