package main

import (
	"fmt"
)

const k = 15811

var matrix [k][k]int

func main() {
	input := 12095831

	for i := 0; i < k; i++ {
		// n is the initial step which is incremented by 1 for each X axys iteration
		n := 2*matrix[i][0] + i + 1
		for j := 0; j < k; j++ {
			// all other values except those in the first position are calculated with the following formula
			if j != 0 {
				matrix[i][j] = matrix[i][j-1] + n + 1
				if printCoordinates(i, j, input) {
					break
				}
				// increment the step by one
				n += 1
				continue
			}

			// if i is zero and j is zero the value is equal to step n
			if i == 0 {
				matrix[i][j] = n
				if printCoordinates(i, j, input) {
					break
				}
				continue
			}

			// if j is zero we calculate the value with the following formula
			matrix[i][j] = matrix[i-1][j] + i
			if printCoordinates(i, j, input) {
				break
			}
		}
	}
}

func printCoordinates(i, j, input int) bool {
	if matrix[i][j] == input {
		fmt.Printf("[%d, %d]\n", j, i)
		return true
	}

	return false
}
