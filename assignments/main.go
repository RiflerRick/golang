package main

import (
	"fmt"
)

func IsEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

func main() {
	n := []int{}

	for i := 0; i < 10; i++ {
		n = append(n, i)
	}

	for _, i := range n {
		if IsEven(i) {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
