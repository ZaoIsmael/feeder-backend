package events

import "deporvillage-feeder-backend/internal/cross-cutting/domain"

type ProductWasAdded struct {
	ProductSKU domain.SKU
}

func (ed ProductWasAdded) EventName() string {
	return ProductWasAddedName()
}

func ProductWasAddedName() string {
	return "event.feeder.product_was_added"
}
