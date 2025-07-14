package maps

import (
	"fmt"
	"strings"
)

func WordCount(s string) {
	words := strings.Fields(s)
	// var m map[string]int
	m := make(map[string]int)
	for _, v := range words {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	fmt.Println(m)
}

func ExecuteMaps() {
	//wc.Test(WordCount)
	WordCount("foo bar foo")
}
