package utils

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Message    string `json:"err_message"`
	StatusCode int    `json:"status_code"`
}

// JSONError function to return error
func JSONError(w http.ResponseWriter, err error, code int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	response := map[string]string{"error": message}
	res, _ := json.Marshal(response)
	w.Write(res)
	// json.NewEncoder(w).Encode(err.Error())
	// er, _ := json.Marshal(err)
	// io.WriteString(w, `{"error": "Wrong request"}`)

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
