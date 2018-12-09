package main

import (
	"container/list"
	"log"
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

		newScore := v.(int) + marble
		*score += newScore

		return newC
	}
	return l.InsertAfter(marble, clockwise(l, c))
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

func main() {
	score := part1(477, 70851)
	log.Printf("High score is:%d", score)
	score = part1(477, 70851*100)
	log.Printf("High score is:%d", score)
}
