package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 142
	actual := getSum("test_input_one.txt")
	assert.Equal(t, expected, actual)
}

func TestPartTwo(t *testing.T) {
	expected := 281
	actual := getSumIncludingLetters("test_input_two.txt")
	assert.Equal(t, expected, actual)
}
