package main

import (
	"fmt"
	"time"
)

func send(clock chan<- int) {
	count := 0
	for {
		time.Sleep(time.Second)
		clock <- count
		count++
	}
}

func receive(clock <-chan int) {
	for {
		c := <-clock
		fmt.Println(c)

	}
}

func main() {
	clock := make(chan int, 1)
	go send(clock)
	go receive(clock)
	time.Sleep(time.Second * 5)
	<-clock
	close(clock)

}
