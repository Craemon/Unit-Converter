package main

import (
	"Units/pkg/length"
	"Units/pkg/mass"
	"Units/pkg/temperature"
	"Units/pkg/time"
	"fmt"
	"os"
	"strconv"
)

var conversionFunctions = map[string]func(float64, string, string) (float64, error){
	"length":      length.Convert,
	"time":        time.Convert,
	"temperature": temperature.Convert,
	"mass":        mass.Convert,
}

func getUnitCategories() map[string][]string {
	return map[string][]string{
		"length":      length.GetUnits(),
		"time":        time.GetUnits(),
		"temperature": temperature.GetUnits(),
		"mass":        mass.GetUnits(),
	}
}

func getUnitCategory(unit string) string {
	unitCategories := getUnitCategories()
	for category, units := range unitCategories {
		for _, u := range units {
			if unit == u {
				return category
			}
		}
	}
	return ""
}
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

	//get and compare unit categories
	fromCategory := getUnitCategory(fromUnit)
	toCategory := getUnitCategory(toUnit)

	if fromCategory == "" || toCategory == "" {
		fmt.Println("Usage: units <amount> <unit> to <target_unit>")
		return
	}
	if fromCategory != toCategory {
		fmt.Println("Units must be of the same type.")
		return
	}
	// Convert and output

	converter, exists := conversionFunctions[fromCategory]
	if !exists {
		fmt.Println("Conversion function not found.")
		return
	}
	result, err := converter(amount, fromUnit, toUnit)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%g %s = %g %s\n", amount, fromUnit, result, toUnit)
}
