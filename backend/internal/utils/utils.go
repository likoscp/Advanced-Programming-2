package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, res interface{}, msg, op string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	slog.Info(msg, "op", op)

	if res != nil {
		json.NewEncoder(w).Encode(res)
	}
}

func ResponseError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
	}
}
