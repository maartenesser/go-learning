package main

import "fmt"

func main() {
	fmt.Println("Hello World!! Learning how to code!")

	//calls function
	foo()

	//looping
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

//function
func foo() {
	fmt.Println("I'm in foo")
}

//control flow:
// (1) Sequence
// (2) loop: Iterative
// (3) conditional
