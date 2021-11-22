package events

type ProductWasInvalid struct {}

func (ed ProductWasInvalid) EventName() string {
	return ProductWasInvalidName()
}

func ProductWasInvalidName() string  {
	return "event.feeder.product_was_invalid"
}