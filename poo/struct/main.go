package main

import "fmt"

//Struct persona
type Persona struct {
	nombre string
	edad   int
}

//metodos
func (p *Persona) imprimir() { //debo mandar una refernacia de la clase ,de lo contrario mandaria copias
	fmt.Printf("Nombre: %s\n Edad:%d\n", p.nombre, p.edad)
}
func (p *Persona) modificarNombre(nuevoNombre string) {

	p.nombre = nuevoNombre
	fmt.Printf("el nuevo nombre es :%s\n", p.nombre)

}

//HERENCIA
type empleado struct {
	Persona  //solo se le pasa la otra clase
	apellido string
}

func main() {
	persona1 := Persona{"gio", 25}
	fmt.Println(persona1)
	persona1.edad = 23 //acceso a objetos
	fmt.Println(persona1)
	//segunda forma para contruir el objeto
	persona2 := Persona{
		nombre: "Gio 2",
		edad:   32,
	}
	fmt.Println(persona2)
	persona2.imprimir()
	persona2.modificarNombre("Giovani Lopez Garcia")
	empleado1 := empleado{

		apellido: "lopes garcia",
	}
	empleado1.nombre = "gioherenciaa"
	fmt.Println(empleado1)

}
