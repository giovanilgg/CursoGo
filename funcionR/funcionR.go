package main

import (
	"fmt"
	"strings"
)

//sacar la funcion factorial de un numero

func factorial(numero int) int {

	if numero == 0 {
		return 1
	} else {
		f := numero * factorial(numero-1)
		return f
	}

}
func main() {

	//funcion anonima se ejecuta automaticamente y no tiene nombre
	//se debe poner al principio si no da error
	func() {
		fmt.Println("funcion anonima")
	}()
	//funcion anonima desde una variable
	funAnonima := func(saludo string) string {
		return saludo
	}
	//clousure funcion que retorna otra funcion
	fmt.Println(funAnonima("Hola desde funcion anonima"))

	fmt.Println(factorial(5))
	funcionRepeat := repeat(3)
	fmt.Println(funcionRepeat("hola \n"))

}

//clousure
func repeat(n int) func(cadena string) string {

	return func(cadena string) string {
		return strings.Repeat(cadena, n) //repeat repite una cadena n veces

	}

}
