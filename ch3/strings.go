package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[len(s)]) // panic: 인덱스가 범위를 벗어남
}
