package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	// forma reduzida

	area := PI * m.Pow(raio, 2)

	fmt.Println("Area é", area)

	// outra forma de constante

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)

	// formas de declarar variável

	// a recebe true, f recebe false
	var e, f bool = true, false

	//g recebe 2, h recebe false, i recebe "epa!"
	g, h, i := 2, false, "epa!"

	//como toda variável tem que ser usada, usamos elas no Println
	fmt.Println(a, b, c, d)
	fmt.Println(g, h, i, e, f)

}
