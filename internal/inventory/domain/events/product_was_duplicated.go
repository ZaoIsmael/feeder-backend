package events

type ProductWasDuplicated struct {
	ProductSKU string
}

func (ed ProductWasDuplicated) EventName() string {
	return ProductWasDuplicatedName()
}

func ProductWasDuplicatedName() string  {
	return "event.feeder.product_was_duplicated"
}