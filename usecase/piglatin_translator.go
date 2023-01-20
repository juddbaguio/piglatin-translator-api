package usecase

import (
	"piglatin-translator-api/contracts"
	"piglatin-translator-api/model"
	"regexp"
	"strings"
)

var vowelRe, _ = regexp.Compile(`(?i)^[aeiou]`)
var consonantRe, _ = regexp.Compile(`(?i)^([b-df-hj-mnp-tvwxy-z]){0,}`)
var honestWordRe, _ = regexp.Compile(`(?i)^(honest)`)

const vowelSuffix = "yay"
const consonantSuffix = "ay"

type Piglatin struct {
	translatorDB contracts.TranslatorDB
}

func NewPiglatinUsecase(db contracts.TranslatorDB) *Piglatin {
	return &Piglatin{
		translatorDB: db,
	}
}

func (p *Piglatin) Translate(input string) *string {
	var translatedSlice []string

	splitInput := cleanSplit(input, `[\.\,\s*]`)
	for _, word := range splitInput {
		transformedConsonant := transformBeginningConsonantSound(word)
		if transformedConsonant != nil {
			translatedSlice = append(translatedSlice, *transformedConsonant)
			continue
		}

		transformedVowel := transformBeginningVowelSound(word)
		if transformedVowel != nil {
			translatedSlice = append(translatedSlice, *transformedVowel)
			continue
		}

		translatedSlice = append(translatedSlice, word)
	}

	translated := strings.Join(translatedSlice, "")
	return &translated
}

func (p *Piglatin) GetTranslationRequests(page int) ([]model.TranslationRequest, error) {
	return p.translatorDB.GetTranslationRequests(page)
}

func transformBeginningVowelSound(word string) *string {
	if honestWordRe.MatchString(word) {
		return nil
	}

	var beginningStr string = vowelRe.FindString(word)
	if beginningStr == "" {
		return nil
	}

	transformedWord := word + vowelSuffix

	return &transformedWord
}

func transformBeginningConsonantSound(word string) *string {
	if honestWordRe.MatchString(word) {
		return nil
	}
	var beginningStr string = consonantRe.FindString(word)
	if beginningStr == "" {
		return nil
	}

	trimmedWord := strings.Replace(word, beginningStr, "", 1)
	transformedWord := trimmedWord + beginningStr + consonantSuffix

	return &transformedWord
}

func cleanSplit(input string, delimiterPattern string) []string {
	var delimiterRe *regexp.Regexp = regexp.MustCompile(delimiterPattern)
	var properSplit []string
	var preWordStr string
	for _, char := range input {
		if delimiterRe.MatchString(string(char)) {
			properSplit = append(properSplit, preWordStr)
			properSplit = append(properSplit, string(char))
			preWordStr = ""
			continue
		}

		preWordStr += string(char)
	}

	properSplit = append(properSplit, preWordStr)
	return properSplit
}
