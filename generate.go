package main

import "fmt"

func generate(packSizes []int, maxCounts []int, targetNumber int, sum int, str string) (string, int) {
	bestNum := 0
	bestCom := ""

	for i := 0; i < maxCounts[0]; i++ {
		a := sum + (i * packSizes[0])
		key := fmt.Sprintf("%s + (%dx%d)", str, i, packSizes[0])

		if a == targetNumber {
			return key, a
		}

		if a > targetNumber {
			if bestNum == 0 || a < bestNum {
				bestNum = a
				bestCom = key
			}

			if len(packSizes) == 1 {
				return key, a
			}
			
			continue
		}

		if len(packSizes) > 1 {
			key, a = generate(packSizes[1:], maxCounts[1:], targetNumber, a, key)
			if a == targetNumber {
				return key, a
			}

			if bestNum == 0 || a < bestNum {
				bestNum = a
				bestCom = key
			}
		}
	}

	return bestCom, bestNum
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
