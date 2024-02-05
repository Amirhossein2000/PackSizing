package main

import (
	"sync"
)

func main() {
	defaultPacks := []int{53, 31, 23}
	packStorage = &storage{
		m: &sync.Mutex{},
	}
	packStorage.update(defaultPacks)

	apiServe()
}
