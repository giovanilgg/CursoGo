package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (ci *Circulo) calcularArea() float64 {
	pi := math.Pi
	return pi * (math.Pow(ci.Radio, 2))

}
func (ci *Circulo) calcularPerimetro() float64 {
	pi := math.Pi
	return 2 * pi * ci.Radio

}
