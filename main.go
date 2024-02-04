package main

import (
	"fmt"
	"sync"
)

func main() {
	defaultPacks := []int{53, 31, 23, 9}
	packStorage = &storage{
		m: &sync.Mutex{},
		p: defaultPacks,
	}

	fmt.Println(calculate(263))
	fmt.Println(calculate(250))
	fmt.Println(calculate(500000))
	fmt.Println(calculate(50000006))

	// tryThis([]int{3, 1}, 10)
}
