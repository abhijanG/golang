package main

import "fmt"

func arr1() {
	var a [7]int
	a[0] = 1
	a[6] = 7

	fmt.Println(a[4])
	fmt.Println(a[6])
	fmt.Println(a)
}

func arr2() {
	a := [3]int{1, 2, 3}

	fmt.Println(a)
	fmt.Println(len(a))
}

func matrix() {
	var m [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			m[i][j] = 1
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
	fmt.Println("length is ", len(m))
	fmt.Println(m)
}

func imatrix(a int) {
	// var m [a][a]int
	var m [4][4]int
	for i := 0; i < a; i++ {
		for j := 0; j < a; j++ {
			if i == j {
				m[i][j] = 1
			} else {
				m[i][j] = 0
			}

			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
}
