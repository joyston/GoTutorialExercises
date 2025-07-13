package main

import (
	s "example/goexercises/Slices"
	"fmt"
)

func main() {
	var d int
	for {
		fmt.Println("1: Slices")
		fmt.Scan(&d)
		switch d {
		case 1:
			s.ExecuteSlices()
		default:
			fmt.Println("Invalid function")
		}
	}
}
