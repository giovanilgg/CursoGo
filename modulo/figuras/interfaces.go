package figuras

import "fmt"

type Operaciones interface {
	calcularArea() float64
	calcularPerimetro() float64
}

func RealizaOperacion(Figura Operaciones) {
	fmt.Println(Figura.calcularArea())
	fmt.Println(Figura.calcularPerimetro())

}
