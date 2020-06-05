package main

import "fmt"

func init() { //por convenção a função init executa antes que tudo.
	fmt.Println("Iicializando...")
}

func main() {
	fmt.Println("Main...")
}
