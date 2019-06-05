package main

import "fmt"

func main() {
	x := []int{1,2,3,4,5}
	y := x[1:3]
	fmt.Println(len(x))
	fmt.Println(x)
	for index,value := range x{
		fmt.Println(index,value)
	}
	fmt.Println(y)
}
