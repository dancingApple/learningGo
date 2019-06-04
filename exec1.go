package main

import "fmt"

var x = 12
var y int
var str = "Is that true?"
var ans = true
type cat int
var k cat
type kitten cat
var f kitten

func main() {
	z := 16
	y = 14
	fmt.Println(x, y, z)
	fmt.Printf("%v\n%b\n", x, x)
	fmt.Printf("%v\n%X\n", y, y)
	fmt.Printf("%v\n%T\n", z, z)
	s := fmt.Sprintf("%v\t%v\t%v\t%v\t%v",x,y,z,str,ans)
	fmt.Println(s)
	fmt.Println(k)
	fmt.Printf("%T\n",k)
	k = 42
	fmt.Println(k)
	f = kitten(k)
	fmt.Println("The number of kitten is",f)
	fmt.Printf("%T\n",f)
	strToByte := []byte(str)
	fmt.Println(strToByte)
}
