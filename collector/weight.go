package collector

import "github.com/spech66/hugo_lspd/helper"

// Measurement contains csv/json data
type Measurement struct {
	Date           string `json:"date"`
	Height         string `json:"height"`
	Weight         string `json:"weight"`
	Bmi            string `json:"bmi"`
	BmiOverweight  string `json:"bmioverweight"`
	BmiUnderweight string `json:"bmiunderweight"`
}

// WeightGetDates returns all defined dates
func WeightGetDates(filename string) []string {
	lines := helper.ReadAllDataFromCSV(filename, ';')
	return helper.DatesFromLines(lines)
}

// GetWeight returns all weight data
func GetWeight(filename string) []Measurement {
	lines := helper.ReadAllDataFromCSV(filename, ';')

	var measurements []Measurement
	firstLine := true
	for _, line := range lines {
		data := Measurement{
			Date:           line[0],
			Height:         line[1],
			Weight:         line[2],
			Bmi:            line[3],
			BmiOverweight:  line[4],
			BmiUnderweight: line[5],
		}
		if !firstLine {
			measurements = append(measurements, data)
		}
		firstLine = false
	}

	return measurements
}
