package main

import "fmt"

//furling
func total(i ...int) int {

	sum := 0
	for _, v := range i {
		sum += v
	}

	return sum

}

func main() {
	y := total(1,2,3,4,5)
	fmt.Println(y)
}