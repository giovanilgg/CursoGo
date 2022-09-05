package main

import (
	"fmt"
)

//fmt mensajes en pantalla

func main() {
	//variables en go
	var nombre string

	nombre = "gio"
	var apellido string
	apellido = "lopez garcia"
	edad := 3 //facilita la sintaxis
	//tipos de dato int,float64,string,bool
	const pi = 3.1416

	//si no defino valores iniciales go asigna 0 int o float vacio string y bool falso
	fmt.Println(nombre+apellido, edad)
	//operadores aritmeticos mism +,-,*,/
	valor1 := 24
	valor2 := 3

	//resultado
	result := valor1 + valor2
	fmt.Println("suma", result)
	//cuando uno quiera divisiones exactas tienes que poner valores acordes
	var division float64 = 10.0 / 2.2
	fmt.Println("division ", division)
	//modulo de una division resto de una division
	modulo := valor1 % valor2
	fmt.Println("modulo ", modulo)

	//funcion fmt %s string y %d entero  \n salto de linea %v cualquier tipo de dato
	fmt.Printf("Hola , %s y su edad es %d \n", nombre, edad)
	//guardando en una variable,%T nos dice que tipo de dato es
	mensaje := fmt.Sprintf("Hola , %s y su edad es %d \n", nombre, edad)
	fmt.Println(mensaje)
	//capturando datos en consola
	fmt.Println("Ingresa un nombre")
	fmt.Scanln(&nombre)
	fmt.Println("El nombre que Ingresaste es ")
	fmt.Println(nombre)
	calculadora(2, 3)
}
func saludo() {
	fmt.Println("Hola desde metodo saludo")

}

//indicar a mi funcion que valor retorna como en typecript
func calculadora(numero1 int, numero2 int) int {
	fmt.Println(numero1 + numero2)
	return numero1 + numero2

}
