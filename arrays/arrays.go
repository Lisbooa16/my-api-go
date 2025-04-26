package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")

	// todos os dados dentro do array precisa ser do mesmo tipo
	var array1 [5]string // -> array quando eu seto a quantidade de posição
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posião 1", "Posião 2", "Posião 3", "Posião 4", "Posião 5"}
	fmt.Println(array2[2])

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17} // -> slice nao tem tamanho fixo porem eu preciso passar um tipo de dado
	var slice_teste []int                          // slice array sem definir tamanho
	fmt.Println(slice_teste)
	fmt.Println()
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)
	array2[1] = "Nova posição"

	slice2 = append(slice2, "Dado novo")
	fmt.Println(slice2)

	// teste solo que eu fiz sobre slice adicionando com a lista
	fmt.Println("teste solo que eu fiz sobre slice adicionando com a lista")
	num := slice[len(slice)-1]
	for len(slice) < 15 {
		num++
		slice = append(slice, num)
	}
	fmt.Println(slice)

	//arrays internos
	// slice3 := make([]float32, 0, 11) -> eu posso começar um slice zerado e assim ir dando o append
	// slice_lista_retorno := []int{10, 11, 12, 13, 14, 15, 16, 17}
	// tamanho := len(slice_lista_retorno)
	// fmt.Println(tamanho)
	// slice3 := make([]float32, tamanho, tamanho) -> consigo criar um slice com um tamnho ex: recebi uma lista fiz a contagem dela ai eu garanto o tamanho dela
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println("-------------------")

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //lenght
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //lenght
	fmt.Println(cap(slice4)) // capacidade

	//teste
	fmt.Println("-------------------")
	fmt.Println("-------------------")
	fmt.Println("-------------------")

	slice_lista_retorno := []int{10, 11, 12, 13, 14, 15, 16, 17}
	// tamanho := len(slice_lista_retorno)
	slice_teste_novo := append([]int{}, slice_lista_retorno...)
	fmt.Println(slice_teste_novo)

	// for _, b := range slice_lista_retorno {
	// 	slice_teste_novo = append(slice_teste_novo, b)
	// }
	// fmt.Println(slice_teste_novo) -> funciona tbm porem o outro jeito é o certo

	for c := 0; c < len(slice_lista_retorno); c++ {
		fmt.Println(c)
	}

}
