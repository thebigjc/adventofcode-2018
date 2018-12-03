package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func mustClose(f io.Closer) {
	err := f.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

type point struct {
	x int
	y int
}

type row struct {
	id int
	l  int
	t  int
	w  int
	h  int
}

func part1() (int, int) {
	f, err := os.OpenFile("day03.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("unable to open input: %v", err)
	}
	defer mustClose(f)

	sc := bufio.NewScanner(f)

	fabric := make(map[point]int)
	owners := make(map[point]int)
	ids := make(map[int]bool)

	for sc.Scan() {
		line := sc.Text()
		var r row
		_, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &r.id, &r.l, &r.t, &r.w, &r.h)

		if err != nil {
			panic(err)
		}
		ids[r.id] = true

		for x := r.l; x < r.l+r.w; x++ {
			for y := r.t; y < r.t+r.h; y++ {
				p := point{x, y}
				fabric[p]++

				if fabric[p] > 1 {
					delete(ids, r.id)
					delete(ids, owners[p])
				}

				owners[p] = r.id
			}
		}
	}

	log.Printf("There are %d ids left", len(ids))

	if len(ids) > 1 {
		panic("Too many ids")
	}

	winner := 0
	for k := range ids {
		winner = k
	}

	cnt := 0
	for _, v := range fabric {
		if v > 1 {
			cnt++
		}
	}

	return cnt, winner
}

func main() {
	cnt, winner := part1()
	log.Printf("Part 1 answer:%d", cnt)
	log.Printf("Part 2 answer:%d", winner)
}
