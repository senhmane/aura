package main

import "fmt"

func main() {
	fmt.Println("hello world", add(5, 6))

}

func add(a, b int) int {
	return a + b
}
