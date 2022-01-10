package route

import (
	"test/controllers"

	"github.com/gorilla/mux"
)

func AuthenRoute(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/token", controllers.GetNewToken).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")
}
