package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/JoeJones649/product-pack-test-go/handlers"
)

var app App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    app.Router.ServeHTTP(rr, req)

    return rr
}


func TestGetProductPacksSuccess(t *testing.T) {
	app.Initialize();

	scenarios := [] struct {
		quantity string
		numberOfPacks int
		extraProducts int
		packConfiguration map[int] int
	} {
		{
			quantity: "1",
			numberOfPacks: 1,
			extraProducts: 249,
			packConfiguration: map[int] int {
				250: 1,
			},
		},
		{
			quantity: "250",
			numberOfPacks: 1,
			extraProducts: 0,
			packConfiguration: map[int] int {
				250: 1,
			},
		},
		{
			quantity: "251",
			numberOfPacks: 1,
			extraProducts: 249,
			packConfiguration: map[int] int {
				250: 0,
				500: 1,
			},
		},
		{
			quantity: "501",
			numberOfPacks: 2,
			extraProducts: 249,
			packConfiguration: map[int] int {
				250: 1,
				500: 1,
			},
		},
		{
			quantity: "12001",
			numberOfPacks: 4,
			extraProducts: 249,
			packConfiguration: map[int] int {
				250: 1,
				2000: 1,
				5000: 2,
			},
		},
	}

	for _, scenario := range scenarios {
		url := fmt.Sprintf("/products/%s/packs?quantity=%s", "e8864473-91a0-4f4b-9ce6-903d15acce4f", scenario.quantity)
		req, _ := http.NewRequest("GET", url, nil)
    	response := executeRequest(req)
		log.Println(response.Body)

		// Ensure a 200 is returned.
		if response.Code != http.StatusOK {
			t.Errorf("Response status NOT OK. URL %s: got %v want %v",
                url, response.Code, http.StatusOK)
		}

		packConfigurationResponse := handlers.PackConfigurationResponse{}
		json.Unmarshal([]byte(response.Body.String()), &packConfigurationResponse)

		// Check the number of packs is as expected.
		if packConfigurationResponse.NumberOfPacks != scenario.numberOfPacks {
			t.Errorf("Incorrect number of packs. URL %s: got %v want %v",
                url, packConfigurationResponse.NumberOfPacks, scenario.numberOfPacks)
		}

		// Check the number of extra products is correct.
		if packConfigurationResponse.ExtraProducts != scenario.extraProducts {
			t.Errorf("Incorrect extra products. URL %s: got %v want %v",
                url, packConfigurationResponse.ExtraProducts, scenario.extraProducts)
		}

		// Check that the pack configuration is as expected.
		for _, packConfiguration := range packConfigurationResponse.PackConfiguration {
			expectedCount, keyPresent := scenario.packConfiguration[packConfiguration.Size]
			if !keyPresent {
				expectedCount = 0
			}

			if packConfiguration.Quantity != expectedCount {
				t.Errorf("Incorrect pack quantity. URL %s: got %v want %v",
                url, packConfiguration.Quantity, expectedCount)
			}
		}
	}

}