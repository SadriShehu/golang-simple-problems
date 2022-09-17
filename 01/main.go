package main

import "fmt"

func segment(x int32, space []int32) int32 {
	maxVal := int32(0)
	spaceLen := len(space)

	for i := 0; i < spaceLen; i++ {
		if int32(spaceLen-i) < x {
			return maxVal
		}

		localMin := space[i]
		for j := int32(i); j < int32(i)+x; j++ {
			if localMin > space[j] {
				localMin = space[j]
			}
		}

		if maxVal < localMin {
			maxVal = localMin
		}
	}

	return maxVal
}

func main() {
	fmt.Println(segment(3, []int32{4, 2, 5, 6, 7, 8}))
}
