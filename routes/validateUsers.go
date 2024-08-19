package routes

import (
	"encoding/json"
	"net/http"
)

func ValidUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		responseJSON := map[string]string{
			"message": "successfully validated",
		}

		responseBytes, err := json.Marshal(responseJSON)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
