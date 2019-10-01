package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)

	fmt.Println("slice with appended 52 =>", x)

	x = append(x, 53, 54, 55)

	fmt.Println("added some more values to the slice => ", x)

	y := []int{56, 57, 58, 59, 60}

	for _, v := range y {

		x = append(x, v)
		// x = append(x, y...)
	}

	fmt.Println(x)
}