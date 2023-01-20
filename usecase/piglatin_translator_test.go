package usecase

import (
	"piglatin-translator-api/mocks"
	"piglatin-translator-api/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PiglatinTranslator(t *testing.T) {
	require := assert.New(t)
	mockPiglatinRepo := &mocks.PiglatinDB{}

	piglatin := Piglatin{
		translatorDB: mockPiglatinRepo,
	}

	testCaseList := []struct {
		input  string
		output string
	}{
		{"pig", "igpay"},
		{"latin", "atinlay"},
		{"banana", "ananabay"},
		{"happy", "appyhay"},
		{"friends", "iendsfray"},
		{"smile", "ilesmay"},
		{"glove", "oveglay"},
		{"store", "orestay"},
		{"we are on a riding,in,tandem", "eway areyay onyay ayay idingray,inyay,andemtay"},
		{"we are on a ridin'g,in,tandem", "eway areyay onyay ayay idin'gray,inyay,andemtay"},
	}

	for _, testCase := range testCaseList {
		mockPiglatinRepo.SupplyFindOneTranslationRequest(func(input string) (*model.TranslationRequest, error) {
			return nil, nil
		})
		mockPiglatinRepo.SupplySaveTranslationRequest(func(input, translated string) error {
			return nil
		})

		transformedWord, err := piglatin.Translate(testCase.input)
		require.NotNil(transformedWord)
		require.Nil(err)
		require.Equal(testCase.output, transformedWord.Translation)
	}
}
