package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
	arg := os.Args
	fmt.Println(arg[1:])
}
