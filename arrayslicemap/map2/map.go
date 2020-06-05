package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"José João":     10000.0,
		"Gabriel Silva": 15000.0,
		"Pedro Santos":  1200.0,
	}

	funcESalarios["rafael Filho"] = 2500.0
	delete(funcESalarios, "Inexistente")

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}
}
