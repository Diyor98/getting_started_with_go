package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 3, 10)
	var num string
	var counter int = 0
	for true {
		fmt.Scanln(&num)
		if num == "X" {
			break
		}
		i, err := strconv.Atoi(num)
		if err == nil {
			if counter >= 3 {
				slice = append(slice, i)
			} else {
				slice[0] = i
			}
			sort.Ints(slice)
			fmt.Println(slice)
			counter++
		}
	}
}
