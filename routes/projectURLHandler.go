package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stratus-api/actions"
	"stratus-api/models"

	"github.com/gorilla/mux"
)

func ProjectURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		username := vars["username"]
		projectname := vars["projectname"]

		fmt.Println(username + ", " + projectname)
		project, err := actions.GetProjectDetails(username, projectname)
		if err!=nil {
			http.Error(w, "Error fetching project", http.StatusNoContent)
			return
		}

		fmt.Println(project.ProjectName)

		response := map[string]models.Project{"projectdetails": project}
		responseJSON, err := json.Marshal(response)
		if err!=nil {
			http.Error(w, "Error generating response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)

	} else {
		http.Error(w, "Methond not allowed", http.StatusMethodNotAllowed)
	}
}