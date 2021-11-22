package domain

import "deporvillage-feeder-backend/pkg/domain"

type Product struct {
	Sku domain.SKU
}

func CreateProduct(sku string) (Product, error) {
	SKU, err := domain.CreateSKU(sku)

	if err != nil {
		return Product{}, err
	}

	return Product{SKU}, nil
}
