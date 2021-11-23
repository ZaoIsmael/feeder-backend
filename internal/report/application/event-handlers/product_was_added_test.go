package event_handlers

import (
	domain2 "deporvillage-feeder-backend/internal/cross-cutting/domain"
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	"deporvillage-feeder-backend/internal/report/domain"
	"deporvillage-feeder-backend/internal/report/infrastructure"
	"testing"
)

func setupProductWasAdded(rm map[string]domain.Report) (ProductWasAddedApplicationService, infrastructure.ReportRepository) {
	repository := infrastructure.NewReportRepository(rm)
	return CreateProductWasAddedApplicationService(repository), *repository
}

func TestProductWasAdded(t *testing.T) {
	// arrange
	rm := make(map[string]domain.Report)
	as, r := setupProductWasAdded(rm)

	sku, _ := domain2.CreateSKU("ABCD-1234")

	as.Execute(events.ProductWasAdded{ProductSKU: sku})

	re, _ := r.Find(domain.ReportId{Value: "1"})

	if re.GetCounterProduct() != 1 {
		t.Errorf("error when running application service")
	}
}
