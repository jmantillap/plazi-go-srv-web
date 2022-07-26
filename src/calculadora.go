package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	fmt.Println(operacion)
	operador := "-"
	valores := strings.Split(operacion, operador)
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])
	operador1, _ := strconv.Atoi(valores[0])
	operador2, err := strconv.Atoi(valores[1])
	if err != nil {
		fmt.Println(err)
	}
	/*if operador == "+" {
		fmt.Println(operador1 + operador2)
	}
	if operador == "-" {
		fmt.Println(operador1 - operador2)
	}*/
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
	case "-":
		fmt.Println(operador1 - operador2)
	case "*":
		fmt.Println(operador1 * operador2)
	case "/":
		fmt.Println(operador1 / operador2)
	default:
		fmt.Println("Operador no soportado")
	}

}
