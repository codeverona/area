package main

import "fmt"

var soma = func(a, b int) int { // passando uma função para uma variável
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(2, 3))
}
