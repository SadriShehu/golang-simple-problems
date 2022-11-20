package main

import "fmt"

type st struct {
	key string
}

func compareStructs() {
	s1 := st{
		key: "1",
	}

	s2 := st{
		key: "1",
	}

	if s1 == s2 {
		fmt.Println("struct true")
		return
	}

	fmt.Println("struct false")
}

func compareArrays() {
	m1 := [2]string{
		"key", "value",
	}

	m2 := [2]string{
		"key", "value",
	}

	if m1 == m2 {
		fmt.Println("array true")
		return
	}

	fmt.Println("array false")
}

func compareSlices() {
	a := []string{"key", "value"}
	b := []string{"key", "value"}

	if len(a) != len(b) {
		fmt.Println("slice false")
		return
	}
	for i, v := range a {
		if v != b[i] {
			fmt.Println("slice false")
			return
		}
	}

	fmt.Println("slice true")
}

func main() {
	compareArrays()
	compareSlices()
	compareStructs()
}
