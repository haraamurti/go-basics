package main

import "fmt"

func main() {
	fmt.Println("hello world")
	
	//basic declaration
	var text = "abil"
	var ayam string = "tod"
	var bagong =5
	ikan := 5

	fmt.Println(text)
	fmt.Println(ayam)
	fmt.Println(bagong)
	fmt.Println(ikan)

	//mutliple declaration
	var a,c,d int = 1 ,3,4
	fmt.Println(a)
	fmt.Println(d)
	fmt.Println(c)

	//box ddeeclaration
	var (
		x int
		y int = 1
		z string = "hello"

	)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//naming rule
	//js use snake case bro

	//go constance
	const phi = 3.14
	fmt.Println(phi)

}