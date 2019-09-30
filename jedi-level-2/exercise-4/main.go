package main

import "fmt"

const (
	_ = iota
)

func main() {
	 x := 35

	 fmt.Printf("%d\n%b\n%#x\n", x, x, x)
	 b := x << 1
	fmt.Printf("%d\n%b\n%#x\n", b, b, b)
}
