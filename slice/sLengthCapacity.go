package main 

import "fmt"

func main (){
	// first 5 is length and second 8 is capacity
	slice := make([]int , 5, 8)
	fmt.Println(slice)

	fmt.Println(len(slice))

	fmt.Println(cap(slice))
}