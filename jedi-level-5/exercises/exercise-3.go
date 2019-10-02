package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	fordTruck := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	volkswagenSedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Red",
		},
		luxury: true,
	}

	fordTruck1 := map[bool]truck{
		fordTruck.fourWheel: fordTruck,
	}

	//fmt.Println(fordTruck1)

	for _, v := range fordTruck1 {
		fmt.Printf("doors: %v\n", v.doors)
		fmt.Printf("color: %s\n", v.color)
		fmt.Printf("fourWheel: %v\n", v.fourWheel)
	}

	volkswagenSedan1 := map[bool]sedan{
		volkswagenSedan.luxury: volkswagenSedan,
	}

	//fmt.Println(volkswagenSedan1)

	for _, v := range volkswagenSedan1 {
		fmt.Printf("doors: %v\n", v.doors)
		fmt.Printf("color: %s\n", v.color)
		fmt.Printf("luxury: %v\n", v.luxury)
	}

}
