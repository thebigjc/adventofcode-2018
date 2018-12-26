package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

type point struct {
	x       int
	y       int
	z       int
	w       int
	cluster map[*point]bool
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func (i *point) distance(j *point) int {
	return abs(i.x-j.x) + abs(i.y-j.y) + abs(i.z-j.z) + abs(i.w-j.w)
}

func (i *point) join(j *point) {
	if i.cluster == nil && j.cluster == nil {
		newCluster := map[*point]bool{i: true, j: true}
		i.cluster = newCluster
		j.cluster = newCluster
	} else if i.cluster == nil && j.cluster != nil {
		i.cluster = j.cluster
		i.cluster[i] = true
	} else if i.cluster != nil && j.cluster == nil {
		j.cluster = i.cluster
		j.cluster[j] = true
	} else {
		log.Printf("Merging a cluster of %d and %d", len(i.cluster), len(j.cluster))
		oldCluster := i.cluster
		for k, v := range oldCluster {
			j.cluster[k] = v
			k.cluster = j.cluster
		}
		log.Printf("Resulting cluster: %d", len(j.cluster))
	}
}

func (i *point) clusterString() string {
	return fmt.Sprintf("%v", reflect.ValueOf(i.cluster).Pointer())
}

func main() {
	bs, err := ioutil.ReadFile("day25.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")
	points := []*point{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		p := point{}

		_, err := fmt.Sscanf(line, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.w)
		if err != nil {
			panic(err)
		}

		points = append(points, &p)
	}

	for i := 0; i < len(points); i++ {
		for j := i; j < len(points); j++ {
			d := points[i].distance(points[j])
			if d <= 3 {
				points[i].join(points[j])
			}
		}
	}

	pointMap := make(map[string]bool)
	for i := 0; i < len(points); i++ {
		str := points[i].clusterString()
		fmt.Printf("Cluster: %s\n", str)
		pointMap[str] = true
	}

	fmt.Printf("Part1: %d\n", len(pointMap))
}
