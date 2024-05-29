package main

import (
	"errors"
	"fmt"
)

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divAndRemainder(num int, den int) (result int, reminder int, err error) {
	//result, reminder = 20, 30
	if den == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / den, num % den, nil
}

// 빈 반환
func divAndRemainder2(num int, den int) (result int, reminder int, err error) {

	if den == 0 {
		err = errors.New("cannot divide by zero")
		return
	}

	result, reminder = num/den, num%den
	return
}

func main() {
	fmt.Println(addTo(5))
	fmt.Println(addTo(5, 2))
	a := []int{4, 3}
	fmt.Println(addTo(5, a...))
	fmt.Println(addTo(5, []int{1, 2, 3, 4, 5}...))

	x, y, z := divAndRemainder(5, 2)
	fmt.Println(x, y, z)

	s, b, c := divAndRemainder2(5, 2)
	fmt.Println(s, b, c)
}
