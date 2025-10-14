package main

import "fmt"

func main (){
	var i int = 90
	var f float32 = float32(i)

	fmt.Printf("%.2f\n", f)
	fmt.Printf("%2f\n", f)
}