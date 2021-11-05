package main

import (
	"fmt"
	"reflect"
)

func main() {

	var i = 10
	var s = "Thailand"
	fmt.Println(i, reflect.TypeOf(i))
	fmt.Println(s, reflect.TypeOf(s))
}