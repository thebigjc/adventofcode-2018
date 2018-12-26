package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/signal"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

func main() {
	f, err := os.Create("proc.prof")
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Got Signal")
		pprof.StopCPUProfile()
		err = f.Close()
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}()

	ticker := time.NewTicker(time.Second * 3)
	go func() {
		for range ticker.C {
			fmt.Printf("%d : %d %d %d %d %d %d. len(r3s): %d\n",
				atomic.LoadInt64(&n),
				atomic.LoadInt64(&r0),
				atomic.LoadInt64(&r1),
				atomic.LoadInt64(&r2),
				atomic.LoadInt64(&r3),
				atomic.LoadInt64(&r4),
				atomic.LoadInt64(&r5),
				len(r3s),
			)
		}
	}()

	bs, err := ioutil.ReadFile("out.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bs), "\n")
	maxn := int64(math.MinInt64)
	bestr0 := int64(math.MinInt64)

	lines = []string{"0"}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		r0, r1, r2, r3, r4, r5 = 0, 0, 0, 0, 0, 0
		n = 0
		r0, err = strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		newn := loop(math.MaxInt64)
		if newn > maxn {
			maxn = newn
			bestr0 = r0
			fmt.Printf("Best r0: %d\n", bestr0)
		}
		fmt.Printf("Best r0: %d\n", bestr0)
	}
	fmt.Printf("Bestest r0: %d\n", bestr0)
}
