package domain

import (
	"deporvillage-backend/internal/inventory/domain/events"
	"deporvillage-backend/pkg/domain"
	"errors"
)

type Inventory struct {
	domain.AggregateRoot
	Id       InventoryId
	Products map[string]Product
}

var ProductDuplicatedError = errors.New("the product with SKU is registered")

func CreateInventory() Inventory {
	i := Inventory{Id: InventoryId{"1"}, Products: make(map[string]Product)}

	i.AggregateRoot.RegisterEvent(
		events.InventoryWasCreated{InventoryId: i.Id.Value},
	)

	return i
}

func (i Inventory) AddProduct(sku string) error {
	p, err := CreateProduct(sku)

	if err != nil {
		i.AggregateRoot.RegisterEvent(
			events.ProductWasInvalid{},
		)

		return err
	}

	_, ok := i.Products[p.Sku.Value]

	if ok {
		i.AggregateRoot.RegisterEvent(
			events.ProductWasDuplicated{ProductSKU: p.Sku.Value},
		)

		return ProductDuplicatedError
	}

	i.Products[p.Sku.Value] = p

	i.AggregateRoot.RegisterEvent(
		events.ProductWasAdded{ProductSKU: p.Sku.Value},
	)

	return nil
}
