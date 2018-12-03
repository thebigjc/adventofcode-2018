package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
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

func part1() (int, int) {
	f, err := os.OpenFile("day03.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("unable to open input: %v", err)
	}
	defer mustClose(f)

	sc := bufio.NewScanner(f)

	// #1 @ 108,350: 22x29
	r := regexp.MustCompile(`\#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

	fabric := make(map[point]int, 0)
	owners := make(map[point]int, 0)
	ids := make(map[int]bool)

	for sc.Scan() {
		line := sc.Text()
		// #1 @ 108,350: 22x29
		s := r.FindStringSubmatch(line)

		id, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatalf("Couldn't parse id: %s, %v", s[1], err)
		}
		ids[id] = true

		l, err := strconv.Atoi(s[2])
		if err != nil {
			log.Fatalf("Couldn't parse left:%s, %v", s[2], err)
		}
		t, err := strconv.Atoi(s[3])
		if err != nil {
			log.Fatalf("Couldn't parse top:%s, %v", s[3], err)
		}
		w, err := strconv.Atoi(s[4])
		if err != nil {
			log.Fatalf("Couldn't parse width:%s, %v", s[4], err)
		}
		h, err := strconv.Atoi(s[5])
		if err != nil {
			log.Fatalf("Couldn't parse height:%s, %v", s[5], err)
		}
		for i := l; i < l+w; i++ {
			for j := t; j < t+h; j++ {
				p := point{i, j}
				fabric[p]++

				if fabric[p] > 1 {
					delete(ids, id)
					delete(ids, owners[p])
				}

				owners[p] = id
			}
		}
	}

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
