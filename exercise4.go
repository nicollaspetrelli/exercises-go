package main

import (
	"fmt"
	"math/rand"
)

/*
A probabilidade de dar um valor em um dado é 1/6 (uma em 6).
Faça um script que simule 1 milhão de lançamentos de dados e mostre a frequência que deu para cada número.
*/

func main() {

	var dado = 0

	var contador1, contador2, contador3, contador4, contador5, contador6 int

	for i := 0; i < 1000000; i++ {
		dado = 1 + rand.Intn(6)
		
		switch dado {
		case 1:
			contador1++
		case 2:
			contador2++
		case 3:
			contador3++
		case 4:
			contador4++
		case 5:
			contador5++
		case 6:
			contador6++
		}
	}
	fmt.Println("1: ", contador1)
	fmt.Println("2: ", contador2)
	fmt.Println("3: ", contador3)
	fmt.Println("4: ", contador4)
	fmt.Println("5: ", contador5)
	fmt.Println("6: ", contador6)

}