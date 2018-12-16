package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type cart struct {
	i    int
	x    int
	y    int
	dx   int
	dy   int
	t    int
	dead bool
}

type point struct {
	x int
	y int
}

func findCart(carts []*cart, cx, cy int) *cart {
	for _, c := range carts {
		if c.x == cx && c.y == cy && !c.dead {
			return c
		}
	}
	return nil
}

func readInput(fn string) ([]*cart, map[point]rune) {
	bs, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return parseInput(string(bs))
}

func parseInput(bs string) ([]*cart, map[point]rune) {
	lines := strings.Split(bs, "\n")
	carts := []*cart{}
	points := map[point]rune{}

	for y, l := range lines {
		for x, c := range l {
			p := point{x, y}
			var r rune
			switch c {
			case '<':
				carts = append(carts, &cart{len(carts), x, y, -1, 0, 0, false})
				r = '-'
			case '>':
				carts = append(carts, &cart{len(carts), x, y, 1, 0, 0, false})
				r = '-'
			case '^':
				carts = append(carts, &cart{len(carts), x, y, 0, -1, 0, false})
				r = '|'
			case 'v':
				carts = append(carts, &cart{len(carts), x, y, 0, 1, 0, false})
				r = '|'
			default:
				r = c
			}
			points[p] = r
		}
	}

	return carts, points
}

func aliveCarts(carts []*cart) int {
	n := 0
	for _, c := range carts {
		if !c.dead {
			n++
		}
	}

	return n
}

func (c *cart) check(points map[point]rune) {
	if c.t > 2 {
		panic("t too big")
	}
	r := points[point{c.x, c.y}]
	if r == 0 {
		panic("off the track")
	}
	if r != '-' && r != '|' && r != '+' && r != '\\' && r != '/' {
		panic("bad track")
	}
}

func findCrashes(carts []*cart, points map[point]rune) {
	for aliveCarts(carts) > 1 {
		log.Printf("Next Step")
		for _, c := range carts {
			if c.dead {
				continue
			}
			log.Printf("ci: %d, x: %d, y: %d, dx: %d, dy: %d, p: %c, t: %d", c.i, c.x, c.y,
				c.dx, c.dy, points[point{c.x, c.y}], c.t)
			c.check(points)
		}

		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y < carts[j].y {
				return true
			}
			if carts[i].y > carts[j].y {
				return false
			}
			return carts[i].x < carts[j].x
		})
		for _, c := range carts {
			if c.dead {
				continue
			}

			cx := c.x + c.dx
			cy := c.y + c.dy
			crashed := findCart(carts, cx, cy)
			if crashed != nil {
				c.dead = true
				crashed.dead = true
				log.Printf("Crash at %d, %d", cx, cy)
			}
			c.x = cx
			c.y = cy

			switch points[point{cx, cy}] {
			case '\\':
				c.dx, c.dy = c.dy, c.dx
			case '/':
				c.dx, c.dy = -c.dy, -c.dx
			case '+':
				if c.t == 0 {
					c.dx, c.dy = c.dy, -c.dx
				} else if c.t == 2 {
					c.dx, c.dy = -c.dy, c.dx
				}
				c.t = (c.t + 1) % 3
			}
		}
	}
	for _, c := range carts {
		if !c.dead {
			log.Printf("Final location: %d, %d", c.x, c.y)
		}
	}
}

func main() {
	carts, points := readInput("day13.txt")
	findCrashes(carts, points)
}
