package main

import "fmt"

func main() {
	//PrintNum()
	//PrintRune()
	//PrintAge()
	//PrintAge1()
	//PrintRemainder()
	//Status()
	//Kitten()
	FavGame()
}

func PrintNum() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func PrintRune() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v\t\n", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d\t%#U\n", j, i)
		}
	}
}

func PrintAge() {
	by := 1989
	age := 0
	for by < 2019 {
		by++
		fmt.Println(by)
		age++
	}
	fmt.Println("You are", age, "years old!")
}

func PrintAge1() {
	by := 1989
	age := 0
	for {
		by++
		fmt.Println(by)
		age++
		if by >= 2019 {
			break
		}
	}
	fmt.Println("You are", age, "years old!")
}

func PrintRemainder() {
	for x := 10; x <= 100; x++ {
		fmt.Println("The number", x, "divided by 4 remainder is", x%4)
	}
}

func Status() {
	x := "Kate"
	if x == "Michael" {
		fmt.Println("You are awesome!")
	} else if x == "Dora" {
		fmt.Println("I guess you are OK!")
	} else {
		fmt.Println("Who are you?")
	}
}

func Kitten() {
	switch {
	case false:
		fmt.Println("Michael is not a cat!")
	case (2 == 1):
		fmt.Println("Dora is also not a cat!")
	case (1 != 1):
		fmt.Println("Siren is a cute kitten!")
	default:
		fmt.Println("Nomi is a kitten Princess")
	}
}

func FavGame(){
	switch "Dota 2" {
	case "CS":
		fmt.Println("My Favourite game is CS:GO")
	case "Dark soul":
		fmt.Println("My Favourite game is Dark soul")
	case "Final Fantasy":
		fmt.Println("My Favourite game is Final Fantasy")
	case "ARK":
		fmt.Println("My Favourite game is ARK")
	default:
		fmt.Println("I don't play game!")
	}
}
