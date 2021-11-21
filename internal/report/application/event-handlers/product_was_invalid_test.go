package event_handlers

import (
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	"deporvillage-feeder-backend/internal/report/domain"
	"deporvillage-feeder-backend/internal/report/infrastructure"
	"testing"
)

func setupProductWasInvalid(rm map[string]domain.Report) (ProductWasInvalidApplicationService, infrastructure.ReportRepository) {
	repository := infrastructure.NewReportRepository(rm)
	return CreateProductWasInvalidApplicationService(repository), repository
}

func TestProductWasInvalid(t *testing.T) {
	// arrange
	rm := make(map[string]domain.Report)
	as, r := setupProductWasInvalid(rm)

	as.Execute(events.ProductWasInvalid{})

	re, _ := r.Find(domain.ReportId{Value: "1"})

	if re.CountProductsInvalid != 1 {
		t.Errorf("error when running application servirce")
	}
}
