package event_handlers

import (
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

	as.Execute(events.ProductWasAdded{ProductSKU: "ABCD-1234"})

	re, _ := r.Find(domain.ReportId{Value: "1"})

	if re.CountProducts != 1 {
		t.Errorf("error when running application service")
	}
}
