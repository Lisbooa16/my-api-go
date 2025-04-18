package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradoro string
	numero    uint8
}

func main() {
	fmt.Println("Arquivo de structs")
	var u usuario
	u.idade = 26
	u.nome = "Guilherme Lisboa"
	fmt.Println(u)

	enderecoex := endereco{"Rua dos bobos", 0}
	usuario2 := usuario{"Guilherme Lisboa", 24, enderecoex}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Guilherme"}
	fmt.Println(usuario3)

}
