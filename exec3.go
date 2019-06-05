package main

import "fmt"

const (
	cyear      = 2019 + iota
	nyear
	nnyear
	nnnyear
	nnnnyear

	byear      = "1989"
	bmonth int = 6
)

func main() {
	x := 232
	chineseStr := "你好"
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	a := (x == 123)
	b := (x <= 123)
	c := (x >= 123)
	d := (x != 123)
	e := (x < 123)
	f := (x > 123)
	fmt.Println(a, b, c, d, e, f)
	fmt.Printf("%T\t%T\n",byear, bmonth)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
	z := `This is       a raw
		  literal String`
	fmt.Println(z)
	fmt.Println(cyear, nyear, nnyear, nnnyear, nnnnyear, byear, bmonth)
	chineseStrByte := []byte(chineseStr)
	fmt.Println(chineseStr)
	fmt.Printf("%b\t%d\n", chineseStrByte,chineseStrByte)
}
