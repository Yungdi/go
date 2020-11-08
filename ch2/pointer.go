package main

import "fmt"

func main() {
	number := 4
	numberP := &number
	fmt.Println(*numberP)
	*numberP = 5
	fmt.Println(*numberP)
}
