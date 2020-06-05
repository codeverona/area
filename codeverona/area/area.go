package area

import "math"

//Pi é uma proporção númerica definida pela relação entre_
//o período de uma circunferência e seu diâmetro
const Pi = 3.1416

//Cinc é responsável por circular a área de circunferência
func Cinc(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por circular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visível...
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
