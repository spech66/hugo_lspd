package collector

import "github.com/spech66/hugo_lspd/helper"

// Strengthtraining contains csv/json data
type Strengthtraining struct {
	Date     string `json:"date"`
	Exercise string `json:"exercise"`
	Reps     string `json:"reps"`
	Weight   string `json:"weight"`
	Notes    string `json:"notes"`
	Rating   string `json:"rating"`
}

// StrengthtrainingGetDates returns all defined dates
func StrengthtrainingGetDates(filename string) []string {
	lines := helper.ReadAllDataFromCSV(filename, ';')
	return helper.DatesFromLines(lines)
}

// GetStrengthtraining returns all strength training data
func GetStrengthtraining(filename string) []Strengthtraining {
	lines := helper.ReadAllDataFromCSV(filename, ';')

	var strengthtrainings []Strengthtraining
	firstLine := true
	for _, line := range lines {
		data := Strengthtraining{
			Date:     line[0],
			Exercise: line[1],
			Reps:     line[2],
			Weight:   line[3],
			Notes:    line[4],
			Rating:   line[5],
		}
		if !firstLine {
			strengthtrainings = append(strengthtrainings, data)
		}
		firstLine = false
	}

	return strengthtrainings
}
