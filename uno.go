package main

import (
	"fmt"
	"path"
)

func Example() {
	fmt.Println("This is my first go program")
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(path.Join("a", "b", "c"))
	Example()
}
