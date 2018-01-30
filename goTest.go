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
	//startMetrics := make(chan bool)

	WaitGroup.Add(1)
	go stompServer(channel) // run this in a separate goroutine
	WaitGroup.Add(1)
	go restServer()

	for {
		select {
		case msg := <-channel:
			if msg.GetState() == "STARTED" {
				WaitGroup.Add(1)
				go allowCustomMetrics()
				WaitGroup.Add(1)
				go sendMetricsToStompServer()

			}

		}

	}

	WaitGroup.Wait()

}

//Writes to Channel
func stompServer(channel chan util.ComponentState) {
	defer WaitGroup.Done()

	//cs := <- channel
	cs := util.NewComponentStateEmpty()
	i:= 1
	for {
		switch i {
		case 1:
			cs.SetState("BOOTSTRAPPED")
			channel <- *cs
			fmt.Println("Getting stuff from STOMP server......")
			fmt.Println("State is set to: " + cs.GetState()) //BOOTSTRAPPED
		case 2:
			cs.SetState("INITIALIZED")
			channel <- *cs
			fmt.Println("Getting stuff from STOMP server......")
			fmt.Println("State is set to: " + cs.GetState()) //INITIALIZED
		case 3:
			cs.SetState("DEPLOYED")
			channel <- *cs
			fmt.Println("Getting stuff from STOMP server......")
			fmt.Println("State is set to: " + cs.GetState()) //DEPLOYED
		case 4:
			cs.SetState("BLOCKED")
			channel <- *cs
			fmt.Println("Getting stuff from STOMP server......")
			fmt.Println("State is set to: " + cs.GetState()) //BLOCKED
		case 5:
			cs.SetState("STARTED")
			channel <- *cs
			fmt.Println("Getting stuff from STOMP server......")
			fmt.Println("State is set to: " + cs.GetState()) //STARTED
		case 6:
			cs.SetState("STOPPED")
			channel <- *cs
			fmt.Println("Getting stuff from STOMP server......")
			fmt.Println("State is set to: " + cs.GetState()) //STOPPED
		case 7:
			cs.SetState("UNDEPLOYED")
			channel <- *cs
			fmt.Println("Getting stuff from STOMP server......")
			fmt.Println("State is set to: " + cs.GetState()) //UNDEPLOYED
		case 8:
			cs.SetState("ERRONEOUS")
			channel <- *cs
			fmt.Println("Getting stuff from STOMP server......")
			fmt.Println("State is set to: " + cs.GetState()) //ERRONEOUS
		case 9:
			cs.SetState("CHAINED")
			channel <- *cs
			fmt.Println("Getting stuff from STOMP server......")
			fmt.Println("State is set to: " + cs.GetState()) //CHAINED
		default:
			//fmt.Println("State is set to: " + cs.GetState())
			//Do nothing

		}

		//r := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(5000) * time.Microsecond)
		//time.Sleep(500 * time.Millisecond)

		i++

	}

	close(channel)

}

//Reads from Channel
func sendMetricsToStompServer() {
	defer WaitGroup.Done()

	for  {
		fmt.Println("Sending stuff to STOMP server......")
		if !BLOCK_CUSTOM_METRICS {
			fmt.Println("Sending extra stuff to STOMP server......")

		}

		r := rand.Intn(5000000-1) + 1
		//r := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(r) * time.Microsecond)
		//time.Sleep(500 * time.Millisecond)

	}

}

//Reads from Channel
func allowCustomMetrics() {
	defer WaitGroup.Done()
	fmt.Println("Executing allowCustomMetrics()......")
	time.Sleep(60 * time.Second)
	BLOCK_CUSTOM_METRICS = false
	fmt.Printf("BLOCK_CUSTOM_METRICS: %t\n", BLOCK_CUSTOM_METRICS)

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