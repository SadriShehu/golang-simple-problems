package main

import (
	"fmt"
	"time"
)

func main() {
	newTime := time.Now().Format("2006-01-02T15:04:05.000Z")

	loc, _ := time.LoadLocation("Europe/Vienna")
	getTimezone := time.Now().In(loc).Format("Z07:00")
	fmt.Println(getTimezone)
	newTime = newTime[:len(newTime)-1] + getTimezone

	finalTime, _ := time.Parse("2006-01-02T15:04:05.000Z07:00", newTime)
	fmt.Println(finalTime)
}
