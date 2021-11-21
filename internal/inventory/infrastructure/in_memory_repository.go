package infrastructure

import (
	"deporvillage-feeder-backend/internal/inventory/domain"
	"errors"
)

type InventoryRepository struct {
	I map[string]domain.Inventory
}

func NewInventoryRepository(i map[string]domain.Inventory) InventoryRepository {
	return InventoryRepository{
		I: i,
	}
}

func (r InventoryRepository) Find(id domain.InventoryId) (domain.Inventory, error) {
	i, ok := r.I[id.Value]

	if ok {
		return i, nil
	}

	return domain.Inventory{}, errors.New("the inventory not exist")
}

func (r *InventoryRepository) Save(is domain.Inventory) {
	r.I[is.Id.Value] = is
}
