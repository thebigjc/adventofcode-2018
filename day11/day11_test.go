package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestData(t *testing.T) {
	x, y, power := part1(18)
	assert.Equal(t, 33, x)
	assert.Equal(t, 45, y)
	assert.Equal(t, 29, power)

	x, y, power = part1(42)
	assert.Equal(t, 21, x)
	assert.Equal(t, 61, y)
	assert.Equal(t, 30, power)
}

func TestI(t *testing.T) {
	level := i(3, 5, 8)
	assert.Equal(t, 4, level)

	level = i(122, 79, 57)
	assert.Equal(t, -5, level)

	level = i(217, 196, 39)
	assert.Equal(t, 0, level)

	level = i(101, 153, 71)
	assert.Equal(t, 4, level)

	level = i(21, 61, 42)
	assert.Equal(t, 4, level)
}
