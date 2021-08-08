package services

import (
	"errors"
	"github.com/JoeJones649/product-pack-test-go/models"
)

type ProductService struct {
	Products []models.Product
}


func (productService *ProductService) GetPacksForQuantity(product models.Product, quantity int) models.ProductPackConfiguration {
	packSizes := product.GetOrderedPackSizes()
	return productService.GetProductPackConfigurationForPackSizes(packSizes, quantity)
}

func (productService *ProductService) GetProductPackConfigurationForPackSizes(packSizes []int, quantity int) models.ProductPackConfiguration {
	productPackConfiguration := models.ProductPackConfiguration{TargetQuantity: quantity};
	remainingQuantity := quantity
	packConfigurations := []models.PackConfiguration{}
	for index, packSize := range packSizes {
		packConfiguration := models.PackConfiguration{Size: packSize, Quantity: 0}

		// Apply full packs while we can.
		for remainingQuantity >= packSize {
			packConfiguration.Quantity++
			remainingQuantity -= packSize
		}

		if remainingQuantity > 0 {
			// If it's not the last one AND the pack size is greater than the remaining quantity.
			if index != len(packSizes)-1 && packSize > remainingQuantity {
				projectedExtraProducts := packSize - remainingQuantity
				// Calculate Waste from children.
				childrenSizesConfiguration := productService.GetProductPackConfigurationForPackSizes(packSizes[index+1:], remainingQuantity)
				// If the children are going to waste the same or more then opt to use the larger pack.
				if childrenSizesConfiguration.GetExtraProductCount() >= projectedExtraProducts {
					for remainingQuantity > 0 {
						packConfiguration.Quantity++
						remainingQuantity -= packSize
					}
				}
			} else {
				// This is the last pack size. We need to use it.
				packConfiguration.Quantity++
				remainingQuantity = 0
			}
		}

		packConfigurations = append(packConfigurations, packConfiguration)
	}
	productPackConfiguration.PackConfigurations = packConfigurations
	return productPackConfiguration
}

func (service *ProductService) GetProduct(id string) (models.Product, error) {
	for _, product := range service.Products {
		if id == product.Id {
			return product, nil
		}
	}
	
	return models.Product{}, errors.New("Cannot find product with given ID")
}