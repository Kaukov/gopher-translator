package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Kaukov/gopher-translator/utils"
)

// TranslatorWord - handles TranslatorWord requests
type TranslatorWord struct {
	logger      *log.Logger
	requestBody struct {
		EnglishWord string `json:"english-word"`
	}
	responseBody struct {
		GopherWord string `json:"gopher-word"`
	}
}

// NewTranslatorWord - returns a new TranslatorWord with bound logger
func NewTranslatorWord(logger *log.Logger) *TranslatorWord {
	return &TranslatorWord{logger: logger}
}

// GetRequestBodyWord - returns the word that was sent with the request
func (tw *TranslatorWord) GetRequestBodyWord() string {
	return tw.requestBody.EnglishWord
}

// GetResponseBodyWord - returns the word that was sent with the response
func (tw *TranslatorWord) GetResponseBodyWord() string {
	return tw.responseBody.GopherWord
}

func (tw *TranslatorWord) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error reading user input word"))

		return
	}

	if err := json.NewDecoder(r.Body).Decode(&tw.requestBody); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Error reading user input word"))

		return
	}

	enc := json.NewEncoder(rw)

	var translateError error
	tw.responseBody.GopherWord, translateError = utils.TranslateWord(tw.requestBody.EnglishWord)

	if translateError != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(translateError.Error()))

		return
	}

	if err := enc.Encode(tw.responseBody); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))

		return
	}
}
