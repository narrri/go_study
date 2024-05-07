package main

import "fmt"

func main() {

	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	fmt.Println(fred)

	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println(julia)

	beth := person{
		age: 20,
	}
	fmt.Println(beth)
}
