package main

import "fmt"

//Não suporta operador ternário
// tem que fazer gambiarra.
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

}
func main() {
	fmt.Println(obterResultado(6.2))
}
