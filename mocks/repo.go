package mocks

import (
	"errors"
	"piglatin-translator-api/model"
)

type PiglatinDB struct {
	saveTranslationRequestFn    func(input string, translated string) error
	findOneTranslationRequestFn func(input string) (*model.TranslationRequest, error)
	getTranslationRequestsFn    func(page int) (*model.TranslationRequestsSummary, error)
}

func (p *PiglatinDB) SaveTranslationRequest(input string, translated string) error {
	if p.saveTranslationRequestFn == nil {
		return errors.New("no function supplied")
	}

	return p.saveTranslationRequestFn(input, translated)
}

func (p *PiglatinDB) FindOneTranslationRequest(input string) (*model.TranslationRequest, error) {
	if p.findOneTranslationRequestFn == nil {
		return nil, errors.New("no function supplied")
	}
	return p.findOneTranslationRequestFn(input)
}

func (p *PiglatinDB) GetTranslationRequests(page int) (*model.TranslationRequestsSummary, error) {
	if p.getTranslationRequestsFn == nil {
		return nil, errors.New("no function supplied")
	}

	return p.getTranslationRequestsFn(page)
}

func (p *PiglatinDB) SupplySaveTranslationRequest(fn func(input string, translated string) error) {
	p.saveTranslationRequestFn = fn
}

func (p *PiglatinDB) SupplyFindOneTranslationRequest(fn func(input string) (*model.TranslationRequest, error)) {
	p.findOneTranslationRequestFn = fn
}

func (p *PiglatinDB) SupplyGetTranslationRequests(fn func(page int) (*model.TranslationRequestsSummary, error)) {
	p.getTranslationRequestsFn = fn
}
