package length

import "fmt"

func Convert(inputAmount float64, from string, to string) (float64, error) {
	var amount float64
	switch from {
	case "mm", "millimeter", "millimeters":
		amount = inputAmount / 1000
	case "cm", "centimeter", "centimeters":
		amount = inputAmount / 100
	case "dm", "decimeter", "decimeters":
		amount = inputAmount / 10
	case "m", "meter", "meters":
		amount = inputAmount
	default:
		return 0, fmt.Errorf("unsupported source unit: %s", from)
	}
	switch to {
	case "mm", "millimeter", "millimeters":
		return amount * 1000, nil
	case "cm", "centimeter", "centimeters":
		return amount * 100, nil
	case "dm", "decimeter", "decimeters":
		return amount * 10, nil
	case "m", "meter", "meters":
		return amount, nil
	default:
		return 0, fmt.Errorf("unsupported target unit: %s", to)
	}
}
