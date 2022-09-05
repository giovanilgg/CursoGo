package main

import "fmt"

func main() {

	//funciones variadicas=recibe argumentos indefinidos
	//sumar(1,3,4)
	//todos lo argumentos que se le pasan los terona en un arreglo
	fmt.Println(sumar(1, 2, 3, 4, 5, 6))
	//al momento de recibir los dos datos se deben crear dos variables

}

//retornar multiples valores
//si voy a recibir mas parametros a parte del indefinido,se pone al principio

func sumar(numeros ...int) (int, string) { //con los tres punytos le indicas que no sabes cuantos num va a recibir

	valorSuma := 0
	for i := 0; i < len(numeros); i++ {
		valorSuma += numeros[i]
	}
	/* o tambien
	   for _,num := range numeros{
		   valorSuma += num
	   }


	*/
	return valorSuma, "gio"
}
