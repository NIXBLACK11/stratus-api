package routes

import (
	"encoding/json"
	"net/http"
)

func TestUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response := map[string]string{
			"message": "Server running...",
		}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to create response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}