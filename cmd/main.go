package main

import (
	"regexp"
)

// var vowelRe, _ = regexp.Compile(`([aAeEiIoOuU]){1}`)
var honestRe, _ = regexp.Compile(`[\.\,\s*]`)

func main() {
	input_word := "we are on a riding,in,tandem"

	var properSplit []string
	var preWordStr string
	for _, char := range input_word {
		if honestRe.MatchString(string(char)) {
			properSplit = append(properSplit, preWordStr)
			properSplit = append(properSplit, string(char))
			preWordStr = ""
			continue
		}

		preWordStr += string(char)
	}

	properSplit = append(properSplit, preWordStr)
}
