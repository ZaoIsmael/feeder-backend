package infrastructure

import (
	"deporvillage-backend/internal/inventory/domain"
	"errors"
	"strconv"
)

type inventoryRepository struct {
	i map[int]domain.Inventory
}

func NewInventoryRepository() inventoryRepository {
	return inventoryRepository{
		i: make(map[int]domain.Inventory),
	}
}

func (r inventoryRepository) Find(id domain.InventoryId) (domain.Inventory, error) {
	idd, err := strconv.Atoi(id.Value)

	if err != nil {
		return domain.Inventory{}, err
	}

	i, ok := r.i[idd]

	if ok {
		return i, nil
	}

	return domain.Inventory{}, errors.New("the inventory not exist")
}

func (r inventoryRepository) Save(i domain.Inventory) {
	idd, err := strconv.Atoi(i.Id.Value)

	if err != nil {
		return
	}

	i, ok := r.i[idd]

	if ok {
		return
	}

	r.i[idd] = i
}
