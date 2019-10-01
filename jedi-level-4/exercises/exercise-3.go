package main

import "fmt"

func main() {
	s := []int {42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("the whole slice:", s)

	fmt.Printf("first slice: %v\n", s[0:5])
	fmt.Printf("first slice: %v\n", s[5:10])
	fmt.Printf("first slice: %v\n", s[2:7])
	fmt.Printf("first slice: %v\n", s[1:6])

}