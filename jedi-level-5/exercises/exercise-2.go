package main

import "fmt"

type person struct {
	first_name, last_name string
	favorite_ice_cream []string
}

func main() {
	maarten := person{
		first_name:         "Maaarten",
		last_name:          "Esser",
		favorite_ice_cream: []string{
			"Strawberry",
			"citron",
			},
	}

	// This is not the right way. Far to static and not at all dynamic
	//m := map[string][]string{
	//	`first_name` : {maarten.first_name},
	//	`last_name` : {maarten.last_name},
	//	`favorite_ice_cream` : maarten.favorite_ice_cream,
	//}

	//This is the best way. super dynamic and much easier to use
	n := map[string]person{
		maarten.last_name: maarten,
	}

	fmt.Println(n)
	//fmt.Println(m)

	for k, v := range n {
		fmt.Println(k)
		fmt.Printf("%s %s\n", v.first_name, v.last_name)
		for i, v2 := range v.favorite_ice_cream {
			fmt.Printf("\t %v: %s\n", i, v2)
		}
	}

}