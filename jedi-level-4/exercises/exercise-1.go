package main

import "fmt"

func main() {
	//print array with 5 int's in array
	var a [5]int

	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	fmt.Println(a)

	for k, v := range a {
		fmt.Printf("key: %v anf the value: %v\n", k, v)
	}

	fmt.Printf("the type of this array is %T\n", a)

}
