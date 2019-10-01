package main

import "fmt"

func main() {
	s := []int {1,2,3,4,5,6,7,8,9,10}

	fmt.Println(s)

	for i, v :=range s {
		fmt.Printf("key: %v, value: %v\n", i, v)
	}

	fmt.Printf("TYPE of the slice: %T\n", s)
}