package main

import (
	"sync"
	"testing"
)

func TestCalc(t *testing.T) {
	defaultPacks := []int{53, 31, 23}
	packStorage = &storage{
		m: &sync.Mutex{},
	}
	packStorage.update(defaultPacks)

	type out struct {
		num int
		com string
	}

	type testCase struct {
		in  int
		out out
	}
	table := []testCase{
		{
			in: 263,
			out: out{
				num: 263,
				com: "(0x53) + (7x31) + (2x23)",
			},
		},
		{
			in: 250,
			out: out{
				num: 251,
				com: "(3x53) + (0x31) + (4x23)",
			},
		},
		{
			in: 500000,
			out: out{
				num: 500000,
				com: "(9429x53) + (7x31) + (2x23)",
			},
		},
	}

	for _, tc := range table {
		com, num := calculate(tc.in)
		if num != tc.out.num {
			t.Fatalf("unexpected: %v want %v", num, tc.out.num)
		}
		if com != tc.out.com {
			t.Fatalf("unexpected: %v want %v", com, tc.out.com)
		}
	}
}
