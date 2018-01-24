package OLD

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 8/11/2017.
 */

import (
	"fmt"
	"sync"
	"time"
)

//Threading variable
var WaitGroup sync.WaitGroup

func main() {
	go sendMetricsToSTOPMServer()
	WaitGroup.Add(1)

	WaitGroup.Wait()

}

func sendMetricsToSTOPMServer() {
	c := time.Tick(2 * time.Second)
	for now := range c {
		fmt.Println("tick", now)

	}

}