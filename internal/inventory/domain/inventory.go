package domain

import (
	"deporvillage-feeder-backend/internal/cross-cutting/domain"
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	"errors"
)

type Inventory struct {
	domain.AggregateRoot
	id       InventoryId
	products Products
}

type Products map[string]Product
type Skus []string

var ProductDuplicatedError = errors.New("the product with SKU is registered")

func CreateInventory(sl Skus) (Inventory, error) {
	lp := make(map[string]Product)

	for _, s := range sl {
		p, err := CreateProduct(s)

		if err != nil {
			return Inventory{}, err
		}

		lp[p.sku.Value()] = p
	}

	return Inventory{id: InventoryId{"1"}, products: lp}, nil
}

func (i Inventory) Id() InventoryId {
	return i.id
}

func (i *Inventory) AddProduct(sku string) error {
	p, err := CreateProduct(sku)

	if err != nil {
		i.RegisterEvent(
			events.ProductWasInvalid{},
		)

		return err
	}

	_, ok := i.products[p.sku.Value()]

	if ok {
		i.RegisterEvent(
			events.ProductWasDuplicated{ProductSKU: p.sku},
		)

		return ProductDuplicatedError
	}

	i.products[p.sku.Value()] = p

	i.RegisterEvent(
		events.ProductWasAdded{ProductSKU: p.sku},
	)

	return nil
}
