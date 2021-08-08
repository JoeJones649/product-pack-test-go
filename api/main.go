package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/JoeJones649/product-pack-test-go/handlers"
)


func main() {
	// Creates a new instance of a mux router
    router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/products/{id}/packs", handlers.GetProductPacksHandler).Methods("GET").Queries("quantity", "{[0-9]*?}")
	log.Fatal(http.ListenAndServe(":8080", router))
}