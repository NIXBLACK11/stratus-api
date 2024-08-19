package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stratus-api/actions"
	"stratus-api/jwt"
	"stratus-api/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		
		exists, err := actions.CheckUserExists(user.Username, user.Password)

		if err != nil {
			http.Error(w, "Error occurred in user authentication", http.StatusBadRequest)
		} else {
			if exists {
				token, err := jwt.CreateToken(user.Username)
				if err != nil {
					http.Error(w, "Error in user authentication", http.StatusBadRequest)
				}

				response := map[string]string{"token": "Bearer " + token}

				responseJSON, err := json.Marshal(response)
				if err != nil {
					http.Error(w, "Failed to create response", http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")

				w.WriteHeader(http.StatusOK)
				w.Write(responseJSON)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "User does not exist")
			}
		}
	} else {
		// Return an error response for unsupported methods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
