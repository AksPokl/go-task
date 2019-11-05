package main

import (
	"log"
	"net/http"

	"github.com/akspokl/go-task/src/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api").Subrouter()
	sub.Methods("POST").Path("/users/register").HandlerFunc(handler.RegisterUser)
	sub.Methods("GET").Path("/users").HandlerFunc(handler.GetUsers)
	sub.Methods("POST").Path("/users/login").HandlerFunc(handler.LoginUser)
	sub.Methods("GET").Path("/github/feed").HandlerFunc(handler.GetGithubEvents)
	sub.Methods("GET").Path("/users/me").Queries("token", "{[0-9]*?}").HandlerFunc(handler.GetCurrentUser)
	log.Fatal(http.ListenAndServe(":3000", router))
}
