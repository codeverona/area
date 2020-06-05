package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferencia de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 3
	i *= 3 // i = i * 3
	i %= 3 // i = i % 2

	fmt.Println(i)

	// atribuir várias variáveis ao mesmo tempo
	x, y := 1, 2
	fmt.Println(x, y)

	// Trocando o valor das variáveis
	x, y = y, x
	fmt.Println(x, y)

}
