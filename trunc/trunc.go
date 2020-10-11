package main

import "fmt"

func main() {
	var mynum float32

	fmt.Printf("Enter a floating point number : \n")
	fmt.Scan(&mynum)
	fmt.Printf("%d\n", int32(mynum))
}
