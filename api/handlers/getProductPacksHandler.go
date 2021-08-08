package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/JoeJones649/product-pack-test-go/models"
	"github.com/JoeJones649/product-pack-test-go/services"
)

type PackConfigurationResponse struct {
	NumberOfPacks int
	ExtraProducts int
	PackConfiguration []models.PackConfiguration
}

type ErrorResponseBody struct {
	Message string
}

func GetProductPacksHandler(w http.ResponseWriter, req *http.Request) {
	productService := services.ProductService{
		[]models.Product{
			models.Product{
				Id: "e8864473-91a0-4f4b-9ce6-903d15acce4f",
				PackSizes: []int{250, 500, 1000, 2000, 5000},
			},
		},
	}

	// Get the variables from the request
	requestVariables := mux.Vars(req)
	id := requestVariables["id"]
	quantity, err := strconv.Atoi(req.FormValue("quantity"))

	// return an error if the quantity is not valid.
	if err != nil {
		log.Println("Invalid Quantity", err)
		w.WriteHeader(http.StatusBadRequest)
		writeErrorResponseBody(w, ErrorResponseBody{Message: "Invalid quantity supplied"})
		return
	}

	product, err := productService.GetProduct(id)
	if err != nil {
		log.Println("Product not found", err)
		w.WriteHeader(http.StatusNotFound)
		writeErrorResponseBody(w, ErrorResponseBody{Message: err.Error()})
		return
	}

	packConfigurationResult := productService.GetPacksForQuantity(product, quantity)

	js, _ := json.Marshal(
		PackConfigurationResponse{
			NumberOfPacks: packConfigurationResult.GetNumberOfPacks(),
			ExtraProducts: packConfigurationResult.GetExtraProductCount(),
			PackConfiguration: packConfigurationResult.PackConfigurations,
		},
	)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func writeErrorResponseBody(w http.ResponseWriter, errorResponseBody ErrorResponseBody) {
	js, _ := json.Marshal(errorResponseBody)
	w.Write(js)
}