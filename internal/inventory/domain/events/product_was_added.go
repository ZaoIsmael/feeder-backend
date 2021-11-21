package events

type ProductWasAdded struct {
	ProductSKU string
}

func (ed ProductWasAdded) EventName() string {
	return ProductWasAddedName()
}

func ProductWasAddedName() string  {
	return "product_was_added"
}