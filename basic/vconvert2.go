package main
import (
	"fmt"
	"strconv"
)

func main (){
	var i int = 42
	var s string = strconv.Itoa(i)

	fmt.Printf("%q\n", s)
}