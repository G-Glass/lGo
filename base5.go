package main

import "fmt"

func main() {
	a := []int{12, 45, 6, 3, 77}
	n := len(a[1:4])
	s := a[2:] // ULTIMOS TRES NUMEROS DO SLICE a
	t := a[:3] // PRIMEIROS TRES NUMEROS DO SLICE a
	r := a[:]  // Toda lista do slice a
	o := a[1]  // O valor da posição 1 em a é 45

	var numero1 = [12]int{}
	nr2 := numero1[3:9]

	fmt.Println("Os valores de 1 a 4 são", n, s, t, r, "\n", o, "Os valores são", nr2)
}
