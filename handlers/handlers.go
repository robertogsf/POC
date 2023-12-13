package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/robertogsf/POC/routers"
	"github.com/robertogsf/POC/security"
	"github.com/rs/cors"
)

func Handlerr() {
	router := mux.NewRouter()
	router.HandleFunc("/products", security.WithRol("admin", routers.PostProduct)).Methods("POST")
	router.HandleFunc("/products/{id}", security.WithRol("admin", routers.UpdateProduct)).Methods("PUT")
	router.HandleFunc("/products/{id}", security.WithRol("admin", routers.DeleteProduct)).Methods("DELETE")
	router.HandleFunc("/products/{id}", security.WithRol("admin", routers.GetProduct)).Methods("GET")
	router.HandleFunc("/products/{id}", security.WithRol("almacenero", routers.UpdateProduct)).Methods("PUT")
	router.HandleFunc("/products/{id}", security.WithRol("almacenero", routers.GetProduct)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
