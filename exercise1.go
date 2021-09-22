package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
1.Escreva  um  script  que  pergunta  ao  usuário  se  ele  deseja  converter
uma temperatura de grau Celsius para Farenheit ou vice-versa. Se ele digitar 1,
é de Celsius para Farenheit, se digitar 2 é de Farenheitpara Celsius,
outro valor mostre uma mensagem de erro. Para cada conversão, chame a função correta.
*/

func CelsiusToFarenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func FarenheitToCelsius(farenheit float64) float64 {
	return (farenheit - 32) / 1.8
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Digite 1 para converter de Celsius para Farenheit")
	fmt.Println("Digite 2 para converter de Farenheit para Celsius")

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	switch text {
		case "1":
			fmt.Println("Digite a temperatura em Celsius")
			text, _ = reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			celsius, _ := strconv.ParseFloat(text, 32)
			fmt.Println(CelsiusToFarenheit(celsius))
		case "2":
			fmt.Println("Digite a temperatura em Farenheit")
			text, _ = reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			farenheit, _ := strconv.ParseFloat(text, 32)
			fmt.Println(FarenheitToCelsius(farenheit))
		default:
			fmt.Println("Opção inválida")
			os.Exit(0)
	}
}