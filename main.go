package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spech66/hugo_lspd/collector"
	"github.com/spech66/hugo_lspd/helper"
)

var flagConfig string

func init() {
	flag.StringVar(&flagConfig, "config", "config.json", "the config file to read from")
	flag.Parse()
}

func generateUniqueDateList(config helper.Config) []string {
	journalDates := collector.JournalGetDates(config.JournalPath)
	enduranceworkoutDates := collector.EnduranceworkoutGetDates(config.EnduranceworkoutData)
	strengthtrainingDates := collector.StrengthtrainingGetDates(config.StrengthtrainingData)
	weightDates := collector.WeightGetDates(config.WeightData)
	daylioDates := collector.WeightGetDates(config.DaylioData)
	habitsDates := collector.WeightGetDates(config.HabitsData)

	// Create a combined date list
	allDates := []string{}
	allDates = append(allDates, journalDates...)
	allDates = append(allDates, enduranceworkoutDates...)
	allDates = append(allDates, strengthtrainingDates...)
	allDates = append(allDates, weightDates...)
	allDates = append(allDates, daylioDates...)
	allDates = append(allDates, habitsDates...)
	allDates = helper.RemoveDuplicatesUnordered(allDates)
	// fmt.Println("Cleaned", allDates)

	return allDates
}

// Check for YEAR directory
func checkCreateYearDirectory(yearPath string) {
	if _, err := os.Stat(yearPath); os.IsNotExist(err) {
		fmt.Println("Create year dir", yearPath)
		err = os.MkdirAll(yearPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	fmt.Println("LifelogSP Hugo Generator")
	fmt.Println("================")
	fmt.Println("Reading config from", flagConfig)

	config := helper.GetConfig(flagConfig)
	hugoPath := "./hugo/content/post/"

	allDates := generateUniqueDateList(config)

	// Get data from all collector in memory if possible (journal is not)
	enduranceworkoutData := collector.GetEnduranceworkout(config.EnduranceworkoutData)
	strengthtrainingData := collector.GetStrengthtraining(config.StrengthtrainingData)
	weightData := collector.GetWeight(config.WeightData)
	daylioData := collector.GetDaylio(config.DaylioData)
	habitsData := collector.GetHabits(config.HabitsData)

	// Create hugo pages for all dates
	for _, date := range allDates {
		fmt.Println("==== Processing:", date, "====")

		// Returns existing journal or empty one
		journalEntry := collector.GetJournalByDate(config.JournalPath, date)
		headlinePage := date
		headlineContent := "# " + date
		if len(journalEntry.Text) > 2 && strings.HasPrefix(journalEntry.Text, "# ") {
			headlinePage = journalEntry.Text[2:strings.Index(journalEntry.Text, "\n")] // Skip "# "
			headlineContent = ""                                                       // journal already contains the headline
		}

		yearPath := filepath.Join(hugoPath, date[:4])
		filename := filepath.Join(yearPath, date) + ".md"
		fmt.Println("  >", filename)

		checkCreateYearDirectory(yearPath)

		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}

		// ===== HEADER =====
		// Write first line and check for errors - then write the rest
		_, err = f.WriteString("---\n")
		if err != nil {
			panic(err)
		}
		f.Sync()

		f.WriteString("title: " + headlinePage + "\n")
		f.WriteString("type: post\n")
		f.WriteString("author: Sebastian\n")
		f.WriteString("date: " + date + "\n")
		f.WriteString("publishdate: " + date + "\n")
		f.WriteString("lastmod: " + date + "\n") // Don't use time.Now().Format("2006-01-02") as this will make the output random
		// f.WriteString("description: Text about this post\n")
		// == TAGS ==
		f.WriteString("tags:\n")
		// habits as tags
		// TODO: make dynamic
		for _, v := range habitsData {
			if v.Date[:10] == date {
				//fmt.Println("Habits Data", v)
				if v.Beweglichkeit > 0 {
					f.WriteString("    - Beweglichkeit\n")
				}
				if v.Fitness > 0 {
					f.WriteString("    - Fitness\n")
				}
				if v.FlowArts > 0 {
					f.WriteString("    - FlowArts\n")
				}
				if v.Schaffen > 0 {
					f.WriteString("    - Schaffen\n")
				}
				if v.Snooze > 0 {
					f.WriteString("    - Snooze\n")
				}
				if v.Wiegen > 0 {
					f.WriteString("    - Wiegen\n")
				}
			}
		}
		// Activities as tags
		for _, v := range daylioData {
			if v.Date[:10] == date {
				//fmt.Println("Daylio Data", v)
				for _, item := range strings.Split(v.Activities, "|") {
					f.WriteString("    - " + item + "\n")
				}
			}
		}
		// == END TAGS ==
		// == CATEGORIES ==
		f.WriteString("categories:\n")
		f.WriteString("    - " + date[:4] + "\n")
		// Weekdays and mood as categories
		for _, v := range daylioData {
			if v.Date[:10] == date {
				//fmt.Println("Daylio Data", v)
				f.WriteString("    - " + v.Mood + "\n")
				f.WriteString("    - " + v.Weekday + "\n")
			}
		}
		// == END CATEGORIES ==
		f.WriteString("---\n\n")
		// ===== END HEADER =====

		// ===== CONTENT =====
		// # Post Header
		f.WriteString(headlineContent + "\n")

		// Content
		f.WriteString(journalEntry.Text + "\n")

		// TODO: Endurance, strenght as hugo blocks
		for _, v := range weightData {
			if v.Date[:10] == date {
				f.WriteString("{{< weight weight=\"" + v.Weight + "\" bmi=\"" + v.Bmi[:5] + "\" >}}\n")
			}
		}

		for _, v := range enduranceworkoutData {
			if v.Date[:10] == date {
				f.WriteString("{{< enduranceworkout exercise=\"" + v.Exercise + "\" duration=\"" + v.Duration + "\" distance=\"" + v.Distance + "\" notes=\"" + v.Notes + "\" rating=\"" + v.Rating + "\" >}}\n")
			}
		}

		stHeader := true
		for _, v := range strengthtrainingData {
			if v.Date[:10] == date {
				if stHeader {
					f.WriteString("{{< strengthtrainingheader >}}\n")
					stHeader = false
				}
				f.WriteString("{{< strengthtraining exercise=\"" + v.Exercise + "\" reps=\"" + v.Reps + "\" weight=\"" + v.Weight + "\" notes=\"" + v.Notes + "\" rating=\"" + v.Rating + "\" >}}\n")
			}
		}
		// ===== END CONTENT =====

		f.Sync()
		f.Close()
	}
}
