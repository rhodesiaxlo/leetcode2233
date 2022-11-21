package main

import (
	"encoding/json"
	"fmt"
)

func main() {

}

func jsonConvert()  {
	jsonStr := `{"id":1058,"name":"RyuGou"}`
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &jsonData)

	sum :=  int(jsonData["id"].(float64)) + 3 // noteice
	// string string
	// int  float64
	// bool bool
	// []interface  array
	// map[string]interface object
	fmt.Println(sum)
}

// Q2
type Father interface {
	Hello()
}

type Child struct {
	Name string
}

func (s Child)Hello()  {

}
func f(out *Father){
	if out != nil{
		fmt.Println("surprise!")
	}
}
func ptrToInterface()  {
	//var buf  Child
	//buf = Child{}
	//f(&buf)
}
