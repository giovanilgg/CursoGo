package main

import "fmt"

type Animal interface {
	mover() string
}

//interfaces manejan un metodo global ,ya que ciertas estructuras lo comparten
type perro struct {
}
type pajaro struct {
}
type conejo struct {
}

func (*perro) mover() string {
	return "Soy perro y camino"

}
func (*pajaro) mover() string {
	return "Soy pajaro y vuelo"

}
func (*conejo) mover() string {
	return "Soy conejo y salto"

}
func moverAnimal(animal Animal) string {
	return animal.mover()

}
func main() {
	doberman := perro{}
	fmt.Println(moverAnimal(&doberman))

}
