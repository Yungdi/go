package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("1. command name: " + os.Args[0])

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("2. ", s)

	join := strings.Join(os.Args[1:], " ")
	fmt.Println("3. ", join)
}
