package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	ns := [][]string{jb, mp}

	fmt.Println("bj slice", jb)
	fmt.Println("mp slice", mp)
	fmt.Println("ns slice", ns)

	for i, nns := range ns {
		fmt.Println("record: ", i)
		for j, val := range nns {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

}
