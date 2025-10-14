package main

import "fmt"

func main (){
	city:= "London"
	{
		country:= "UK"
		fmt.Println(country)
		fmt.Println(city)
	}
	// fmt.Println(country) // declare in inner block
	fmt.Println(city)
}