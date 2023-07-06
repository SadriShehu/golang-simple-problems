package main

import (
	"fmt"
	"time"
)

const THIS_IS_CONST string = "test_this_print"

func main() {
	fmt.Println(time.Now().Format("2006-01-02T15:04:05.000-0700"))

}

func toPrint(this string) {
	err := fmt.Errorf("%s we printed", this)

	fmt.Println(err)
}
