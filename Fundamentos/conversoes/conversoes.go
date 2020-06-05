package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado... retorna o valor da tabela ASC.
	fmt.Println("Teste " + string(97))

	//int para String
	fmt.Println("Teste " + strconv.Itoa(123))

	//String para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("True")

	// para retornar Verdadeiro somente usando "True" ou "1" qualquer outro valor retorna Falso
	if b {
		fmt.Println("Verdadeiro")
	}
}
