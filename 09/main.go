package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "{countryCode} 1234test"

	s = regexp.MustCompile(`\{countryCode\}`).ReplaceAllString(s, "mx")

	fmt.Println(s)
}
