package event_handlers

import (
	pkg "deporvillage-feeder-backend/internal/cross-cutting/domain"
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	"deporvillage-feeder-backend/internal/report/domain"
)

type ProductWasInvalidApplicationService struct {
	repository domain.ReportRepository
}

func CreateProductWasInvalidApplicationService(repository domain.ReportRepository) ProductWasInvalidApplicationService {
	return ProductWasInvalidApplicationService{repository}
}

func (as ProductWasInvalidApplicationService) EventSubscriberName() string {
	return events.ProductWasInvalidName()
}

func (as ProductWasInvalidApplicationService) Execute(_ pkg.EventDomain) {
	report, err := as.repository.Find(domain.ReportId{Value: "1"})

	if err != nil {
		report = domain.CreateNewReport()
	}

	report.IncrementProductInvalid()

	as.repository.Save(report)
}
