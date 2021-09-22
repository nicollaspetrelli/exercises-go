package main

import "fmt"

/*

A  série  de  Fibonacci  é  uma  sequência  de  números,  cujos  dois  primeiros são 0 e 1.
O termo seguinte da sequênciaé obtido somando os dois anteriores.
Faça um script que solicite um inteiro positivo ao usuário,
Então uma função exibe todos os termos da sequência até on-ésimotermo. Use recursividade.

*/

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}

func main() {
	fmt.Println("Funções Recursivas");
	
	// 1 1 2 3 5 8 13

	posicao := uint(12)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i));
		
	}

	fmt.Println(fibonacci(posicao));
}
