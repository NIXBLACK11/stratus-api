package routes

import (
	"encoding/json"
	"net/http"
	"stratus-api/actions"

	"github.com/gorilla/mux"
)

func UserProjectsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		username := vars["username"]

		projects, err := actions.GetProjects(username)
		if err!=nil {
			http.Error(w, "Failed to fetch projects", http.StatusInternalServerError)
		}

		if len(projects)>0 {
			response := map[string][]string{"projects": projects}
			responseJSON, err := json.Marshal(response)
			if err!=nil {
				http.Error(w, "Failed to create response", http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(responseJSON)
		} else {
			http.Error(w, "Projects not found", http.StatusNotFound)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}