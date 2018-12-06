package main

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(x1, y1, x2, y2 int) int {
	return abs(x2-x1) + abs(y2-y1)
}

func main() {
	bs, err := ioutil.ReadFile("day06.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")
	points := []point{}
	minX := math.MaxInt32
	minY := math.MaxInt32
	maxX := math.MinInt32
	maxY := math.MinInt32

	for _, l := range lines {
		ps := strings.Split(l, ",")
		x, err := strconv.Atoi(strings.Trim(ps[0], "\r "))
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(strings.Trim(ps[1], "\r "))
		if err != nil {
			panic(err)
		}

		p := point{x, y}
		points = append(points, p)
		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	log.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d", minX, minY, maxX, maxY)
	l := minX - (maxX - minX)
	t := minY - (maxY - minY)
	r := maxX + (maxX - minX)
	b := maxY + (maxY - minY)
	grid := make(map[point]int)
	regionSize := 0

	log.Printf("Points: %d, Width: %d, Height: %d", len(points), r-l, b-t)
	for x := l; x < r; x++ {
		for y := t; y < b; y++ {
			gp := point{x, y}
			minD := math.MaxInt32
			totalD := 0
			for i, p := range points {
				d := distance(x, y, p.x, p.y)
				totalD += d
				if d == minD {
					grid[gp] = -1
				}
				if d < minD {
					minD = d
					grid[gp] = i
				}
			}
			if totalD < 10000 {
				regionSize++
			}
		}
	}

	areas := make(map[int]int)
	log.Printf("Grid points:%d", len(grid))
	for k, v := range grid {
		if v == -1 {
			// Tie
			continue
		}
		// Blacklisted areas in 'infinite' space
		if areas[v] == -1 {
			continue
		}
		if k.x < minX || k.y < minY || k.x > maxX || k.y > maxY {
			log.Printf("Point %d is out", v)
			areas[v] = -1
			continue
		}
		areas[v]++
	}

	maxArea := 0
	for _, v := range areas {
		if v > maxArea {
			maxArea = v
		}
	}
	log.Printf("Max area is:%d", maxArea)
	log.Printf("Max region size is:%d", regionSize)
}
