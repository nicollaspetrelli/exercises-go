package main

import "fmt"

/*
Faça um programa que leia três numeros inteiros,
e que forneça a soma desses três numeros através de uma função.
Seu script também deve fornecer a média dos três números,
através de uma segunda função que chama a primeira.
*/

func soma(numeros []int) int {
	var soma int
	for _, num := range numeros {
		soma += num
	}
	return soma
}

func media(numeros []int) float64 {
	return float64(soma(numeros)) / float64(len(numeros))
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

	fmt.Println("A soma dos numeros é: ", soma(numeros))
	fmt.Println("A media dos numeros é: ", media(numeros))
}