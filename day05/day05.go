package main

import (
	"container/list"
	"io/ioutil"
	"log"
	"math"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func collapse(e *list.Element) bool {
	if e == nil {
		return false
	}
	pe := e.Prev()
	if pe == nil {
		return false
	}
	b := int(e.Value.(byte))
	pb := int(pe.Value.(byte))
	//log.Printf("%c - %c = %d", b, pb, abs(b-pb))
	return abs(b-pb) == 32
}

func collapsedWithFilter(bs []byte, f byte) *list.List {
	l := list.New()

	for _, b := range bs {
		if b == f || b-32 == f {
			continue
		}
		l.PushBack(b)

		for collapse(l.Back()) {
			l.Remove(l.Back())
			l.Remove(l.Back())
		}
	}

	return l
}

func main() {
	bs, err := ioutil.ReadFile("day05.txt")

	if err != nil {
		panic(err)
	}

	l := collapsedWithFilter(bs, 0)

	log.Printf("There are %d items left", l.Len())

	minLen := math.MaxInt32
	for i := 'A'; i <= 'Z'; i++ {
		l := collapsedWithFilter(bs, byte(i))
		listLen := l.Len()
		if listLen < minLen {
			minLen = listLen
		}
	}

	log.Printf("The shortest polymer is %d", minLen)
}
