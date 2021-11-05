package main

import (
	"fmt"
	"reflect"
)

func main() {
	var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000
	fmt.Println(fname , " " , lname)
	fmt.Println(m , n , o)
	fmt.Println(reflect.TypeOf(item), item, "_", price)
}