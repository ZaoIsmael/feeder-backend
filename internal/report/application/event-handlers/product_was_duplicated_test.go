package event_handlers

import (
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
	rm := make(map[string]domain.Report)
	as, r := setupProductWasDuplicated(rm)

	as.Execute(events.ProductWasDuplicated{ProductSKU: "ABCD-1234"})

	re, _ := r.Find(domain.ReportId{Value: "1"})

	if re.CountProductsDuplicated != 1 {
		t.Errorf("error when running application service")
	}
}
