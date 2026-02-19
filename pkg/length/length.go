package length

import "fmt"

const (
	millimetersPerMeter = 1000
	centimetersPerMeter = 100
	decimetersPerMeter  = 10
	dekametersPerMeter  = 0.1
	hectometersPerMeter = 0.01
	kilometersPerMeter  = 0.001
)

func Convert(inputAmount float64, from string, to string) (float64, error) {
	var amount float64
	switch from {
	case "mm", "millimeter", "millimeters":
		amount = inputAmount / millimetersPerMeter
	case "cm", "centimeter", "centimeters":
		amount = inputAmount / centimetersPerMeter
	case "dm", "decimeter", "decimeters":
		amount = inputAmount / decimetersPerMeter
	case "m", "meter", "meters":
		amount = inputAmount
	case "dam", "dekameter", "dekameters":
		amount = inputAmount / dekametersPerMeter
	case "hm", "hectometer", "hectometers":
		amount = inputAmount / hectometersPerMeter
	case "km", "kilometer", "kilometers":
		amount = inputAmount / kilometersPerMeter
	default:
		return 0, fmt.Errorf("unsupported source unit: %s", from)
	}
	switch to {
	case "mm", "millimeter", "millimeters":
		return amount * millimetersPerMeter, nil
	case "cm", "centimeter", "centimeters":
		return amount * centimetersPerMeter, nil
	case "dm", "decimeter", "decimeters":
		return amount * decimetersPerMeter, nil
	case "m", "meter", "meters":
		return amount, nil
	case "dam", "dekameter", "dekameters":
		return amount * dekametersPerMeter, nil
	case "hm", "hectometer", "hectometers":
		return amount * hectometersPerMeter, nil
	case "km", "kilometer", "kilometers":
		return amount * kilometersPerMeter, nil
	default:
		return 0, fmt.Errorf("unsupported target unit: %s", to)
	}
}
