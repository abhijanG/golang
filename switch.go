package main

import "fmt"

func whatday(n int) {
	if n != 0 && n <= 7 {
		switch n {

		case 1:
			fmt.Println("Monday")

		case 2:
			fmt.Println("Tuesday")

		case 6:
			fmt.Println("Saturday")

		case 7:
			fmt.Println("Hola - Sunday")
		default:
			fmt.Println("It is weekday")
		}
	} else {
		fmt.Println("Wrong Input - 1 -7 Accepted")
	}

}
