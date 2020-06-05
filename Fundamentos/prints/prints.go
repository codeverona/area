package main

import "fmt"

func main() {
	fmt.Print("mesma")
	fmt.Print("linha")

	fmt.Println(" Nova")
	fmt.Println(" Linha")

	//concatenando tipos diferentes
	x := 3.141516
	xs := fmt.Sprint(x)

	fmt.Println("O valor de X é:" + xs)

	// Outraas formas de concatenar tipos diferentes

	fmt.Println("O valor de X é", x)

	//Truncando casas decimais (o ponto depois do f é só para finalizar a frase com um ponto)
	fmt.Printf("O valor de X é %.2f.", x)

	//ganchos % no texto para valores de variáveis
	a := 1
	b := 1.99999
	c := false
	d := "Opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v %v", a, b, b, c, d)

}
