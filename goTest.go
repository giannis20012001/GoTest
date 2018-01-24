package main

/**
 * Created by John Tsantilis (A.K.A lumi) on 24/1/2018.
 * @author John Tsantilis <i.tsantilis [at] yahoo [dot] com>
 */

import (
	"fmt"
	"time"
)

//Boolean flag used for custom metrics
var BLOCK_CUSTOM_METRICS = true

func main() {
	channel := make(chan ComponentState)// create channel

	go stompServer(channel) // run this in a separate goroutine
	go restServer(channel)
	go allowCustomMetrics(channel)
	go sendMetricsToStompServer("Hello!", channel)

	for {
		fmt.Printf("%s\n", <-channel) // read from channel and print out message

	}

	fmt.Println("Cool, that's all I wanted to say")

}

func stompServer(channel chan ComponentState) {
	for {
		channel <- fmt.Sprintf("stompServer")
		time.Sleep(500 * time.Millisecond)

	}

}

func restServer(channel chan ComponentState) {
	for {
		channel <- fmt.Sprintf("restServer")
		time.Sleep(500 * time.Millisecond)

	}

}

func allowCustomMetrics(channel chan ComponentState) {
	time.Sleep(60 * time.Second)
	BLOCK_CUSTOM_METRICS = false

}

func sendMetricsToStompServer(channel chan ComponentState) {
	for i := 0; i < 5; i++ {
		channel <- fmt.Sprintf("I wanted to say '%s' for the %dth time", s, i)
		time.Sleep(500 * time.Millisecond)

	}

}