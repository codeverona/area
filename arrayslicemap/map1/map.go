package main

import "fmt"

func main() {

	//var aprovados map[int]string
	//mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12345678100] = "Pedro"
	aprovados[12345678682] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados { // percorre recuperando chave e valor
		fmt.Printf("%s (CPF: % d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678978])
	delete(aprovados, 12345678682)
	fmt.Println(aprovados[12345678682])

}
