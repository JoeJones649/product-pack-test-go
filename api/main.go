package main
import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"github.com/gorilla/mux"
)

type PackConfiguration struct {
	size int
	quantity int
}

type ProductPackConfiguration struct {
	packConfigurations []PackConfiguration
	targetQuantity int
}

func (productPackConfiguration *ProductPackConfiguration) GetWaste() int {
	totalProducts := 0
	for _, packConfiguration := range productPackConfiguration.packConfigurations {
		totalProducts += packConfiguration.size * packConfiguration.quantity
	}
	return totalProducts - productPackConfiguration.targetQuantity
}

func (productPackConfiguration *ProductPackConfiguration) GetNumberOfPacks() int {
	return len(productPackConfiguration.packConfigurations)
}

type Product struct {
	id string
	packSizes []int
}

func (product *Product) GetOrderedPackSizes() []int {
	packSizes := product.packSizes
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
	return packSizes
}

type ProductPackService struct {

}

func (productPackService *ProductPackService) GetPacksForQuantity(product Product, quantity int) ProductPackConfiguration {
	packSizes := product.GetOrderedPackSizes()
	return productPackService.GetProductPackConfigurationForPackSizes(packSizes, quantity)
}

func (productPackService *ProductPackService) GetProductPackConfigurationForPackSizes(packSizes []int, quantity int) ProductPackConfiguration {
	log.Println("packSizes", packSizes)
	productPackConfiguration := ProductPackConfiguration{targetQuantity: quantity};
	remainingQuantity := quantity
	packConfigurations := []PackConfiguration{}
	for index, packSize := range packSizes {
		packConfiguration := PackConfiguration{size: packSize, quantity: 0}
		for remainingQuantity >= packSize {
			log.Println("Inner for", packSize)
			packConfiguration.quantity++
			remainingQuantity -= packSize
		}

		if remainingQuantity > 0 {
			// If it's not the last one
			if index != len(packSizes)-1 && packSize > remainingQuantity {
				projectedWaste := packSize - remainingQuantity
				// Calculate Waste from children.
				childrenSizesConfiguration := productPackService.GetProductPackConfigurationForPackSizes(packSizes[index+1:], remainingQuantity)
				log.Println("childrenSizesConfiguration", childrenSizesConfiguration)
				log.Println(projectedWaste, childrenSizesConfiguration.GetWaste())
				// If the children are going to waste the same or more then opt to use the larger pack.
				if childrenSizesConfiguration.GetWaste() >= projectedWaste {
					log.Println("using big one")
					for remainingQuantity > 0 {
						packConfiguration.quantity++
						remainingQuantity -= packSize
					}
				}
			} else {
				log.Println("It's the last one, using it.")
				packConfiguration.quantity++
				remainingQuantity = 0
			}
		}

		packConfigurations = append(packConfigurations, packConfiguration)
	}
	productPackConfiguration.packConfigurations = packConfigurations
	return productPackConfiguration
}

type ProductService struct {
	products []Product
}

func (service *ProductService) GetProduct(id string) Product {
	return service.products[0]
}

var productService = ProductService{[]Product{Product{id: "x", packSizes: []int{250, 500, 1000, 2000, 5000}}}}
var productPackService = ProductPackService{}

func getProductPacksHandler(w http.ResponseWriter, req *http.Request) {
	requestVariables := mux.Vars(req)
	id := requestVariables["id"]
	quantity, err := strconv.Atoi(req.FormValue("quantity"))

	if err != nil {
		log.Println("fail", err)
	}
	product := productService.GetProduct(id)

	log.Println(productPackService.GetPacksForQuantity(product, quantity))
	log.Println(id)
	js, _ := json.Marshal(map[string]string{
		"test": "test",
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}


func main() {
	// creates a new instance of a mux router
    router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/products/{id}/packs", getProductPacksHandler).Methods("GET").Queries("quantity", "{[0-9]*?}")

	log.Println("Runnings")
	log.Fatal(http.ListenAndServe(":8080", router))
}