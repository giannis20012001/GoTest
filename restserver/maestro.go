package restserver

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 5/7/2017.
 */

import (
	"log"
	"net/http"

)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}