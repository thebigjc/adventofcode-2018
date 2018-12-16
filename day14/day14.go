package main

import (
	"log"
)

func indexOf(l []int, m []int) int {
	for i := 0; i < len(l)-len(m); i++ {
		found := true
		for j := 0; j < len(m); j++ {
			if l[i+j] != m[j] {
				found = false
			}
		}
		if found {
			return i
		}
	}

	return -1
}

func intToInts(x int) []int {
	origX := x
	a := make([]int, 0)
	for x > 0 {
		a = append(a, x%10)
		x /= 10
	}

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	log.Printf("%d == %v", origX, a)
	return a
}

func doElves(recipes int) string {
	l := []int{3, 7}

	first := 0
	second := 1

	needed := recipes*100 + 10

	for needed > 0 {
		if (needed % 10000) == 0 {
			log.Printf("Needed:%d", needed)
		}

		newScore := l[first] + l[second]
		if newScore >= 10 {
			l = append(l, 1)
			needed--
		}
		l = append(l, newScore%10)

		needed--
		first = (first + l[first] + 1) % len(l)
		second = (second + l[second] + 1) % len(l)
	}

	out := make([]rune, 10)
	for i := 0; i < 10; i++ {
		out[i] = rune(l[recipes+i] + '0')
	}

	i := indexOf(l, intToInts(recipes))

	log.Printf("Found at %d", i)

	return string(out)
}

func main() {
	out := doElves(765071)
	log.Printf("Output: %s", out)
}
