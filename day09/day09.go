package main

import (
	"container/list"
	"container/ring"
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

func clockwise(l *list.List, c *list.Element) *list.Element {
	if c.Next() == nil {
		return l.Front()
	}
	return c.Next()
}

func counterClockwise(l *list.List, c *list.Element) *list.Element {
	if c.Prev() == nil {
		return l.Back()
	}
	return c.Prev()
}

func insertMarble(l *list.List, c *list.Element, marble int, score *int) *list.Element {
	if (marble % 23) == 0 {
		for i := 0; i < 7; i++ {
			c = counterClockwise(l, c)
		}
		newC := clockwise(l, c)
		v := l.Remove(c)

		*score += (v.(int) + marble)

		return newC
	}
	return l.InsertAfter(marble, clockwise(l, c))
}

func insertMarbleRing(c *ring.Ring, marble *ring.Ring, score *int) *ring.Ring {
	mv := marble.Value.(int)
	if (mv % 23) == 0 {
		c = c.Move(-8)

		v := c.Unlink(1).Value.(int)

		*score += (v + mv)

		return c.Next()
	}
	c = c.Move(1)
	c.Link(marble)

	return marble
}

func part1(players, last int) int {
	circle := list.New()
	scores := make([]int, players)

	current := circle.PushBack(0)
	for i := 1; i <= last; i++ {
		current = insertMarble(circle, current, i, &scores[i%players])
	}

	maxScore := 0
	for _, s := range scores {
		if s > maxScore {
			maxScore = s
		}
	}

	return maxScore
}

func part1Ring(players, last int) int {
	marbles := make([]ring.Ring, last+1)
	scores := make([]int, players)
	marbles[0].Value = 0

	current := &marbles[0]
	for i := 1; i <= last; i++ {
		marbles[i].Value = i
		current = insertMarbleRing(current, &marbles[i], &scores[i%players])
	}

	maxScore := 0
	for _, s := range scores {
		if s > maxScore {
			maxScore = s
		}
	}

	return maxScore
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

	score := part1(477, 70851)
	log.Printf("High score is:%d", score)
	score = part1Ring(477, 70851*100)
	log.Printf("High score is:%d", score)
}
