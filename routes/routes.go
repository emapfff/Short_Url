package routes

import (
	"github.com/gorilla/mux"
	"test_ozon/controllers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	//router.HandleFunc("/api/url", _).Methods("POST")
	router.HandleFunc("/api/url", controllers.GetOriginalUrl).Methods("GET")

	return router
}
