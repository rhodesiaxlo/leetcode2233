package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	//takeAddressOfMapEle()

	mapIsNotOrdered()
}

// Q1
func takeAddressOfMapEle() *string {
	m := make(map[string]string, 10)
	m["name"] = "sennalu"
	//return &m["name"] // can not take address of map element
	return nil
}

// Q2
func mapIsNotOrdered() {
	m := make(map[string]string, 10)
	for i := 0; i < 10; i++ {
		tmp := strconv.Itoa(i)
		m[tmp] = tmp
	}
	for i := 0; i < 10; i++ {
		tmp := strconv.Itoa(i)
		fmt.Println(m[tmp])
	}

	data, _ := json.Marshal(m)
	fmt.Println(string(data))
}

