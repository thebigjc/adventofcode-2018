package main

import (
	"container/list"
	"io/ioutil"
	"log"
	"math"
)

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
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

func collapseSlice(bs []byte) bool {
	l := len(bs)
	if l < 2 {
		return false
	}

	b := int(bs[l-1])
	pb := int(bs[l-2])

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

func collapsedWithFilterSlice(bs []byte, f byte) []byte {
	l := make([]byte, 0, len(bs))

	for _, b := range bs {
		if b == f || b-32 == f {
			continue
		}
		l = append(l, b)

		for collapseSlice(l) {
			l = l[:len(l)-2]
		}
	}

	return l
}

func main() {
	bs, err := ioutil.ReadFile("day05.txt")

	if err != nil {
		panic(err)
	}

	l := collapsedWithFilterSlice(bs, 0)

	log.Printf("There are %d items left", len(l))

	minLen := math.MaxInt32
	for i := 'A'; i <= 'Z'; i++ {
		l := collapsedWithFilterSlice(bs, byte(i))
		listLen := len(l)
		if listLen < minLen {
			minLen = listLen
		}
	}

	log.Printf("The shortest polymer is %d", minLen)
}
