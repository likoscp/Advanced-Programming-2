package response

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

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
