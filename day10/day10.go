package main

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func mustClose(f io.Closer) {
	err := f.Close()
	if err != nil {
		log.Fatalf("Failed to close file: %v", err)
	}
}

func parseInt(s string) int {
	x, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return x
}

type point struct {
	x  int
	y  int
	dx int
	dy int
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func distance(a, b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	f, err := os.OpenFile("day10.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("unable to open input: %v", err)
	}
	defer mustClose(f)

	sc := bufio.NewScanner(f)

	points := make([]point, 0)

	for sc.Scan() {
		line := sc.Text()
		// 00000000001111111111222222222233333333334444444444
		// 01234567890123456789012345678901234567890123456789
		// position=< -9951, -50547> velocity=< 1,  5>
		var p point
		p.x = parseInt(line[10:16])
		p.y = parseInt(line[18:24])
		p.dx = parseInt(line[36:38])
		p.dy = parseInt(line[40:42])

		points = append(points, p)
	}

	minD := math.MaxInt64
	seconds := 0

	for {
		for i := 0; i < len(points); i++ {
			points[i].x += points[i].dx
			points[i].y += points[i].dy
		}
		d := 0
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				d += distance(points[i], points[j])
			}
		}

		if d < minD {
			seconds++
			minD = d
		}
		if d > minD {
			for i := 0; i < len(points); i++ {
				points[i].x -= points[i].dx
				points[i].y -= points[i].dy
			}
			break
		}
	}

	minX := math.MaxInt64
	minY := math.MaxInt64
	maxX := math.MinInt64
	maxY := math.MinInt64

	for i := 0; i < len(points); i++ {
		minX = min(minX, points[i].x)
		minY = min(minY, points[i].y)
		maxX = max(maxX, points[i].x)
		maxY = max(maxY, points[i].y)
	}

	r := image.Rect(minX, minY, maxX+1, maxY+1)
	img := image.NewRGBA(r)

	for i := 0; i < len(points); i++ {
		img.Set(points[i].x, points[i].y, color.Black)
	}

	f, err = os.OpenFile("day10.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer mustClose(f)
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	log.Printf("Seconds: %d", seconds)
}
