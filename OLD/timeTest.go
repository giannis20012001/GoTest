package OLD

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 25/10/2017.
 */

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today+0:
		fmt.Println("Today")
	case today+1:
		fmt.Println("Tomorrow")
	case today+2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	//time.Sleep(3000 * time.Millisecond)
	time.Sleep(30 * time.Second)
	fmt.Println("Just woke up")

}