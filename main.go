package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	firstYear = 1900
	lastYear  = 2999
)

// readDate reads a date from the standard input and returns its component in year, month and day,
// or an error if the date is invalid.
func readDate(rd io.Reader) (string, error) {
	var reader = bufio.NewReader(rd)

	date, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(date), nil
}

func processDates(rd io.Reader) (int, error) {
	fmt.Printf("Enter the first date in d/m/Y format (like 3/1/1989):")

	date1, err := readDate(rd)
	if err != nil {
		return 0, err
	}

	y1, m1, d1, err := parseDate(date1)
	if err != nil {
		return 0, err
	}

	fmt.Printf("Enter the second date in d/m/Y format (like 31/12/1989):")
	date2, err := readDate(rd)
	if err != nil {
		return 0, err
	}

	y2, m2, d2, err := parseDate(date2)
	if err != nil {
		return 0, err
	}

	return calcDiffDays(y1, m1, d1, y2, m2, d2)
}

func main() {
	diffDays, err := processDates(os.Stdin)

	if err != nil {
		fmt.Println("Error during execution:", err)
		return
	}

	if diffDays == 1 {
		fmt.Printf("The difference between the given dates is %d day.", diffDays)
	} else {
		fmt.Printf("The difference between the given dates is %d days.", diffDays)
	}

}
