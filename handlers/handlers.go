package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertogsf/POC/routers"
	"github.com/robertogsf/POC/security"
	"github.com/rs/cors"
)

func Handlerr() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	router := mux.NewRouter()
	router.Use(security.Authenticate)
	router.HandleFunc("/login", routers.Login).Methods("POST")
	router.HandleFunc("/register", routers.PostUser).Methods("POST")
	router.HandleFunc("/searchuser", routers.GetUser).Methods("GET")
	router.HandleFunc("/searchusers", routers.GetUsers).Methods("GET")
	router.HandleFunc("/products", security.WithRol("admin", routers.PostProduct)).Methods("POST")

	router.HandleFunc("/products/{id}", security.WithRol("admin", routers.UpdateProduct)).Methods("PUT")
	router.HandleFunc("/products/{id}", security.WithRol("admin", routers.DeleteProduct)).Methods("DELETE")
	router.HandleFunc("/products/{id}", security.WithRol("admin", routers.GetProduct)).Methods("GET")
	router.HandleFunc("/products/{id}", security.WithRol("almacenero", routers.UpdateProduct)).Methods("PUT")
	router.HandleFunc("/products/{id}", security.WithRol("almacenero", routers.GetProduct)).Methods("GET")

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe("localhost:8000", handler))

}
