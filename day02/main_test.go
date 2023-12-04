package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImpl(t *testing.T) {
	expected := 8
	actual := getSumOfIdsOfValidGames("test_input.txt")
	assert.Equal(t, expected, actual)
}

func TestActual(t *testing.T) {
	expected := 1853
	actual := getSumOfIdsOfValidGames("input.txt")
	assert.Equal(t, expected, actual)
}
