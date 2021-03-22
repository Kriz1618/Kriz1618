package main

import (
	"log"
	"net/http"

	"github.com/Kriz1618/crud-api-rest-go/commons"
	"github.com/Kriz1618/crud-api-rest-go/routes"
	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Server running on port 9000")
	log.Println(server.ListenAndServe())
}
