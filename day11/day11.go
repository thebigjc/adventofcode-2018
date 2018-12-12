package main

import "log"

func i(x, y, input int) int {
	rackID := x + 10
	power := rackID * y
	power += input
	power *= rackID
	power /= 100
	power %= 10
	power -= 5
	return power
}

const size = 300

func main() {
	x, y, power := part1(8868)
	log.Printf("%d,%d @ %d", x, y, power)
	x, y, size, power := part2(8868)
	log.Printf("%d,%d,%d @ %d", x, y, size, power)
}

func firstPass(input int) [301][301]int {
	var sum [301][301]int
	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			sum[y][x] = i(x, y, input) + sum[y-1][x] + sum[y][x-1] - sum[y-1][x-1]
		}
	}

	return sum
}

func bestForSize(input, s int, sum [301][301]int) (int, int, int) {
	best := 0
	bx := -1
	by := -1

	for y := s; y <= 300; y++ {
		for x := s; x <= 300; x++ {
			total := sum[y][x] - sum[y-s][x] - sum[y][x-s] + sum[y-s][x-s]
			if total > best {
				best = total
				bx = x
				by = y
			}
		}
	}

	return bx - s + 1, by - s + 1, best
}

func part1(input int) (int, int, int) {
	sum := firstPass(input)

	return bestForSize(input, 3, sum)
}

func part2(input int) (int, int, int, int) {
	sum := firstPass(input)

	best := 0
	bx := -1
	by := -1
	bs := -1

	for size := 1; size <= 300; size++ {
		x, y, power := bestForSize(input, size, sum)
		if power > best {
			best = power
			bx = x
			by = y
			bs = size
		}
	}

	return bx, by, bs, best
}
