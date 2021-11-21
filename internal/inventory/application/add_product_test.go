package application

import (
	"deporvillage-backend/internal/inventory/domain"
	"deporvillage-backend/internal/inventory/infrastructure"
	pkg "deporvillage-backend/pkg/infrastructure"
	"testing"
)

func setupTest() (AddProductApplicationService, domain.InventoryRepository) {
	repository := infrastructure.NewInventoryRepository()
	return CreateAddProductApplicationService(repository, pkg.InMemoryEventBus{}), repository
}

func TestAddProduct(t *testing.T) {
	as, r := setupTest()

	as.execute("LPOS-3241")

	_, err := r.Find(domain.InventoryId{Value: "1"})

	if err != nil {
		t.Errorf("application servirce error to execute")
	}
}
