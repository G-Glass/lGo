package main

import "fmt"

func valorePrimarios(v1, v2, v3 int) {

	var s int
	if v1 > v2 {
		s = v1 + v2
		fmt.Println("O valor será substituido por ", s)
	}
	fmt.Printf("Os valores acima cidada constam os segintes \n %v Primeiro valor"+
		"\n %v segundo valor \n %v Terceiro valor", v1, v2, v3)
}

func valorSegundario() {

}
func main() {
	valorePrimarios(21, 4, 55)
}
