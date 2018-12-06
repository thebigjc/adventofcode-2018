package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math"
	"os"
	"runtime/pprof"
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

	//log.Printf("minX: %d, minY: %d, maxX: %d, maxY: %d", minX, minY, maxX, maxY)
	l := minX - len(points)
	t := minY - len(points)
	r := maxX + len(points)
	b := maxY + len(points)
	grid := make([][]int, (r - l))
	regionSize := 0

	//log.Printf("Points: %d, Width: %d, Height: %d", len(points), r-l, b-t)
	for x := l; x < r; x++ {
		grid[x-l] = make([]int, b-t)
		for y := t; y < b; y++ {
			minD := math.MaxInt32
			totalD := 0
			currentBest := 0
			for i, p := range points {
				d := distance(x, y, p.x, p.y)
				totalD += d
				if d == minD {
					currentBest = -1
				}
				if d < minD {
					minD = d
					currentBest = i
				}
			}
			grid[x-l][y-t] = currentBest
			if totalD < 10000 {
				regionSize++
			}
		}
	}

	areas := make([]int, len(points))
	//log.Printf("Grid points:%d", len(grid))
	for x, row := range grid {
		for y, v := range row {
			if v == -1 {
				// Tie
				continue
			}
			// Blacklisted areas in 'infinite' space
			if areas[v] == -1 {
				continue
			}
			if x+l < minX || y+t < minY || x+l > maxX || y+t > maxY {
				//log.Printf("Point %d is out", v)
				areas[v] = -1
				continue
			}
			areas[v]++
		}
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
