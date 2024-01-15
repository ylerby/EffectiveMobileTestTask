package app

import (
	"EffectiveMobileTask/schemas"
	"encoding/json"
	"net/http"
	"os"
)

func (a *App) ErrorResponseWriter(w http.ResponseWriter, httpStatusCode int, errorMessage string) {
	response, err := json.Marshal(schemas.ErrorResponse{
		Error: errorMessage,
	})
	if err != nil {
		os.Exit(1)
	}

	w.WriteHeader(httpStatusCode)
	_, err = w.Write(response)
	if err != nil {
		os.Exit(1)
	}
}

func (a *App) CorrectResponseWriter(w http.ResponseWriter, httpStatusCode int, data interface{}) {
	response, err := json.Marshal(schemas.CorrectResponse{
		Data: data,
	})
	if err != nil {
		os.Exit(1)
	}

	w.WriteHeader(httpStatusCode)
	_, err = w.Write(response)
	if err != nil {
		os.Exit(1)
	}
}
