package main

import (
	"io/ioutil"
	"log"
	"math"
	"strings"
)

const size = 50

type yard [size][size]rune

func safeCmp(yd yard, x, y int, r rune) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= size || y >= size {
		return false
	}

	return yd[x][y] == r
}

func countAdjacent(yd yard, x, y int, r rune) int {
	cnt := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if safeCmp(yd, x+i, y+j, r) {
				cnt++
			}
		}
	}

	return cnt
}

func runYard(y yard) yard {
	out := y
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			switch y[i][j] {
			case '.':
				trees := countAdjacent(y, i, j, '|')
				if trees >= 3 {
					out[i][j] = '|'
				}
			case '|':
				lumber := countAdjacent(y, i, j, '#')
				if lumber >= 3 {
					out[i][j] = '#'
				}
			case '#':
				lumber := countAdjacent(y, i, j, '#')
				trees := countAdjacent(y, i, j, '|')
				if lumber < 1 || trees < 1 {
					out[i][j] = '.'
				}
			}
		}
	}

	return out
}

func scoreYard(yd yard) int {
	wood := 0
	lumber := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			switch yd[i][j] {
			case '|':
				wood++
			case '#':
				lumber++
			}
		}
	}
	return wood * lumber
}

func main() {
	bs, err := ioutil.ReadFile("day18.txt")
	if err != nil {
		panic(err)
	}

	var yd yard
	var start yard
	lines := strings.Split(string(bs), "\n")
	for x, line := range lines {
		for y, r := range strings.TrimSpace(line) {
			yd[x][y] = r
		}
	}
	start = yd

	for i := 0; i < 10; i++ {
		yd = runYard(yd)
	}
	score := scoreYard(yd)
	log.Printf("Part1: %d", score)

	yd = start

	minScore := math.MaxUint32

	minIs := []int{}
	values := []int{}
	lastValues := []int{}

	for i := 1; i < 1200; i++ {
		yd = runYard(yd)
		if i < 1000 {
			continue
		}

		score := scoreYard(yd)
		if score < minScore {
			minScore = score
			minIs = []int{}
		}
		if score == minScore {
			minIs = append(minIs, i)
			if len(values) == len(lastValues) && len(values) > 1 {
				break
			}
			lastValues = values
			values = []int{}
		}
		values = append(values, score)
	}

	period := len(values)
	startI := minIs[0]
	target := 1000000000
	shifted := target - startI
	offset := shifted % period

	output := values[offset]
	log.Printf("Part2: %d", output)
}
