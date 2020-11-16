package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Kaukov/gopher-translator/utils"
)

// TranslatorHistory - struct for the history of translated words and sentences
//
// Has logger (*log.Logger) and data (utils.Storage)
type TranslatorHistory struct {
	logger *log.Logger
	data   utils.Storage
}

type responseData struct {
	History map[string]string `json:"history"`
}

// NewTranslatorHistory - returns a new TranslatorHistory with bound logger
func NewTranslatorHistory(logger *log.Logger, storedData utils.Storage) *TranslatorHistory {
	return &TranslatorHistory{logger: logger, data: storedData}
}

func (th *TranslatorHistory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(rw)

	responseBody := responseData{map[string]string{}}

	for k, v := range th.data.Words {
		responseBody.History[k] = v
	}

	for k, v := range th.data.Sentences {
		responseBody.History[k] = v
	}

	if err := enc.Encode(responseBody); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))

		return
	}
}
