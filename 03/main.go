package main

import (
	"fmt"
	"log"
	"time"
)

func Solution(P []int, D []string) int {
	skipPayment := 0
	countPayment := 0
	payment := 0
	income := 0
	monthOccurred := map[time.Month]bool{}
	for i, d := range D {
		t, err := time.Parse("2006-01-02", d)
		if err != nil {
			log.Fatalln("time parsing failed")
		}

		monthOccurred[t.Month()] = true
		ok := monthOccurred[t.Month()]

		if P[i] < 0 {
			payment += abs(P[i])
			if ok {
				countPayment++
			}
		}

		if P[i] >= 0 {
			income += P[i]
		}

		if ok && payment >= 100 && countPayment >= 3 {
			skipPayment++
		}

	}

	total := income - payment - (12-skipPayment)*5

	return total
}

func abs(n int) int {
	return n * -1
}

func main() {
	s := Solution([]int{180, -50, -25, -25},
		[]string{"2020-01-01", "2020-01-01", "2020-01-01", "2020-01-31"})

	fmt.Println(s)
}
