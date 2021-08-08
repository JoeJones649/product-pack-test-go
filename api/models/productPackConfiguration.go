package models

type ProductPackConfiguration struct {
	PackConfigurations []PackConfiguration
	TargetQuantity int
}

func (productPackConfiguration *ProductPackConfiguration) GetExtraProductCount() int {
	totalProducts := 0
	for _, packConfiguration := range productPackConfiguration.PackConfigurations {
		totalProducts += packConfiguration.Size * packConfiguration.Quantity
	}
	return totalProducts - productPackConfiguration.TargetQuantity
}

func (productPackConfiguration *ProductPackConfiguration) GetNumberOfPacks() int {
	totalPacks := 0
	for _, packConfiguration := range productPackConfiguration.PackConfigurations {
		totalPacks += packConfiguration.Quantity
	}
	return totalPacks
}
