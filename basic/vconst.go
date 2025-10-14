package main

import "fmt"

func main (){
	const name = "rajendra"
	// name = "hello" // its shows error because of const variable

	fmt.Printf("%v %T \n", name, name)
}