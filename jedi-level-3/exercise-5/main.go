package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("If you devide %v by 4, the modulo is %v\n ", i, i%4)
	}
}
