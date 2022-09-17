package main

import "sort"

func Solution(A []int) int {
	minVal := 1

	if len(A) == 0 {
		return minVal
	}

	N := []int{}
	for i, n := range A {
		if A[i] > 0 {
			N = append(N, n)
		}
	}

	sort.Ints(N)
	for _, n := range N {
		if minVal < n {
			return minVal
		}
		minVal = n + 1
	}

	return minVal
}

func main() {
	A := []int{1, 3, 6, 4, 1, 2}
	Solution(A)
}
