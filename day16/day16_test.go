package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunOps(t *testing.T) {
	before := [4]int{3, 2, 1, 1}
	op := [4]int{9, 2, 1, 2}
	after := [4]int{3, 2, 2, 1}
	matches := runAllOps(before, op, after)
	cnt := len(matches)
	assert.Equal(t, 3, cnt)

	reg := mulr(2, 1, 2, before)
	assert.Equal(t, 2, reg[2])

	reg = muli(2, 1, 2, before)
	assert.NotEqual(t, 2, reg[2])

	reg = addi(2, 1, 2, before)
	assert.Equal(t, 2, reg[2])

	reg = seti(2, 1, 2, before)
	assert.Equal(t, 2, reg[2])
}

func TestParsing(t *testing.T) {
	before := parseReg("Before: [3, 2, 1, 1]")
	assert.Equal(t, [4]int{3, 2, 1, 1}, before)

	after := parseReg("After:  [3, 2, 2, 1]")
	assert.Equal(t, [4]int{3, 2, 2, 1}, after)

	op := parseOp("9 2 1 2")
	assert.Equal(t, [4]int{9, 2, 1, 2}, op)

	op2 := parseOp("9 2 11 2")
	assert.Equal(t, [4]int{9, 2, 11, 2}, op2)
}

func TestOps(t *testing.T) {
	assert.Len(t, opFuncs, 16)
}

func TestGtri(t *testing.T) {
	before := [4]int{3, 0, 0, 3}
	after := [4]int{3, 1, 0, 3}
	reg := gtri(3, 0, 1, before)
	assert.Equal(t, after, reg)

	before = [4]int{1, 0, 3, 1}
	after = [4]int{1, 0, 0, 1}
	reg = gtri(2, 3, 2, before)
	assert.Equal(t, after, reg)
}

func TestGtrr(t *testing.T) {
	before := [4]int{2, 0, 0, 0}
	after := [4]int{2, 0, 1, 0}
	reg := gtrr(0, 3, 2, before)
	assert.Equal(t, after, reg)

}
