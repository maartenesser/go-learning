package main

import (
	"fmt"
)

type person struct {
	first, last string
}

type secretAgent struct {
	person
	ltk bool
}

//method
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person who codes")
}

//interface
// a value can be of more than one type
//keyword, identifier, type
type human interface {
	// Hey baby if you get this message you are my type
	speak()
}

//func that calls human
func call (h human ) {
	fmt.Println("I Called human", h)
}


// simple function with one parameter string s
func bar(s string)  {
	fmt.Printf("Hello, %v\n", s)
}

//function with parameter and a return value of the type string
func woo(st string) string  {
	return fmt.Sprintln("Hello from woo, ", st)
}

//function with multiple returns
func mouse(fn string, ln string ) (string, bool)  {
	x := fmt.Sprint(fn, ln, `, say "Hello"`)
	y := true
	return x, y

}

func main() {

	// func (r receiver) identifier(parameters ) (returns(s)) { code }

	bar("james")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)

	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk:    true,
	}
	fmt.Println(sa1)
	sa1.speak()

	p1 := person{
		first: "Maarten",
		last:  "Esser",
	}

	//polymorphisem
	call(sa1)
	call(p1)
	p1.speak()

	//function expression
	f := func(x int) {
		fmt.Println("I'm ", x, "years old")
	}
	f(26)

}

