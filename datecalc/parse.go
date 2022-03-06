package datecalc

import (
	"fmt"
	"regexp"
	"strconv"

	"errors"
)

// ParseDate converts a date written in "d/m/yyyy" form into year, month and day integers.
func ParseDate(s string) (int, int, int, error) {
	re := regexp.MustCompile(`^([0-9]+)/([0-9]+)/([0-9]+)$`)

	parts := re.FindStringSubmatch(s)

	if parts == nil {
		return 0, 0, 0, errors.New("date is not in format DD/MM/YYYY")
	}

	year, err := parseInt(parts[3], firstYear, lastYear)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("year part of date cannot be parsed: %w", err)
	}

	month, err := parseInt(parts[2], 1, 12)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("month part of date cannot be parsed: %w", err)
	}

	lom := lengthOfMonth(year, month)

	day, err := parseInt(parts[1], 1, lom)

	if err != nil {
		return 0, 0, 0, fmt.Errorf("day part of date cannot be parsed: %w", err)
	}

	return year, month, day, nil
}

// parseInt parses a string and returns an int if it is a valid integer between min and max,
// otherwise it returns an error.
func parseInt(s string, min int, max int) (int, error) {
	value, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("parsing error: %w", err)
	}

	if value < int64(min) || value > int64(max) {
		return 0, fmt.Errorf("value \"%d\" is out of range [%d,%d]", value, min, max)
	}

	return int(value), nil
}
