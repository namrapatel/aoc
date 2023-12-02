package main

import (
	"bufio"
	_ "embed"
	"os"
	"unicode"
)

func getFirstDigit(s string) int {
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - '0')
		}
	}
	panic("no digit found in " + s)
}

func getSecondDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - '0')
		}
	}
	panic("no digit found in " + s)
}

func getSum(name string) int {
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
