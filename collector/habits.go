package collector

import (
	"strconv"

	"github.com/spech66/hugo_lspd/helper"
)

// Habit contains csv/json data
// TODO: Make the fields dynamic!!!!!
// Date,Snooze,Wiegen,Beweglichkeit,Lesen,Flow Arts,Fitness,Schaffen,
// 2019-02-17,2,2,1,1,0,1,2,
type Habit struct {
	Date          string `json:"Date"`
	Snooze        int64  `json:"Snooze"`
	Wiegen        int64  `json:"Wiegen"`
	Beweglichkeit int64  `json:"Beweglichkeit"`
	FlowArts      int64  `json:"Flow Arts"`
	Fitness       int64  `json:"activities"`
	Schaffen      int64  `json:"note"`
}

// HabitsGetDates returns all defined dates
func HabitsGetDates(filename string) []string {
	lines := helper.ReadAllDataFromCSV(filename, ',')
	return helper.DatesFromLines(lines)
}

// GetHabits returns all habit data
func GetHabits(filename string) []Habit {
	lines := helper.ReadAllDataFromCSV(filename, ',')

	var habits []Habit
	firstLine := true
	for _, line := range lines {
		snooze, _ := strconv.ParseInt(line[1], 10, 64)
		wiegen, _ := strconv.ParseInt(line[2], 10, 64)
		beweglichkeit, _ := strconv.ParseInt(line[3], 10, 64)
		flowArts, _ := strconv.ParseInt(line[4], 10, 64)
		fitness, _ := strconv.ParseInt(line[5], 10, 64)
		schaffen, _ := strconv.ParseInt(line[6], 10, 64)
		data := Habit{
			Date:          line[0],
			Snooze:        snooze,
			Wiegen:        wiegen,
			Beweglichkeit: beweglichkeit,
			FlowArts:      flowArts,
			Fitness:       fitness,
			Schaffen:      schaffen,
		}
		if !firstLine {
			habits = append(habits, data)
		}
		firstLine = false
	}

	return habits
}
