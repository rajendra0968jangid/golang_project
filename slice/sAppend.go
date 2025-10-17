package main

import "fmt"

func main (){
	arr := [4]int{10,20,30,40}

	slice := arr[1:3]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// i append one then no cap increase 
	slice = append(slice,900)
	
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// when i increase more 
	slice = append(slice,-900,90)
	
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	
}