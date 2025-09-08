package main

import "fmt"

func main () {

	for i := 0; i <10; i++{
		fmt.Println("loops")
	}

	for j := 0; j <10; j ++{
		if j == 8{
		continue
		}
				fmt.Println(j);


	}

}