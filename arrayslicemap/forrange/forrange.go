package main

import "fmt"

func main() {

	numeros := [...]int{1, 2, 3, 4, 5} //O compilador conta

	// retorna o indece [x] e o valor do indece
	// Ex.: numeros[0] = 1
	// Ex.: numeros[2] = 5
	for i, numero := range numeros {
		fmt.Printf("Indece numeros %d Valor numeros %d\n", i+1, numero)
	}

	a := 1
	b := 1
	c := 1
	d := 1
	e := 1
	numeros1 := [...]int{a, b, c, d, e} //O compilador conta
	for i, numero := range numeros1 {
		fmt.Printf("Indece numeros1 %d Valor numeros1 %d\n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Println("Pegando só os Indeces", num)
	}

}
