package main

import (
	"bufio"
	"log"
	"os"
)

func part1() int {
	f, err := os.OpenFile("day02.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("unable to open input: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	twos := 0
	threes := 0

	for sc.Scan() {
		l := sc.Text()
		freq := make(map[rune]int)
		for _, i := range l {
			freq[i]++
		}
		var two bool
		var three bool
		for _, v := range freq {
			if v == 2 {
				two = true
			}
			if v == 3 {
				three = true
			}
		}

		if two {
			twos++
		}
		if three {
			threes++
		}
	}

	return twos * threes
}

func part2() string {
	f, err := os.OpenFile("day02.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("unable to open input: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	ids := make(map[string]string)

	for sc.Scan() {
		l := sc.Text()
		uniq := make(map[string]bool)
		for i := range l {
			newl := l[:i] + l[i+1:]
			uniq[newl] = true
		}

		for k := range uniq {
			if v, ok := ids[k]; ok {
				log.Printf("Found %s in %s and %s", k, v, l)
				return k
			}
			ids[k] = l
		}
	}

	return ""
}

func main() {
	log.Printf("Part 1 answer:%d", part1())
	log.Printf("Part 2 answer:%s", part2())
}
