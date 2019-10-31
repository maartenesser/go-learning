package main

import "fmt"

func main() {
	function := returnFunc()
	fmt.Println(function())
}

func returnFunc() func() string {
	return func() string {
		return "Hello Wrold"
	}
}
