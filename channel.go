package main

import (
	"fmt"
	"sync"
	"time"
)

func recvChan(ch chan int) {
	ch <- 299
}

func main() {
	ch := (chan int)(nil)

	// ch := make(chan int)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		fmt.Println(<-ch)
	}()
	time.Sleep(time.Second)
	// recvChan(ch)
	ch <- 299

	sync.Map{}
	sync.RWMutex
}
