package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := []int{3,4,5}

	//slicing the first half of the x slice
	z := x[0:3]
	y = append(y, z...)

	//slicing the last half of the x slice
	z = x[6:10]
	y = append(y, z...)

	fmt.Println("z slice", z)

	fmt.Println("y slice: ", y)

}