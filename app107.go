package main

import "fmt"

// const product string = "Mama"
// const price = 8

const (
	product string = "Mama"
	quantity int = 50
	price int = 8
	stock = true
)

func main() {
	fmt.Println(product)
	fmt.Println(quantity)
	fmt.Println(price)
	fmt.Println(stock)
}