package main

import (
	"fmt"
)

func main() {
	sl := make([]int, 10)

	for i := 0; i < 10; i++ {
		sl[i] = i
	}
	for i := 0; i < 10; i++ {
		fmt.Println(sl[i:i : i+1])
	}
}
