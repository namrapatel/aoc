package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var remainingRGB = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func getSumOfIdsOfValidGames(name string) int {
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
		if id, validity := checkLineValidity(line); validity {
			sum += id
		}
	}
	return sum
}

func checkLineValidity(line string) (id int, valid bool) {
	parts := strings.Fields(line)
	id, _ = strconv.Atoi(getTrimmedStringAtIndex(parts, 1))
	for i := 2; i < len(parts); i += 2 {
		amount, err := strconv.Atoi(parts[i])
		if err != nil {
		}
		var color string
		if strings.HasSuffix(parts[i+1], ",") || strings.HasSuffix(parts[i+1], ";") {
			color = getTrimmedStringAtIndex(parts, i+1)
		} else {
			color = parts[i+1]
		}

		if remainingRGB[color]-amount < 0 {
			return id, false
		}
	}
	return id, true
}

func getTrimmedStringAtIndex(stringsArr []string, index int) string {
	temp := stringsArr[index]
	return temp[:len(temp)-1]
}
