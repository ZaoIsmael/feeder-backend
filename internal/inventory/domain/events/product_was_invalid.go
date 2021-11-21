package events

type ProductWasInvalid struct {}

func (ed ProductWasInvalid) EventName() string {
	return ProductWasInvalidName()
}

func ProductWasInvalidName() string  {
	return "product_was_invalid"
}