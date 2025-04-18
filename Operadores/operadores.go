package main

import "fmt"

func main() {
	// Aritimeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// var numero1 int16 = 10
	// var numero2 int32 =  25
	// var soma2 int16 = numero1 + int16(numero2)
	// fmt.Print(soma2) -> impossivel somar ou comprar tipos diferentes tipo int16 com int32

	var numero1 int16 = 10
	var numero2 int16 = 25
	var soma2 int16 = numero1 + int16(numero2)
	fmt.Println(soma2)

	// FIM DOS ARITIMETICOS E ATRIBUIÇÃO
	var variavel1 string = "string"
	variavel := "String"
	fmt.Println(variavel, variavel1)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	fmt.Println()
	varedadeiro, falso := true, false
	fmt.Println(varedadeiro && falso)
	fmt.Println(varedadeiro || falso)
	fmt.Println(!varedadeiro)
	fmt.Println(!falso)

	//unarios
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)
	numero--
	numero -= 20
	fmt.Println(numero)
}
