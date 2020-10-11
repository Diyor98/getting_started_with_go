package main

import (
	"encoding/json"
	f "fmt"
)

func main() {
	var name string
	var address string

	f.Println("Enter your name: ")
	f.Scanln(&name)
	f.Println("Enter you address: ")
	f.Scanln(&address)

	mymap := make(map[string]string)
	mymap[name] = address

	myjson, _ := json.Marshal(mymap)

	f.Println(string(myjson))
}
