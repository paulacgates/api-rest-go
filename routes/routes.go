package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paulacgates/api-rest-go/controllers"
	"github.com/paulacgates/api-rest-go/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.AlteraUmaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
