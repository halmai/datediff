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

func processDates() error {
	var reader = bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the first date in d/m/Y format (like 3/1/1989):")
	date1, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	y1, m1, d1, err := parseDate(strings.TrimSpace(date1))
	if err != nil {
		return err
	}

	fmt.Printf("Enter the second date in d/m/Y format (like 3/1/1989):")
	date2, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	y2, m2, d2, err := parseDate(strings.TrimSpace(date2))
	if err != nil {
		return err
	}

	diffDays, err := calcDiffDays(y1, m1, d1, y2, m2, d2)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("The diff is %d days", diffDays))

	return nil
}

func main() {
	err := processDates()

	if err != nil {
		fmt.Println("Error during execution:", err)
	}
}
