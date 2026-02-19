package main

import (
	"Units/pkg/length"
	"fmt"
	"os"
	"strconv"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: units <amount> <unit> to <target_unit>")
		return
	}

	// Parse amount from args
	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	// Set unit and target_unit from args
	fromUnit := os.Args[2]
	toUnit := os.Args[4]

	// Determine conversion package based on unit type
	switch fromUnit {
	case "mm", "millimeter", "millimeters", "cm", "centimeter", "centimeters", "m", "meter", "meters":
		result, err := length.Convert(amount, fromUnit, toUnit)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%g %s = %g %s\n", amount, fromUnit, result, toUnit)
	default:
		fmt.Println("Invalid unit")
	}
}
