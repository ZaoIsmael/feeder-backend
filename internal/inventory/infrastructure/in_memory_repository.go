package infrastructure

import (
	"deporvillage-feeder-backend/internal/inventory/domain"
	"errors"
)

type InventoryRepository struct {
	i map[string]domain.Inventory
}

func NewInventoryRepository(i map[string]domain.Inventory) *InventoryRepository {
	return &InventoryRepository{
		i: i,
	}
}

func (r InventoryRepository) Find(id domain.InventoryId) (domain.Inventory, error) {
	i, ok := r.i[id.Value]

	if ok {
		return i, nil
	}

	return domain.Inventory{}, errors.New("the inventory not exist")
}

func (r *InventoryRepository) Save(is domain.Inventory) {
	r.i[is.Id().Value] = is
}
