package main

import "fmt"

func soma(a, b int) int {
	return a * b
}
func main() {
	s := soma(4, 2)
	fmt.Println("A soma é ", s)
}
