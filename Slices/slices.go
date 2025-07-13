package slices

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	m := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			m[y] = make([]uint8, dx)

			for x := range m[y] {
				m[y][x] = uint8(imagefunction(x, y))
			}
		}
	}
	return m
}

func imagefunction(x, y int) int {
	return x * y
}

func ExecuteSlices() {
	pic.Show(Pic)
}
