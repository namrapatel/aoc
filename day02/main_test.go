package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestActual(t *testing.T) {
	expectedSum, expectedPower := 8, 2286
	actualSum, actualPower := getSumOfIdsOfValidGames("test_input.txt")
	assert.Equal(t, expectedSum, actualSum)
	assert.Equal(t, expectedPower, actualPower)
}
