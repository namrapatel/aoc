package main

import (
	"bufio"
	_ "embed"
	"os"
	"strings"
	"unicode"
)

var spelledDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func containsLetterNumber(s string) (bool, int) {
	for k, v := range spelledDigits {
		if strings.Contains(s, k) {
			return true, v
		}
	}
	return false, 0
}

func getFirstDigit(s string) int {
	for i := 0; i < len(s); i++ {
		if found, d := containsLetterNumber(s[:i]); found {
			return d
		} else if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - '0')
		}
	}
	panic("no digit found in " + s)
}

func getSecondDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if found, d := containsLetterNumber(s[i:]); found {
			return d
		} else if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - '0')
		}
	}
	panic("no digit found in " + s)
}

func getNumbersIncludingLetters(name string) int {
	sum := 0

	inputFile, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	defer func(inputFile *os.File) {
		err = inputFile.Close()
		if err != nil {
			panic(err)
		}
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		combined := 0

		combined = getFirstDigit(line)
		secondDigit := getSecondDigit(line)
		combined = combined*10 + secondDigit
		sum += combined
	}
	return sum
}
