package routers

import (
	"erpvietnam/crm/controllers"
	"erpvietnam/crm/middleware"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// InitRoutes creates the routes for handling requests.
// This function returns an *mux.Router.
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// Setup static file serving
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))

	//API router
	api := router.PathPrefix("/api").Subrouter()
	api = api.StrictSlash(true)
	api.HandleFunc("/login", controllers.Login).Methods("POST")
	api.Handle("/refresh-token",
		negroni.New(
			negroni.HandlerFunc(middleware.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RefreshToken),
		)).Methods("GET")

	return router
}
