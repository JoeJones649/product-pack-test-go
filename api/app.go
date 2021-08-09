package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/JoeJones649/product-pack-test-go/handlers"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Router.HandleFunc(
		"/products/{id}/packs",
		handlers.GetProductPacksHandler,
	).Methods("GET").Queries("quantity", "{[0-9]*?}")
}

func (a *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, a.Router))
}