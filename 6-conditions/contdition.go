package main

import "fmt"

func main (){
	fmt.Println("hi")

	IPK := 4

	if IPK < 4 {
		fmt.Println("lo tolol")

	}else if IPK == 4 {
		fmt.Println("ya ngga goblok lah")
	}else {
		fmt.Println("gua akuin lo pinter")
	}

	//=================================================
	//go switch condition

	olinedanneval := 5
	switch olinedanneval{
	case 1:
		fmt.Println("ril")
	
	case 2:
		fmt.Println("bohong")
	
	case 5:
		fmt.Println("rill parah")
	
}
}