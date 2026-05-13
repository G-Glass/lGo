package main

import "fmt"

var client1 int
var client2 int

func aleatorio(client1, client2 int) {
	if client1 >= client2 {
		fmt.Println("Valor não processado")
	}
	if client1 == client2 {
		fmt.Println("Nova acção")
		for ale:=1;ale <= client1;ale++ {
			fmt.Println(ale)
		}
	}
}
func main() {

	fmt.Println("Digite o valor do primeiro cliente:")
	fmt.Scan(&client1)

	fmt.Println("Digite o valor do segundo cliente:")
	fmt.Scan(&client2)

	aleatorio(client1, client2)
	

	fmt.Println("Funcionando em LiteIDE")
}
