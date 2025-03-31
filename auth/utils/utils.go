package utils

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Set("ContentType", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	Response(w, r, code, map[string]string{"error": err.Error()})
}