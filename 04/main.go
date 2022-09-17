package main

import "fmt"

func Solution(A []string) int {
	m := map[int]map[rune]bool{}
	for i, s := range A {
		l := map[rune]bool{}
		canBe := true
		for _, r := range s {
			ok := l[r]
			if ok {
				canBe = false
				break
			}
			l[r] = true
		}

		if canBe == true {
			m[i] = l
		}
	}

	longestWord := map[rune]bool{}
	for _, str := range m {
		localstr := map[rune]bool{}
		for j, m := range longestWord {
			localstr[j] = m
		}

		canConcat := true
		for s := range str {
			ok := localstr[s]
			if ok {
				canConcat = false
				break
			}
			localstr[s] = true
		}

		if canConcat {
			longestWord = localstr
		}
	}

	return len(longestWord)
}

func main() {
	i := Solution([]string{"abc", "yyy", "def", "csv"})
	j := Solution([]string{"hac", "ker", "ra", "nk"})
	fmt.Println(i)
	fmt.Println(j)
}
