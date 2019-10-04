package main

import "fmt"

type person struct {
	first, last string
	age int
}

type user struct{
	person
}

func (u user) speak () {
	fmt.Println("Hello i'm", u.first, u.last, "and i'm ", u.age, "years old")
}


func main() {
	//creating a user with person values
	u := user{person{
		first: "Maarten",
		last:  "Esser",
		age:   26,
	}}

	//calling the speak function
	u.speak()
}