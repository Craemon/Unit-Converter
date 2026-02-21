package temperature

import "fmt"

func Convert(inputAmount float64, from string, to string) (float64, error) {
	var amount float64
	switch from {
	case "K", "kelvin":
		amount = inputAmount
	case "°C", "celsius":
		amount = inputAmount + 273.15
	case "°F", "fahrenheit":
		amount = (inputAmount + 459.67) * 1.8
	default:
		return 0, fmt.Errorf("unsupported source unit: %s", from)
	}
	switch to {
	case "K", "kelvin":
		return amount, nil
	case "°C", "celsius":
		return amount - 273.15, nil
	case "°F", "fahrenheit":
		return (amount * 1.8) - 459.67, nil
	default:
		return 0, fmt.Errorf("unsupported target unit: %s", to)
	}
}
