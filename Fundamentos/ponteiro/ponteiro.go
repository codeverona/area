package main

import "fmt"

func main() {

	i := 1

	var p *int = nil
	p = &i //pegando o enderço da variável
	*p++   // desreferenciando (pegando o valor da variável)
	i++
	//GO não tem aritimética de ponteiros
	fmt.Println(p, *p, i, &i)
}
