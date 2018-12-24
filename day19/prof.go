package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"sync/atomic"
	"time"
)

func profile() {
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

	r0 = 1

	ticker := time.NewTicker(time.Second * 3)
	go func() {
		for range ticker.C {
			fmt.Printf("%d %d %d %d %d %d\n",
				atomic.LoadInt64(&r0),
				atomic.LoadInt64(&r1),
				atomic.LoadInt64(&r2),
				atomic.LoadInt64(&r3),
				atomic.LoadInt64(&r4),
				atomic.LoadInt64(&r5),
			)
		}
	}()
}
