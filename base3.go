package main

import (
	"fmt"
	"time"
)

func valormaior(a, b int) int {
	if a <= 45 {
		return a * b
	}
	return a + b
}
func main() {
	s := valormaior(101, 4)
	var tempo = time.Now()
	fmt.Println("Valor é ", s, "e são", tempo.Format("20:54:20"))
}
