package domain

import (
	"deporvillage-backend/internal/inventory/domain/events"
	"deporvillage-backend/pkg/domain"
)

type Inventory struct {
	domain.AggregateRoot
	Id       InventoryId
	products map[string]Product
}

func CreateInventory() Inventory {
	i := Inventory{Id: InventoryId{"1"}, products: make(map[string]Product)}

	i.AggregateRoot.RegisterEvent(
		events.InventoryWasCreated{InventoryId: i.Id.Value},
	)

	return i
}

func (i Inventory) AddProduct(sku string) {
	p, err := CreateProduct(sku)

	if err != nil {
		i.AggregateRoot.RegisterEvent(
			events.ProductWasInvalid{},
		)

		return
	}

	_, ok := i.products[p.sku.value]

	if ok {
		i.AggregateRoot.RegisterEvent(
			events.ProductWasDuplicated{ProductSKU: p.sku.value},
		)

		return
	}

	i.products[p.sku.value] = p

	i.AggregateRoot.RegisterEvent(
		events.ProductWasAdded{ProductSKU: p.sku.value},
	)
}
