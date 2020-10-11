package main

import (
	"fmt"
	"strings"
)

func main() {
	var mystr string

	fmt.Println("Enter a string")
	_, err := fmt.Scanln(&mystr)

	mystr = strings.ToLower(mystr)

	if err == nil {
		if strings.HasPrefix(mystr, "i") && strings.HasSuffix(mystr, "n") && strings.Contains(mystr, "a") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not found!")
		}
	}

}
