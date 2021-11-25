package event_handlers

import (
	"deporvillage-feeder-backend/internal/cross-cutting/domain"
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	"deporvillage-feeder-backend/internal/inventory/infrastructure"
	"testing"
)

func setupProductWasAdded(d map[string]struct{}) (ProductWasAddedApplicationService, infrastructure.InMemoryRegisterProduct) {
	rp := infrastructure.CreateInMemoryRegisterProduct(d)
	return CreateProductWasAddedEventHandler(rp), *rp
}

func TestProductWasAdded(t *testing.T) {
	// arrange
	rm := map[string]struct{}{}
	as, r := setupProductWasAdded(rm)

	sku, _ := domain.CreateSKU("ABCD-1234")

	as.Execute(events.ProductWasAdded{ProductSKU: sku})

	se := r.Exist(sku)

	if !se {
		t.Errorf("error when running application service")
	}
}
