package main

import "fmt"

type testSt struct {
	foo string
	bar string
}

func main() {
	var data []*testSt

	myArr := []string{"foo", "foo", "bar"}

	for _, v := range myArr {
		if v == "foo" {
			data = append(data, &testSt{
				bar: "bar_test",
				foo: v,
			})
		} else {
			data = append(data, &testSt{
				bar: v,
				foo: "foo_test",
			})
		}
	}

	for _, v := range data {
		p := &testSt{}
		p = v
		fmt.Printf("%p\n", &p)
	}

	for _, v := range data {
		p := testSt{}
		p = *v
		fmt.Printf("%p\n", &p)
	}
}
