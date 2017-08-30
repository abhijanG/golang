package main

import "fmt"

func recursion() {
	z := fact(5)
	fmt.Printf("Factorial : %d\n", z)

}

func fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

// Recursion pattern

func recursionpattern(n int) {
	if n == 1 {
		fmt.Println("*")
	} else {
		for i := 1; i <= n; i++ {
			fmt.Print("* ")
		}
		fmt.Println()
		recursionpattern(n - 1)
	}
}
