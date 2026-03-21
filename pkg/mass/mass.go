package mass

import "fmt"

const (
	// Metric Units
	microgramsPerKilogram = 1000000000
	milligramsPerKilogram = 1000000
	gramsPerKilogram      = 1000
	tonnesPerKilogram     = 0.001

	// Misc Units
	poundsPerKilogram = 2.2046226218
	ouncesPerKilogram = 35.27396195
)

func Convert(inputAmount float64, from string, to string) (float64, error) {
	var amount float64
	switch from {
	case "µg", "microgram", "micrograms":
		amount = inputAmount / microgramsPerKilogram
	case "mg", "milligram", "milligrams":
		amount = inputAmount / milligramsPerKilogram
	case "g", "gram", "grams":
		amount = inputAmount / gramsPerKilogram
	case "kg", "kilogram", "kilograms":
		amount = inputAmount
	case "t", "tonne", "tonnes":
		amount = inputAmount / tonnesPerKilogram
	case "lb", "lbs", "pound", "pounds":
		amount = inputAmount / poundsPerKilogram
	case "oz", "ounce", "ounces":
		amount = inputAmount / ouncesPerKilogram
	default:
		return 0, fmt.Errorf("unsupported source unit: %s", from)
	}
	switch to {
	case "µg", "microgram", "micrograms":
		return amount * microgramsPerKilogram, nil
	case "mg", "milligram", "milligrams":
		return amount * milligramsPerKilogram, nil
	case "g", "gram", "grams":
		return amount * gramsPerKilogram, nil
	case "kg", "kilogram", "kilograms":
		return amount, nil
	case "t", "tonne", "tonnes":
		return amount * tonnesPerKilogram, nil
	case "lb", "lbs", "pound", "pounds":
		return amount * poundsPerKilogram, nil
	case "oz", "ounce", "ounces":
		return amount * ouncesPerKilogram, nil
	default:
		return 0, fmt.Errorf("unsupported target unit: %s", to)
	}
}

func GetUnits() []string {
	return []string{
		"µg", "ug", "microgram", "micrograms",
		"mg", "milligram", "milligrams",
		"g", "gram", "grams",
		"kg", "kilogram", "kilograms",
		"t", "tonne", "tonnes",
		"lb", "lbs", "pound", "pounds",
		"oz", "ounce", "ounces",
	}
}
