package main

import "fmt"

func main() {

	x := "sunny"
	switch x {
	case "sunny":
		{
			fmt.Println("it's going to be sunny so please bring your sunglasses")
		}
	case "rainy":
		{
			fmt.Println("It's going to rain so please bring a raincote ")
		}
	default:
		{
			fmt.Println("Weather is neutral so you don't need to bring anything")
		}
	}
}
