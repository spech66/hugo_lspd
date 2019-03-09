package collector

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Journal contains csv/json data
type Journal struct {
	Date string `json:"date"`
	Text string `json:"text"`
}

// JournalGetDates returns all defined dates
func JournalGetDates(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var dates []string
	for _, f := range files {
		if f.IsDir() {
			subPath := filepath.Join(path, f.Name())
			fmt.Println(subPath)

			subPathFiles, err := ioutil.ReadDir(subPath)
			if err != nil {
				panic(err)
			}

			for _, sf := range subPathFiles {
				if !sf.IsDir() {
					dates = append(dates, strings.Replace(sf.Name(), ".md", "", -1))
				}
			}
		}
	}

	return dates
}

// GetJournalByDate returns specific journal data
func GetJournalByDate(path string, date string) Journal {
	filename := filepath.Join(path, date[:4], date) + ".md"
	fmt.Println("Journal data from", filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Journal not found, returning empty one", filename)
		data := Journal{
			Date: date,
			Text: "",
		}
		return data
	}

	data := Journal{
		Date: date,
		Text: string(content[:]),
	}
	return data
}
