package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)

}

func main() {
	if i := numeroAleatorio(); i > 5 { //tembem é suportado pelo switch
		fmt.Println("Ganhou !!!")
		fmt.Println("O número sorteado foi", i)
	} else {
		fmt.Println("Perdeu !!!")
		fmt.Println("O número sorteado foi", i)
	}
}
