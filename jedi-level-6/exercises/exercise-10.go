package main

import "fmt"

func main() {
	function := Closure()
	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
	fmt.Println(function())
}

func Closure() func() int {
	number := 5
	//clusure cuntion doesn't know what the function beneeth here is doing.
	return func() int {
		number++
		return number
	}

}