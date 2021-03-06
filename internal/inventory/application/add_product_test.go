package application

import (
	domain2 "deporvillage-feeder-backend/internal/cross-cutting/domain"
	pkg "deporvillage-feeder-backend/internal/cross-cutting/infrastructure"
	"deporvillage-feeder-backend/internal/inventory/domain"
	"deporvillage-feeder-backend/internal/inventory/infrastructure"
	"errors"
	"testing"
)

func setupTest(im map[string]domain.Inventory) AddProductApplicationService {
	repository := infrastructure.NewInventoryRepository(im)
	return CreateAddProductApplicationService(repository, pkg.InMemoryEventBus{})
}

func TestAddProduct(t *testing.T) {
	// arrange
	im := map[string]domain.Inventory{}
	as := setupTest(im)

	err := as.Execute("LPOS-3241")

	if err != nil {
		t.Errorf("error when running application servirce")
	}
}

func TestAddProductWithSkuInvalid(t *testing.T) {
	// arrange
	im := map[string]domain.Inventory{}
	as := setupTest(im)

	err := as.Execute("LPOS-32411")

	if !errors.Is(err, domain2.SkuInvalidError) {
		t.Errorf("error when running application service")
	}
}

func TestAddProductDuplicated(t *testing.T) {
	// arrange
	sku := "LPOS-3241"
	im := map[string]domain.Inventory{}
	i, _ := domain.InventoryDTO{Id: "1", Products: []string{sku}}.ToDomain()

	im[i.Id().Value] = i

	as := setupTest(im)
	err := as.Execute(sku)

	if !errors.Is(err, domain.ProductDuplicatedError) {
		t.Errorf("error when running application service")
	}
}
