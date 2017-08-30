package main

import "fmt"

func mapdemo() {
	nm := map[string]int{"first": 1, "second": 2}
	fmt.Println("Map nm :", nm)

	m := make(map[string]int)

	m["a"] = 12
	m["b"] = 13
	m["c"] = 14
	m["a"] = 15

	fmt.Println("Map m :", m)

	A := m["a"]
	fmt.Println("Value of m['a'] is ", A)

	delete(m, "b")
	fmt.Println("Deleted m['b'] from Map m :", m)

	B, is_B_present := m["b"]
	C, is_C_present := m["c"]

	fmt.Println("Value of B is :", B)
	fmt.Println("is_B_present :", is_B_present)
	fmt.Println("Value of C is :", C)
	fmt.Println("is_C_present :", is_C_present)

}
