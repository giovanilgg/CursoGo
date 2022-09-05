package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expr string) int {
	arrayValores := strings.Split(expr, "+")

	sumaValores := 0
	for _, valores := range arrayValores {

		num, error := strconv.Atoi(valores) //se pasan dos valores para controlar el parseo
		//se invoca strconv ,tambien puede ser parseInt

		if error != nil {
			fmt.Println(error)
			fmt.Println("digita bien ")

		} else {
			sumaValores += num

		}

	}

	return sumaValores

}

func main() {
	var expresion string

	var resultado int
	fmt.Println("=>")
	fmt.Scanln(&expresion)

	resultado = sumar(expresion)
	fmt.Println(resultado)

}

//go no maneja try y catch
