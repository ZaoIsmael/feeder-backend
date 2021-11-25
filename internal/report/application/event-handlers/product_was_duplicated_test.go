package event_handlers

import (
	domain2 "deporvillage-feeder-backend/internal/cross-cutting/domain"
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	"deporvillage-feeder-backend/internal/report/domain"
	"deporvillage-feeder-backend/internal/report/infrastructure"
	"testing"
)

func setupProductWasDuplicated(rm map[string]domain.Report) (ProductWasDuplicatedApplicationService, infrastructure.ReportRepository) {
	repository := infrastructure.NewReportRepository(rm)
	return CreateProductWasDuplicatedApplicationService(repository), *repository
}

func TestProductWasDuplicated(t *testing.T) {
	// arrange
	rm := map[string]domain.Report{}
	as, r := setupProductWasDuplicated(rm)
	sku, _ := domain2.CreateSKU("ABCD-1234")

	as.Execute(events.ProductWasDuplicated{ProductSKU: sku})

	re, _ := r.Find(domain.ReportId{Value: "1"})

	if re.GetCounterProductDuplicated() != 1 {
		t.Errorf("error when running application service")
	}
}
