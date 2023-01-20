package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PiglatinTranslator(t *testing.T) {
	require := assert.New(t)

	piglatin := Piglatin{}
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
		transformedWord := piglatin.Translate(testCase.input)
		require.NotNil(transformedWord)
		require.Equal(testCase.output, *transformedWord)
	}
}
