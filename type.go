package main

import "fmt"

func main() {
	var nome string
	var (
		senha int
	)

	fmt.Println("Digite o seu nome")
	fmt.Scan(&nome)

	fmt.Println("Digite a senha")
	fmt.Scan(&senha)

	fmt.Printf("Nome: %s \n Senha: %d", nome, senha)
}
