package events

type ProductWasAdded struct {
	ProductSKU string
}

func (ed ProductWasAdded) EventName() string {
	return ProductWasAddedName()
}

func ProductWasAddedName() string {
	return "event.feeder.product_was_added"
}
