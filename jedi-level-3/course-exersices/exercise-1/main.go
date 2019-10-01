package main

import "fmt"

func main() {
	//print out all the numbers in letters from 33 to 122
	for i := 33; i <= 122; i++ {
		fmt.Printf("%#U\n", i)
		i++
	}
}
