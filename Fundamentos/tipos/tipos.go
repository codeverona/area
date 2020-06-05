package main

import (
	"fmt"
	"math"
	r "reflect"
)

func main() {

	// númros inteiros
	fmt.Println(1, 2, 100)
	// o pacote reflect.TypeOf() devolve o tipo do literal
	// o r antes de .TupeOf é o alias do import reflect
	fmt.Println("Literal inteiro é", r.TypeOf(32000))

	//sem sinal (só positivos)... unit8 unit16 unit32 unit64
	var b byte = 255
	fmt.Println("O byte é", r.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo do i1 é", r.TypeOf(i1))

	var i2 rune = 'a' // rune representa o mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", r.TypeOf(i2))
	fmt.Println(i2)

	//números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", r.TypeOf(x))
	fmt.Println("O tipo litral 49.99 é", r.TypeOf(49.99)) // foat64

	//boolean
	bo := true
	fmt.Println("O tipo de variável bo é", r.TypeOf(bo))
	fmt.Println(!bo)

	//String
	s1 := "olá meu nome é Leo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// não é com '' é com ``(crase)
	s2 := `Olá
	meu
	nome
	é
	Leo`
	fmt.Println("O tamanho da string é", len(s2))

	//char ????
	char := 'a'
	fmt.Println("O tipo chae é", r.TypeOf(char))
	fmt.Println(char) // não tem char ele é convertido para tabela unicode (VALOR 97)

}
