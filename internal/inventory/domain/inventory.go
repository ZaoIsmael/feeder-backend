package domain

import (
	"deporvillage-feeder-backend/internal/cross-cutting/domain"
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	"errors"
)

type Inventory struct {
	domain.AggregateRoot
	Id       InventoryId
	Products map[string]Product
}

var ProductDuplicatedError = errors.New("the product with SKU is registered")

func CreateInventory() Inventory {
	return Inventory{Id: InventoryId{"1"}, Products: make(map[string]Product)}
}

func (i *Inventory) AddProduct(sku string) error {
	p, err := CreateProduct(sku)

	if err != nil {
		i.RegisterEvent(
			events.ProductWasInvalid{},
		)

		return err
	}

	_, ok := i.Products[p.Sku.Value]

	if ok {
		i.RegisterEvent(
			events.ProductWasDuplicated{ProductSKU: p.Sku.Value},
		)

		return ProductDuplicatedError
	}

	i.Products[p.Sku.Value] = p

	i.RegisterEvent(
		events.ProductWasAdded{ProductSKU: p.Sku.Value},
	)

	return nil
}
