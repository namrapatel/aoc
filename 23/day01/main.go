package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func GetNumbers() {
	sum := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	counter := 0
	// For each line
	for scanner.Scan() {
		if counter > 2 {
			break
		}
		line := scanner.Text()
		lineChar := 0
		print(strings.Fields(line))
		// For each character in the line
		for _, char := range line {

			if unicode.IsDigit(char) {
				print(int(char - '0'))
				lineChar = int(char - '0')
				break
			}
		}

		// Go backwards in the same line and find the first digit, add to lineSum
		for i := len(line) - 1; i >= 0; i-- {
			char := line[i]
			if unicode.IsDigit(rune(char)) {
				fmt.Printf("Backward: %d\n", int(char-'0'))
				lineChar = lineChar*10 + int(char-'0')
				break
			}
		}
		sum += lineChar
		counter++
		// Print the total sum for the line
		fmt.Printf("Total Sum for Line: %d\n", lineChar)

	}
	fmt.Println(sum)
}
