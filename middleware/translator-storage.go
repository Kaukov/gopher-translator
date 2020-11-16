package middleware

import (
	"log"
	"net/http"

	"github.com/Kaukov/gopher-translator/handlers"
	"github.com/Kaukov/gopher-translator/utils"
)

var storedData utils.Storage = utils.Storage{
	Words:     make(map[string]string),
	Sentences: make(map[string]string),
}

// NewTranslatorStorage - returns a storage middleware that caches each word
// and sentence that gets translated by the API
func NewTranslatorStorage(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(rw, r)

			switch next.(type) {
			case *handlers.TranslatorWord:
				tw, _ := next.(*handlers.TranslatorWord)
				requestedWord := tw.GetRequestBodyWord()

				if _, ok := storedData.Words[requestedWord]; !ok {
					storedData.Words[requestedWord] = tw.GetResponseBodyWord()
				}

			case *handlers.TranslatorSentence:
				ts, _ := next.(*handlers.TranslatorSentence)

				requestedSentence := ts.GetRequestBodySentence()

				if _, ok := storedData.Words[requestedSentence]; !ok {
					storedData.Words[requestedSentence] = ts.GetResponseBodySentence()
				}
			}
		})
	}
}

// GetStoredData - returns the cached translated words and sentences
func GetStoredData() utils.Storage {
	return storedData
}
