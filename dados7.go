package main

import (
	"fmt"
)

func main() {
	var multiplicando int
	var total int
	var calc int
	var adicacao = "+"
	var vezes = "x"
	var igual = "="

	numeros := []int{4, 2, 5, 3}
	for total = range numeros {
		calc = total * numeros[2]
		multiplicando = numeros[2]
		fmt.Println("A multiplicação de  ", total, "x", multiplicando, "=", calc)
	}
	for i := 0; i <= 9; i++ {
		//fmt.Println(i)
		res := i * 5
		if multiplicando <= 5 {
			fmt.Println(i, vezes, multiplicando, igual, res)
		}

	}
	//ADICÇAÕ
	fmt.Println("ADIÇÃO")
	for s := 0; s <= 7; s++ {
		add := s + 7
		fmt.Println(s, adicacao, 7, igual, add)
	}

	//SUBTRAÇÃO
	fmt.Println("SUBTRAÇÃO")
	for n := 0; n <= 7; n++ {
		sub := 7 - n
		fmt.Println(7, "-", n, igual, sub)
	}
	//DIVISÃO
	fmt.Println("DIVISÃO")
	for d := 1; d <= 18; d++ {
		divis := 18 / d
		fmt.Println(18, ":", d, igual, divis)
	}

	//fmt.Println("Os números são ", numeros)
}
