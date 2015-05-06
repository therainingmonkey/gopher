package main

import (
	"strings"
)

var bodyLines = make([]string, 3)

// TODO Error handling

func loadPage(url string) (bodyLines []string, err error) {
	url = strings.TrimSpace(url) // Catch any pesky trailing newlines \n
	bodyLines, err = retrieve(url, "\n")
	return bodyLines, err
}

func loadHomepage() (parsedLines [][]string, err error) {
	// TODO update to work with parsed body structure
	if homepage == "welcome" {
		bodyLines = append(bodyLines, "Welcome", "to", "Gopher!")
		parsedLines := parseBody(bodyLines)
		err = nil
		return parsedLines, err
	} else {
		bodyLines, err = retrieve(homepage, "\n")
		parsedLines := parseBody(bodyLines)
		return parsedLines, err
	}
}

func parseBody(bodyLines []string) (parsedLines [][]string) {
	for _, line := range(bodyLines) {
		parsedLine := strings.Split(line, "\t")
		parsedLines = append(parsedLines, parsedLine)
	}
	return parsedLines
}
