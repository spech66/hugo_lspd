package collector

import "github.com/spech66/hugo_lspd/helper"

// Enduranceworkout contains csv/json data
type Enduranceworkout struct {
	Date     string `json:"date"`
	Exercise string `json:"exercise"`
	Distance string `json:"distance"`
	Duration string `json:"duration"`
	Notes    string `json:"notes"`
	Rating   string `json:"rating"`
}

// EnduranceworkoutGetDates returns all defined dates
func EnduranceworkoutGetDates(filename string) []string {
	lines := helper.ReadAllDataFromCSV(filename, ';')
	return helper.DatesFromLines(lines)
}

// GetEnduranceworkout returns all endurance workout data
func GetEnduranceworkout(filename string) []Enduranceworkout {
	lines := helper.ReadAllDataFromCSV(filename, ';')

	var enduranceworkouts []Enduranceworkout
	firstLine := true
	for _, line := range lines {
		data := Enduranceworkout{
			Date:     line[0],
			Exercise: line[1],
			Distance: line[2],
			Duration: line[3],
			Notes:    line[4],
			Rating:   line[5],
		}
		if !firstLine {
			enduranceworkouts = append(enduranceworkouts, data)
		}
		firstLine = false
	}

	return enduranceworkouts
}
