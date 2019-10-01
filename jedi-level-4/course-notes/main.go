package main

import (
	"fmt"
)

func main() {
	// COMPOSITE LITERAL
	x := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(x)
	fmt.Printf("The length of the slice = %v\n", len(x))
	fmt.Printf("The first item from the slice = %v\n", x[0])

	// FOR RANGE
	for i, v := range x {
		fmt.Printf("for %v the value is %v\n", i, v)
	}

	// SLICING A SLICE
	fmt.Printf("slicing this slice %v in half %v\n", x, x[3:])
	fmt.Printf("slicing from %v to %v => %v\n", x[1], x[5], x[1:6])

	// MAP
	//type[key] => Composite literal

	m := map[string]int{
		"james":           32,
		"miss Moneypenny": 27,
	}
	fmt.Printf("printing the map: %v\n", m)
	fmt.Printf("printing the age of james %v\n", m["james"])

	fmt.Println(m["Barnabas"])

	// Comma OK idiom checking if value is ok in the map
	if v, ok := m["miss Moneypenny"]; ok {
		fmt.Println("This is the if print", v)
	}

	// Loop through the slice and print out the keys
	for k, v := range m {
		fmt.Printf("Printing the key with the value: key: %v, value: %v\n", k, v)
	}
	fmt.Println("")

	// Loop trough the slice and print out the index
	for i, v := range x {
		fmt.Printf("Printing the index for every value: index %v, value: %v\n", i, v)
	}

	// Delete map
	//checking if james has a value if jes then delete
	if v, ok := m["james"]; ok {
		fmt.Println("value:", v)
		delete(m, "james")
	}
	fmt.Printf("removing %v", m["james"])
	fmt.Println(m)
}

// a SLICE allows you to group together VALUES of the same TYPE
