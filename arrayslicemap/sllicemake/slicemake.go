package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)        // tamanho do Slice(10) e tamanho do Array(20)
	fmt.Println(s, len(s), cap(s)) // tamanho e capacidade do slice.

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) //peenchendo o slice com mais 10, total 20 de 20
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) //adcionando mais 1 no slice cheio. o slice fica com 21 e o array interno fica com 40
	fmt.Println(s, len(s), cap(s))
}
