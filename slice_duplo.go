package main

import "fmt"

func main() {
	var valor int
	var vTotal int
	var sTotal int
	var are int
	fmt.Println("Digite valor")

	fmt.Scan(&valor)
	fmt.Println("Multiplicaçaõ")

	for vTotal = range valor {
		sTotal = valor * vTotal
		fmt.Println(valor, "x", vTotal, "=", sTotal)
	}

	for arr := range sTotal {
		are = sTotal + arr
		fmt.Println("O valor é ", are, arr)
	}

}
