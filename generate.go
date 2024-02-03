package main

import "fmt"

func generate(packSizes []int, maxCounts []int, targetNumber int, sum int, str string) (string, int) {
	bestValue := 0
	bestCombination := ""

	for i := 0; i < maxCounts[0]; i++ {
		a := sum + (i * packSizes[0])
		key := fmt.Sprintf("%s + (%dx%d)", str, i, packSizes[0])

		if targetNumber == a {
			return key, a
		}

		if bestValue != 0 && a >= bestValue {
			break
		}

		if a >= targetNumber && (a < bestValue || bestValue == 0) {
			bestCombination = key
			bestValue = a
		}

		if len(packSizes) > 1 {
			str, sum := generate(packSizes[1:], maxCounts[1:], targetNumber, a, key)
			if sum >= targetNumber && (sum < bestValue || bestValue == 0) {
				bestCombination = str
				bestValue = sum
			}
		}
	}

	return bestCombination, bestValue
}

func calculateMaxCounts(packSizes []int, targetNumber int) []int {
	maxCounts := make([]int, len(packSizes))
	for i, size := range packSizes {
		maxCounts[i] = targetNumber / size
	}

	return maxCounts
}

func tryThis(input []int, number int) {
	if len(input) < 1 {
		return
	}

	for i := 0; i < input[0]; i++ {
		fmt.Println(number)
		tryThis(input[1:], number)
	}
}
