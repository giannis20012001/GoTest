package restserver

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 5/7/2017.
 */

import (
	"io"
	"fmt"
	"strconv"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"

)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	//go exe_cmd("ping","8.8.8.8")

}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)

	}

}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)

}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)

	}

	if err := r.Body.Close(); err != nil {
		panic(err)

	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)

		}

	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)

	}

}

func GetTotalMem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	v, _ := mem.VirtualMemory()
	//Return total RAM in Mega bytes
	if err := json.NewEncoder(w).Encode(strconv.Itoa(int(v.Total) / 1048576)); err != nil {
		panic(err)

	}


}

func GetFreeMem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	v, _ := mem.VirtualMemory()
	//Return free RAM in Mega bytes
	if err := json.NewEncoder(w).Encode(strconv.Itoa(int(v.Free) / 1048576)); err != nil {
		panic(err)

	}

}

func GetUsedMem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	v, _ := mem.VirtualMemory()
	//Return used RAM in Mega bytes
	if err := json.NewEncoder(w).Encode(strconv.Itoa(int(v.Used) / 1048576)); err != nil {
		panic(err)

	}

}

func GetFreeMemPercentage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	v, _ := mem.VirtualMemory()
	//Return free RAM in percentage
	freepercentage := ((float64(v.Free) / 1048576) * 100) / (float64(v.Total) / 1048576)
	if err := json.NewEncoder(w).Encode(strconv.FormatFloat(freepercentage, 'f', 5, 64)); err != nil {
		panic(err)

	}

}

func GetUsedMemPercentage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	v, _ := mem.VirtualMemory()
	//Return used RAM in percentage
	if err := json.NewEncoder(w).Encode(strconv.FormatFloat(v.UsedPercent, 'f', 5, 64)); err != nil {
		panic(err)

	}

}

func GetUptime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	v, _ := host.Info()
	//Return upTime in milliseconds
	if err := json.NewEncoder(w).Encode(strconv.Itoa(int(v.Uptime))); err != nil {
		panic(err)

	}

}

func GetCpuLoad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	v, _ := load.Avg()
	//Return CPU loaf in percentage
	if err := json.NewEncoder(w).Encode(strconv.FormatFloat(v.Load1 * 100, 'f', 0, 64)); err != nil {
		panic(err)

	}

}