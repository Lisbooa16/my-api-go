package main

import "fmt"

func somar(valor int8, valor2 int8) int8 {
	return valor + valor2
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)
}
