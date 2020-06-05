package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough // faz com que ele ler a 10 e lê a 9.
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"

	}
}
func main() {
	fmt.Println(notaParaConceito(9.1))
	fmt.Println(notaParaConceito(7.0))
	fmt.Println(notaParaConceito(1.0))
}
