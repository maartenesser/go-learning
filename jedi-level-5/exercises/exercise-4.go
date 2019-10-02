package main

import "fmt"

func main() {
	computer := struct {
		cpu    string
		ram    string
		hdd    string
		screen string
	}{
		cpu:    "intel",
		ram:    "DDR4",
		hdd:    "SSD",
		screen: "4k",
	}

	fmt.Printf(" CPU: %s\n RAM: %s\n HDD: %s\n Screen: %s\n", computer.cpu, computer.ram, computer.hdd, computer.screen)
}
