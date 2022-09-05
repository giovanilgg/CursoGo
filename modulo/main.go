package main

import (
	"fmt"
	"paquetes/figuras"
	"paquetes/models"

	"github.com/donvito/hellomod"
)

func main() {

	figura1 := figuras.Cuadrado{
		Lado: 5,
	}
	figuras.RealizaOperacion(&figura1)
	persona1 := models.Persona{}
	persona1.Contructor("Giovani", 25)
	fmt.Println(persona1)
	//persona1.nombre="sdsd" no se puede porque es un atributo privado se hace mediante el contructor
	persona1.GetNombre()
	fmt.Println(persona1.GetNombre())
	hellomod.SayHello() //paquete de tercero

}
