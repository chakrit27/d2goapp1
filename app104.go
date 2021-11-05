package main

import "fmt"

var s = "Thailand"
//s = "Thailand"

func main() {
	var x bool = true
	fmt.Println(s)
	if x {
		var  y int = 1
		if x != true {
			fmt.Println(s)
			fmt.Println(x)
			fmt.Println(y)
		}
	}
	fmt.Println(x)
}