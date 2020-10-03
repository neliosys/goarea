package goarea

import "math"

// Pi definido como circunferencia dividia pelo diametro
const Pi = 3.1416

// Circulo retorna a área do circulo
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect retorna a área do retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// função privada
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
