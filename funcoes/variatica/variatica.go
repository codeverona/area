package main

import "fmt"

func media(numeros ...float64) float64 { //recebe quantos prametros forem passados
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(7, 6, 8)) //passar quantos parametros quiser
}
