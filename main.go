package main

import "fmt"

func main() {
	fmt.Println("hello world", add(5, 6), sub(5, 4))

}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
