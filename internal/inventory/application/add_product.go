package application

import (
	pkg "deporvillage-feeder-backend/internal/cross-cutting/domain"
	"deporvillage-feeder-backend/internal/inventory/domain"
)

type AddProductApplicationService struct {
	repository domain.InventoryRepository
	bus        pkg.EventBus
}

func CreateAddProductApplicationService(repository domain.InventoryRepository, bus pkg.EventBus) AddProductApplicationService {
	return AddProductApplicationService{repository, bus}
}

func (as AddProductApplicationService) Execute(sku string) error {
	inventory, err := as.repository.Find(domain.InventoryId{Value: "1"})

	if err != nil {
		inventory = domain.CreateInventory()
	}

	de := inventory.AddProduct(sku)

	as.bus.Publish(inventory.Pull())
	as.repository.Save(inventory)

	return de
}
