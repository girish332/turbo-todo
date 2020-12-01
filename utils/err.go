package utils

import (
	"encoding/json"
	"net/http"
)

// JSONError function to return error
func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

// JSONOk func for status 200 requeests
func JSONOk(w http.ResponseWriter, body interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if body == nil {
		w.Write([]byte("Task Completed"))
	} else {
		json.NewEncoder(w).Encode(body)
	}

}
