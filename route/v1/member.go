package route

import (
	"test/controllers"

	"github.com/gorilla/mux"
)

func MemberRoute(router *mux.Router) {

	router.HandleFunc("/member", controllers.VerifyAccess(controllers.GetAllMember)).Methods("GET")
	router.HandleFunc("/member/{id}", controllers.VerifyAccess(controllers.GetMember)).Methods("GET")
	router.HandleFunc("/member", controllers.VerifyAccess(controllers.PostMember)).Methods("POST")
	router.HandleFunc("/member/{id}", controllers.VerifyAccess(controllers.PutMember)).Methods("PUT")
	router.HandleFunc("/member/{id}", controllers.VerifyAccess(controllers.DelMember)).Methods("DELETE")
}

// func LoginMember(router *mux.Router) {
// 	router.HandleFunc("/login", controllers.LoginMember).Methods("POST")
// 	router.HandleFunc("/logout", controllers.LogoutMember).Methods("POST")
// }
