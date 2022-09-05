package main

import (
	"fmt"
)

func main() {
	//suma de dos numeros enteros

	numero1 := 0
	numero2 := 0

	fmt.Println("Introduce un numero")
	fmt.Scanln(&numero1)
	fmt.Println("Introduce otro numero")
	fmt.Scanln(&numero2)
	fmt.Println("la suma de esos dos numero es: ", sumaNumero(numero1, numero2))

}
func sumaNumero(num1, num2 int) int {
	return num1 + num2

}
