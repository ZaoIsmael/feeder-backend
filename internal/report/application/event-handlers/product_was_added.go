package event_handlers

import (
	pkg "deporvillage-feeder-backend/internal/cross-cutting/domain"
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	"deporvillage-feeder-backend/internal/report/domain"
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

func (as ProductWasAddedApplicationService) Execute(_ pkg.EventDomain) {
	report, err := as.repository.Find(domain.ReportId{Value: "1"})

	if err != nil {
		report = domain.CreateNewReport()
	}

	report.IncrementProduct()

	as.repository.Save(report)
}
