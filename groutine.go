package main

import (
	"fmt"
	"time"
)

func test(ch chan int) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int)
	go test(ch)
	go test(ch)
	close(ch)
	time.Sleep(time.Second)
}
