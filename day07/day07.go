package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type node struct {
	name           rune
	outgoing       []*node
	incoming       int
	completionTime int
}

func (n *node) start(time int) {
	n.completionTime = time + int(n.name-'A'+1) + 60
	log.Printf("Starting %c at %d until %d", n.name, time, n.completionTime)
}
func makeNode(name rune) *node {
	return &node{name, nil, 0, 0}
}

func (n *node) addOutgoingEdge(b *node) {
	n.outgoing = append(n.outgoing, b)
	b.incoming++
}

func sortNodes(nodes []*node) {
	sort.Slice(nodes, func(i, j int) bool { return nodes[i].name < nodes[j].name })
}

func noincomingNodes(nodes map[rune]*node) []*node {
	s := []*node{}
	for _, v := range nodes {
		if v.incoming == 0 {
			s = append(s, v)
		}
	}
	return s
}

func part1() {
	nodes := readNodes()

	L := []rune{}
	S := noincomingNodes(nodes)

	for len(S) > 0 {
		sortNodes(S)

		n := S[0]
		S = S[1:]

		L = append(L, n.name)

		for _, nb := range n.outgoing {
			nb.incoming--
			if nb.incoming == 0 {
				S = append(S, nb)
			}
		}
	}
	log.Printf("Order: %s", string(L))
}

func readNodes() map[rune]*node {
	bs, err := ioutil.ReadFile("day07.txt")

	if err != nil {
		panic(err)
	}

	nodes := make(map[rune]*node)
	lines := strings.Split(string(bs), "\n")
	for _, l := range lines {
		l = strings.TrimSpace(l)
		var a rune
		var b rune

		_, err := fmt.Sscanf(l, "Step %c must be finished before step %c can begin.", &a, &b)
		if err != nil {
			panic(err)
		}
		var na *node
		var nb *node
		var ok bool
		if na, ok = nodes[a]; !ok {
			na = makeNode(a)
			nodes[a] = na
		}
		if nb, ok = nodes[b]; !ok {
			nb = makeNode(b)
			nodes[b] = nb
		}
		na.addOutgoingEdge(nb)
	}

	return nodes
}

func part2() {
	nodes := readNodes()

	S := noincomingNodes(nodes)
	time := -1
	workers := 5
	var workQueue []*node

	for len(S) > 0 || len(workQueue) > 0 {
		time++

		if len(workQueue) > 0 {
			newWQ := []*node{}
			for _, wq := range workQueue {
				if wq.completionTime > time {
					newWQ = append(newWQ, wq)
					continue
				}
				for _, nb := range wq.outgoing {
					nb.incoming--
					if nb.incoming == 0 {
						S = append(S, nb)
					}
				}
			}
			workQueue = newWQ
		}

		sortNodes(S)

		for len(S) > 0 && len(workQueue) < workers {
			n := S[0]

			workQueue = append(workQueue, n)
			n.start(time)
			S = S[1:]
		}

		sort.Slice(workQueue, func(i, j int) bool { return workQueue[i].completionTime < workQueue[i].completionTime })
	}

	log.Printf("Time: %d", time)
}

func main() {
	part1()
	part2()
}
