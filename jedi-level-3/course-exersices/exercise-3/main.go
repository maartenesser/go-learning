package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("This should not print fasle")
	default:
		fmt.Println("hello this is default")

	}
}
