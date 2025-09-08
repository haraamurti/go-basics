package main

import "fmt"

type film struct {
	title string
	director string
	sound string
	duration int
}

func main (){
	fmt.Println("hello world ")
	
	var film1 film
	var film2 film


	//setting untuk film1
	film1.title = "hollow moving's castle"
	film1.director = "ghibli studio"
	film1.duration = 1
	film1.sound = "joe hisashi"

	//setting untuk film2
	film2.title = "the wind Rises"
	film2.director = "ghibli studio"
	film2.duration = 2
	film2.sound = "joe hisashi"


	fmt.Println("Film title : " + film1.title)
	//how to access to func main

}