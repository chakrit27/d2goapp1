package main

import (
	"fmt"
	"reflect"
)

//Global varible declaration
var m,n int

func main() {
	var x int = 1 //Integer Data Type
	var y int	  //Integer Data Type
	fmt.Println(x)
	fmt.Println(y)

	var a,b,c = 5, 25.25, 14.0 //Multiple float32
	fmt.Println(reflect.TypeOf(a), a, reflect.TypeOf(b), b, reflect.TypeOf(c), c)
}