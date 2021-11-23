package event_handlers

import (
	domain2 "deporvillage-feeder-backend/internal/cross-cutting/domain"
	pkg "deporvillage-feeder-backend/internal/cross-cutting/infrastructure"
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	"testing"
)

func setupProductWasAdded(d map[string]struct{}) (ProductWasAddedApplicationService, pkg.InMemoryRegisterProduct) {
	rp := pkg.CreateInMemoryRegisterProduct(d)
	return CreateProductWasAddedEventHandler(rp), *rp
}

func TestProductWasAdded(t *testing.T) {
	// arrange
	rm := make(map[string]struct{})
	as, r := setupProductWasAdded(rm)

	sku, _ := domain2.CreateSKU("ABCD-1234")

	as.Execute(events.ProductWasAdded{ProductSKU: sku})

	se := r.Exist(sku)

	if !se {
		t.Errorf("error when running application service")
	}
}
