package routes

import (
	"encoding/json"
	"net/http"
	action "stratus-api/actions"
	"stratus-api/models"

	"github.com/gorilla/mux"
)

func AddProjectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		vars := mux.Vars(r)
		username := vars["username"]

		var project models.Project
		err := json.NewDecoder(r.Body).Decode(&project)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		success, err := action.AddProject(username, project)
		if err != nil {
			http.Error(w, "Failed to update project", http.StatusInternalServerError)
			return
		}

		if success==true {
			response := map[string]string{"message": "Update Successfull"}
			responseJSON, err := json.Marshal(response)
			if err!=nil {
				http.Error(w, "Failed to create response", http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseJSON)
			return
		}
	} else {
		http.Error(w, "Mathod not allowed", http.StatusMethodNotAllowed)
	}
}