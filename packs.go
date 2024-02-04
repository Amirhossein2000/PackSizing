package main

import (
	"sort"
	"sync"
)

type storage struct {
	m *sync.Mutex
	p []int
}

var packStorage *storage

func (s *storage) getPacks() []int {
	output := make([]int, len(s.p))
	copy(output, s.p)
	return output
}

func (p *storage) update(input []int) {
	p.m.Lock()
	defer p.m.Unlock()

	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	p.p = input
}
