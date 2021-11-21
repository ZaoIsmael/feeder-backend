package event_handlers

import (
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	"deporvillage-feeder-backend/internal/report/domain"
	pkg "deporvillage-feeder-backend/pkg/domain"
)

type ProductWasAddedApplicationService struct {
	repository domain.ReportRepository
}

func CreateProductWasAddedApplicationService(repository domain.ReportRepository) ProductWasAddedApplicationService {
	return ProductWasAddedApplicationService{repository}
}

func (as ProductWasAddedApplicationService) EventSubscriberName() string {
	return events.ProductWasAddedName()
}

func (as ProductWasAddedApplicationService) Execute(e pkg.EventDomain) {
	report, err := as.repository.Find(domain.ReportId{Value: "1"})

	if err != nil {
		report = domain.CreateReport()
	}

	report.IncrementProduct()

	as.repository.Save(report)
}
