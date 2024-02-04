package main

import "fmt"

func moreLenThree(packSizes []int, maxCounts []int, targetNumber int, sum int, str string) (string, int) {
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
			key, a = moreLenThree(packSizes[1:], maxCounts[1:], targetNumber, a, key)
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

func maxCounts(packSizes []int, targetNumber int) []int {
	maxCounts := make([]int, len(packSizes))
	for i, size := range packSizes {
		maxCounts[i] = targetNumber / size
	}

	return maxCounts
}

func calculate(targetNum int) (string, int) {
	packs := packStorage.getPacks()

	if len(packs) <= 3 {
		return lenThreeSolution(packs, targetNum)
	}

	com, num := moreLenThree(packs, maxCounts(packs, targetNum), targetNum, 0, "")
	return com[2:], num
}
