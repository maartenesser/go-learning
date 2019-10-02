package main

import "fmt"

type person struct {
	first_name, last_name, favorite_ice_cream string
}

func main() {
	maarten := person{
		first_name:         "Maaarten",
		last_name:          "Esser",
		favorite_ice_cream: "Strawberry",
	}

	fmt.Println(maarten.first_name)
	fmt.Printf("First name: %s\n", maarten.first_name)
	fmt.Printf("Last name: %s\n", maarten.last_name)
	fmt.Printf("Fav ice cream: %s\n", maarten.favorite_ice_cream)
}
