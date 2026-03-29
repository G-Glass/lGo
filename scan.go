package main

import "fmt"

func main() {

	var valor1 int
	var valor2 int
	var valor3 int
	var val1 int

	var opcao1 string

	fmt.Println("Digite valor1")
	fmt.Scan(&valor1)

	fmt.Println("Digite valor2")
	fmt.Scan(&valor2)

	fmt.Println("Digite outros valores")
	fmt.Scan(&valor3)

	//fmt.Printf("Todos valores %d , %d e %d", valor1, valor2, valor3)

	//tValores := []int{valor1, valor2, valor3}

	fmt.Println("Escolha sinal matematico, +,*,/ ou -")
	fmt.Scan(&opcao1)

	if opcao1 == "+" {
		fmt.Println("Adição")
		for val1 = 0; val1 <= valor1; val1++ {
			aVal1 := val1 + valor1
			fmt.Println(valor1, "+", val1, "=", aVal1)
		}

	} else if opcao1 == "*" {
		fmt.Println("Multiplicação")
		for val1 = range valor1 {
			sVal1 := val1 * valor1
			fmt.Println(valor1, "x", val1, "=", sVal1)
		}
	} else if opcao1 == "/" {
		fmt.Println("Divisão")
		for val1 = 1; val1 <= valor1; val1++ {
			dVal1 := valor1 / val1
			fmt.Println(valor1, "÷", val1, "=", dVal1)
		}
	} else if opcao1 == "-" {
		fmt.Println(`Subtraçaõ`)
		for val1 = range valor1 {
			sVal1 := valor1 - val1
			fmt.Println(valor1, "-", val1, "=", sVal1)
		}

	}

	//fmt.Println(tValores)
}
