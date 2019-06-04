package main

import "fmt"

const (
	x1 = iota + 1
	y1 = iota
	z
)

const (
	//kb = 1024
	//mb = kb * 1024
	//gb = mb * 1024
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(z)
	x := 2323
	fmt.Printf("%b\t\t%d\n", x, x)
	y := x << 1
	fmt.Printf("%b\t\t%d\n", y, y)

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
