package main

import "fmt"

var x = 12
var y int

func main() {
	z := 16
	y = 14
	fmt.Println(x, y, z)
	fmt.Printf("%v\n%b\n", x, x)
	fmt.Printf("%v\n%X\n", y, y)
	fmt.Printf("%v\n%T\n", z, z)
}
