package main

import (
	"fmt"
)

var nice = []int{1, 2, 4, 8, 16, 32, 64, 128}

func simpleRangeMain() {
	for i, v := range nice {
		fmt.Printf("2**%d=%d\n", i, v)
	}
}

func ssMain() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = i + 2
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
