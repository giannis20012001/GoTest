package main

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
)

/**
 * Created by John Tsantilis (A.K.A lumi) on 1/3/2018.
 * @author John Tsantilis <i.tsantilis [at] yahoo [dot] com>
 */

func main() {
	yolo := make( []string, 1 )
	yolo[0] = "tcp"
	v, _ := net.IOCounters(false)

	// almost every return value is a struct
	//fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v[0].PacketsRecv)
	fmt.Println(v[0].PacketsSent)

}