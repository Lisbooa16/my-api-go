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

}
