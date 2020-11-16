package tests

import (
	"testing"

	"github.com/Kaukov/gopher-translator/utils"
)

func TestTranslator(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "hello"},
		{"my", "ymogo"},
		{"xray", "gexray"},
		{"square", "aresquogo"},
	}

	for _, test := range tests {
		if output, err := utils.TranslateWord(test.input); output != test.expected || err != nil {
			if err != nil {
				t.Error(err)

				return
			}

			t.Errorf("Test failed: %v input, %v expected, got: %v\n", test.input, test.expected, output)
		}
	}
}
