package event_handlers

import (
	"deporvillage-backend/internal/inventory/domain/events"
	"deporvillage-backend/internal/report/domain"
	pkg "deporvillage-backend/pkg/domain"
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

func (as ProductWasDuplicatedApplicationService) Execute(e pkg.EventDomain) {
	report, err := as.repository.Find(1)

	if err != nil {
		report = domain.CreateReport()
	}

	report.IncrementProductDuplicated()

	as.repository.Save(report)
}
