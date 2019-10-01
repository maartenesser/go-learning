package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["octo_pussy"] = []string{`Bieng Evil`, `Revenge on Bond`, `suducton`}
	delete(m, `bond_james`)

	for k, v := range m {
		fmt.Printf("Actor: %v\n", k)
		for _, v2 := range v {
			fmt.Printf("\t interest in, %v\n", v2)
		}
	}
}
