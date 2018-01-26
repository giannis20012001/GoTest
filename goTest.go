package main

/**
 * Created by John Tsantilis (A.K.A lumi) on 24/1/2018.
 * @author John Tsantilis <i.tsantilis [at] yahoo [dot] com>
 */

import (
	"fmt"
	"sync"
	"time"
	"runtime"
	"math/rand"
	"github.com/giannis20012001/GoTest/util"
)

//Boolean flag used for custom metrics
var BLOCK_CUSTOM_METRICS = true
//Threading variable
var WaitGroup sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

}

func main() {
	channel := make(chan util.ComponentState)// create channel
	cs := util.NewComponentState("BOOTSTRAPPED")

	WaitGroup.Add(1)
	go stompServer(channel) // run this in a separate goroutine
	WaitGroup.Add(1)
	go restServer()
	WaitGroup.Add(1)
	go allowCustomMetrics(channel)
	WaitGroup.Add(1)
	go sendMetricsToStompServer(channel)

	channel <- *cs

	WaitGroup.Wait()

}

//Writes to Channel
func stompServer(channel chan util.ComponentState) {
	defer WaitGroup.Done()

	cs := <- channel
	for i:=1; i < 20; i++{
		fmt.Println("Getting stuff from STOMP server......")

		switch i {
			case 1:
				cs.SetState("INITIALIZED")
				channel <- cs
				fmt.Println("State is set to: " + cs.GetState()) //INITIALIZED
			case 2:
				cs.SetState("DEPLOYED")
				channel <- cs
				fmt.Println("State is set to: " + cs.GetState()) //DEPLOYED
			case 3:
				cs.SetState("BLOCKED")
				channel <- cs
				fmt.Println("State is set to: " + cs.GetState()) //BLOCKED
			case 4:
				cs.SetState("STARTED")
				channel <- cs
				fmt.Println("State is set to: " + cs.GetState()) //STARTED
			case 5:
				cs.SetState("STOPPED")
				channel <- cs
				fmt.Println("State is set to: " + cs.GetState()) //STOPPED
			case 6:
				cs.SetState("UNDEPLOYED")
				channel <- cs
				fmt.Println("State is set to: " + cs.GetState()) //UNDEPLOYED
			case 7:
				cs.SetState("ERRONEOUS")
				channel <- cs
				fmt.Println("State is set to: " + cs.GetState()) //ERRONEOUS
			case 8:
				cs.SetState("CHAINED")
				channel <- cs
				fmt.Println("State is set to: " + cs.GetState()) //CHAINED
			default:
				fmt.Println("State is set to: " + cs.GetState())

		}

		//r := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(5000) * time.Microsecond)
		//time.Sleep(500 * time.Millisecond)

	}

	close(channel)

}

//Doesn't need Channel
func restServer() {
	defer WaitGroup.Done()

	for {
		fmt.Println("Doing REST stuff......")
		r := rand.Intn(5000000-1) + 1
		//r := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(r) * time.Microsecond)
		//time.Sleep(500 * time.Millisecond)

	}

}

//Reads from Channel
func allowCustomMetrics(channel <-chan util.ComponentState) {
	defer WaitGroup.Done()

	cs := <-channel
	if cs.GetState() == "STARTED" {
		fmt.Println("Executing allowCustomMetrics()......")
		time.Sleep(60 * time.Second)
		BLOCK_CUSTOM_METRICS = false

	}

}

//Reads from Channel
func sendMetricsToStompServer(channel <-chan util.ComponentState) {
	defer WaitGroup.Done()

	cs := <- channel
	for  {
		if cs.GetState() == "STARTED" {
			fmt.Println("Sending stuff to STOMP server......")
			if !BLOCK_CUSTOM_METRICS {
				fmt.Println("Sending extra stuff to STOMP server......")

			}

			r := rand.Intn(500-0) + 0
			//r := rand.New(rand.NewSource(time.Now().UnixNano()))
			time.Sleep(time.Duration(r) * time.Microsecond)
			//time.Sleep(500 * time.Millisecond)

		}

	}

}