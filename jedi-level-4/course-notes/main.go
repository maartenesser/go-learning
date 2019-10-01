package main

import "fmt"

func main() {
	// COMPOSITE LITERAL
	x := []int{1,2,3,4,5,6,7}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])

	// FOR RANGE
	for i, v := range x  {
		fmt.Println(i,v)
	}

}

// a SLICE allows you to group together VALUES of the same TYPE
