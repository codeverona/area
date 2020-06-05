package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo, não precisa informar os retornos
}

func main() {
	x, y := trocar(1, 2)
	fmt.Println(x, y)

	segundo, primeiro := trocar(1, 7)
	fmt.Println(segundo, primeiro)
}
