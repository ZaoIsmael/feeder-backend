package event_handlers

import (
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	"deporvillage-feeder-backend/internal/report/domain"
	pkg "deporvillage-feeder-backend/pkg/domain"
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

func (as ProductWasInvalidApplicationService) Execute(e pkg.EventDomain) {
	report, err := as.repository.Find(1)

	if err != nil {
		report = domain.CreateReport()
	}

	report.IncrementProductInvalid()

	as.repository.Save(report)
}
