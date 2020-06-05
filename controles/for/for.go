package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par ", i)
		} else {
			fmt.Println("Impar ", i)
		}
	}
	for { // loop infinito (para parar o comando é; Ctrl + Alt + m)
		fmt.Println("Para Sempre...")
		time.Sleep(time.Second * 5) // repete a cada 5 segundos
	}
}
