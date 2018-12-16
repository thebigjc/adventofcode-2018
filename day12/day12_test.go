package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToBools(t *testing.T) {
	b := stringToBools("#")
	assert.Len(t, b, 1)
	assert.Equal(t, true, b[0])

	b = stringToBools(".")
	assert.Len(t, b, 1)
	assert.Equal(t, false, b[0])

	b = stringToBools(".#")
	assert.Len(t, b, 2)
	assert.Equal(t, false, b[0])
	assert.Equal(t, true, b[1])
}

func TestExtractPlants(t *testing.T) {
	plants := stringToBools("##.#..")

	ep1 := extractPlants(0, plants)
	assert.Equal(t, stringToBools("..##."), ep1)

	ep2 := extractPlants(1, plants)
	assert.Equal(t, stringToBools(".##.#"), ep2)

	ep3 := extractPlants(2, plants)
	assert.Equal(t, stringToBools("##.#."), ep3)

	ep4 := extractPlants(3, plants)
	assert.Equal(t, stringToBools("#.#.."), ep4)

	ep5 := extractPlants(4, plants)
	assert.Equal(t, stringToBools(".#..."), ep5)
}

func TestEdgePlants(t *testing.T) {
	plants := stringToBools("..###")

	ep1 := extractPlants(4, plants)
	assert.Equal(t, stringToBools("###.."), ep1)

	ep2 := extractPlants(5, plants)
	assert.Equal(t, stringToBools("##..."), ep2)

	ep3 := extractPlants(6, plants)
	assert.Equal(t, stringToBools("#...."), ep3)
}

func TestSampleInputRow(t *testing.T) {
	input := stringToBools("#..#.#..##......###...###")
	patterns := [][]bool{
		stringToBools("...##"),
		stringToBools("..#.."),
		stringToBools(".#..."),
		stringToBools(".#.#."),
		stringToBools(".#.##"),
		stringToBools(".##.."),
		stringToBools(".####"),
		stringToBools("#.#.#"),
		stringToBools("#.###"),
		stringToBools("##.#."),
		stringToBools("##.##"),
		stringToBools("###.."),
		stringToBools("###.#"),
		stringToBools("####.")}
	output := generation(input, patterns)
	expected := stringToBools("#...#....#.....#..#..#..#")
	assert.Equal(t, expected, trimRight(output))
}

func trimRight(b []bool) []bool {
	var i int
	for i = 1; i < len(b); i++ {
		if b[len(b)-i] {
			break
		}
	}

	return b[:len(b)-i+1]
}

func TestScore(t *testing.T) {
	plants, patterns := readInput("day12.txt")
	score := part1(plants, patterns)

	assert.Equal(t, 325, score)
}

func TestScore2(t *testing.T) {
	plants, patterns := readInput("day12.txt")
	score := part2(plants, patterns)

	assert.Equal(t, 999999999374, score)
}

func TestReading(t *testing.T) {
	plants, patterns := readInput("day12.txt")
	assert.Len(t, patterns, 14)
	score := part1(plants, patterns)

	assert.Equal(t, 325, score)
}
