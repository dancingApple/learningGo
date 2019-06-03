package main

import "fmt"

func main() {
	/* "_" means throw away returns */
	n, _ := fmt.Println("This is a String", 28, true)
	fmt.Println("The number of bytes written is",n)

	/* Declare and assign value using short declaration operator */
	num := 78
	fmt.Println("Declare variable num and assign its value to",num)
	num = 56
	fmt.Println("Re-assign num value to",num)
}
