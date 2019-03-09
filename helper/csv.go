package helper

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadAllDataFromCSV reads all fields to interface
func ReadAllDataFromCSV(filename string, comma rune) [][]string {
	fmt.Println("Read all data from", filename)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = comma
	r.LazyQuotes = true
	lines, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	return lines
}

// DatesFromLines creates dates arrays from first csv field
func DatesFromLines(lines [][]string) []string {
	var dates []string
	firstLine := true
	for _, line := range lines {
		date := line[0]

		if !firstLine {
			dates = append(dates, date[:10])
		}

		firstLine = false
	}

	return dates
}
