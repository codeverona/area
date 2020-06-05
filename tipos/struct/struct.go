package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//Método: função com receiver (recebedor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{ //declarar produto
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1989.90, 0.10} // outra forma de declarar produto
	fmt.Println(produto2.precoComDesconto())
}
