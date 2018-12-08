package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// States is the list of possible states
type States int

const (
	numChildren States = iota
	numMetadata
	metadata
)

func (s States) String() string {
	names := [...]string{
		"numChildren",
		"numMetadata",
		"metadata",
	}
	return names[s]
}

type state struct {
	state      States
	childNodes int
	metadatas  int
	value      int
	children   []*state
}

func newState() *state {
	return &state{numChildren, 0, 0, 0, nil}
}

func main() {
	bs, err := ioutil.ReadFile("day08.txt")
	if err != nil {
		panic(err)
	}
	numStrs := strings.Split(string(bs), " ")

	curState := newState()
	root := curState
	sum := 0

	var states []*state

	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(nil)
		}
		switch curState.state {
		case numChildren:
			curState.childNodes = num
			curState.state = numMetadata

		case numMetadata:
			curState.metadatas = num
			curState.state = metadata

			if curState.childNodes > 0 {
				// push
				states = append(states, curState)
				nextState := newState()
				curState.children = append(curState.children, nextState)
				curState = nextState
			}
		case metadata:
			sum += num
			if len(curState.children) == 0 {
				curState.value += num
			} else {
				offset := num - 1
				if offset >= 0 && offset < len(curState.children) {
					v := curState.children[offset].value
					curState.value += v
				}
			}
			curState.metadatas--
			if curState.metadatas == 0 {
				if len(states) > 0 {
					// peek
					curState = states[len(states)-1]
					curState.childNodes--

					if curState.childNodes > 0 {
						nextState := newState()
						curState.children = append(curState.children, nextState)
						curState = nextState
					} else {
						// pop
						states = states[:len(states)-1]
					}
				} else {
					curState = newState()
					states = nil
				}

			}
		}
	}
	log.Printf("Sum: %d", sum)
	log.Printf("Root value: %d", root.value)
}
