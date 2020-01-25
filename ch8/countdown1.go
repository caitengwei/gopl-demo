package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Rocket launching!...")
}
func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("commencing countdown...")
	//tick := time.Tick(1 * time.Second)
	//for countdown := 10; countdown > 0; countdown-- {
	//	println(countdown)
	//	<-tick
	//}
	tick := time.NewTicker(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-tick.C:
			println(countdown)
			// do nothing
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	tick.Stop()

	launch()
}