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

type InvalidRequestPayload struct {
	message string
}

func GetProductPacksHandler(w http.ResponseWriter, req *http.Request) {
	productService := services.ProductService{[]models.Product{models.Product{Id: "x", PackSizes: []int{250, 500, 1000, 2000, 5000}}}}

	requestVariables := mux.Vars(req)
	id := requestVariables["id"]
	quantity, err := strconv.Atoi(req.FormValue("quantity"))

	if err != nil {
		log.Println("Invalid Quantity", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(InvalidRequestPayload{message: "Invalid quantity supplied"})
	}
	product := productService.GetProduct(id)

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