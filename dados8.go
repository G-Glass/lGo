package main

import (
	"fmt"
	"os"
)

func main() {
	nome := os.Args[0]
	fmt.Println("O nome é :", nome)
	nur := [12]int{}
	fmt.Scan(&nome)
	r := make([]int, 0, 9)
	s := cap(r)
	fmt.Printf("Os valores %d e %d ", nur, s)
	if len(os.Args) > 1 {
		preOsarg := os.Args[1]
		fmt.Println("Visivel", nome)
		fmt.Println("primeiro argumento", preOsarg)
	} else {
		fmt.Println("Não foi passado nenhum argumento ")
	}
}
