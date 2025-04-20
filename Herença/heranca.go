package main

import "fmt"

type pessoa struct {
	nome      string
	idade     uint8
	altura    uint8
	sobrenome string
}

type estudante struct {
	pessoa    // nao precisa passar o tipo apenas a struct (se eu passar o tipo ele cria os campos no estudante e nao nao seria com uma herança e sim como campos novos)
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herenca")
	// pessoa := pessoa{nome: "Guilherme", idade: 25, altura: 185, sobrenome: "Lisboa"} se eu deixar a variavel com o mesmo nome do struct ele substitui o struct ai da ruim
	p1 := pessoa{nome: "Guilherme", idade: 25, altura: 185, sobrenome: "Lisboa"}
	fmt.Println(p1) // ex de pessoa 1

	p2 := pessoa{nome: "Guilherme", idade: 25, altura: 185, sobrenome: "Lisboa"}
	fmt.Println(p2) // ex de pessoa 2

	// estudante := estudante{pessoa: p1, curso: "Tecnologia da informação", faculdade: "Unip"} mesma coisa do de cima assim a struct estudante acaba sumindo
	el := estudante{pessoa: p1, curso: "Tecnologia da informação", faculdade: "Unip"}
	// fmt.Println(el.pessoa.altura) assim tambem funciona
	fmt.Println(el.altura)
}
