package main

import "fmt"

func main() {
	var x, y int
	var pointer *int
	fmt.Println(&x == &x, &x == &y, &x == nil, pointer == nil)
	fmt.Println(x, y)
}
