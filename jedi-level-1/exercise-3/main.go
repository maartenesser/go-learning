package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	//fmt.Printf("%T\n" , x, y, z)
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
