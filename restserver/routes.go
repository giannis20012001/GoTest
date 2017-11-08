package restserver

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 5/7/2017.
 */

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc

}

type Routes struct {
	routes []Route
}

func NewRoutes() *Routes {
	r := new(Routes)
	r.routes = make([]Route, 0)
	r.routes = append(r.routes, Route{
		"Index",
		"GET",
		"/",
		Index,
	})

	r.routes = append(r.routes, Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	})

	r.routes = append(r.routes, Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	})

	r.routes = append(r.routes, Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	})

	//==================================================================================================================
	//Add metrics routes
	//==================================================================================================================
	r.routes = append(r.routes, Route{
		"osMemoryTotal",
		"GET",
		"/os/memory/total",
		GetTotalMem,
	})

	r.routes = append(r.routes, Route{
		"osMemoryFree",
		"GET",
		"/os/memory/free",
		GetFreeMem,
	})

	r.routes = append(r.routes, Route{
		"osMemoryUsed",
		"GET",
		"/os/memory/used",
		GetUsedMem,
	})

	r.routes = append(r.routes, Route{
		"osMemoryFreePercentage",
		"GET",
		"/os/memory/freepercentage",
		GetFreeMemPercentage,
	})

	r.routes = append(r.routes, Route{
		"osMemoryUsedPercentage",
		"GET",
		"/os/memory/usedpercentage",
		GetUsedMemPercentage,
	})

	r.routes = append(r.routes, Route{
		"osUpTime",
		"GET",
		"/os/uptime",
		GetUptime,
	})

	r.routes = append(r.routes, Route{
		"osCpuLoad",
		"GET",
		"/os/cpu_load",
		GetCpuLoad,
	})

	return r

}