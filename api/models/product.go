package models

import (
	"sort"
)

type Product struct {
	Id string
	PackSizes []int
}

func (product *Product) GetOrderedPackSizes() []int {
	packSizes := product.PackSizes
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
	return packSizes
}