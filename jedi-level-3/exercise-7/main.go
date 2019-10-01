package main

import (
	"fmt"
)

func main() {
	weather := ""

	if weather == "raining" {
		fmt.Println("Please bring a raincoat or umbrella")
	} else if weather == "sunny" {
		fmt.Println("Please bring yourt sunglasses")
	} else {
		fmt.Println("The weather is neutral so don't bring anything")
	}
}
