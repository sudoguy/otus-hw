package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

type WordCounter struct {
	Word  string
	Count int
}

func Top10(text string) []string {
	const MaxResult = 10

	r := regexp.MustCompile(`[-\w]+`)

	// replace all newlines on whitespaces
	words := strings.Split(strings.ReplaceAll(text, "\n", " "), " ")

	counters := make(map[string]*WordCounter)

	for _, word := range words {
		cleaned := strings.ToLower(r.FindString(word))
		// skip empty word and only dash
		if cleaned == "" || cleaned == "-" {
			continue
		}

		// increase word counter on every match
		if _, ok := counters[cleaned]; ok {
			counter := counters[cleaned]
			counter.Count++
		} else {
			counters[cleaned] = &WordCounter{cleaned, 1}
		}
	}

	values := []WordCounter{}

	// convert map to slice of values
	for _, value := range counters {
		values = append(values, *value)
	}
	// sort values by count
	sort.Slice(values, func(i, j int) bool { return values[i].Count > values[j].Count })

	result := []string{}

	for _, value := range values {
		result = append(result, value.Word)
	}

	if len(result) > MaxResult {
		result = result[:MaxResult]
	}

	return result
}
