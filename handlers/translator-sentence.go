package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Kaukov/gopher-translator/utils"
)

// TranslatorSentence - handles TranslatorSentence requests
type TranslatorSentence struct {
	logger      *log.Logger
	requestBody struct {
		EnglishSentence string `json:"english-sentence"`
	}
	responseBody struct {
		GopherSentence string `json:"gopher-sentence"`
	}
}

// NewTranslatorSentence - returns a new TranslatorSentence with bound logger
func NewTranslatorSentence(logger *log.Logger) *TranslatorSentence {
	return &TranslatorSentence{logger: logger}
}

// GetRequestBodySentence - returns the sentence that was sent with the request
func (ts *TranslatorSentence) GetRequestBodySentence() string {
	return ts.requestBody.EnglishSentence
}

// GetResponseBodySentence - returns the sentence that was sent with the response
func (ts *TranslatorSentence) GetResponseBodySentence() string {
	return ts.responseBody.GopherSentence
}

func (ts *TranslatorSentence) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&ts.requestBody); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error reading user input word"))

		return
	}

	enc := json.NewEncoder(rw)

	translatedWords := []string{}

	requestSentence := ts.requestBody.EnglishSentence

	// Go through each word until the last sign of the sentence
	for _, w := range strings.Split(requestSentence[:len(requestSentence)-1], " ") {
		word, err := utils.TranslateWord(w)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Error reading user input word"))

			return
		}

		translatedWords = append(translatedWords, word)
	}

	// Merge all translated words in a single string and append the last sign of the sentence
	ts.responseBody.GopherSentence = strings.Join(translatedWords, " ") + string(requestSentence[len(requestSentence)-1])

	if err := enc.Encode(ts.responseBody); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))

		return
	}
}
