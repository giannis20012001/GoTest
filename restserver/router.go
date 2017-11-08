package restserver

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 5/7/2017.
 */

import (
	"log"
	"time"
	"net/http"
	"github.com/gorilla/mux"

)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	r := NewRoutes()

	for _, route := range r.routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router

}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(r.Method, " ", r.RequestURI, " ", name, " ", time.Since(start))

	})

}