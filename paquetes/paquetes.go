package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string) string {
	//split recibe una cadena y lo convierte en un arreglo

	arrayCadena := strings.Split(cadena, "") //recibe la cadena y corta cada carcater al detectar algo vacio
	arrayInvertida := make([]string, 0)
	// Obteniendo los indices del ultimo elemento al primer elemento
	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
		//almacena en array invertida lo que hay desde la ultima posicion hasta la primera

	}

	return strings.Join(arrayInvertida, "") //de arreglo a cadena de caracteres ,el entrecomillado es la separacion entre cada elemenyto

}

func palabra(pal string) bool {
	pal = strings.ToUpper(pal)             //mayusculs s
	pal = strings.ReplaceAll(pal, " ", "") //remplaza las 0 por las q y solo 2 palabras
	//con raplaceAll aplica a todos los elementos replace solito debes especificar la letras que va a cambiar

	//metodo split
	fmt.Println(reverse(pal))
	palabraInvertida := reverse(pal)

	if palabraInvertida == pal {
		return true
	} else {
		return false
	}

}

func main() {
	//identificar palabra palindromo
	if palabra("amor a roma") == true {
		fmt.Println("Si es un palindromo")

	} else {
		fmt.Println("No es un palindromo")
	}

}
