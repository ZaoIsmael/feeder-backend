package application

import "deporvillage-backend/internal/inventory/domain"
import pkg "deporvillage-backend/pkg/domain"

type AddProductApplicationService struct {
	repository domain.InventoryRepository
	bus        pkg.EventBus
}

func CreateAddProductApplicationService(repository domain.InventoryRepository, bus pkg.EventBus) AddProductApplicationService {
	return AddProductApplicationService{repository, bus}
}

func (as AddProductApplicationService) execute(sku string) {
	inventory, err := as.repository.Find(domain.InventoryId{Value: `1`})

	if err != nil {
		inventory = domain.CreateInventory()
	}

	inventory.AddProduct(sku)

	as.repository.Save(inventory)

	as.bus.Publish(inventory.Pull())
}
