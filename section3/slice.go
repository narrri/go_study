package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 []int

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice1 == nil)
	fmt.Println(slice2 == nil)

	//len, cap
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//append
	slice2 = append(slice2, 3, 4, 5, 6)
	slice1 = append(slice1, 6)

	fmt.Println(slice1)
	fmt.Println(slice2)

	//make
	slice3 := make([]int, 3)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 1)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println("=========")

	var slice4 = []int{4, 5, 6, 7, 8, 9}
	fmt.Println(slice4)
	fmt.Println(slice4[:3])  //4,5,6
	fmt.Println(slice4[3:])  //7,8,9
	fmt.Println(slice4[2:4]) //6,7

}
