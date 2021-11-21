package domain

import (
	"deporvillage-backend/internal/inventory/domain/events"
	"deporvillage-backend/pkg/domain"
	"sync"
)

type Inventory struct {
	domain.AggregateRoot
	mu       *sync.Mutex
	id       InventoryId
	products map[string]Product
}

func CreateInventory() *Inventory {
	i := &Inventory{id: InventoryId{`1`}, products: make(map[string]Product)}
	// Register domain event

	i.AggregateRoot.RegisterEvent(
		events.InventoryWasCreated{InventoryId: i.id.Value},
	)

	return i
}

func (i Inventory) AddProduct(sku string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	p, err := CreateProduct(sku)

	if err != nil {
		// Register domain event of invalid product
		i.AggregateRoot.RegisterEvent(
			events.ProductWasInvalid{},
		)

		return
	}

	_, ok := i.products[p.sku.value]

	if ok {
		// Register domain event of to exist product
		i.AggregateRoot.RegisterEvent(
			events.ProductWasDuplicated{ProductSKU: p.sku.value},
		)

		return
	}

	i.products[p.sku.value] = p
	// Register domain event of the add product
	i.AggregateRoot.RegisterEvent(
		events.ProductWasAdded{ProductSKU: p.sku.value},
	)
}
