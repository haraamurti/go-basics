package main

import "fmt"

//this is a function
func im_the_best_function(){
	fmt.Println("hi Im the best function")
	fmt.Println("i just got excecuted")
}
//with param
func withparameter(name string ){
	fmt.Println("hello user, dear "+ name + ", i just got excecuted ")
}

//with return value
func with_return_value(){


}

func main (){
	fmt.Println("hello world")
	im_the_best_function()
	withparameter("neval")
}