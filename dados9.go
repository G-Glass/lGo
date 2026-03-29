package main

import (
	"fmt"
)

func main() {
	var nome string
	var nr int

	fmt.Println("Qual é o seu nome?")
	fmt.Scan(&nome)

	fmt.Println("olá:", nome, "tenho ", nr)
	if nome != "Helio" {
		fmt.Println("O nome é invalido")
	} else {
		fmt.Println("continue")
	}

	fmt.Scan(&nr)
	if nr >= 4 {
		fmt.Println("Okay")
		for res := 0; res <= nr; res++ {
			s := res * nr
			fmt.Println(res, "x", nr, "=", s)
		}
	}

}
