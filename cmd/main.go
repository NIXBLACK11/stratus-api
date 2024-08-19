package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"stratus-api/color"
	"stratus-api/database"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	PORT := os.Getenv("PORT")
	err := database.InitMongoDB()
	if err != nil {
		log.Fatal(fmt.Sprintf("%s%v%s", color.Red, err.Error(), color.Reset))
	}

	mux := mux.NewRouter()

	// Route to test server
	mux.HandleFunc("/", routes.TestUp)

	mux.HandleFunc("/signup", routes.Signuphandler)

	mux.HandleFunc("/login", routes.LoginHandler)
	
	mux.HandleFunc("/{username}/projects", routes.UserProjects)

	mux.HandleFunc("/{username}/{project}", routes.ProjectURLs)

	mux.HandleFunc("/{username}/addProject", middlewares.AuthorizationMiddleware(routes.AddProject))

	mux.HandleFunc("/{username}/removeProject", middlewares.AuthorizationMiddleware(routes.RemoveProject))

	mux.HandleFunc("/{username}/checkProject", middlewares.AuthorizationMiddleware(routes.CheckProject))

	mux.HandleFunc("/{username}/validate", middlewares.AuthorizationMiddleware(routes.ValidUser))

	// CORS middleware
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	fmt.Printf("%sServer listening on port %d%s", color.Green, PORT, color.Reset)

	err = http.ListenAndServe(":"+PORT, handlers.CORS(headers, origins, methods)(mux))
	if err != nil {
		log.Fatal(fmt.Sprint("%sError starting server%s", color.Red, color.Reset))
	}
}