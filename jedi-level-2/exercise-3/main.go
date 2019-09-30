package main

import (
	"fmt"
)

//create TYPED and UNTYPED constant\

const (
	maarten = 26
	name string = "Maarten"
)

func main() {
	const name = "Maarten"
	fmt.Println(maarten)
	fmt.Println(name)
}