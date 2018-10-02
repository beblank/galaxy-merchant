package calculator

import "strings"

var (
	I float64 = 1
	V float64 = 5
	X float64 = 10
	L float64 = 50
	C float64 = 100
	D float64 = 500
	M float64 = 1000

	Iron   float64 = 195.5
	Silver float64 = 17
	Gold   float64 = 14450
)

func Calculate(currency string, metal string) float64 {
	var result float64
	list := strings.Split(currency, "")
	for _, str := range list {
		switch str {
		case "I":
			result = result + I
			break
		case "V":
			result = result + V
			break
		case "X":
			result = result + X
			break
		case "L":
			result = result + L
			break
		}
	}
	switch metal {
	case "Iron":
		result = result * Iron
		break
	case "Silver":
		result = result * Silver
		break
	case "Gold":
		result = result * Gold
		break
	}
	return result
}
