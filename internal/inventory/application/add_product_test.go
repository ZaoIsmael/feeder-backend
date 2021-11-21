package application

import (
	"deporvillage-backend/internal/inventory/domain"
	"deporvillage-backend/internal/inventory/infrastructure"
	pkg "deporvillage-backend/pkg/infrastructure"
	"errors"
	"testing"
)

func setupTest(im map[string]domain.Inventory) AddProductApplicationService {
	repository := infrastructure.NewInventoryRepository(im)
	return CreateAddProductApplicationService(repository, pkg.InMemoryEventBus{})
}

func TestAddProduct(t *testing.T) {
	// arrange
	im := make(map[string]domain.Inventory)
	as := setupTest(im)

	err := as.execute("LPOS-3241")

	if err != nil {
		t.Errorf("error when running application servirce")
	}
}

func TestAddProductWithSkuInvalid(t *testing.T) {
	// arrange
	im := make(map[string]domain.Inventory)
	as := setupTest(im)

	err := as.execute("LPOS-32411")

	if !errors.Is(err, domain.SkuInvalidError) {
		t.Errorf("error when running application servirce")
	}
}

func TestAddProductDuplicated(t *testing.T) {
	// arrange
	sku := "LPOS-3241"
	im := make(map[string]domain.Inventory)
	id := domain.InventoryId{Value: "1"}
	i := domain.Inventory{
		Id: id,
		Products: map[string]domain.Product{
			sku: {Sku: domain.SKU{Value: sku}},
		},
	}

	im[id.Value] = i

	as := setupTest(im)
	err := as.execute(sku)

	if !errors.Is(err, domain.ProductDuplicatedError) {
		t.Errorf("error when running application servirce")
	}
}