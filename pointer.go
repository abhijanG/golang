package main

import "fmt"

func pointer1() {
	var a *int
	b := 6

	a = &b
	fmt.Println("Address ", a)
	fmt.Println("value ", *a)
}

func pointer2() {
	c := 5
	// increment(c)
	// increment(c)
	// fmt.Println("Called From Pointer2 Func :", c)

	incrementbyptr(&c)
	incrementbyptr(&c)
	fmt.Println("Called From Pointer2 Func :", c)
}

func increment(c int) {
	c++
	fmt.Println("Called within increment", c)
}

func incrementbyptr(c *int) {
	*c++
	fmt.Println("Called within incrementbyptr", *c)
}
