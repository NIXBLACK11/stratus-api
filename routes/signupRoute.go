package routes

import (
	"encoding/json"
	"net/http"
	"stratus-api/actions"
	"stratus-api/models"
)

func Signuphandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		exists, err := actions.CheckUserExists(user.Username, user.Password)

		if err != nil {
			http.Error(w, "User already exists1", http.StatusBadRequest)
			return
		} else {
			if exists {
				http.Error(w, "User already exists", http.StatusBadRequest)
				return
			} else {
				success, err := actions.CreateUser(user)

				if err != nil {
					http.Error(w, "Failed to create user", http.StatusInternalServerError)
					return
				}

				if success {
					w.WriteHeader(http.StatusOK)
					response := map[string]string{"message": "User created successfully"}
					json.NewEncoder(w).Encode(response)
					return
				} else {
					http.Error(w, "Failed to create user", http.StatusInternalServerError)
					return
				}
			}
		}
		
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
