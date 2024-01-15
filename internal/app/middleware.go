package app

import (
	"EffectiveMobileTask/schemas"
	"encoding/json"
	"net/http"
	"os"
)

func (a *App) LoggingMiddleware(httpMethod string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != httpMethod {
			jsonResponse, err := json.Marshal(schemas.ErrorResponse{
				Error: "Method not allowed",
			})
			if err != nil {
				a.Logger.Info("ошибка при сериализации")
				os.Exit(1)
			}

			w.WriteHeader(http.StatusMethodNotAllowed)
			_, err = w.Write(jsonResponse)
			if err != nil {
				a.Logger.Info("ошибка при ответе")
				os.Exit(1)
			}
		} else {
			a.Logger.Debugf("method: %s, request: %v", r.Method, r)
			next.ServeHTTP(w, r)
		}
	}
}
