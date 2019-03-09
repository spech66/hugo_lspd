package collector

import "github.com/spech66/hugo_lspd/helper"

// Daylio contains csv/json data
// full_date,date,weekday,time,mood,activities,note
// 2019-02-17,17. Februar,Sonntag,19:05,Ok,"Lesen | Yoga | Meditieren | Etwas geschaffen | Programmiert | Zuviel falsches gegessen",""
type Daylio struct {
	/*FullDate   string `json:"full_date"`
	Date       string `json:"date"`*/
	Date       string `json:"date"` // full_date field das date
	Weekday    string `json:"weekday"`
	Time       string `json:"time"`
	Mood       string `json:"mood"`
	Activities string `json:"activities"`
	Note       string `json:"note"`
}

// DaylioGetDates returns all defined dates
func DaylioGetDates(filename string) []string {
	lines := helper.ReadAllDataFromCSV(filename, ',')
	return helper.DatesFromLines(lines)
}

// GetDaylio returns all daylio data
func GetDaylio(filename string) []Daylio {
	lines := helper.ReadAllDataFromCSV(filename, ',')

	var daylios []Daylio
	firstLine := true
	for _, line := range lines {
		data := Daylio{
			/*FullDate:   line[0],
			Date:       line[1],*/
			Date:       line[0],
			Weekday:    line[2],
			Time:       line[3],
			Mood:       line[4],
			Activities: line[5],
			Note:       line[6],
		}
		if !firstLine {
			daylios = append(daylios, data)
		}
		firstLine = false
	}

	return daylios
}
