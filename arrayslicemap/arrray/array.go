package main

import "fmt"

func main() {
	//Array é homogênea (mesmo tipo) e estática (fixo)

	var notas [3]float64
	fmt.Println(notas)

	// inicializando array
	// posição do array......... = valor para cada posição
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)

}
