package main

import "fmt"

// passando funcção como parametro
func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int { // não entendi essa parte
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	fmt.Println(resultado)
}
