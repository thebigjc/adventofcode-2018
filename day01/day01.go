package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func part1() int64 {
	f, err := os.OpenFile("day01.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("unable to open input: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	freq := int64(0)

	for sc.Scan() {
		l := sc.Text()

		i, err := strconv.ParseInt(l, 10, 32)
		if err != nil {
			log.Fatalf("Couldn't parse line: %s, %v", l, err)
		}
		freq += i
	}

	return freq
}

func part2() int64 {
	freq := int64(0)
	freqTable := make(map[int64]int)

	for {
		f, err := os.OpenFile("day01.txt", os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatalf("unable to open input: %v", err)
		}
		defer f.Close()

		sc := bufio.NewScanner(f)

		for sc.Scan() {
			l := sc.Text()

			i, err := strconv.ParseInt(l, 10, 32)
			if err != nil {
				log.Fatalf("Couldn't parse line: %s, %v", l, err)
			}
			freq += i
			seen := freqTable[freq]
			seen++
			freqTable[freq] = seen

			if seen > 1 {
				return freq
			}
		}
		log.Printf("Looping:%d, %d", freq, len(freqTable))
	}
}

func main() {
	log.Printf("Part 1 answer:%d", part1())
	log.Printf("Part 2 answer:%d", part2())
}
