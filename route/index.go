package route

import (
	"net/http"
	"os"

	v1 "free_credit/route/v1"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func IndexRoute() {

	router := mux.NewRouter()

	v1.AuthenRoute(router.PathPrefix("/v1").Subrouter())
	v1.MemberRoute(router.PathPrefix("/v1").Subrouter())

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(headers, methods, origins)(router))

}
