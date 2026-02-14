package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

const whenDateLayout = "2006/01/02"

type ScheduleItem struct {
	When      string `yaml:"when"`
	Milestone string `yaml:"milestone"`
	Who       string `yaml:"who"`
	What      string `yaml:"what"`
	Notes     string `yaml:"notes"`
}

var (
	schedulePath string
)

func main() {
	flag.StringVar(&schedulePath, "schedule-path", "", "path to the folder that contains the yaml file and where the output will be generated")
	flag.Parse()

	scheduleDatafilePath := filepath.Join(schedulePath, "schedule.yaml")

	yamlFile, err := os.ReadFile(scheduleDatafilePath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var schedule []ScheduleItem
	err = yaml.Unmarshal(yamlFile, &schedule)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	// Calculate column widths
	headers := []string{"When", "Milestone", "Who", "What", "Notes"}
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(h) + 4 // +4 for "**" and spaces
	}

	for _, item := range schedule {
		if len(item.When) > widths[0] {
			widths[0] = len(item.When)
		}
		if len(item.Milestone) > widths[1] {
			widths[1] = len(item.Milestone)
		}
		if len(item.Who) > widths[2] {
			widths[2] = len(item.Who)
		}
		if len(item.What) > widths[3] {
			widths[3] = len(item.What)
		}
		if len(item.Notes) > widths[4] {
			widths[4] = len(item.Notes)
		}
	}

	// Adjust width for month names
	months := make(map[time.Month]bool)
	for _, item := range schedule {
		t, err := time.Parse(whenDateLayout, item.When)
		if err == nil {
			months[t.Month()] = true
		}
	}
	for m := range months {
		monthName := m.String()
		if len(monthName)+4 > widths[0] {
			widths[0] = len(monthName) + 4
		}
	}

	var builder strings.Builder

	// Write header
	for i, h := range headers {
		builder.WriteString(fmt.Sprintf("| **%s**", h))
		builder.WriteString(strings.Repeat(" ", widths[i]-len(h)-3))
	}
	builder.WriteString("|\n")

	// Write separator
	for _, w := range widths {
		builder.WriteString(fmt.Sprintf("|%s", strings.Repeat("-", w+2)))
	}
	builder.WriteString("|\n")

	// Write rows
	var currentMonth time.Month
	for _, item := range schedule {
		t, err := time.Parse(whenDateLayout, item.When)
		if err == nil {
			if t.Month() != currentMonth {
				currentMonth = t.Month()
				monthStr := currentMonth.String()

				// Month separator row
				// Col 1
				builder.WriteString(fmt.Sprintf("| **%s**", monthStr))
				builder.WriteString(strings.Repeat(" ", widths[0]-len(monthStr)-3))

				// Other columns (empty)
				for i := 1; i < len(headers); i++ {
					builder.WriteString("|   ") // `|   ` is 3 chars.
					builder.WriteString(strings.Repeat(" ", widths[i]-1))
				}
				builder.WriteString("|\n")
			}
		}
		fields := []string{item.When, item.Milestone, item.Who, item.What, item.Notes}
		for i, f := range fields {
			if f == "" {
				f = " "
			}
			builder.WriteString(fmt.Sprintf("| %s ", f))
			builder.WriteString(strings.Repeat(" ", widths[i]-len(f)))
		}
		builder.WriteString("|\n")
	}

	// Create a buffer to write the formatted markdown to
	var buf bytes.Buffer
	buf.WriteString(builder.String())

	// Write the buffer to the file
	mdOutputPath := filepath.Join(schedulePath, "schedule.md")
	err = os.WriteFile(mdOutputPath, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Error writing Markdown file: %v", err)
	}

	log.Printf("%s generated successfully", mdOutputPath)
}
