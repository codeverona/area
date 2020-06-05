package main

import "fmt"

func obterValorAprovado(numero int) int {
	//"defer" retarda a execução da linha até o momento imediatamente antes do retorno
	defer fmt.Println("Fim!")

	if numero > 5000 {
		fmt.Println("Valor Alto...")
		return 5000
	}
	fmt.Println("Valor Baixo...")
	return numero
}
func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
