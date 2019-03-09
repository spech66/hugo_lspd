package helper

import (
	"encoding/json"
	"io/ioutil"
)

// Config holds all configuration data
type Config struct {
	WeightData           string `json:"weightdata"`
	JournalPath          string `json:"journalpath"`
	EnduranceworkoutData string `json:"enduranceworkoutdata"`
	StrengthtrainingData string `json:"strengthtrainingdata"`
	DaylioData           string `json:"dayliodata"`
	HabitsData           string `json:"habitsdata"`
}

// GetConfig to read config from json
func GetConfig(filename string) Config {
	config := Config{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}
