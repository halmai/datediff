package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

const (
	firstYear = 1900
	lastYear  = 2999
)

// calcDiffDays calculates the difference between two dates.
func calcDiffDays(y1, m1, d1, y2, m2, d2 int) (int, error) {
	i1 := indexOfDay(y1, m1, d1)
	i2 := indexOfDay(y2, m2, d2)

	if i1 == i2 {
		return 0, errors.New("dates must be different")
	}

	if i1 > i2 {
		return i1 - i2 - 1, nil
	}

	return i2 - i1 - 1, nil
}

// readDate reads a date from the standard input and returns its component in year, month and day.
func readDate() (int, int, int, error) {
	var reader = bufio.NewReader(os.Stdin)

	date, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, 0, err
	}

	date = strings.TrimSpace(date)

	return parseDate(date)
}

func processDates() error {
	fmt.Printf("Enter the first date in d/m/Y format (like 3/1/1989):")
	y1, m1, d1, err := readDate()
	if err != nil {
		return err
	}

	fmt.Printf("Enter the second date in d/m/Y format (like 31/12/1989):")
	y2, m2, d2, err := readDate()
	if err != nil {
		return err
	}

	diffDays, err := calcDiffDays(y1, m1, d1, y2, m2, d2)
	if err != nil {
		return err
	}

	fmt.Printf("The difference between %d/%d/%d and %d/%d/%d days is %d day(s).", d1, m1, y1, d2, m2, y2, diffDays)

	return nil
}

func main() {
	err := processDates()

	if err != nil {
		fmt.Println("Error during execution:", err)
	}
}
