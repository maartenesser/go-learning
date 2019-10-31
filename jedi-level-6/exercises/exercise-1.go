package main

import "fmt"

func foo(i int) int {

	return i
}

func bar(i int, s string) (int, string) {

	return i, s
}

func main() {
	 x := foo(2)
	 fmt.Println(x)

	 x, y := bar(2, "Hello")
	 fmt.Println(x, y)

}