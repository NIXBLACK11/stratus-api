package routes

import (
	"encoding/json"
	"net/http"
	"stratus-api/actions"

	"github.com/gorilla/mux"
)

func RemoveProjectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		vars := mux.Vars(r)
		username := vars["username"]
		projectname := vars["projectname"]

		success, err := actions.RemoveProject(username, projectname)
		if err!=nil || success==false{
			http.Error(w, "Error in deleting project", http.StatusInternalServerError)
			return
		}
		
		response := map[string]string{"message": "Successfully deleted: "+projectname}
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