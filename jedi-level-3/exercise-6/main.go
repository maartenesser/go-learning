package main

import "fmt"

func main() {
	x := "noon"

	if x == "morning" {
		fmt.Println("Time for a fresh brewed coffee")
	} else if x == "noon" {
		fmt.Println("Time for lunch")
	} else {
		fmt.Println("please go back to work")
	}
}
