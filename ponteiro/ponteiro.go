package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variaval1 int = 10
	var variavel2 int = variaval1

	fmt.Println(variaval1, variavel2)

	variaval1++
	fmt.Println(variaval1, variavel2)

	// PONTEIRO REFERENCIA DE MEMORIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3
	// ponteiro = variavel3 -> forma errada no ponteiro

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) // desreferenciação
	variavel3 = 150
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)

}
