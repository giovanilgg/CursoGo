package main

import (
	"fmt"
	"math"
)

type Operaciones interface {
	calcularArea() float64
	calcularPerimetro() float64
}

type Cuadrado struct {
	lado float64
}
type Circulo struct {
	radio float64
}

func (c *Cuadrado) calcularArea() float64 {
	return c.lado * c.lado

}
func (c *Cuadrado) calcularPerimetro() float64 {
	return (2.0 * c.lado) * (2.0 * c.lado)

}
func (ci *Circulo) calcularArea() float64 {
	pi := math.Pi
	return pi * (math.Pow(ci.radio, 2))

}
func (ci *Circulo) calcularPerimetro() float64 {
	pi := math.Pi
	return 2 * pi * ci.radio

}
func realizaOperacion(figura Operaciones) {
	fmt.Println(figura.calcularArea())
	fmt.Println(figura.calcularPerimetro())

}

func main() {

	cuadradoFigura := Cuadrado{lado: 3.5}
	circuloFigura := Circulo{radio: 3.6}
	realizaOperacion(&cuadradoFigura)
	realizaOperacion(&circuloFigura)

	fmt.Println()

}
