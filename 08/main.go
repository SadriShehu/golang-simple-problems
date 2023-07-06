package main

import (
	"fmt"
	"time"
)

func stringToTime(st string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05.000Z", st)
	if err != nil {
		panic(err)
	}

	return t
}

func main() {
	t := stringToTime("2022-10-13T08:14:52.718Z")
	fmt.Println(t)
}
