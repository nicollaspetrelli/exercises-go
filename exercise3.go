package main

import (
	"fmt"
)

/*
Faça um programa que recebe três números do usuário, e identifica
o maior através de uma função e o menor número através de outra função.
*/

func maiorEMenor(numeros []int) (int, int) {
	maior := numeros[0]
	menor := numeros[0]

	for _, num := range numeros {
		if num > maior {
			maior = num
		}
		if num < menor {
			menor = num
		}
	}

	return maior, menor
}

func main() {
	var num1, num2, num3 int
	fmt.Print("Digite o primeiro numero: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo numero: ")
	fmt.Scan(&num2)
	fmt.Print("Digite o terceiro numero: ")
	fmt.Scan(&num3)

	numeros := []int{num1, num2, num3}

	fmt.Println(numeros)

	maior, menor := maiorEMenor(numeros)
	fmt.Printf("O maior numero é %d e o menor é %d\n", maior, menor)
}