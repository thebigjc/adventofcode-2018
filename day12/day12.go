package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

var padding = []bool{false, false, false, false, false}

func extractPlants(n int, inplants []bool) []bool {
	n += len(padding)
	plants := padding[:]
	plants = append(plants, inplants...)
	plants = append(plants, padding...)

	i := n - 2
	j := n + 1

	left := plants[i : i+2]
	right := plants[j : j+2]

	return []bool{left[0], left[1], plants[n], right[0], right[1]}
}

func stringToBools(s string) []bool {
	b := make([]bool, 0, len(s))
	for _, c := range s {
		if c == '#' {
			b = append(b, true)
		} else {
			b = append(b, false)
		}
	}

	return b
}

func readInput(fn string) ([]bool, [][]bool) {
	bs, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")
	plants := []bool{false, false, false}
	plants = append(plants, stringToBools(strings.TrimSpace(lines[0][15:]))...)

	patterns := [][]bool{}

	for _, l := range lines[2:] {
		pattern := l[0:5]
		plant := l[9]
		if plant == '#' {
			patterns = append(patterns, stringToBools(pattern))
		}
	}

	return plants, patterns
}

func main() {
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	plants, patterns := readInput("day12_input.txt")
	score := part1(plants, patterns)
	log.Printf("Part1: %d", score)
	score = part2(plants, patterns)
	log.Printf("Part2: %d", score)
}

func compareSlice(a, b []bool) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func generation(plants []bool, patterns [][]bool) []bool {
	newPlants := make([]bool, len(plants)+3)

	for j := 0; j < len(plants)+3; j++ {
		curPlant := extractPlants(j, plants)

		for _, pattern := range patterns {
			if compareSlice(pattern, curPlant) {
				newPlants[j] = true
				break
			}
		}
	}

	return newPlants
}

func runGenerations(plants []bool, patterns [][]bool, generations int) int {
	for i := 0; i < generations; i++ {
		plants = generation(plants, patterns)
	}

	score := 0
	for i, v := range plants {
		if v {
			score += i - 3
		}
	}

	return score
}

func part1(plants []bool, patterns [][]bool) int {
	score := runGenerations(plants, patterns, 20)
	return score
}

func part2(plants []bool, patterns [][]bool) int {
	x1 := 300
	x2 := 500
	y1 := runGenerations(plants, patterns, x1)
	y2 := runGenerations(plants, patterns, x2)

	m := (y2 - y1) / (x2 - x1)
	b := y2 - m*x2

	log.Printf("y1: %d, y2: %d", y1, y2)
	log.Printf("m: %d, b: %d", m, b)

	return m*50000000000 + b
}
