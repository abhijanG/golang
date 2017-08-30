package main

import "fmt"

func variadic() {
	// sum(1, 2, 3)
	n := []int{5, 6, 7}
	sum(n...)
}

func sum(nums ...int) {
	fmt.Println("Nums Received", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Total : %d\n", total)
}
