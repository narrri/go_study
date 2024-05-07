package main

import (
	"fmt"
)

func main() {
	//var nilMap map[string]int
	//nilMap["nil1"] = 1 //패닉

	test1 := map[string]int{}
	test1["test"] = 1

	fmt.Println(test1)

	ages := make(map[int][]string, 10)
	ages[10] = []string{"Sarah", "Peter"}
	ages[20] = []string{"Fred", "Ralph", "Ze"}

	fmt.Println(ages[20])

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	delete(m, "GO") //존재하지 않는 키를 삭제해도 아무일도 일어나지 않음
	fmt.Println(m)

	intSet := map[int]bool{}
	vals := []int{3, 6, 1, 9, 3, 4, 5, 10, 7, 2, 7}
	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet)) //11, 9
	fmt.Println(intSet[500])            //false
	if intSet[10] {
		fmt.Println("10 is in the set")
	}
}
