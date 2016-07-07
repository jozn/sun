package models

import (
	"strings"
)

type TextParser struct {
	Tags      []string
	UserNames []string
}

func NewTextParser() *TextParser {
	return &TextParser{}
}

func (t *TextParser) Parse(text string) {
	tagsMap := make(map[string]int, 30)
	userNamesMap := make(map[string]int, 30)

	words := strings.Fields(text)
	for _, w := range words {
		if len(w) > 0 {
			switch string(w[0]) {
			case `#`:
				tagsMap[w] += 1
			case `@`:
				userNamesMap[w] += 1
			}
		}

	}

	//TODO: combine functinality below in `for` above
	for k, _ := range tagsMap {
		t.Tags = append(t.Tags, k[1:len(k)])
	}

	for k, _ := range userNamesMap {
		t.UserNames = append(t.UserNames, k[1:len(k)])
	}

	// t.Tags = tagsMap
	// t.UserNames = userNamesMap
}
