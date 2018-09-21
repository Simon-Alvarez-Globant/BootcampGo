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

//    • Write a program with three functions. One will send stuff(whatever you'd like) over a channel every one second and one will receive it and print it. The third function will tell the other two functions to stop and return(it could be the main func) after 5 seconds. NOTE: the program can not end until the sender and receiver have returned.
