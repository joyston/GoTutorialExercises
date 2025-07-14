package main

import (
	m "example/goexercises/Maps"
	s "example/goexercises/Slices"

	"fmt"
)

func main() {
	var d int
	for {
		fmt.Println("1: Slices")
		fmt.Println("2: Maps")

		fmt.Scan(&d)
		switch d {
		case 1:
			s.ExecuteSlices()
		case 2:
			m.ExecuteMaps()
		default:
			fmt.Println("Invalid function")
		}
	}
}
