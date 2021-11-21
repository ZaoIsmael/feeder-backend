package domain

type InventoryRepository interface {
	Find(id InventoryId) (Inventory, error)
	Save(i Inventory)
}
