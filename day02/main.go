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

func getSumOfIdsOfValidGames(name string) (int, int) {
	sum := 0
	powerOfMins := 0

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
		id, validity, power := checkLineValidity(line)
		if validity {
			sum += id
		}
		powerOfMins += power
	}
	return sum, powerOfMins
}

func checkLineValidity(line string) (id int, valid bool, power int) {
	valid = true
	power = 1
	parts := strings.Fields(line)
	id, _ = strconv.Atoi(getTrimmedStringAtIndex(parts, 1))

	maxOfColor := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}
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
		maxOfColor[color] = max(maxOfColor[color], amount)
		if remainingRGB[color]-amount < 0 {
			valid = false
		}
	}
	power = calcPower(maxOfColor)
	return id, valid, power
}

func calcPower(mins map[string]int) int {
	power := 1
	for _, v := range mins {
		power *= v
	}
	return power
}

func getTrimmedStringAtIndex(stringsArr []string, index int) string {
	temp := stringsArr[index]
	return temp[:len(temp)-1]
}
