package routes

import (
	"github.com/Kriz1618/crud-api-rest-go/controller"
	"github.com/gorilla/mux"
)

func SetPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/persona/api").Subrouter()
	subRoute.HandleFunc("/all", controller.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controller.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controller.Delete).Methods("DELETE")
	subRoute.HandleFunc("/find/{id}", controller.Get).Methods("GET")
}
