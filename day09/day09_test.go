package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	e := part1(9, 25)
	assert.Equal(t, e, 32)
	e = part1Ring(9, 25)
	assert.Equal(t, e, 32)
}

func TestExamples(t *testing.T) {
	e1 := part1(10, 1618)
	assert.Equal(t, e1, 8317)

	e2 := part1(13, 7999)
	assert.Equal(t, e2, 146373)

	e3 := part1(17, 1104)
	assert.Equal(t, e3, 2764)

	e4 := part1(21, 6111)
	assert.Equal(t, e4, 54718)

	e5 := part1(30, 5807)
	assert.Equal(t, e5, 37305)

	e1 = part1Ring(10, 1618)
	assert.Equal(t, e1, 8317)

	e2 = part1Ring(13, 7999)
	assert.Equal(t, e2, 146373)

	e3 = part1Ring(17, 1104)
	assert.Equal(t, e3, 2764)

	e4 = part1Ring(21, 6111)
	assert.Equal(t, e4, 54718)

	e5 = part1Ring(30, 5807)
	assert.Equal(t, e5, 37305)
}

func BenchmarkInputList(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		part1(477, 70851)
	}
}

func BenchmarkInputRing(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		part1Ring(477, 70851)
	}
}
