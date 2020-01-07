package api

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func ErrorJSON(w http.ResponseWriter, r *http.Request, code int, err error) {
	RespondJSON(w, r, code, map[string]string{"error": err.Error()})
}

func handleDefault() http.HandlerFunc {

	type response struct {
		APIVersion string `json:"api_version"`
		AppName    string `json:"app_name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response{
			APIVersion: "1",
			AppName:    "go rest api template",
		}
		RespondJSON(w, r, http.StatusOK, resp)
	}
}
