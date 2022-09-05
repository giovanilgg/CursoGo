package main

import (
	"fmt"
)

func main() {
	a := 4
	b := 5

	var valor bool = false

	if a > b {
		valor = true
		fmt.Println(valor)

	} else if a < b {
		valor = false

		fmt.Println(valor)

	}

	//igualda   r= a==b  false
	//distinto  r= a!=b true
	//mayor q    r= a>b
	//menor q   a<b
	//mayor o igula a>=b

	//operadores logicos
	//on  || and && not !

	//variables en IF

	if nombre, edad := "gio", 25; nombre == "gio" && edad == 25 {
		fmt.Println("el nombre es correcto")

	} else {
		fmt.Println("el nombre es incorrecto")

	}

	//obtener valores de un mapa
	//si llamamos a una clave que no existe llama a un campo vacio
	users := make(map[string]string)
	users["nombre"] = "giovani"
	users["apellido"] = "lopez"

	nom, existe := users["nombre"] //existe toma por defecto porque existe
	fmt.Println(nom, existe)
	//uso de switch
	var caracter string

	fmt.Println("Es vocal o no ???")
	fmt.Scanln(&caracter)
	//tambien puede ser asi
	/*
		case caracter==a || caracter==A si  inicalizaelo al inicio del swtich
		se puede usar sin break pero jalo con break
		tambien se puede utilizar todo en un v¿break

	*/
	switch caracter {

	case "a", "A":
		fmt.Println("Es vocal")

	case "e", "E":
		fmt.Println("Es vocal")

	case "i", "I":
		fmt.Println("Es vocal")

	case "o", "O":
		fmt.Println("Es vocal")

	case "u", "U":
		fmt.Println("Es vocal")

	default:
		fmt.Println("no es vocal")

	}
	//asignaciones contadores
	/*
		a+=2   a=a+2
		a-=2  a=a-2
		a*=2  a=a*2
		a/=2  a=a/2
		a++   o a--




	*/
	//bucles en go,no hay while
	//bucle infinito
	/*for {
		fmt.Println("hola")
	}*/
	//bucle tiipo while
	//determinar la cantidad de digitos que tiene numeros
	numeros := 10000
	contador := 0
	for numeros > 0 {
		numeros /= 10
		contador++
		//fmt.Println(contador)
	}
	//bucle for normal
	for i := 0; i < 4; i++ {

		//imprime solo numeros pares
		if i%2 == 0 {
			//fmt.Println(i)

		}

	}
	fmt.Println("------------------")
	//break y continue
	for i := 0; i <= 10; i++ {
		if i == 5 {
			fmt.Println("salta la iteracion 5")
			continue //salta la iteracion

		}
		if i == 8 {
			fmt.Println("el bucle termina en la iteracion 8")
			break //rompe la ejecucion

		}
		//fmt.Println(i)

	}
	//foreach
	fmt.Println("------------------")
	//iterando meses del año
	meses := [...]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto"}
	/*for i := 0; i < len(meses); i++ {
		fmt.Println(meses[i])

	}*/
	//si no ocupo el indice se coloca _ barra baja
	for _, elementos := range meses { //primer elemento indice,segundo son los elementos

		fmt.Println(elementos)
	}

}
