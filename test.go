package main

import (
	"strings"
	"fmt"
)

/**
 * Created by John Tsantilis 
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 5/9/2017.
 */

const ORCHESTRATOR_URI string = "http://arcadia-sc.euprojects.net"
const CONFIGURATION_API string = "/api/v1/node/%s/config"
const COMPONENT_API string = "/api/v1/node/%s/component"

func main()  {
	id := "192"
	result1 := strings.Replace(CONFIGURATION_API, "%s", id, -1)
	fmt.Println(ORCHESTRATOR_URI + result1)

	fmt.Println()
	result2 := strings.Replace(COMPONENT_API, "%s", id, -1)
	fmt.Println(ORCHESTRATOR_URI + result2)

	fmt.Println()
	value := "http://%s:%q"
	r := strings.NewReplacer("%s", "192.168.30.545", "%q", "8080")
	result := r.Replace(value)
	fmt.Println(result)

}