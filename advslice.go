package main

import "fmt"

func adslice() {
	s := make([]int, 3)
	for i := 0; i < 3; i++ {
		s[i] = i + 1
	}
	s = append(s, 4, 5, 6, 7, 8, 9)

	fmt.Println("s = ", s)

	// copy slice

	// b := make([]int, len(s))
	// copy(b, s)
	// fmt.Println("b = ", b)

	// cut slice - 3,4 out
	// s = append(s[:2], s[4:]...)
	// fmt.Println("cut(3,4) s = ", s)

	//Del by Index
	s = delbyindex(2, s)
	fmt.Println("Deleted index 2 in s = ", s)

}

func delbyindex(i int, a []int) []int {
	a = append(a[:i], a[i+1:]...)
	return a
}
