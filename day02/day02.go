package main

import (
	"bufio"
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

func part1() int {
	f, err := os.OpenFile("day02.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("unable to open input: %v", err)
	}
	defer mustClose(f)

	sc := bufio.NewScanner(f)
	twos := 0
	threes := 0

	for sc.Scan() {
		l := sc.Text()
		freq := make(map[rune]int)
		for _, i := range l {
			freq[i]++
		}
		var two int
		var three int
		for _, v := range freq {
			if v == 2 {
				two = 1
			}
			if v == 3 {
				three = 1
			}
		}
		twos += two
		threes += three
	}

	return twos * threes
}

func part2() string {
	f, err := os.OpenFile("day02.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("unable to open input: %v", err)
	}
	defer mustClose(f)

	sc := bufio.NewScanner(f)

	ids := make(map[string]bool)

	for sc.Scan() {
		l := sc.Text()
		uniq := make(map[string]bool)
		for i := range l {
			newl := l[:i] + l[i+1:]
			uniq[newl] = true
		}

		for k := range uniq {
			if ids[k] {
				return k
			}
			ids[k] = true
		}
	}

	return ""
}

func main() {
	log.Printf("Part 1 answer:%d", part1())
	log.Printf("Part 2 answer:%s", part2())
}
