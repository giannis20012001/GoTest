package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

/**
 * Created by John Tsantilis 
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 26/9/2017.
 */

type Person struct {
	Name string
	Phone string ",omitempty"
}

func main() {
	data, err := bson.Marshal(&Person{Name:"Bob"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", data)
}