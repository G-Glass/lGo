package main

import "fmt"

func imprimirDados(nome string, idade int, nivel string) {
	fmt.Printf("%s tem %d anos e é do %s", nome, idade, nivel)
}
func main() {
	imprimirDados("Jose", 32, "segundario")
}
