package main

import "fmt"

// código refatorado do arquivo ifelseif

func notaParaConceito(n float64) string {
	switch { // aqui tem o valor True implícito
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n <= 9:
		return "B"
	case n >= 5 && n <= 8:
		return "C"
	case n >= 3 && n <= 5:
		return "D"
	default:
		return "E"
	}

}
func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(8.0))
	fmt.Println(notaParaConceito(3.8))
}
