package main

import "fmt"

func main() {
	function := func(xs string) string {
		return xs
	}
	test := callback(function)
	fmt.Println(test)

}
//callback function
func callback (theCallback func(xs string) string) string {
	return theCallback("Hello World")
}
