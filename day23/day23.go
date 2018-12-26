package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type nanobot struct {
	x int
	y int
	z int
	r int
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func (i *nanobot) distance(j *nanobot) int {
	return abs(i.x-j.x) + abs(i.y-j.y) + abs(i.z-j.z)
}

func min(x int, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

type oned struct {
	dist  int
	delta int
}

func main() {
	bs, err := ioutil.ReadFile("day23.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")
	var bots []nanobot

	maxBot := nanobot{}
	maxRange := -1

	for _, line := range lines {
		line = strings.TrimSpace(line)
		bot := nanobot{}
		_, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &bot.x, &bot.y, &bot.z, &bot.r)
		if err != nil {
			panic(err)
		}
		if bot.r > maxRange {
			maxBot = bot
			maxRange = bot.r
		}

		bots = append(bots, bot)
	}

	inrange := 0

	for _, bot := range bots {
		if maxBot.distance(&bot) <= maxBot.r {
			inrange++
		}
	}

	fmt.Printf("Part1: %d\n", inrange)

	pq := make([]oned, 0, len(bots)*2)

	for _, bot := range bots {
		d := abs(bot.x) + abs(bot.y) + abs(bot.z)
		pq = append(pq, oned{max(0, d-bot.r), 1})
		pq = append(pq, oned{d + bot.r + 1, -1})
	}

	count := 0
	maxCount := 0
	result := 0

	sort.Slice(pq, func(i, j int) bool { return pq[i].dist < pq[j].dist })

	for _, item := range pq {
		dist := item.dist
		count += item.delta
		if count > maxCount {
			result = dist
			maxCount = count
		}
	}
	fmt.Printf("Part2: %d", result)
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
