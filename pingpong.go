package main

import (
	"fmt"
	"time"
)

func main() {
	pinger := make(chan string, 1)
	ponger := make(chan string, 1)

	go ping(pinger, ponger)
	go pong(ponger, pinger)

	pinger <- "s"

	var input string
	fmt.Scanln(&input)
}

func ping(pinger chan<- string, ponger <-chan string) {
	for {
		pinger <- "ping"
		fmt.Println("pinging...")
		time.Sleep(1 * time.Second)
		<-ponger
	}
}

func pong(ponger chan<- string, pinger <-chan string) {
	for {
		ponger <- "pong"
		fmt.Println("	...ponging")
		time.Sleep(1 * time.Second)
		<-pinger
	}
}