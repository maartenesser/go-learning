package main

import "fmt"

//struct
type person struct {
	first, last string
	age         int
}

// embedded struct
type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "Maarten",
		last:  "Esser",
	}

	//anonamous struct
	p2 := struct {
		first, last string
		age         int
	}{
		first: "Bart",
		last:  "Fokker",
		age:   25,
	}

	fmt.Println(sa1)
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(sa1.first, sa1.last)
	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)

}
