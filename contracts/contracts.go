package contracts

import "piglatin-translator-api/model"

type TranslatorDB interface {
	SaveTranslationRequest(input string, translated string) error
	GetTranslationRequests(page int) ([]model.TranslationRequest, error)
}
