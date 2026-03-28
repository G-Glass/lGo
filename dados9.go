package main

import (
	"fmt"
)

func main() {
	var nome string
	var nr int
	fmt.Scan(&nome)

	fmt.Println("olá:", nome, "tenho ", nr)
	if nome != "Helio" {
		fmt.Println("O nome é invalido")
	} else {
		fmt.Println("continue")
	}

	fmt.Scan(&nr)
	for res := 0; res <= nr; res++ {
		s := res * nr
		fmt.Println(res, "x", nr, "=", s)
	}
	
}
