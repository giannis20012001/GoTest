package OLD

import (
	"fmt"
)

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 26/10/2017.
 */

func main() {
	var data []string
	var finalData map[string][]string

	temp := make(map[string][]string)
	finalData = temp

	data = append(data, "192.168.7.50")
	data = append(data, "45045")

	finalData["mysqluri"] = data
	finalData["mysqlport"] = data

	fmt.Println(finalData)

}
