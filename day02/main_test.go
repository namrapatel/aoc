package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestActual(t *testing.T) {
	expectedSum, expectedPower := 1853, 72706
	actualSum, actualPower := getSumOfIdsOfValidGames("input.txt")
	assert.Equal(t, expectedSum, actualSum)
	assert.Equal(t, expectedPower, actualPower)
}
