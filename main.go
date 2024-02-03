package main

import (
	"fmt"
)

func generateCombinations(packSizes []int, targetNumber int) (string, int) {
	maxCounts := []int{targetNumber / packSizes[0], targetNumber / packSizes[1], targetNumber / packSizes[2]}

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

func main() {
	packetSize := []int{53, 31, 23}

	fmt.Println(generateCombinations(packetSize, 263))
	fmt.Println(generateCombinations(packetSize, 250))
	fmt.Println(generateCombinations(packetSize, 500000))
	fmt.Println(generateCombinations(packetSize, 50000000))

	fmt.Println(generate(packetSize, calculateMaxCounts(packetSize, 263), 263, 0, ""))
	fmt.Println(generate(packetSize, calculateMaxCounts(packetSize, 250), 250, 0, ""))
	fmt.Println(generate(packetSize, calculateMaxCounts(packetSize, 500000), 500000, 0, ""))
	fmt.Println(generate(packetSize, calculateMaxCounts(packetSize, 50000000), 50000000, 0, ""))

	// tryThis([]int{3, 1}, 10)
}
