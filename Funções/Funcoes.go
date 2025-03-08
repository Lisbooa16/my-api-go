package main

import "fmt"

func somar(valor int8, valor2 int8) int8 {
	return valor + valor2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		return txt
	}
	resultado := f("teste da função")
	fmt.Println(resultado)

	resultadosSoma, resultadosSubtracao := calculosMatematicos(10, 15)
	// _, resultadosSubtracao := calculosMatematicos(10, 15)

	fmt.Println(resultadosSoma, resultadosSubtracao)
}
