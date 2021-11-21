package events

type InventoryWasCreated struct {
	InventoryId string
}

func (ed InventoryWasCreated) EventName() string {
	return `inventory_was_created`
}
