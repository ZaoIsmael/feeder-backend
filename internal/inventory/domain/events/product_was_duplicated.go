package events

import (
	"deporvillage-feeder-backend/internal/cross-cutting/domain"
)

type ProductWasDuplicated struct {
	ProductSKU domain.SKU
}

func (ed ProductWasDuplicated) EventName() string {
	return ProductWasDuplicatedName()
}

func ProductWasDuplicatedName() string  {
	return "event.feeder.product_was_duplicated"
}
