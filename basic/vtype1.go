package main

import (
	"fmt"
	"reflect"
)

func main (){

	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("Rajendra"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(46.0))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))
}