package main

import (
	"fmt"
	"regexp"
)

func main() {
	match := regexp.MustCompile(`Order with id= ([\w|0-9]+) already exists.`)

	s := "Order with id= a0CS000000DYOjqMAH already exists."

	btlrID := match.FindStringSubmatch(s)

	if len(btlrID) > 1 && btlrID[1] == "a0CS000000DYOjqMAH" {
		fmt.Println("we got a match")
	}
}
