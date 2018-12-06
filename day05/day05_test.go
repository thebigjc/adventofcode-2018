package main

import (
	"io/ioutil"
	"testing"
)

func BenchmarkLinkedList(b *testing.B) {
	bs, err := ioutil.ReadFile("day05.txt")
	if err != nil {
		panic(nil)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		collapsedWithFilter(bs, 0)
	}
}
func BenchmarkSlice(b *testing.B) {
	bs, err := ioutil.ReadFile("day05.txt")
	if err != nil {
		panic(nil)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		collapsedWithFilterSlice(bs, 0)
	}
}
