package time

import "fmt"

const (
	nanosecondsPerSecond  = 1000000000
	microsecondsPerSecond = 1000000
	millisecondsPerSecond = 1000
	secondsPerMinute      = 60
	secondsPerHour        = 3600
	secondsPerDay         = 86400
	secondsPerWeek        = 604800
)

func Convert(inputAmount float64, from string, to string) (float64, error) {
	var amount float64
	switch from {
	case "ns", "nanosecond", "nanoseconds":
		amount = inputAmount / nanosecondsPerSecond
	case "μs", "microsecond", "microseconds":
		amount = inputAmount / microsecondsPerSecond
	case "ms", "millisecond", "milliseconds":
		amount = inputAmount / millisecondsPerSecond
	case "s", "second", "seconds":
		amount = inputAmount
	case "min", "minute", "minutes":
		amount = inputAmount * secondsPerMinute
	case "h", "hour", "hours":
		amount = inputAmount * secondsPerHour
	case "day", "days":
		amount = inputAmount * secondsPerDay
	case "week", "weeks":
		amount = inputAmount * secondsPerWeek
	default:
		return 0, fmt.Errorf("unsupported source unit: %s", from)
	}
	switch to {
	case "ns", "nanosecond", "nanoseconds":
		return amount * nanosecondsPerSecond, nil
	case "μs", "microsecond", "microseconds":
		return amount * microsecondsPerSecond, nil
	case "ms", "millisecond", "milliseconds":
		return amount * millisecondsPerSecond, nil
	case "s", "second", "seconds":
		return amount, nil
	case "min", "minute", "minutes":
		return amount / secondsPerMinute, nil
	case "h", "hour", "hours":
		return amount / secondsPerHour, nil
	case "day", "days":
		return amount / secondsPerDay, nil
	case "week", "weeks":
		return amount / secondsPerWeek, nil
	default:
		return 0, fmt.Errorf("unsupported target unit: %s", to)
	}
}

func GetUnits() []string {
	return []string{
		"ns", "nanosecond", "nanoseconds",
		"μs", "us", "microsecond", "microseconds",
		"ms", "millisecond", "milliseconds",
		"s", "second", "seconds",
		"min", "minute", "minutes",
		"h", "hour", "hours",
		"day", "days",
		"week", "weeks",
	}
}
