package main

import "fmt"

func lenThreeSolution(packSizes []int, targetNumber int) (string, int) {
	maxCounts := maxCounts(packSizes, targetNumber)

	bestCombination := ""
	bestValue := 0

	for i := 0; i <= maxCounts[0]; i++ {
		a := i * packSizes[0]
		if bestValue != 0 && a > bestValue {
			break
		}

		for j := 0; j <= maxCounts[1]; j++ {
			b := j * packSizes[1]
			if bestValue != 0 && a+b > bestValue {
				break
			}

			for k := 0; k <= maxCounts[2]; k++ {
				c := k * packSizes[2]
				value := a + b + c
				if bestValue != 0 && value > bestValue {
					break
				}

				if value >= targetNumber && (value < bestValue || bestValue == 0) {
					bestCombination = fmt.Sprintf("(%dx%d) + (%dx%d) + (%dx%d)", i, packSizes[0], j, packSizes[1], k, packSizes[2])
					bestValue = value
				}
				if targetNumber == value {
					goto end
				}
			}
		}
	}

end:
	return bestCombination, bestValue
}
