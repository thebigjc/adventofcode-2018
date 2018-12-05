package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type guard struct {
	id    int
	total int
	sleep [60]int
}

func part() (int, int) {
	b, err := ioutil.ReadFile("day04.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	sort.Strings(lines)
	// [1518-11-23 00:00] Guard #3271 begins shift
	r := regexp.MustCompile(`\[\d{4}-\d{2}-\d{2} \d{2}:(\d{2})\] (Guard #(\d+)|falls|wakes)`)
	guards := make(map[int]*guard)
	var curGuard *guard
	var sleep int

	for _, l := range lines {
		s := r.FindStringSubmatch(l)
		if len(s[3]) > 0 {
			guardId, err := strconv.Atoi(s[3])

			if err != nil {
				panic(err)
			}

			curGuard = guards[guardId]
			if curGuard == nil {
				curGuard = &guard{guardId, 0, [60]int{}}
				guards[guardId] = curGuard
			}
		}
		if s[2] == "falls" {
			sleep, err = strconv.Atoi(s[1])
			if err != nil {
				panic(err)
			}
		}
		if s[2] == "wakes" {
			wake, err := strconv.Atoi(s[1])
			if err != nil {
				panic(err)
			}
			curGuard.total += wake - sleep
			for i := sleep; i < wake; i++ {
				curGuard.sleep[i]++
			}
		}
	}

	var maxGuard *guard
	maxSleep := 0
	maxSleepMin := 0
	maxSleepGuard := -1

	for _, guard := range guards {
		if maxGuard == nil {
			maxGuard = guard
			continue
		}
		if guard.total > maxGuard.total {
			maxGuard = guard
		}
		for i, m := range guard.sleep {
			if m > maxSleep {
				maxSleepMin = i
				maxSleep = m
				maxSleepGuard = guard.id
			}
		}
	}

	maxMin := 0
	maxIdx := 0

	for i, m := range maxGuard.sleep {
		if m > maxMin {
			maxIdx = i
			maxMin = m
		}
	}

	return maxIdx * maxGuard.id, maxSleepGuard * maxSleepMin
}

func main() {
	score1, score2 := part()
	log.Printf("Part 1 answer:%d", score1)
	log.Printf("Part 2 answer:%d", score2)
}
