/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright the KubeVirt Authors.
 *
 */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
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
		log.Fatalf("error unmarshalling YAML: %v", err)
	}

	var data [][]string

	var currentMonth time.Month
	for _, item := range schedule {
		t, err := time.Parse(whenDateLayout, item.When)
		if err != nil {
			log.Fatalf("failed to parse date %s as %s", item.When, whenDateLayout)
		}

		// add a Month separator for each new month
		if t.Month() != currentMonth {
			currentMonth = t.Month()
			data = append(data, []string{mdBold(currentMonth.String()), "", "", "", ""})
		}
		data = append(
			data,
			[]string{item.When, item.Milestone, item.Who, item.What, item.Notes},
		)
	}

	output, err := markdown.NewTableFormatterBuilder().
		WithPrettyPrint().
		Build(mdBolds("When", "Milestone", "Who", "What", "Notes")...).
		Format(data)
	if err != nil {
		log.Fatalf("error formatting table: %v", err)
	}

	var buf bytes.Buffer
	buf.WriteString(output)

	mdOutputPath := filepath.Join(schedulePath, "schedule.md")
	err = os.WriteFile(mdOutputPath, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Error writing Markdown file: %v", err)
	}

	log.Printf("%s generated successfully", mdOutputPath)
}

func mdBolds(items ...string) []string {
	var result []string
	for _, item := range items {
		result = append(result, mdBold(item))
	}
	return result
}

func mdBold(item string) string {
	return fmt.Sprintf("**%s**", item)
}
