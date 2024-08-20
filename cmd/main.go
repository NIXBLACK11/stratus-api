package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"stratus-api/color"
	"stratus-api/database"
	"stratus-api/middleware"
	"stratus-api/routes"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	PORTSTR := os.Getenv("PORT")
	PORT, err := strconv.Atoi(PORTSTR)
    if err != nil {
        log.Fatal(fmt.Sprintf("%sError converting string to int%s", color.Red, color.Reset))
    }

	err = database.InitMongoDB()
	if err != nil {
		log.Fatal(fmt.Sprintf("%s%v%s", color.Red, err.Error(), color.Reset))
	}

	mux := mux.NewRouter()

	mux.HandleFunc("/", routes.TestUp)

	mux.HandleFunc("/signup", routes.Signuphandler)

	mux.HandleFunc("/login", routes.LoginHandler)
	
	mux.HandleFunc("/{username}/validate", middleware.AuthorizationMiddleware(routes.ValidUser))

	// mux.HandleFunc("/{username}/projects", routes.UserProjects)

	// mux.HandleFunc("/{username}/{project}", routes.ProjectURLs)

	mux.HandleFunc("/{username}/addProject", middleware.AuthorizationMiddleware(routes.AddProjectHandler))

	// mux.HandleFunc("/{username}/removeProject", middlewares.AuthorizationMiddleware(routes.RemoveProject))

	// mux.HandleFunc("/{username}/checkProject", middlewares.AuthorizationMiddleware(routes.CheckProject))

	// CORS middleware
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	fmt.Printf("%sServer listening on port: %d\n%s", color.Green, PORT, color.Reset)

	err = http.ListenAndServe(":"+PORTSTR, handlers.CORS(headers, origins, methods)(mux))
	if err != nil {
		log.Fatal(fmt.Sprint("%sError starting server%s", color.Red, color.Reset))
	}
}