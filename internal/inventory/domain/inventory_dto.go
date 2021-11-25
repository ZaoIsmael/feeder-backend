package domain

import "deporvillage-feeder-backend/internal/cross-cutting/domain"

type InventoryDTO struct {
	Id       string
	Products []string
}

func (i InventoryDTO) ToDomain() (Inventory, error) {
	pd := map[string]Product{}

	for _, s := range i.Products {
		sku, err := domain.CreateSKU(s)
		if err != nil {
			return Inventory{}, err
		}

		pd[s] = Product{sku: sku}
	}

	return Inventory{id: InventoryId{i.Id}, products: pd}, nil
}
