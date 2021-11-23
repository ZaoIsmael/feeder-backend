package event_handlers

import (
	pkg "deporvillage-feeder-backend/internal/cross-cutting/domain"
	"deporvillage-feeder-backend/internal/inventory/domain/events"
	"deporvillage-feeder-backend/internal/report/domain"
)

type ProductWasDuplicatedApplicationService struct {
	repository domain.ReportRepository
}

func CreateProductWasDuplicatedApplicationService(repository domain.ReportRepository) ProductWasDuplicatedApplicationService {
	return ProductWasDuplicatedApplicationService{repository}
}

func (as ProductWasDuplicatedApplicationService) EventSubscriberName() string {
	return events.ProductWasDuplicatedName()
}

func (as ProductWasDuplicatedApplicationService) Execute(_ pkg.EventDomain) {
	report, err := as.repository.Find(domain.ReportId{Value: "1"})

	if err != nil {
		report = domain.CreateNewReport()
	}

	report.IncrementProductDuplicated()

	as.repository.Save(report)
}
