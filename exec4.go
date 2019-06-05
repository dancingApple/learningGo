package main

import "fmt"

func main() {
	x := 33
	for {
		fmt.Printf("%d\t%#U\n", x, x)
		x++
		if x >= 123 {
			break
		}
	}

	switch "Michael"{
	case "Dora", "Michael":
		fmt.Println("I wanna learn everything about computer science")
	case "Mike", "coke":
		fmt.Println("this is not print out")
	default:
		fmt.Println("Just here in case")
	}
}
