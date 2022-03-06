package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/halmai/datediff/datecalc"
)

type outputWriter interface {
	Printf(s string)
}

type outputColl struct{}

func (o outputColl) Printf(s string) {
	fmt.Printf(s)
}

// readDate reads a date from the standard input and returns its component in year, month and day,
// or an error if the date is invalid.
func readDate(rd io.Reader) (string, error) {
	var reader = bufio.NewReader(rd)

	d, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(d), nil
}

func processDates(rd io.Reader, out outputWriter) (int, error) {
	out.Printf("Enter the first date in d/m/Y format (like 3/1/1989):")

	date1, err := readDate(rd)
	if err != nil {
		return 0, err
	}

	y1, m1, d1, err := datecalc.ParseDate(date1)
	if err != nil {
		return 0, err
	}

	out.Printf("Enter the second date in d/m/Y format (like 31/12/1989):")
	date2, err := readDate(rd)
	if err != nil {
		return 0, err
	}

	y2, m2, d2, err := datecalc.ParseDate(date2)
	if err != nil {
		return 0, err
	}

	return datecalc.CalcDiffDays(y1, m1, d1, y2, m2, d2)
}

func main() {
	var oc outputColl
	diffDays, err := processDates(os.Stdin, oc)

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
