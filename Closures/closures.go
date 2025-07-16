package closures

import "fmt"

func fibonacci() func() int {
	current := 0
	first := 0
	second := 1
	return func() int {
		if first == 0 {
			//current = first + second
			first = 1
			second = 1
			return 0
		} else {
			current = first
			first = second
			second = second + current
			return current
		}
	}

}

func ExecuteClosures() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
