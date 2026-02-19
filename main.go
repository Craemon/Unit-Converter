package main

import (
	"Units/pkg/length"
	"Units/pkg/time"
	"fmt"
	"os"
	"strconv"
)

var unitCategories = map[string][]string{
	"length": {
		"mm", "millimeter", "millimeters",
		"cm", "centimeter", "centimeters",
		"dm", "decimeter", "decimeters",
		"m", "meter", "meters",
		"dam", "dekameter", "dekameters",
		"hm", "hectometer", "hectometers",
		"km", "kilometer", "kilometers"},
	"time": {
		"ns", "nanosecond", "nanoseconds",
		"μs", "microsecond", "microseconds",
		"ms", "millisecond", "milliseconds",
		"s", "second", "seconds",
		"min", "minute", "minutes",
		"h", "hour", "hours",
		"day", "days",
		"week", "weeks",
	},
}

func getUnitCategory(unit string) string {
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
	// Determine conversion package based on unit type

	switch fromCategory {
	case "length":
		result, err := length.Convert(amount, fromUnit, toUnit)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%g %s = %g %s\n", amount, fromUnit, result, toUnit)
	case "time":
		result, err := time.Convert(amount, fromUnit, toUnit)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%g %s = %g %s\n", amount, fromUnit, result, toUnit)
	default:
		fmt.Println("This should never happen.")
	}
}
