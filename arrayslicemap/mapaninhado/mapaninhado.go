package main

import "fmt"

func main() {

	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 1000.0,
			"Guga Pereira":  2000.0,
		},
		"J": {
			"José João": 3000.0,
		},
		"P": {
			"Pedro Junior": 4000.0,
		},
	}
	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
