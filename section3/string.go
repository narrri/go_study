package main

import "fmt"

func main() {
	var a rune = 'X'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)

	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(s2)

	var x int = 65
	var y string = string(x)

	fmt.Println(x)
	fmt.Println(y)
}
